package rights

import (
	"chat/service/chat/api/internal/service"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckRightsTempLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	service *service.RiskService
}

func NewCheckRightsTempLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckRightsTempLogic {
	return &CheckRightsTempLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		service: service.NewRiskService(ctx, svcCtx),
	}
}

func (l *CheckRightsTempLogic) CheckRightsTemp(req *types.CheckRightsTempRequest) (resp *types.CheckRightsResponse, err error) {
	enabled, remainTimes, err := l.service.Check(req.User, req.OpenKfID)
	if err != nil {
		return
	}
	resp = &types.CheckRightsResponse{
		RemainTimes: remainTimes,
		Enabled:     enabled,
	}
	return
}
