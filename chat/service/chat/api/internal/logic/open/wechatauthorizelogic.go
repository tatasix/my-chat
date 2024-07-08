package open

import (
	"chat/service/chat/api/internal/service"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WechatAuthorizeLogic struct {
	logx.Logger
	ctx               context.Context
	svcCtx            *svc.ServiceContext
	wechatUserService *service.WechatUserService
}

func NewWechatAuthorizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WechatAuthorizeLogic {
	return &WechatAuthorizeLogic{
		Logger:            logx.WithContext(ctx),
		ctx:               ctx,
		svcCtx:            svcCtx,
		wechatUserService: service.NewWechatUserService(ctx, svcCtx),
	}
}

func (l *WechatAuthorizeLogic) WechatAuthorize(req *types.WechatRedirectRequest) (resp *types.Response, err error) {

	err = l.wechatUserService.WechatAuthorizeCallBack(req.State, req.Code)
	if nil != err {
		return
	}
	return &types.Response{Message: "ok"}, nil

}
