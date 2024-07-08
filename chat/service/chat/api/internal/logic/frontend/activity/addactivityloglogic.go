package activity

import (
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/model"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddActivityLogLogic struct {
	logx.Logger
	ctx                context.Context
	svcCtx             *svc.ServiceContext
	activityLogService *service.ActivityLogService
}

func NewAddActivityLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddActivityLogLogic {
	return &AddActivityLogLogic{
		Logger:             logx.WithContext(ctx),
		ctx:                ctx,
		svcCtx:             svcCtx,
		activityLogService: service.NewActivityLogService(ctx, svcCtx),
	}
}

func (l *AddActivityLogLogic) AddActivityLog(req *types.SaveActivityLogRequest) (resp *types.Response, err error) {
	return assembler.Return(l.activityLogService.Insert(&model.ActivityLog{
		User:      req.User,
		Type:      req.Type,
		BeginTime: req.BeginTime,
		EndTime:   req.EndTime,
		Duration:  req.Duration,
	}))
}
