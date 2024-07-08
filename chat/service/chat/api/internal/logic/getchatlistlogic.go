package logic

import (
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatListLogic struct {
	logx.Logger
	ctx         context.Context
	svcCtx      *svc.ServiceContext
	chatService *service.ChatService
}

func NewGetChatListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatListLogic {
	return &GetChatListLogic{
		Logger:      logx.WithContext(ctx),
		ctx:         ctx,
		svcCtx:      svcCtx,
		chatService: service.NewChatService(ctx, svcCtx),
	}
}

func (l *GetChatListLogic) GetChatList(req *types.GetChatListRequest) (resp *types.GetChatListPageResult, err error) {
	return l.chatService.GetChatList(req.User, req.ChatRecordUser, req.Customer, req.Agent, req.StartCreatedAt, req.EndCreatedAt, req.ChatType, "created_at desc", req.Page, req.PageSize)
}
