package ai

import (
	"context"
	"core/common/redis"
	"core/internal/config"
	"core/internal/service/structure"
	"core/vars"
	"encoding/json"
	"errors"
	"fmt"
	copenai "github.com/sashabaranov/go-openai"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/proxy"
)

type OpenAi struct {
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

func (c *OpenAi) Set() {
	fmt.Println("ConcreteStrategyA do something")
}

const TextModel = "text-davinci-003"
const ChatModel = "gpt-3.5-turbo"
const ChatModelNew = "gpt-3.5-turbo-0301"
const ChatModel4 = "gpt-4"

const MaxToken = 2000
const Temperature = 0.8

const NeedLoopErrorMessage = "Incorrect API key provided"

func NewOpenAi(ctx context.Context, conf config.Config) *OpenAi {
	return &OpenAi{
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

func (c *OpenAi) SetConfig(model string, temperature float32, maxToken int) {
	if model != "" {
		c.Model = model
		c.PostModel = model
	}

	if temperature != 0 {
		c.Temperature = temperature
	}

	if maxToken != 0 {
		c.MaxToken = maxToken
	}
	return
}

func (c *OpenAi) Completion(req string) (string, error) {
	// 打印请求信息
	c.Logger.Info("Completion req: ", req)
	request := copenai.CompletionRequest{
		Model:       copenai.GPT3TextDavinci003,
		Prompt:      req,
		MaxTokens:   c.MaxToken,
		Temperature: c.Temperature,
		TopP:        1,
	}
	var completion copenai.CompletionResponse
	origin, err1 := c.MakeOpenAILoopRequest(&OpenAIRequest{
		FuncName: "CreateCompletion",
		Request:  request,
	})
	if err1 != nil {
		return "", err1
	}
	origin1, ok := origin.(copenai.CompletionResponse)
	if !ok {
		return "", errors.New("conversion failed")
	}
	completion = origin1

	txt := ""
	for _, choice := range completion.Choices {
		txt += choice.Text
	}
	c.Logger.Info("Completion reps: ", txt)
	return txt, nil
}

func (c *OpenAi) buildConfig() copenai.ClientConfig {
	c.WithOpenAIKey()

	config := copenai.DefaultConfig(c.APIKey)
	if c.Origin == "azure" || c.Origin == "azure_ad" {
		config = copenai.DefaultAzureConfig(c.APIKey, c.BaseHost)

		// 默认只使用 一个部署的模型
		config.AzureModelMapperFunc = func(model string) string {
			//azureModelMapping := map[string]string{
			//	defaultModelType: c.Engine,
			//}
			//return azureModelMapping[model]
			return c.Engine
		}
	}
	if c.HttpProxy != "" {
		proxyUrl, _ := url.Parse(c.HttpProxy)
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
		config.HTTPClient = &http.Client{
			Transport: transport,
		}
	} else if c.Socks5Proxy != "" {
		socks5Transport := &http.Transport{}
		dialer, _ := proxy.SOCKS5("tcp", c.Socks5Proxy, nil, proxy.Direct)
		socks5Transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		}
		config.HTTPClient = &http.Client{
			Transport: socks5Transport,
		}
	}

	if c.BaseHost != "" && c.Origin == "open_ai" {
		// trim last slash
		config.BaseURL = strings.TrimRight(c.BaseHost, "/") + "/v1"
	}
	return config
}

func (c *OpenAi) WithKey(key string) *OpenAi {
	if key != "" {
		c.APIKey = key
	}
	return c
}

func (c *OpenAi) WithOpenAIKey() *OpenAi {
	keys := c.APIKeys
	openAiKey, err := redis.Rdb.Get(c.Ctx, vars.OpenAIUsedKey).Result()
	var currentKey int
	if err == nil {
		currentKey, _ = strconv.Atoi(openAiKey)
	}
	c.APIKey = keys[currentKey]
	return c
}

func (c *OpenAi) WithNextOpenAIKey() *OpenAi {
	ctx := c.Ctx
	keys := c.APIKeys
	openAiKey, err := redis.Rdb.Get(ctx, vars.OpenAIUsedKey).Result()
	var currentKey int
	if err == nil {
		currentKey, _ = strconv.Atoi(openAiKey)
	}
	if len(keys) > currentKey+1 {
		redis.Rdb.Incr(ctx, vars.OpenAIUsedKey)
		c.APIKey = keys[currentKey+1]
	} else {
		redis.Rdb.Del(ctx, vars.OpenAIUsedKey)
		c.APIKey = keys[0]
	}
	return c
}

type OpenAIRequest struct {
	Error    error
	FuncName string
	Request  interface{}
}

func (c *OpenAi) MakeOpenAILoopRequest(req *OpenAIRequest) (interface{}, error) {
	//账号有问题
	loopTimes := len(c.APIKeys)
	var resultError error
	for {
		if loopTimes < len(c.APIKeys) {
			if loopTimes < 0 {
				c.Logger.Errorf("MakeOpenAILoopRequest 循环达到最大次数 err:%+v", resultError)
				return "", resultError
			}
			c.Logger.Info(fmt.Sprintf("MakeOpenAILoopRequest loopTimes:%d \n", loopTimes))

			c.WithNextOpenAIKey()
		}
		requestConfig := c.buildConfig()
		requestConfig.APIVersion = "2022-12-01"
		cli := copenai.NewClientWithConfig(requestConfig)

		c.Logger.Info(" MakeOpenAILoopRequest start request:" + fmt.Sprintf("%v", req.Request))

		var result interface{}
		switch req.FuncName {
		case "CreateTranscription":
			result, resultError = cli.CreateTranscription(c.Ctx, req.Request.(copenai.AudioRequest))

		case "CreateCompletion":
			result, resultError = cli.CreateCompletion(c.Ctx, req.Request.(copenai.CompletionRequest))

		case "CreateChatCompletion":
			result, resultError = cli.CreateChatCompletion(c.Ctx, req.Request.(copenai.ChatCompletionRequest))

		case "CreateChatCompletionStream":
			resultStream, resultStreamError := cli.CreateChatCompletionStream(c.Ctx, req.Request.(copenai.ChatCompletionRequest))
			if nil != resultStreamError {
				c.Logger.Error("cli.CreateChatCompletionStream err:" + resultStreamError.Error())
				loopTimes--
				continue
			}
			_, resultError = resultStream.Recv()
			if resultError == nil || !strings.Contains(resultError.Error(), NeedLoopErrorMessage) {
				result = resultStream
			}
		case "CreateEmbeddings":
			result, resultError = cli.CreateEmbeddings(c.Ctx, req.Request.(copenai.EmbeddingRequest))

		default:
			c.Logger.Info(fmt.Sprintln("没有匹配到对应方法" + req.FuncName))
			return nil, fmt.Errorf("没有匹配到对应方法")
		}
		c.Logger.Info(fmt.Sprintf("MakeOpenAILoopRequest dealMethod:%s,params: %+v ,err:%+v , response:%+v \n\n", req.FuncName, req.Request, resultError, result))
		if resultError != nil {
			if strings.Contains(resultError.Error(), NeedLoopErrorMessage) {
				loopTimes--
				continue
			} else {
				return "", resultError
			}
		}

		return result, nil
	}

}

//估算长度

func (c *OpenAi) GetNumTokens(message string) int {
	var messages []structure.ChatModelMessage
	messages = append(messages, structure.ChatModelMessage{
		Role:    "user",
		Content: message,
	})

	return NumTokensFromMessages(messages, c.PostModel)
}

// ChatStream 数据流式传输
func (c *OpenAi) ChatStream(req []structure.ChatModelMessage, channel chan string) (string, error) {
	first := 0
	for i, _ := range req {

		if i%2 == 0 {
			continue
		}
		//估算长度
		if NumTokensFromMessages(req[len(req)-i-1:], c.PostModel) < (3900 - c.MaxToken) {
			first = len(req) - i - 1
		} else {
			break
		}
	}
	var messages []copenai.ChatCompletionMessage

	if first != 0 {
		messages = c.DealMultiplePrompt(req)
	}

	for _, message := range req[first:] {
		messages = append(messages, copenai.ChatCompletionMessage{
			Role:    message.Role,
			Content: message.Content,
		})
	}
	if c.Model == "" || (c.Model != ChatModel && c.Model != ChatModelNew && c.Model != ChatModel4) {
		c.Model = ChatModel
	}
	request := copenai.ChatCompletionRequest{
		Model:       c.PostModel,
		Messages:    messages,
		MaxTokens:   c.MaxToken,
		Temperature: c.Temperature,
		TopP:        1,
	}
	var stream *copenai.ChatCompletionStream
	stream1, err1 := c.MakeOpenAILoopRequest(&OpenAIRequest{
		FuncName: "CreateChatCompletionStream",
		Request:  request,
	})
	if err1 != nil {
		fmt.Println(err1)
		return "", err1
	}
	stream2, ok := stream1.(*copenai.ChatCompletionStream)
	if !ok {
		return "", errors.New("conversion failed")
	}
	stream = stream2
	messageText := ""
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			logx.Info("Stream finished")
			return messageText, nil
		}

		if err != nil {
			fmt.Printf("Stream error: %v\n", err)
			close(channel)
			return messageText, err
		}

		if len(response.Choices) > 0 {
			content := response.Choices[0].Delta.Content
			channel <- content
			// 如果是对内容的进行补充
			messageText += content
			// 结算
			if response.Choices[0].FinishReason != "" {
				close(channel)
				return messageText, nil
			}
		}

		//logx.Info("Stream response:", response)
	}
}

