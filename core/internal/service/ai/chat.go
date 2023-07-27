package ai

import "core/internal/service/structure"

type ChatService struct {
	strategy ChatAi
}

func NewChatService() *ChatService {
	return &ChatService{}
}

func (c *ChatService) Set(strategy ChatAi) {
	c.strategy = strategy
}

func (c *ChatService) Chat(request []structure.ChatModelMessage) (string, error) {
	return c.strategy.Chat(request)
}

func (c *ChatService) ChatStream(request []structure.ChatModelMessage, channel chan string) (string, error) {
	return c.strategy.ChatStream(request, channel)
}

func (c *ChatService) Completion(request string) (string, error) {
	return c.strategy.Completion(request)
}
func (c *ChatService) SetConfig(model string, temperature float32, maxToken int) {
	c.strategy.SetConfig(model, temperature, maxToken)
	return
}
