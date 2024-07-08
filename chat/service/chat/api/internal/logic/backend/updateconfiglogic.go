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

type UpdateConfigLogic struct {
	logx.Logger
	ctx           context.Context
	svcCtx        *svc.ServiceContext
	configService *service.ConfigService
}

func NewUpdateConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateConfigLogic {
	return &UpdateConfigLogic{
		Logger:        logx.WithContext(ctx),
		ctx:           ctx,
		svcCtx:        svcCtx,
		configService: service.NewConfigService(ctx, svcCtx),
	}
}

// UpdateConfig 更新配置
func (l *UpdateConfigLogic) UpdateConfig(req *types.UpdateConfigRequest) (resp *types.Response, err error) {
	return assembler.Return(l.configService.Update(req.Id, &model.Config{
		Id:          req.Id,
		ConfigType:  req.ConfigType,
		Name:        req.Name,
		Description: req.Description,
		Value:       req.Value,
		UpdatedBy:   req.UpdateBy,
	}))
}
