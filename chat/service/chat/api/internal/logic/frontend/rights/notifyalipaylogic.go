package rights

import (
	"chat/service/chat/api/internal/service/rights"
	"context"
	"net/http"

	"chat/service/chat/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyAlipayLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	service *rights.RightsService
}

func NewNotifyAlipayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyAlipayLogic {
	return &NotifyAlipayLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		service: rights.NewRightsService(ctx, svcCtx),
	}
}

func (l *NotifyAlipayLogic) NotifyAlipay(r *http.Request) (res string, err error) {
	return l.service.Notify(1, r)
}
