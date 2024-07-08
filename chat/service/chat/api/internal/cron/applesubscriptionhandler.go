package cron

import (
	"chat/service/chat/api/internal/service/pay"
	"chat/service/chat/api/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

func AppleSubscriptionHandler(ctx context.Context, svcCtx *svc.ServiceContext) {

	applePayService := pay.NewApplePayService(ctx, svcCtx)

	err := applePayService.CheckApplePaySubscription()
	if err != nil {
		logx.WithContext(ctx).Infof("AppleSubscriptionHandler res:%+v err:%+v", err)
	}
	return
}
