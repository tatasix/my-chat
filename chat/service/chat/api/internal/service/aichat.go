package service

import (
	"chat/common/openai"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type AiChatService struct {
	logx.Logger
	ctx           context.Context
	svcCtx        *svc.ServiceContext
	model         string
	baseHost      string
	configService *ConfigService
	chatService   *ChatService
}

func NewAiChatService(ctx context.Context, svcCtx *svc.ServiceContext) *AiChatService {
	return &AiChatService{
		Logger:        logx.WithContext(ctx),
		ctx:           ctx,
		svcCtx:        svcCtx,
		configService: NewConfigService(ctx, svcCtx),
		chatService:   NewChatService(ctx, svcCtx),
	}
}

func (l *AiChatService) Chat(customerConfig *model.CustomerConfig, summary []openai.ChatModelMessage, user string) (res string, err error) {
	l.setBaseHost()

	openAiKey, err := l.configService.GetConfigKey()
	if err != nil {
		return
	}

	// openai client
	c := openai.NewChatClient(l.ctx, openAiKey).
		WithModel(customerConfig.PostModel).
		WithBaseHost(l.baseHost).
		WithOrigin(l.svcCtx.Config.OpenAi.Origin).
		WithEngine(l.svcCtx.Config.OpenAi.Engine).
		WithPostModel(customerConfig.PostModel)
	if l.svcCtx.Config.Proxy.Enable {
		c = c.WithHttpProxy(l.svcCtx.Config.Proxy.Http).WithSocks5Proxy(l.svcCtx.Config.Proxy.Socket5)
	}

	res, tokens, err := c.Chat(summary)
	go l.chatService.SaveTimesAndToken(user, int64(tokens))
	//NewRiskService(l.ctx, l.svcCtx).Reduce(user, req.OpenKfID)

	return res, err
}

func (l *AiChatService) setBaseHost() (ls *AiChatService) {
	if l.svcCtx.Config.OpenAi.Host == "" {
		l.svcCtx.Config.OpenAi.Host = "https://api.openai.com"
	}
	l.baseHost = l.svcCtx.Config.OpenAi.Host
	return l
}

func (l *AiChatService) setModelName(baseModel string) (ls *AiChatService) {
	m := l.svcCtx.Config.WeCom.Model
	if "" != baseModel {
		m = baseModel
	}
	if m == "" {
		m = openai.TextModel
	}
	l.model = m
	return l
}

// setPrompt
func (l *AiChatService) setPrompt(basePrompt string) []openai.ChatModelMessage {
	var summary []openai.ChatModelMessage
	summary = append(summary, openai.ChatModelMessage{
		Role:    "system",
		Content: "",
	})
	summary = append(summary, openai.ChatModelMessage{
		Role:    "system",
		Content: "",
	})
	return summary
}

func (l *AiChatService) CustomerChatStream(customerConfig *model.CustomerConfig, summary []openai.ChatModelMessage, channel chan string, user string) (messageText string, err error) {

	l.setBaseHost()

	openAiKey, err := l.configService.GetConfigKey()
	if err != nil {
		return
	}
	// openai client
	c := openai.NewChatClient(l.ctx, openAiKey).
		WithModel(customerConfig.PostModel).
		WithBaseHost(l.baseHost).
		WithOrigin(l.svcCtx.Config.OpenAi.Origin).
		WithEngine(l.svcCtx.Config.OpenAi.Engine).WithPostModel(customerConfig.PostModel)
	if l.svcCtx.Config.Proxy.Enable {
		c = c.WithHttpProxy(l.svcCtx.Config.Proxy.Http).WithSocks5Proxy(l.svcCtx.Config.Proxy.Socket5)
	}
	messageText, err = c.ChatStream(summary, channel)

	go func() {
		tokens := openai.NumTokensFromMessagesV2(summary, messageText, customerConfig.PostModel)
		l.chatService.SaveTimesAndToken(user, tokens)
	}()
	return
}

func (l *AiChatService) SetSystemAndUserPrompt(prompt string, user ...string) (summary []openai.ChatModelMessage) {

	summary = append(summary, openai.ChatModelMessage{
		Role:    "system",
		Content: prompt,
	})
	for _, v := range user {
		summary = append(summary, openai.ChatModelMessage{
			Role:    "user",
			Content: v,
		})
	}
	return summary
}

func (l *AiChatService) SetSystemPrompt(summary []openai.ChatModelMessage, prompt string) []openai.ChatModelMessage {

	return append(summary, openai.ChatModelMessage{
		Role:    "system",
		Content: prompt,
	})
}

func (l *AiChatService) SetUserPrompt(summary []openai.ChatModelMessage, user ...string) []openai.ChatModelMessage {
	for _, v := range user {
		summary = append(summary, openai.ChatModelMessage{
			Role:    "user",
			Content: v,
		})
	}
	return summary
}

func (l *AiChatService) SetAssistantPrompt(summary []openai.ChatModelMessage, assistant ...string) []openai.ChatModelMessage {
	for _, v := range assistant {
		summary = append(summary, openai.ChatModelMessage{
			Role:    "assistant",
			Content: v,
		})
	}
	return summary
}

func (l *AiChatService) UpdateSystemPrompt(summary []openai.ChatModelMessage, prompt string) []openai.ChatModelMessage {

	for i, v := range summary {
		if v.Role == "system" {
			summary[i].Content = prompt
		}
	}
	return summary
}
