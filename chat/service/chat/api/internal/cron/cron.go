package cron

import (
	"chat/service/chat/api/internal/svc"
	"context"
	"github.com/robfig/cron/v3"
)

func RegisterCrontab(serverCtx *svc.ServiceContext) {
	c := cron.New()
	ctx := context.Background()
	//统计日活月活 0 1 0 * * *
	//_, _ = c.AddFunc("3 0 * * *", func(ctx context.Context, svcCtx *svc.ServiceContext) func() {
	//	return func() {
	//		StatisticsHandler(ctx, serverCtx)
	//	}
	//}(ctx, serverCtx))

	_, _ = c.AddFunc("1 0 * * *", func(ctx context.Context, svcCtx *svc.ServiceContext) func() {
		return func() {
			RightsExpireHandler(ctx, serverCtx)
		}
	}(ctx, serverCtx))

	//苹果支付自动订阅续费/取消订阅/过期
	var spec string
	if serverCtx.Config.ApplePay.IsProd {
		spec = "0 4 * * *" //每天凌晨4点
	} else {
		spec = "*/3 * * * *" //每 3 分钟
	}
	_, _ = c.AddFunc(spec, func(ctx context.Context, svcCtx *svc.ServiceContext) func() {
		return func() {
			AppleSubscriptionHandler(ctx, serverCtx)
		}
	}(ctx, serverCtx))

	c.Start()

}
