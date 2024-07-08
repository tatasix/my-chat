package tool

import (
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type MigrateChatLogLogic struct {
	logx.Logger
	ctx         context.Context
	svcCtx      *svc.ServiceContext
	chatService *service.ChatService
}

func NewMigrateChatLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MigrateChatLogLogic {
	return &MigrateChatLogLogic{
		Logger:      logx.WithContext(ctx),
		ctx:         ctx,
		svcCtx:      svcCtx,
		chatService: service.NewChatService(ctx, svcCtx),
	}
}

func (l *MigrateChatLogLogic) MigrateChatLog(req *types.MigrateChatLogRequest) (resp *types.Response, err error) {
	old, _, err := l.chatService.GetAllOld(uint64(req.Page), uint64(req.PageSize), "id asc")
	if err != nil {
		return
	}
	if old == nil || len(old) <= 0 {
		resp.Message = "没有需要迁移的数据"
		return
	}
	var ids []int64
	for _, vv := range old {
		ids = append(ids, vv.Id)
	}
	origin, err := l.chatService.GetByIds(ids)
	if err != nil {
		return
	}
	exist := make(map[int64]int64)
	if origin != nil && len(origin) > 0 {
		for _, vvv := range origin {
			exist[vvv.Id] = vvv.Id
		}
	}

	var data []*model.ChatRecord
	for _, v := range old {
		if _, ok := exist[v.Id]; ok {
			continue
		}
		chatType := repository.ChatTypeCustomer
		if v.AgentId != 0 {
			chatType = repository.ChatTypeApplication
		}
		data = append(data, &model.ChatRecord{
			Id:               v.Id,
			RelationId:       v.Id,
			User:             v.User,
			MessageId:        v.MessageId,
			OpenKfId:         v.OpenKfId,
			AgentId:          v.AgentId,
			Content:          v.ReqContent,
			Emoji:            0,
			ChatType:         chatType,
			AnswerOrQuestion: repository.AnswerOrQuestionQuestion,
			MessageType:      repository.MessageTypeText,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
		})
		data = append(data, &model.ChatRecord{
			Id:               v.Id + 100000,
			RelationId:       v.Id,
			User:             v.User,
			MessageId:        v.MessageId,
			OpenKfId:         v.OpenKfId,
			AgentId:          v.AgentId,
			Content:          v.ResContent,
			Emoji:            0,
			ChatType:         chatType,
			AnswerOrQuestion: repository.AnswerOrQuestionAnswer,
			MessageType:      repository.MessageTypeText,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
		})
		if len(data) > 50 {
			err = l.chatService.BatchInsert(data)
			if err != nil {
				return
			}
			data = []*model.ChatRecord{}
		}
	}
	if len(data) > 0 {
		err = l.chatService.BatchInsert(data)
		if err != nil {
			return
		}
	}
	return &types.Response{Message: "ok"}, nil

}
