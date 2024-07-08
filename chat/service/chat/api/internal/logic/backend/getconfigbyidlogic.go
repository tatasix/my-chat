package backend

import (
	"chat/service/chat/api/internal/logic/assembler"

	"chat/service/chat/api/internal/service"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigByIdLogic struct {
	logx.Logger
	ctx           context.Context
	svcCtx        *svc.ServiceContext
	configService *service.ConfigService
}

func NewGetConfigByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigByIdLogic {
	return &GetConfigByIdLogic{
		Logger:        logx.WithContext(ctx),
		ctx:           ctx,
		svcCtx:        svcCtx,
		configService: service.NewConfigService(ctx, svcCtx),
	}
}

func (l *GetConfigByIdLogic) GetConfigById(req *types.IdV3Request) (resp *types.Config, err error) {
	info, err := l.configService.GetById(req.Id)
	if err != nil {
		return
	}
	result := assembler.POTODTOGetOne(info)
	resp = &result

	return
}
