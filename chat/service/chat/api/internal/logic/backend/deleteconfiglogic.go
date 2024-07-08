package backend

import (
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/model"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteConfigLogic struct {
	logx.Logger
	ctx           context.Context
	svcCtx        *svc.ServiceContext
	configService *service.ConfigService
}

func NewDeleteConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteConfigLogic {
	return &DeleteConfigLogic{
		Logger:        logx.WithContext(ctx),
		ctx:           ctx,
		svcCtx:        svcCtx,
		configService: service.NewConfigService(ctx, svcCtx),
	}
}

func (l *DeleteConfigLogic) DeleteConfig(req *types.IdV2Request) (resp *types.Response, err error) {

	return assembler.Return(l.configService.Update(req.Id, &model.Config{
		Id:        req.Id,
		IsDeleted: 1,
		UpdatedBy: req.UpdateBy,
	}))
}
