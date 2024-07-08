package ai

import (
	"bufio"
	"bytes"
	"context"
	"core/common"
	"core/internal/config"
	"core/internal/service/structure"
	. "core/internal/service/structure"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

// https://open.bigmodel.cn/doc/api#chatglm_std
// chatglm_std, chatglm_lite
// https://open.bigmodel.cn/api/paas/v3/model-api/chatglm_std/invoke
// https://open.bigmodel.cn/api/paas/v3/model-api/chatglm_std/sse-invoke

type ZhipuMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ZhipuRequest struct {
	Prompt      []ZhipuMessage `json:"prompt"`
	Temperature float64        `json:"temperature,omitempty"`
	TopP        float64        `json:"top_p,omitempty"`
	RequestId   string         `json:"request_id,omitempty"`
	Incremental bool           `json:"incremental,omitempty"`
}

type ZhipuResponseData struct {
	TaskId     string         `json:"task_id"`
	RequestId  string         `json:"request_id"`
	TaskStatus string         `json:"task_status"`
	Choices    []ZhipuMessage `json:"choices"`
	Usage      `json:"usage"`
}

type ZhipuResponse struct {
	Code    int               `json:"code"`
	Msg     string            `json:"msg"`
	Success bool              `json:"success"`
	Data    ZhipuResponseData `json:"data"`
}

type ZhipuStreamMetaResponse struct {
	RequestId  string `json:"request_id"`
	TaskId     string `json:"task_id"`
	TaskStatus string `json:"task_status"`
	Usage      `json:"usage"`
}

type zhipuTokenData struct {
	Token      string
	ExpiryTime time.Time
}

var zhipuTokens sync.Map
var expSeconds int64 = 24 * 3600

type ZhiPu struct {
	APIKeys     []string `json:"api_keys"`
	APIKey      string   `json:"api_key"`
	Origin      string   `json:"origin"`
	Engine      string   `json:"engine"`
	HttpProxy   string   `json:"http_proxy"`
	Socks5Proxy string   `json:"socks5_proxy"`
	Model       string   `json:"model"`
	PostModel   string   `json:"post_model"`
	BaseHost    string   `json:"base_host"`
	MaxToken    int      `json:"max_token"`
	Temperature float32  `json:"temperature"`
	Ctx         context.Context
	Config      config.Config
	logx.Logger
}

func NewZhiPu(ctx context.Context, conf config.Config) *ZhiPu {
	return &ZhiPu{
		Config:   conf,
		APIKeys:  conf.OpenAi.Key,
		MaxToken: MaxToken,
		Ctx:      context.Background(),
		Logger:   logx.WithContext(ctx),
		Origin:   conf.OpenAi.Origin,
		Engine:   conf.OpenAi.Engine,
		BaseHost: conf.OpenAi.Host,
	}
}

func (z *ZhiPu) getZhiPuToken(apikey string) string {
	data, ok := zhipuTokens.Load(apikey)
	if ok {
		tokenData := data.(zhipuTokenData)
		if time.Now().Before(tokenData.ExpiryTime) {
			return tokenData.Token
		}
	}

	split := strings.Split(apikey, ".")
	if len(split) != 2 {
		z.Logger.Error("invalid zhipu key: " + apikey)
		return ""
	}

	id := split[0]
	secret := split[1]

	expMillis := time.Now().Add(time.Duration(expSeconds)*time.Second).UnixNano() / 1e6
	expiryTime := time.Now().Add(time.Duration(expSeconds) * time.Second)

	timestamp := time.Now().UnixNano() / 1e6

	payload := jwt.MapClaims{
		"api_key":   id,
		"exp":       expMillis,
		"timestamp": timestamp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	token.Header["alg"] = "HS256"
	token.Header["sign_type"] = "SIGN"

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return ""
	}

	zhipuTokens.Store(apikey, zhipuTokenData{
		Token:      tokenString,
		ExpiryTime: expiryTime,
	})

	return tokenString
}

func (z *ZhiPu) requestOpenAI2Zhipu(request GeneralOpenAIRequest) *ZhipuRequest {
	messages := make([]ZhipuMessage, 0, len(request.Messages))
	for _, message := range request.Messages {
		if message.Role == "system" {
			messages = append(messages, ZhipuMessage{
				Role:    "system",
				Content: message.Content,
			})
			messages = append(messages, ZhipuMessage{
				Role:    "user",
				Content: "Okay",
			})
		} else {
			messages = append(messages, ZhipuMessage{
				Role:    message.Role,
				Content: message.Content,
			})
		}
	}
	return &ZhipuRequest{
		Prompt:      messages,
		Temperature: request.Temperature,
		TopP:        request.TopP,
		Incremental: false,
	}
}

func (z *ZhiPu) responseZhipu2OpenAI(response *ZhipuResponse) *OpenAITextResponse {
	fullTextResponse := OpenAITextResponse{
		Id:      response.Data.TaskId,
		Object:  "chat.completion",
		Created: time.Now().Unix(),
		Choices: make([]OpenAITextResponseChoice, 0, len(response.Data.Choices)),
		Usage:   response.Data.Usage,
	}
	for i, choice := range response.Data.Choices {
		openaiChoice := OpenAITextResponseChoice{
			Index: i,
			Message: Message{
				Role:    choice.Role,
				Content: strings.Trim(choice.Content, "\""),
			},
			FinishReason: "",
		}
		if i == len(response.Data.Choices)-1 {
			openaiChoice.FinishReason = "stop"
		}
		fullTextResponse.Choices = append(fullTextResponse.Choices, openaiChoice)
	}
	return &fullTextResponse
}

func (z *ZhiPu) streamResponseZhipu2OpenAI(zhipuResponse string) *ChatCompletionsStreamResponse {
	var choice ChatCompletionsStreamResponseChoice
	choice.Delta.Content = zhipuResponse
	choice.FinishReason = ""
	response := ChatCompletionsStreamResponse{
		Object:  "chat.completion.chunk",
		Created: time.Now().Unix(),
		Model:   "chatglm",
		Choices: []ChatCompletionsStreamResponseChoice{choice},
	}
	return &response
}

func (z *ZhiPu) streamMetaResponseZhipu2OpenAI(zhipuResponse *ZhipuStreamMetaResponse) (*ChatCompletionsStreamResponse, *Usage) {
	var choice ChatCompletionsStreamResponseChoice
	choice.Delta.Content = ""
	choice.FinishReason = "stop"
	response := ChatCompletionsStreamResponse{
		Id:      zhipuResponse.RequestId,
		Object:  "chat.completion.chunk",
		Created: time.Now().Unix(),
		Model:   "chatglm",
		Choices: []ChatCompletionsStreamResponseChoice{choice},
	}
	return &response, &zhipuResponse.Usage
}

func (z *ZhiPu) zhipuStreamHandler(c *gin.Context, resp *http.Response) (*OpenAIErrorWithStatusCode, *Usage) {
	var usage *Usage
	scanner := bufio.NewScanner(resp.Body)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := strings.Index(string(data), "\n\n"); i >= 0 && strings.Index(string(data), ":") >= 0 {
			return i + 2, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	})
	dataChan := make(chan string)
	metaChan := make(chan string)
	stopChan := make(chan bool)
	go func() {
		for scanner.Scan() {
			data := scanner.Text()
			lines := strings.Split(data, "\n")
			for i, line := range lines {
				if len(line) < 5 {
					continue
				}
				if line[:5] == "data:" {
					dataChan <- line[5:]
					if i != len(lines)-1 {
						dataChan <- "\n"
					}
				} else if line[:5] == "meta:" {
					metaChan <- line[5:]
				}
			}
		}
		stopChan <- true
	}()
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")
	c.Writer.Header().Set("X-Accel-Buffering", "no")
	c.Stream(func(w io.Writer) bool {
		select {
		case data := <-dataChan:
			response := z.streamResponseZhipu2OpenAI(data)
			jsonResponse, err := json.Marshal(response)
			if err != nil {
				z.Logger.Error("error marshalling stream response: " + err.Error())
				return true
			}
			c.Render(-1, common.CustomEvent{Data: "data: " + string(jsonResponse)})
			return true
		case data := <-metaChan:
			var zhipuResponse ZhipuStreamMetaResponse
			err := json.Unmarshal([]byte(data), &zhipuResponse)
			if err != nil {
				z.Logger.Error("error unmarshalling stream response: " + err.Error())
				return true
			}
			response, zhipuUsage := z.streamMetaResponseZhipu2OpenAI(&zhipuResponse)
			jsonResponse, err := json.Marshal(response)
			if err != nil {
				z.Logger.Error("error marshalling stream response: " + err.Error())
				return true
			}
			usage = zhipuUsage
			c.Render(-1, common.CustomEvent{Data: "data: " + string(jsonResponse)})
			return true
		case <-stopChan:
			c.Render(-1, common.CustomEvent{Data: "data: [DONE]"})
			return false
		}
	})
	err := resp.Body.Close()
	if err != nil {
		return errorWrapper(err, "close_response_body_failed", http.StatusInternalServerError), nil
	}
	return nil, usage
}

