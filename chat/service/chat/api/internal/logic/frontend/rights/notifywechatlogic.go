package rights

import (
	"chat/service/chat/api/internal/service/rights"
	"context"
	"net/http"

	"chat/service/chat/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyWechatLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	service *rights.RightsService
}

func NewNotifyWechatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyWechatLogic {
	return &NotifyWechatLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		service: rights.NewRightsService(ctx, svcCtx),
	}
}

func (l *NotifyWechatLogic) NotifyWechat(r *http.Request) (string, error) {
	return l.service.Notify(2, r)
}