type T struct {
	Type       string
	Properties struct {
		MovieTitle struct {
			Type        string
			Description string
		}
		ReleaseDate struct {
			Type        string
			Description string
		}
		Director struct {
			Type        string
			Description string
		}
	}
	Required []string
}

func (c *OpenAi) Chat(req []structure.ChatModelMessage) (string, error) {

	// 打印请求信息
	logx.Info("req: ", req)

	first := 0
	for i, _ := range req {

		if i%2 == 0 {
			continue
		}
		//估算长度
		if NumTokensFromMessages(req[len(req)-i-1:], c.PostModel) < (3900 - c.MaxToken) {
			first = len(req) - i - 1
		} else {
			break
		}
	}

	var messages []copenai.ChatCompletionMessage
	if first != 0 {
		messages = c.DealMultiplePrompt(req)
	}
	for _, message := range req[first:] {
		messages = append(messages, copenai.ChatCompletionMessage{
			Role:    message.Role,
			Content: message.Content,
		})
	}

	paramsss := T{
		Type: "",
		Properties: struct {
			MovieTitle struct {
				Type        string
				Description string
			}
			ReleaseDate struct {
				Type        string
				Description string
			}
			Director struct {
				Type        string
				Description string
			}
		}{
			MovieTitle: struct {
				Type        string
				Description string
			}{
				Type:        "",
				Description: "",
			},
			ReleaseDate: struct {
				Type        string
				Description string
			}{
				Type:        "",
				Description: "",
			},
			Director: struct {
				Type        string
				Description string
			}{
				Type:        "",
				Description: "",
			},
		},
		Required: []string{
			"ReleaseDate",
			"Director",
			"MovieTitle",
		},
	}
	request := copenai.ChatCompletionRequest{
		Model:       c.PostModel,
		Messages:    messages,
		MaxTokens:   c.MaxToken,
		Temperature: c.Temperature,
		TopP:        1,
		Functions: []copenai.FunctionDefinition{
			{
				Name:        "aaa",
				Description: "aaa",
				Parameters:  paramsss,
			},
		},
	}
	var chat copenai.ChatCompletionResponse
	chatOrigin, err1 := c.MakeOpenAILoopRequest(&OpenAIRequest{
		FuncName: "CreateChatCompletion",
		Request:  request,
	})
	if err1 != nil {
		return "", err1
	}
	chat1, ok := chatOrigin.(copenai.ChatCompletionResponse)
	if !ok {
		return "", errors.New("conversion failed")
	}
	chat = chat1

	txt := ""
	for _, choice := range chat.Choices {
		txt += choice.Message.Content
	}

	return txt, nil
}

