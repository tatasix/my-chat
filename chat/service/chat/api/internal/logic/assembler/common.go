package assembler

import (
	"chat/service/chat/api/internal/types"
)

func Return(err error) (*types.Response, error) {
	if nil != err {
		return nil, err
	}
	return &types.Response{Message: "ok"}, nil
}
