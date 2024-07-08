package chatroom

import (
	"chat/service/chat/api/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChatRoomSocketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatRoomSocketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatRoomSocketLogic {
	return &ChatRoomSocketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//
//func (l *ChatRoomSocketLogic) ChatRoomSocket(req *types.ChatRoomSocketRequest, w http.ResponseWriter, r *http.Request) {
//	im.SendMessageToUid(req.User, req.User, req.Content, req.Type)
//	return
//}
