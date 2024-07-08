package chatroom

import (
	"chat/common/util"
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatRoomByIdLogic struct {
	logx.Logger
	ctx             context.Context
	svcCtx          *svc.ServiceContext
	chatRoomService *service.ChatRoomService
}

func NewGetChatRoomByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatRoomByIdLogic {
	return &GetChatRoomByIdLogic{
		Logger:          logx.WithContext(ctx),
		ctx:             ctx,
		svcCtx:          svcCtx,
		chatRoomService: service.NewChatRoomService(ctx, svcCtx),
	}
}

func (l *GetChatRoomByIdLogic) GetChatRoomById(req *types.GetChatRoomByIdRequest) (resp *types.ChatRoomAddResponse, err error) {

	info, err := l.chatRoomService.GetById(util.ConvertToInt64(req.Id))
	if nil != err {
		return
	}
	resp = assembler.POTODTOGetChatRoomById(info)
	return
}
