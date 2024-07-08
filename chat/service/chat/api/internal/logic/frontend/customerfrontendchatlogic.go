package frontend

import (
	"chat/common/milvus"
	"chat/common/openai"
	"chat/common/redis"
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/util/sensitive"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CustomerFrontendChatLogic struct {
	logx.Logger
	ctx                   context.Context
	svcCtx                *svc.ServiceContext
	model                 string
	baseHost              string
	basePrompt            string
	message               string
	postModel             string
	wechatUserRepository  *repository.WechatUserRepository
	chatService           *service.ChatService
	configService         *service.ConfigService
	customerConfigService *service.CustomerConfigService
	riskService           *service.RiskService
}

func NewCustomerFrontendChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustomerFrontendChatLogic {
	return &CustomerFrontendChatLogic{
		Logger:                logx.WithContext(ctx),
		ctx:                   ctx,
		svcCtx:                svcCtx,
		wechatUserRepository:  repository.NewWechatUserRepository(ctx, svcCtx),
		chatService:           service.NewChatService(ctx, svcCtx),
		configService:         service.NewConfigService(ctx, svcCtx),
		customerConfigService: service.NewCustomerConfigService(ctx, svcCtx),
		riskService:           service.NewRiskService(ctx, svcCtx),
	}
}

