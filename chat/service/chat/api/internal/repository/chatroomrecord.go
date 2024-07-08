package repository

import (
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChatRoomRecordRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatRoomRecordRepository(ctx context.Context, svcCtx *svc.ServiceContext) *ChatRoomRecordRepository {
	return &ChatRoomRecordRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatRoomRecordRepository) GetByDialogueIdAndStatus(dialogueId string, status uint8) (ChatRoomRecordPo []*model.ChatRoomRecord, err error) {
	return l.svcCtx.ChatRoomRecordModel.FindAll(l.ctx,
		l.svcCtx.ChatRoomRecordModel.RowBuilder().Where(squirrel.Eq{"dialogue_id": dialogueId}).Where(squirrel.Eq{"status": status}),
	)
}

func (l *ChatRoomRecordRepository) GetByDialogueId(dialogueId string) (ChatRoomRecordPo []*model.ChatRoomRecord, err error) {
	return l.svcCtx.ChatRoomRecordModel.FindAll(l.ctx,
		l.svcCtx.ChatRoomRecordModel.RowBuilder().Where(squirrel.Eq{"dialogue_id": dialogueId}),
	)
}

func (l *ChatRoomRecordRepository) GetById(id int64) (ChatRoomRecord *model.ChatRoomRecord, err error) {
	return l.svcCtx.ChatRoomRecordModel.FindOne(context.Background(), id)
}

func (l *ChatRoomRecordRepository) Insert(ChatRoomRecord *model.ChatRoomRecord) (lastId int64, err error) {
	l.svcCtx.ChatRoomRecordModel.BuildFiled(nil, ChatRoomRecord)
	res, err := l.svcCtx.ChatRoomRecordModel.Insert(l.ctx, ChatRoomRecord)
	if err != nil {
		return
	}
	return res.LastInsertId()
}

func (l *ChatRoomRecordRepository) GetAll(dialogueId, user, kfId string, status int8, channelId int64) (ChatRoomRecordPo []*model.ChatRoomRecord, err error) {

	rowBuilder := l.svcCtx.ChatRoomRecordModel.RowBuilder()
	if user != "" {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"send_user_id": user})
	}

	if kfId != "" {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"kf_id": kfId})
	}

	if dialogueId != "" {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"dialogue_id": dialogueId})
	}

	if status != 0 {
		if status == -1 {
			rowBuilder = rowBuilder.Where(squirrel.Eq{"status": 0})
		} else {
			rowBuilder = rowBuilder.Where(squirrel.Eq{"status": status})
		}
	}

	if channelId != 0 {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"channel_id": channelId})
	}

	return l.svcCtx.ChatRoomRecordModel.FindAll(l.ctx, rowBuilder)
}
