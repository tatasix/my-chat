package backend

import (
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddConfigLogic struct {
	logx.Logger
	ctx           context.Context
	svcCtx        *svc.ServiceContext
	configService *service.ConfigService
}

func NewAddConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddConfigLogic {
	return &AddConfigLogic{
		Logger:        logx.WithContext(ctx),
		ctx:           ctx,
		svcCtx:        svcCtx,
		configService: service.NewConfigService(ctx, svcCtx),
	}
}

func (l *AddConfigLogic) AddConfig(in *types.AddConfigRequest) (resp *types.Response, err error) {

	return assembler.Return(l.configService.Insert(&model.Config{
		ConfigType:  in.ConfigType,
		Name:        in.Name,
		Description: in.Description,
		Value:       in.Value,
		CreatedBy:   in.CreatedBy,
		UpdatedBy:   in.CreatedBy,
	}))
}