// CustomerChat 疗愈喵
func (l *CustomerFrontendChatLogic) CustomerChat(req *types.CustomerChatRequest, channel chan string) (response types.CustomerChatResponse, err error) {
	response = types.CustomerChatResponse{}
	if req.Heartbeat {
		//心跳检测
		return
	}
	l.Logger.Info(" CustomerChat start messageId:" + req.MessageID)

	if req.Message == "" || req.OpenKfID == "" || req.MessageID == "" || req.User == "" {
		err = fmt.Errorf("缺少必传参数")
		return
	}
	if req.Message != "" {
		if !sensitive.NewSensitive(l.ctx, l.svcCtx).Check(req.Message) {
			response = types.CustomerChatResponse{
				Message:      sensitive.ErrorMessage,
				MessageID:    req.MessageID,
				IsEnd:        true,
				ErrorMessage: sensitive.ErrorMessage,
			}
			l.Logger.Errorf("SensitiveError user:%s message:%s", req.User, req.Message)
			return response, nil
		}
	}

	enable, _, err := l.riskService.Check(req.User, req.OpenKfID)
	if err != nil {
		return
	}
	if enable == 2 {
		response.Code = xerr.RightsNotHaveTimesError
		return
	} else if enable == 3 {
		response.Code = xerr.RightsNotVip
		return
	}
	l.setModelNameV2().setBasePrompt().setBaseHost()

	// 幂等
	exist, existEmoji, err := l.chatService.Idempotent(req.MessageID, req.User)
	if err != nil {
		return
	}

	if exist != "" {
		return SendToWebsocket(exist, req.MessageID, existEmoji, 0), nil
	}
	questionModel := &model.ChatRecord{
		User:             req.User,
		MessageId:        req.MessageID,
		OpenKfId:         req.OpenKfID,
		Content:          req.Message,
		ChatType:         repository.ChatTypeCustomerFrontend,
		AnswerOrQuestion: repository.AnswerOrQuestionQuestion,
		MessageType:      repository.MessageTypeText,
	}
	_ = l.chatService.Insert(questionModel)

	embeddingEnable := true
	embeddingMode := vars.EMBEDDING_MODEMBEDDING
	var baseScore float32
	var baseTopK int
	var clearContextTime int64
	var postModel string
	//get prompt
	customerConfigPo, err := l.customerConfigService.GetPrompt(req.OpenKfID, req.User)
	if err != nil {
		return
	}
	if customerConfigPo != nil {
		if customerConfigPo.Prompt != "" {
			l.basePrompt = customerConfigPo.Prompt
		}
		embeddingEnable = customerConfigPo.EmbeddingEnable
		embeddingMode = customerConfigPo.EmbeddingMode
		if customerConfigPo.Score.Valid {
			baseScore = float32(customerConfigPo.Score.Float64)
		}
		if customerConfigPo.TopK != 0 {
			baseTopK = int(customerConfigPo.TopK)
		}
		clearContextTime = customerConfigPo.ClearContextTime
		postModel = customerConfigPo.PostModel
	}
	l.setPostModel(postModel)
	openAiKey, err := l.configService.GetConfigKey()
	if err != nil {
		return
	}
	// openai client
	c := openai.NewChatClient(l.ctx, openAiKey).
		WithModel(l.model).
		WithBaseHost(l.baseHost).
		WithOrigin(l.svcCtx.Config.OpenAi.Origin).
		WithEngine(l.svcCtx.Config.OpenAi.Engine).WithPostModel(postModel)
	if l.svcCtx.Config.Proxy.Enable {
		c = c.WithHttpProxy(l.svcCtx.Config.Proxy.Http).WithSocks5Proxy(l.svcCtx.Config.Proxy.Socket5)
	}
	// context
	collection := openai.NewUserContext(l.ctx,
		openai.GetUserUniqueID(req.User, req.OpenKfID),
	).WithPrompt(l.basePrompt).WithModel(l.model).WithClient(c)

	//判断是否需要清除聊天记录
	if clearContextTime > 0 {
		duration := time.Duration(clearContextTime) * time.Minute
		formattedTime := time.Now().Add(-duration).Format("2006-01-02 15:04:05")
		clearStatus, err1 := l.CheckClearContext(l.ctx, req.OpenKfID, req.User, formattedTime)
		if err1 != nil {
			util.Error("CustomerFrontendChatLogic:CustomerChat:读取 stream 失败:" + err1.Error())
			err = err1
			return
		}
		if clearStatus {
			collection.Clear()
			collection.Messages = []openai.ChatModelMessage{}
			collection.Summary = []openai.ChatModelMessage{}
		}
	}

	// 然后 把 消息 发给 openai
	//go func() {
	// 基于 summary 进行补充
	messageText := ""
	if embeddingEnable {

		// md5 this req.MSG to key
		key := md5.New()
		_, _ = io.WriteString(key, req.Message)
		keyStr := fmt.Sprintf("%x", key.Sum(nil))
		type EmbeddingCache struct {
			Embedding []float64 `json:"embedding"`
		}
		milvusService, errMilvus := milvus.InitMilvus(l.svcCtx.Config.Embeddings.Milvus.Host, l.svcCtx.Config.Embeddings.Milvus.Username, l.svcCtx.Config.Embeddings.Milvus.Password)
		if errMilvus != nil {
			fmt.Println(err.Error())
			return response, errMilvus
		}
		defer milvusService.CloseClient()
		embeddingRes, err := redis.Rdb.Get(l.ctx, fmt.Sprintf(redis.EmbeddingsCacheKey, keyStr)).Result()
		var embedding []float64
		if err == nil {
			tmp := new(EmbeddingCache)
			_ = json.Unmarshal([]byte(embeddingRes), tmp)
			embedding = tmp.Embedding
		} else {
			res, err := c.CreateOpenAIEmbeddings(req.Message)
			if err == nil {
				embedding = res.Data[0].Embedding
				// 去将其存入 redis
				embeddingCache := EmbeddingCache{
					Embedding: embedding,
				}
				redisData, err := json.Marshal(embeddingCache)
				if err == nil {
					redis.Rdb.Set(l.ctx, fmt.Sprintf(redis.EmbeddingsCacheKey, keyStr), string(redisData), -1*time.Second)
				}
			}
		}

		if embeddingMode == "QA" {
			// 去通过 embeddings 进行数据匹配
			type EmbeddingData struct {
				Q string `json:"q"`
				A string `json:"a"`
			}
			var embeddingData []EmbeddingData
			result := milvusService.SearchFromQA(embedding, baseTopK)
			tempMessage := ""
			for _, qa := range result {
				if qa.Score > 0.3 {
					continue
				}
				if len(embeddingData) < 2 {
					embeddingData = append(embeddingData, EmbeddingData{
						Q: qa.Q,
						A: qa.A,
					})
				} else {
					tempMessage += qa.Q + "\n"
				}
			}
			//if tempMessage != "" {
			//	go sendToUser(req.OpenKfID, "", req.CustomerID, "正在思考中，也许您还想知道"+"\n\n"+tempMessage, l.svcCtx.Config)
			//go SendToWebsocket("正在思考中，也许您还想知道"+"\n\n"+tempMessage, client)
			//}
			for _, chat := range embeddingData {
				collection.Set(chat.Q, chat.A, false)
			}
			collection.Set(req.Message, "", false)
		} else if embeddingMode == "ARTICLE" {
			//如果是article模式，清理掉上下午，因为文章内容可能会很长
			collection.Clear()
			collection.Messages = []openai.ChatModelMessage{}
			collection.Summary = []openai.ChatModelMessage{}
			// 去通过 embeddings 进行数据匹配
			type EmbeddingData struct {
				Text string `json:"text"`
			}
			var embeddingData []EmbeddingData
			result := milvusService.SearchFromArticle(embedding, baseTopK)

			for _, item := range result {
				fmt.Println("text:", item.CnText)
				fmt.Println("score:", item.Score)
				if item.Score < baseScore {
					continue
				}
				embeddingData = append(embeddingData, EmbeddingData{
					Text: item.CnText,
				})
			}

			if len(embeddingData) > 0 {
				messageText += "When given CONTEXT you answer questions using only that information,and you always format your output in markdown.Answer with chinese.\n\n"
				messageText += "CONTEXT:"
				for _, chat := range embeddingData {
					messageText += chat.Text + "\n\n"
					if c.WithModel(l.model).WithBaseHost(l.baseHost).GetNumTokens(messageText) > vars.MaxToken {
						break
					}
				}
				messageText += "USER QUESTION:" + req.Message
				collection.Set(messageText, "", false)
			} else {
				collection.Set(req.Message, "", false)
			}
		} else {
			collection.Set(req.Message, "", false)
		}
	}

	if req.Assistant != "" {
		collection.Set("", req.Assistant, true)
	}
	collection.Set(req.Message, "", false)

	prompts := collection.GetChatSummary()
	l.Logger.Info(" CustomerChat start openai  messageId:" + req.MessageID)

	messageText, err = c.ChatStream(prompts, channel)

	if err != nil {
		l.Logger.Error(" CustomerFrontendChatLogic:CustomerChat:error::" + err.Error())
		return
	}

	l.Logger.Info(" CustomerChat end openai  messageId:" + req.MessageID)

	//randle := util.GenerateSnowflakeString()
	//baseEmoji := util.GetTheEmojiString()

	//emojiPrompt := fmt.Sprintf("你需要根据得到的问题，总结出一个与之相关的情绪词。你可以从以下情绪词中选择：%s。请用一个词来回答，例如：开心", baseEmoji)
	//collectionEmoji := openai.NewUserContext(l.ctx,
	//	openai.GetUserUniqueID(req.User+randle, req.OpenKfID),
	//).WithPrompt(emojiPrompt).WithModel(l.model).WithClient(c)
	//collectionEmoji.Set(req.Message, "", false)
	//prompts1 := collectionEmoji.GetChatSummary()
	//
	//emojiText, err1 := c.Chat(prompts1)
	//if err1 != nil {
	//	l.Logger.Error("CustomerFrontendChatLogic:CustomerChat:error:" + err1.Error())
	//	return SendToWebsocket("系统错误:"+err1.Error(), req.MessageID, 0), nil
	//}
	//emoji := util.GetKeyFromEmojiMap(emojiText)
	//collectionEmoji.Clear()

	go func() {
		tokens := openai.NumTokensFromMessagesV2(prompts, messageText, postModel)
		l.chatService.SaveTimesAndToken(req.User, tokens)
	}()
	collection.Set("", messageText, true)

	insertId, _ := l.chatService.InsertReturnInsertId(&model.ChatRecord{
		User:             req.User,
		RelationId:       questionModel.Id,
		MessageId:        req.MessageID,
		OpenKfId:         req.OpenKfID,
		Content:          messageText,
		Emoji:            0,
		ChatType:         repository.ChatTypeCustomerFrontend,
		AnswerOrQuestion: repository.AnswerOrQuestionAnswer,
		MessageType:      repository.MessageTypeText,
	})
	go l.riskService.Reduce(req.User, req.OpenKfID, insertId)

	l.wechatUserRepository.InsertWechatUserIfNotExist(req.User, false)
	//}()
	// 然后把数据 发给对应的客户
	l.Logger.Info(" CustomerChat end  messageId:" + req.MessageID)

	return SendToWebsocket(messageText, req.MessageID, 0, insertId), nil

}

