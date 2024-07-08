package logic

import (
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLastChatRecordLogic struct {
	logx.Logger
	ctx            context.Context
	svcCtx         *svc.ServiceContext
	chatService    *service.ChatService
	chatRepository *repository.ChatRepository
}

func NewGetLastChatRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLastChatRecordLogic {
	return &GetLastChatRecordLogic{
		Logger:         logx.WithContext(ctx),
		ctx:            ctx,
		svcCtx:         svcCtx,
		chatService:    service.NewChatService(ctx, svcCtx),
		chatRepository: repository.NewChatRepository(ctx, svcCtx),
	}
}

func (l *GetLastChatRecordLogic) GetLastChatRecord(req *types.GetLastChatInfoRequest) (resp *types.GetLastChatInfoReply, err error) {
	chat, _, err := l.chatService.GetAll(req.UserNickname, req.KfName, "", "", "", req.ChatType, "created_at asc", 1, 1, 1)
	if err != nil {
		return
	}

	if chat != nil && len(chat) > 0 && chat[0].Id > 0 {
		answer, err1 := l.chatService.GetAnswerByQuestion([]int64{chat[0].Id})
		if err1 != nil {
			err = err1
			return
		}
		var response string
		if answer != nil && len(answer) > 0 && answer[0].Id > 0 {
			response = answer[0].Content
		}
		return &types.GetLastChatInfoReply{
			ResContent:  response,
			ReqContent:  chat[0].Content,
			CreatedTime: chat[0].CreatedAt.Format("2006-01-02 15:04:05"),
		}, nil
	}
	return
}
