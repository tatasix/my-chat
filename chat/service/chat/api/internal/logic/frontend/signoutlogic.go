package frontend

import (
	"chat/service/chat/api/internal/service"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	*service.WechatUserService
}

func NewSignoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignoutLogic {
	return &SignoutLogic{
		Logger:            logx.WithContext(ctx),
		ctx:               ctx,
		svcCtx:            svcCtx,
		WechatUserService: service.NewWechatUserService(ctx, svcCtx),
	}
}

func (l *SignoutLogic) Signout(req *types.SignoutRequest) (resp *types.Response, err error) {
	err = l.WechatUserService.Signout(req.Authorization)
	if nil != err {
		return
	}
	return &types.Response{Message: "ok"}, nil

}