func (l *CustomerFrontendChatLogic) setModelName() (ls *CustomerFrontendChatLogic) {
	m := l.svcCtx.Config.WeCom.Model
	if m == "" {
		m = openai.TextModel
	}
	l.svcCtx.Config.WeCom.Model = m
	return l
}

func (l *CustomerFrontendChatLogic) setBasePrompt() (ls *CustomerFrontendChatLogic) {
	p := l.svcCtx.Config.WeCom.BasePrompt
	if p == "" {
		p = "你是 ChatGPT, 一个由 OpenAI 训练的大型语言模型, 你旨在回答并解决人们的任何问题，并且可以使用多种语言与人交流。\n"
	}
	l.basePrompt = p
	return l
}

func (l *CustomerFrontendChatLogic) setBaseHost() (ls *CustomerFrontendChatLogic) {
	if l.svcCtx.Config.OpenAi.Host == "" {
		l.svcCtx.Config.OpenAi.Host = "https://api.openai.com"
	}
	l.baseHost = l.svcCtx.Config.OpenAi.Host
	return l
}

func (l *CustomerFrontendChatLogic) setModelNameV2() (ls *CustomerFrontendChatLogic) {
	m := l.svcCtx.Config.WeCom.Model
	if m == "" {
		m = openai.TextModel
	}
	l.svcCtx.Config.WeCom.Model = m
	l.model = m
	return l
}

func (l *CustomerFrontendChatLogic) CheckClearContext(ctx context.Context, openKfId, userId, formattedTime string) (bool, error) {
	chatPo, _, err := l.chatService.GetAllRecord(0, openKfId, userId, "", formattedTime, "", "", 0, 0, 0)
	if err != nil {
		return false, err
	}
	if len(chatPo) > 0 {
		return false, nil
	}
	return true, nil
}

func (l *CustomerFrontendChatLogic) setPostModel(postModel string) (ls *CustomerFrontendChatLogic) {
	var m string
	if "" != postModel {
		m = postModel
	}
	if m == "" {
		m = openai.ChatModel
	}
	l.postModel = m
	return l
}

func SendToWebsocket(message, messageId string, emoji int64, chatRecordId int64) types.CustomerChatResponse {
	// 然后把数据 发给对应的客户
	return types.CustomerChatResponse{
		Message:      message,
		MessageID:    messageId,
		IsEnd:        true,
		Emoji:        emoji,
		ChatRecordId: strconv.FormatInt(chatRecordId, 10),
	}
}
