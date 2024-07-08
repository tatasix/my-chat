package frontend

import (
	"chat/service/chat/api/internal/service/wechat"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRedirectUrlLogic struct {
	logx.Logger
	ctx           context.Context
	svcCtx        *svc.ServiceContext
	wechatService *wechat.OfficialAccount
}

func NewGetRedirectUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRedirectUrlLogic {
	return &GetRedirectUrlLogic{
		Logger:        logx.WithContext(ctx),
		ctx:           ctx,
		svcCtx:        svcCtx,
		wechatService: wechat.NewOfficialAccount(ctx, svcCtx),
	}
}

func (l *GetRedirectUrlLogic) GetRedirectUrl(req *types.GetRedirectUrlRequest) (resp *types.GetRedirectUrlResponse, err error) {
	resp = &types.GetRedirectUrlResponse{}
	resp.Url, err = l.wechatService.GetRedirectURL(req.User, req.ReturnUrl)
	return
}