func (z *ZhiPu) zhipuHandler(c *gin.Context, resp *http.Response) (*OpenAIErrorWithStatusCode, *Usage) {
	var zhipuResponse ZhipuResponse
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return errorWrapper(err, "read_response_body_failed", http.StatusInternalServerError), nil
	}
	err = resp.Body.Close()
	if err != nil {
		return errorWrapper(err, "close_response_body_failed", http.StatusInternalServerError), nil
	}
	err = json.Unmarshal(responseBody, &zhipuResponse)
	if err != nil {
		return errorWrapper(err, "unmarshal_response_body_failed", http.StatusInternalServerError), nil
	}
	if !zhipuResponse.Success {
		return &OpenAIErrorWithStatusCode{
			OpenAIError: OpenAIError{
				Message: zhipuResponse.Msg,
				Type:    "zhipu_error",
				Param:   "",
				Code:    zhipuResponse.Code,
			},
			StatusCode: resp.StatusCode,
		}, nil
	}
	fullTextResponse := z.responseZhipu2OpenAI(&zhipuResponse)
	jsonResponse, err := json.Marshal(fullTextResponse)
	if err != nil {
		return errorWrapper(err, "marshal_response_body_failed", http.StatusInternalServerError), nil
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(resp.StatusCode)
	_, err = c.Writer.Write(jsonResponse)
	return nil, &fullTextResponse.Usage
}

