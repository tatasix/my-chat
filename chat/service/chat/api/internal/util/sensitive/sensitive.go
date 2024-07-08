package sensitive

import (
	"chat/service/chat/api/internal/svc"
	"context"
)

const ErrorMessage = "您的消息中含有敏感词信息，请重新输入"

// Sensitive 对象存储接口
type Sensitive interface {
	Check(input string) bool
}

func NewSensitive(ctx context.Context, svcCtx *svc.ServiceContext) Sensitive {
	switch svcCtx.Config.Sensitive.Type {
	case "local":
		return &Tencent{ctx, svcCtx}
	default:
		return &Tencent{ctx, svcCtx}
	}
}
