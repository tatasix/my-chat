package cron

import (
	"chat/service/chat/api/internal/logic/tool"
	"chat/service/chat/api/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

func RightsExpireHandler(ctx context.Context, svcCtx *svc.ServiceContext) {
	res, err := tool.NewRightsExpireLogic(ctx, svcCtx).RightsExpire()
	logx.WithContext(ctx).Infof("RightsExpireHandler res:%+v err:%+v", res, err)
	return
}
