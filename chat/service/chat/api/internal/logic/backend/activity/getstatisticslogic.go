package activity

import (
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/service"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStatisticsLogic struct {
	logx.Logger
	ctx               context.Context
	svcCtx            *svc.ServiceContext
	StatisticsService *service.StatisticsService
}

func NewGetStatisticsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStatisticsLogic {
	return &GetStatisticsLogic{
		Logger:            logx.WithContext(ctx),
		ctx:               ctx,
		svcCtx:            svcCtx,
		StatisticsService: service.NewStatisticsService(ctx, svcCtx),
	}
}

func (l *GetStatisticsLogic) GetStatistics(req *types.GetStatisticsRequest) (resp *types.GetStatisticsResponse, err error) {
	resp = &types.GetStatisticsResponse{
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	pos, count, err := l.StatisticsService.GetAll(req.StartCreatedAt, req.EndCreatedAt, uint64(req.Page), uint64(req.PageSize))
	if err != nil {
		return
	}
	resp.List = assembler.POTODTOGetStatistics(pos)
	resp.Total = count

	return
}
