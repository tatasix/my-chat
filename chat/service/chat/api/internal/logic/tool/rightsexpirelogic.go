package tool

import (
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/service/rights"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type RightsExpireLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	service *rights.RightsService
}

func NewRightsExpireLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RightsExpireLogic {
	return &RightsExpireLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		service: rights.NewRightsService(ctx, svcCtx),
	}
}

func (l *RightsExpireLogic) RightsExpire() (resp *types.Response, err error) {
	return assembler.Return(l.service.ExpireRights())
}
