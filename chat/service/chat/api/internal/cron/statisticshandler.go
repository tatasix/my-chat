package cron

import (
	"chat/service/chat/api/internal/logic/tool"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

func StatisticsHandler(ctx context.Context, svcCtx *svc.ServiceContext) {
	res, err := tool.NewStatisticsLogic(ctx, svcCtx).Statistics(&types.StatisticsRequest{})
	logx.WithContext(ctx).Infof("StatisticsHandler res:%+v err:%+v", res, err)
	return
}
