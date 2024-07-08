package myerror

import "errors"

var (
	ParamMissError = errors.New(" 参数缺失")
	RecordNotFound = errors.New(" record not found")
	//ParamMissError = errors.New(" 参数缺失")
)
