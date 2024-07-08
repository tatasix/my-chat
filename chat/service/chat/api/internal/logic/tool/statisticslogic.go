package tool

import (
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/service"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StatisticsLogic struct {
	logx.Logger
	ctx               context.Context
	svcCtx            *svc.ServiceContext
	StatisticsService *service.StatisticsService
}

func NewStatisticsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatisticsLogic {
	return &StatisticsLogic{
		Logger:            logx.WithContext(ctx),
		ctx:               ctx,
		svcCtx:            svcCtx,
		StatisticsService: service.NewStatisticsService(ctx, svcCtx),
	}
}

func (l *StatisticsLogic) Statistics(r *types.StatisticsRequest) (resp *types.Response, err error) {
	return assembler.Return(l.StatisticsService.Statistics(r.Now))
}