func (z *ZhiPu) Completion(request string) (string, error) {

	return "", nil
}

func (z *ZhiPu) SpeakToTxt(voiceUrl string) (string, error) {

	return "", nil

}

func (z *ZhiPu) CreateOpenAIEmbeddings(input string) (structure.EmbeddingResponse, error) {

	return structure.EmbeddingResponse{}, nil

}

func (z *ZhiPu) Chat(request []structure.ChatModelMessage) (string, error) {
	method := "invoke"
	//if textRequest.Stream {
	//	method = "sse-invoke"
	//}
	fullRequestURL := fmt.Sprintf("https://open.bigmodel.cn/api/paas/v3/model-api/%s/%s", z.Model, method)
	zhipuRequest := z.requestOpenAI2Zhipu(textRequest)
	jsonStr, err := json.Marshal(zhipuRequest)
	if err != nil {
		return "", err
	}
	requestBody := bytes.NewBuffer(jsonStr)

	req, err := http.NewRequest(c.Request.Method, fullRequestURL, requestBody)
	if err != nil {
		return "", err
	}

	return "", nil

}

func (z *ZhiPu) ChatStream(request []structure.ChatModelMessage, channel chan string) (string, error) {

	return "", nil

}

func (z *ZhiPu) SetConfig(model string, temperature float32, maxToken int) {

	return
}
