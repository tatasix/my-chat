package activity

import (
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/service"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeUserIdLogic struct {
	logx.Logger
	ctx                context.Context
	svcCtx             *svc.ServiceContext
	activityLogService *service.ActivityLogService
}

func NewChangeUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeUserIdLogic {
	return &ChangeUserIdLogic{
		Logger:             logx.WithContext(ctx),
		ctx:                ctx,
		svcCtx:             svcCtx,
		activityLogService: service.NewActivityLogService(ctx, svcCtx),
	}
}

func (l *ChangeUserIdLogic) ChangeUserId(req *types.ChangeUserRequest) (resp *types.Response, err error) {
	return assembler.Return(l.activityLogService.UpdateUser(req.User, req.OldUser))
}
