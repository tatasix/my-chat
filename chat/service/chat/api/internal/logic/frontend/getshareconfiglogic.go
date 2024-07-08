package frontend

import (
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/service/wechat"
	"context"
	"net/url"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShareConfigLogic struct {
	logx.Logger
	ctx           context.Context
	svcCtx        *svc.ServiceContext
	wechatService *wechat.OfficialAccount
}

func NewGetShareConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShareConfigLogic {
	return &GetShareConfigLogic{
		Logger:        logx.WithContext(ctx),
		ctx:           ctx,
		svcCtx:        svcCtx,
		wechatService: wechat.NewOfficialAccount(ctx, svcCtx),
	}
}

func (l *GetShareConfigLogic) GetShareConfig(req *types.GetShareConfigRequest) (resp *types.GetShareConfigResponse, err error) {
	resp = &types.GetShareConfigResponse{}
	uri := req.Url
	if req.Encode {
		uri = url.QueryEscape(uri)
	}
	appId, nonceStr, signature, timestamp, err := l.wechatService.GetConfig(uri)
	if err != nil {
		return
	}
	resp = assembler.POTODTOGetShareConfig(appId, nonceStr, signature, timestamp)
	return
}
