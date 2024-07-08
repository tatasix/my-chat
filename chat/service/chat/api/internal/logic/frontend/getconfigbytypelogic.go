package frontend

import (
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigByTypeLogic struct {
	logx.Logger
	ctx           context.Context
	svcCtx        *svc.ServiceContext
	configService *service.ConfigService
}

func NewGetConfigByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigByTypeLogic {
	return &GetConfigByTypeLogic{
		Logger:        logx.WithContext(ctx),
		ctx:           ctx,
		svcCtx:        svcCtx,
		configService: service.NewConfigService(ctx, svcCtx),
	}
}

func (l *GetConfigByTypeLogic) GetConfigByType(req *types.GetConfigRequest) (resp *types.GetConfigResponse, err error) {
	info, err := l.configService.GetAll(req.ConfigType)
	if err != nil {
		return
	}
	resp = assembler.POTODTOGetConfig(info)

	return
}
