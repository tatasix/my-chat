package frontend

import (
	"context"

	"core/internal/svc"
	"core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WechatChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWechatChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WechatChatLogic {
	return &WechatChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WechatChatLogic) WechatChat(req *types.WechatChatRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
