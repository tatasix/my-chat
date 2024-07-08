package frontend

import (
	"chat/common/util"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateUniqueIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateUniqueIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateUniqueIdLogic {
	return &GenerateUniqueIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateUniqueIdLogic) GenerateUniqueId() (resp *types.GenerateUniqueIdResponse, err error) {
	resp = &types.GenerateUniqueIdResponse{Id: util.GenerateSnowflakeString()}
	return
}
