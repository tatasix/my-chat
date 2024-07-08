package myerror

import "errors"

var (
	ConfigExist = errors.New(" 当前配置在系统中已存在")

	ConfigOpenAiKeyEmpty = errors.New(" 未配置key")
)
