package chatroom

import (
	"chat/common/util"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatRoomDeleteLogic struct {
	logx.Logger
	ctx             context.Context
	svcCtx          *svc.ServiceContext
	chatRoomService *service.ChatRoomService
}

func NewChatRoomDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatRoomDeleteLogic {
	return &ChatRoomDeleteLogic{
		Logger:          logx.WithContext(ctx),
		ctx:             ctx,
		svcCtx:          svcCtx,
		chatRoomService: service.NewChatRoomService(ctx, svcCtx),
	}
}

func (l *ChatRoomDeleteLogic) ChatRoomDelete(req *types.ChatRoomDeleteRequest) (resp *types.Response, err error) {
	err = l.chatRoomService.Delete(util.ConvertToInt64(req.ChannelId))
	return
}
