package frontend

import (
	"context"

	"chat/service/chat/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChatIndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatIndexLogic {
	return &ChatIndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatIndexLogic) ChatIndex() error {
	// todo: add your logic here and delete this line

	return nil
}
