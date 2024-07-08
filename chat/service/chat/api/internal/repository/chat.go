package repository

import (
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChatRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatRepository(ctx context.Context, svcCtx *svc.ServiceContext) *ChatRepository {
	return &ChatRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetAll 仅仅用来迁移旧数据
func (l *ChatRepository) GetAll(agentId int64, openKfId, user, startTime, endTime, order string, page, limit uint64, chatType int32) (chatPos []*model.Chat, count int64, err error) {

	countBuilder := l.svcCtx.ChatModel.CountBuilder("id")
	rowBuilder := l.svcCtx.ChatModel.RowBuilder()
	if user != "" {
		countBuilder = countBuilder.Where(squirrel.Eq{"user": user})
		rowBuilder = rowBuilder.Where(squirrel.Eq{"user": user})
	}
	if agentId != 0 {
		countBuilder = countBuilder.Where(squirrel.Eq{"agent_id": agentId})
		rowBuilder = rowBuilder.Where(squirrel.Eq{"agent_id": agentId})
	}

	if openKfId != "" {
		countBuilder = countBuilder.Where(squirrel.Eq{"open_kf_id": openKfId})
		rowBuilder = rowBuilder.Where(squirrel.Eq{"open_kf_id": openKfId})
	}

	if agentId == 0 && openKfId == "" {
		if chatType == 1 {
			countBuilder = countBuilder.Where("agent_id <> 0")
			rowBuilder = rowBuilder.Where("agent_id <> 0")
		} else if chatType == 2 {
			countBuilder = countBuilder.Where("open_kf_id <> ''")
			rowBuilder = rowBuilder.Where("open_kf_id <> ''")
		}
	}

	if startTime != "" {
		countBuilder = countBuilder.Where("created_at >= ?", startTime)
		rowBuilder = rowBuilder.Where("created_at >= ?", startTime)
	}

	if endTime != "" {
		countBuilder = countBuilder.Where("created_at < ?", endTime)
		rowBuilder = rowBuilder.Where("created_at < ?", endTime)
	}

	count, err = l.svcCtx.ChatModel.FindCount(context.Background(), countBuilder)
	if err != nil {
		return
	}
	if count <= 0 {
		return nil, 0, nil
	}

	rowBuilder = rowBuilder.OrderBy(order)
	if limit != 0 {
		offset := (page - 1) * limit
		rowBuilder = rowBuilder.Limit(limit).Offset(offset)
	}
	chatPos, err = l.svcCtx.ChatModel.FindAll(context.Background(), rowBuilder)
	if err != nil {
		return
	}
	return
}

func (l *ChatRepository) GetById(id int64) (CustomerConfigPo *model.Chat, err error) {
	return l.svcCtx.ChatModel.FindOne(l.ctx, id)
}
