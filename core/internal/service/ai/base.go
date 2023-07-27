package ai

import "core/internal/service/structure"

type ChatAi interface {
	Completion(request string) (string, error)
	SpeakToTxt(voiceUrl string) (string, error)
	CreateOpenAIEmbeddings(input string) (structure.EmbeddingResponse, error)
	Chat(request []structure.ChatModelMessage) (string, error)
	ChatStream(request []structure.ChatModelMessage, channel chan string) (string, error)
	SetConfig(model string, temperature float32, maxToken int)
}
