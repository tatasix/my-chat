package service

import (
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChatRoomRecordService struct {
	logx.Logger
	ctx                      context.Context
	svcCtx                   *svc.ServiceContext
	chatRoomRecordRepository *repository.ChatRoomRecordRepository
}

func NewChatRoomRecordService(ctx context.Context, svcCtx *svc.ServiceContext) *ChatRoomRecordService {
	return &ChatRoomRecordService{
		Logger:                   logx.WithContext(ctx),
		ctx:                      ctx,
		svcCtx:                   svcCtx,
		chatRoomRecordRepository: repository.NewChatRoomRecordRepository(ctx, svcCtx),
	}
}

func (l *ChatRoomRecordService) Insert(ChatRoomRecord *model.ChatRoomRecord) (lastId int64, err error) {
	return l.chatRoomRecordRepository.Insert(ChatRoomRecord)
}

func (l *ChatRoomRecordService) GetByDialogueIdAndStatus(dialogueId string, status uint8) (ChatRoomRecord []*model.ChatRoomRecord, err error) {
	return l.chatRoomRecordRepository.GetByDialogueIdAndStatus(dialogueId, status)
}

func (l *ChatRoomRecordService) GetByDialogueId(dialogueId string) (ChatRoomRecord []*model.ChatRoomRecord, err error) {
	return l.chatRoomRecordRepository.GetByDialogueId(dialogueId)
}

func (l *ChatRoomRecordService) Idempotent(dialogueId, user string, status int8) (flag bool, err error) {
	originInfo, err := l.chatRoomRecordRepository.GetAll(dialogueId, user, "", status, 0)
	if err != nil {
		return
	}

	if originInfo != nil && len(originInfo) > 0 {
		flag = true
	}
	return
}
