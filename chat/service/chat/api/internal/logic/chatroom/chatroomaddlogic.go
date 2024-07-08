package chatroom

import (
	"chat/common/xerr"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/model"
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatRoomAddLogic struct {
	logx.Logger
	ctx             context.Context
	svcCtx          *svc.ServiceContext
	chatRoomService *service.ChatRoomService
}

func NewChatRoomAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatRoomAddLogic {
	return &ChatRoomAddLogic{
		Logger:          logx.WithContext(ctx),
		ctx:             ctx,
		svcCtx:          svcCtx,
		chatRoomService: service.NewChatRoomService(ctx, svcCtx),
	}
}

func (l *ChatRoomAddLogic) ChatRoomAdd(req *types.ChatRoomAddRequest) (resp *types.ChatRoomAddResponse, err error) {
	resp = &types.ChatRoomAddResponse{}
	enable, _, err := service.NewRiskService(l.ctx, l.svcCtx).Check(req.User, req.OpenKfID)
	if err != nil {
		return
	}
	if enable == 2 {
		resp.Code = xerr.RightsNotHaveTimesError
		return
	} else if enable == 3 {
		resp.Code = xerr.RightsNotVip
		return
	}
	chatRoomModel := &model.ChatRoom{
		User: req.User,
		KfId: req.OpenKfID,
	}
	id, err := l.chatRoomService.Insert(chatRoomModel)
	if err != nil {
		return
	}

	err = l.chatRoomService.AddChatRoomUsers(chatRoomModel.ChannelId, 1, req.User)
	resp.Id = strconv.Itoa(int(id))
	resp.User = req.User
	resp.ChannelId = strconv.Itoa(int(chatRoomModel.ChannelId))

	return
}
