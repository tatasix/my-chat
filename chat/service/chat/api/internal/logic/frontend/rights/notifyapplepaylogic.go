package rights

import (
	"chat/service/chat/api/internal/service/pay"
	"context"
	"net/http"

	"chat/service/chat/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyApplePayLogic struct {
	logx.Logger
	ctx             context.Context
	svcCtx          *svc.ServiceContext
	applePayService *pay.ApplePayService
}

func NewNotifyApplePayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyApplePayLogic {
	return &NotifyApplePayLogic{
		Logger:          logx.WithContext(ctx),
		ctx:             ctx,
		svcCtx:          svcCtx,
		applePayService: pay.NewApplePayService(ctx, svcCtx),
	}
}

func (l *NotifyApplePayLogic) NotifyApplePay(r *http.Request) error {
	return l.applePayService.Notify(r)
}