func (c *OpenAi) DealMultiplePrompt(req []structure.ChatModelMessage) (res []copenai.ChatCompletionMessage) {

	for _, msg := range req {
		if msg.Role == "system" {
			res = append(res, copenai.ChatCompletionMessage{
				Role:    msg.Role,
				Content: msg.Content,
			})
		}
	}

	return
}

func (c *OpenAi) CreateOpenAIEmbeddings(input string) (response structure.EmbeddingResponse, err error) {

	requestBody := copenai.EmbeddingRequest{
		Model: copenai.AdaEmbeddingV2,
		Input: []string{input},
	}
	var res copenai.EmbeddingResponse

	origin, err1 := c.MakeOpenAILoopRequest(&OpenAIRequest{
		FuncName: "CreateEmbeddings",
		Request:  requestBody,
	})
	if err1 != nil {
		return
	}
	stream2, ok := origin.(copenai.EmbeddingResponse)
	if !ok {
		err = errors.New(" Conversion failed ")
		return
	}
	res = stream2

	var arr []structure.Embedding
	for i, v := range res.Data {
		var arr2 []float64
		for _, embedding := range v.Embedding {
			arr2 = append(arr2, float64(embedding))
		}
		arr = append(arr, structure.Embedding{
			Index:     i,
			Object:    v.Object,
			Embedding: arr2,
		})

	}

	return structure.EmbeddingResponse{
		Object: res.Object,
		Data:   arr,
		Model:  res.Model.String(),
		Usage: structure.EmbeddingUsage{
			PromptTokens: res.Usage.PromptTokens,
			TotalTokens:  res.Usage.TotalTokens,
		},
	}, nil
}

