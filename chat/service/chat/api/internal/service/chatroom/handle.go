package chatroom

import (
	"chat/common/chatroom"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type Handler interface {
	Handle(msg chatroom.Message, customerInfo *model.CustomerConfig, channel chan string) (res chatroom.Message, err error) // 处理方法
}

// NewChatRoomHandler Handler 工厂函数，用于创建不同类型的处理程序
func NewChatRoomHandler(ctx context.Context, svcCtx *svc.ServiceContext, customer string) Handler {
	switch customer {
	case "uqTIN13j6HKg2nYSyuTay6mHRQULNRSU":
		return &AdventureHandler{
			logger: logx.WithContext(ctx),
			ctx:    ctx,
			svcCtx: svcCtx,
		}
	default:
		return nil
	}
}
