package chatroom

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatRoomJoinLogic struct {
	logx.Logger
	ctx             context.Context
	svcCtx          *svc.ServiceContext
	chatRoomService *service.ChatRoomService
}

func NewChatRoomJoinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatRoomJoinLogic {
	return &ChatRoomJoinLogic{
		Logger:          logx.WithContext(ctx),
		ctx:             ctx,
		svcCtx:          svcCtx,
		chatRoomService: service.NewChatRoomService(ctx, svcCtx),
	}
}

func (l *ChatRoomJoinLogic) ChatRoomJoin(req *types.ChatRoomJoinRequest) (resp *types.Response, err error) {
	channelId := util.ConvertToInt64(req.ChannelId)
	info, err := l.chatRoomService.GetChatRoomByChannelId(channelId)
	if err != nil {
		return
	}
	if info == nil || info.Id <= 0 {
		err = util.ReturnError(xerr.ChatRoomChannelNotFound)
		return
	}
	chatRoomUsers, err := l.chatRoomService.GetChatRoomUsers(channelId)
	if err != nil {
		return nil, err
	}
	if len(chatRoomUsers) >= 2 {
		err = util.ReturnError(xerr.ChatRoomFull)
		return
	}

	for _, v := range chatRoomUsers {
		if req.User == v.User {
			err = util.ReturnError(xerr.ChatRoomAlreadyInRoom)
			return
		}
	}

	return assembler.Return(l.chatRoomService.AddChatRoomUsers(channelId, 0, req.User))
}