type Speaker interface {
	SpeakToTxt(voiceUrl string) (string, error)
}

func (c *OpenAi) SpeakToTxt(voiceUrl string) (string, error) {
	// 打印文件信息
	logx.Info("File: ", voiceUrl)
	info, err := os.Stat(voiceUrl)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	logx.Info("FileInfo: ", info)

	req := copenai.AudioRequest{
		Model:       copenai.Whisper1,
		FilePath:    voiceUrl,
		Prompt:      "使用简体中文",
		Temperature: 0.5,
		Language:    "zh",
	}
	var resp copenai.AudioResponse
	origin, err1 := c.MakeOpenAILoopRequest(&OpenAIRequest{
		FuncName: "CreateTranscription",
		Request:  req,
	})
	if err1 != nil {
		return "", err1
	}
	origin2, ok := origin.(copenai.AudioResponse)
	if !ok {
		return "", errors.New("conversion failed")
	}
	resp = origin2

	// 用完就删掉
	_ = os.Remove(voiceUrl)

	return resp.Text, nil
}

// GetUsageByKey 获取key的使用情况
func GetUsageByKey(key string, proxyEnable bool, proxyHttp string, proxySocket5 string) (*structure.UsageInfo, error) {
	reqUrl := "https://api.openai.com/v1/dashboard/billing/subscription"
	method := "GET"

	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	if proxyEnable {
		if proxyHttp != "" {
			client.Transport = &http.Transport{
				Proxy: http.ProxyURL(&url.URL{Host: proxyHttp}),
			}

		} else if proxySocket5 != "" {
			dialer, err := proxy.SOCKS5("tcp", proxySocket5, nil, proxy.Direct)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			client.Transport = &http.Transport{
				DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
					return dialer.Dial(network, addr)
				},
			}
		}
	}
	req, err := http.NewRequest(method, reqUrl, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+key)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	//parse body to UsageSubscription
	var usage structure.UsageSubscription
	err = json.Unmarshal(body, &usage)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	//parse UsageSubscription to UsageInfo
	var usageInfo structure.UsageInfo
	usageInfo.HardLimitUsd = usage.HardLimitUsd
	usageInfo.AccountName = usage.AccountName
	//usageInfo.UsedAmountUsd = usage.HardLimitUsd - usage.SoftLimitUsd
	//usageInfo.RemainingAmountUsd = usage.SoftLimitUsd

	//  usage.AccessUntil is a timestamp
	startTime := ""
	endTime := ""
	if time.Now().Local().Unix() > int64(usage.AccessUntil) {
		endTime = time.Unix(int64(usage.AccessUntil), 0).Format("2006-01-02")
		usageInfo.AccessUntil = endTime + "-已过期"
		// 开始时间，在到期时间的前三个月
		startTime = time.Unix(int64(usage.AccessUntil), 0).AddDate(0, -3, 0).Format("2006-01-02")
	} else {
		usageInfo.AccessUntil = time.Unix(int64(usage.AccessUntil), 0).Format("2006-01-02 15:04:05")
		endTime = time.Now().Local().Format("2006-01-02")
		// 开始时间，在当前时间的前三个月
		startTime = time.Now().Local().AddDate(0, -3, 0).Format("2006-01-02")
	}

	//2023-04-01
	reqUrl = fmt.Sprintf("https://api.openai.com/v1/dashboard/billing/usage?start_date=%s&end_date=%s", startTime, endTime)
	req, err = http.NewRequest(method, reqUrl, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+key)

	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(body))
	// parse body to UsageDailyList
	var usageDailyList structure.UsageDailyList
	err = json.Unmarshal(body, &usageDailyList)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// parse UsageDailyList to UsageInfo
	usageInfo.UsedAmountUsd, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", usageDailyList.TotalUsage/100.00), 64)
	usageInfo.RemainingAmountUsd = usageInfo.HardLimitUsd - usageInfo.UsedAmountUsd
	return &usageInfo, nil
}
