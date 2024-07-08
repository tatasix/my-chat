package tool

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/service/rights"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserVipLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	service *rights.RightsService
}

func NewDeleteUserVipLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserVipLogic {
	return &DeleteUserVipLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		service: rights.NewRightsService(ctx, svcCtx),
	}
}

func (l *DeleteUserVipLogic) DeleteUserVip(req *types.DeleteUserVipRequest) (resp *types.Response, err error) {
	if l.svcCtx.Config.Env == "prod" {
		return nil, util.ReturnError(xerr.SystemError)
	}
	return assembler.Return(l.service.DeleteUserRights(req.User))
}
