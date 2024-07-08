package logic

import (
	"chat/service/chat/api/internal/service"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLikeListLogic struct {
	logx.Logger
	ctx             context.Context
	svcCtx          *svc.ServiceContext
	userLikeService *service.UserLikeService
}

func NewGetUserLikeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLikeListLogic {
	return &GetUserLikeListLogic{
		Logger:          logx.WithContext(ctx),
		ctx:             ctx,
		svcCtx:          svcCtx,
		userLikeService: service.NewUserLikeService(ctx, svcCtx),
	}
}

func (l *GetUserLikeListLogic) GetUserLikeList(req *types.GetUserLikeListRequest) (resp *types.GetUserLikeListPageResult, err error) {
	return l.userLikeService.GetUserLikeList(req.ChatRecordUser, req.StartCreatedAt, req.EndCreatedAt, "", req.Page, req.PageSize)
}
