package client

import (
	"cron/cron/internal/config"
	"cron/cron/util"
)

type HttpService struct {
}

func NewHttpService() *HttpService {
	return &HttpService{}
}

func (h *HttpService) RunScript() (string, error) {
	url := config.ScriptService + config.RunScriptUri
	return util.Get(url)
}

func (h *HttpService) RunStatistics() (string, error) {
	url := config.ChatService + config.RunStatistics
	return util.Get(url)
}
