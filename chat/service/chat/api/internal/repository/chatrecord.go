package repository

import (
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	ChatTypeApplication      int64 = 1
	ChatTypeCustomer         int64 = 2
	ChatTypeCustomerFrontend int64 = 3

	MessageTypeText int64 = 1 //文本消息

	AnswerOrQuestionQuestion  int64 = 1 //问题
	AnswerOrQuestionAnswer    int64 = 2 //答案
	AnswerOrQuestionSummarize int64 = 3 //总结

	MessageTypeUser      int64 = 1
	MessageTypeCustomer  int64 = 2
	MessageTypeSummarize int64 = 3
)

type ChatRecordRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatRecordRepository(ctx context.Context, svcCtx *svc.ServiceContext) *ChatRecordRepository {
	return &ChatRecordRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatRecordRepository) GetAll(agentId int64, openKfId, user, chatRecordUser, startTime, endTime, order string, page, limit uint64, chatType int32, aq int64) (chatRecordPos []*model.ChatRecord, count int64, err error) {
	l.Logger.Info(" GetCustomerChatRecord ChatRecordRepository start ")

	countBuilder := l.svcCtx.ChatRecordModel.CountBuilder("id")
	rowBuilder := l.svcCtx.ChatRecordModel.RowBuilder()
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

	if aq != 0 {
		countBuilder = countBuilder.Where(squirrel.Eq{"answer_or_question": aq})
		rowBuilder = rowBuilder.Where(squirrel.Eq{"answer_or_question": aq})
	}

	if chatRecordUser != "" {
		countBuilder = countBuilder.Where("user = ?", chatRecordUser)
		rowBuilder = rowBuilder.Where("user = ?", chatRecordUser)
	}
	count, err = l.svcCtx.ChatRecordModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return
	}
	if count <= 0 {
		return nil, 0, nil
	}
	if order != "" {
		rowBuilder = rowBuilder.OrderBy(order)
	}
	if limit != 0 {
		offset := (page - 1) * limit
		rowBuilder = rowBuilder.Limit(limit).Offset(offset)
	}
	chatRecordPos, err = l.svcCtx.ChatRecordModel.FindAll(l.ctx, rowBuilder)
	if err != nil {
		return
	}
	l.Logger.Info(" GetCustomerChatRecord ChatRecordRepository end ")

	return
}

func (l *ChatRecordRepository) GetById(id int64) (*model.ChatRecord, error) {
	return l.svcCtx.ChatRecordModel.FindOne(l.ctx, id)
}

func (l *ChatRecordRepository) GetByMessageAndCustomer(messageId, customerId string) ([]*model.ChatRecord, error) {
	return l.svcCtx.ChatRecordModel.FindAll(l.ctx,
		l.svcCtx.ChatRecordModel.RowBuilder().Where(squirrel.Eq{"user": customerId}).Where(squirrel.Eq{"message_id": messageId}),
	)
}

func (l *ChatRecordRepository) GetByMessageAndState(relationId int64, state int) (*model.ChatRecord, error) {
	return l.svcCtx.ChatRecordModel.FindOneByQuery(l.ctx,
		l.svcCtx.ChatRecordModel.RowBuilder().Where(squirrel.Eq{"state_id": relationId}, squirrel.Eq{"state": state}).Where(squirrel.Eq{"state": state}).Where(squirrel.Eq{"chat_type": 1}),
	)
}

func (l *ChatRecordRepository) Insert(chatRecord *model.ChatRecord) (sql.Result, error) {
	l.svcCtx.ChatRecordModel.BuildFiled(nil, chatRecord)
	return l.svcCtx.ChatRecordModel.Insert(l.ctx, chatRecord)
}

func (l *ChatRecordRepository) BatchInsert(chatRecord []*model.ChatRecord) (sql.Result, error) {
	return l.svcCtx.ChatRecordModel.BatchInsert(l.ctx, chatRecord)
}

func (l *ChatRecordRepository) GetByIds(ids []int64) ([]*model.ChatRecord, error) {
	return l.svcCtx.ChatRecordModel.FindAll(l.ctx,
		l.svcCtx.ChatRecordModel.RowBuilder().Where(squirrel.Eq{"id": ids}),
	)
}

func (l *ChatRecordRepository) GetAnswerByQuestion(relationIds []int64) (chatRecordPos []*model.ChatRecord, err error) {

	return l.svcCtx.ChatRecordModel.FindAll(l.ctx,
		l.svcCtx.ChatRecordModel.RowBuilder().Where(squirrel.Eq{"relation_id": relationIds}).Where(squirrel.Eq{"answer_or_question": AnswerOrQuestionAnswer}),
	)

}

func (l *ChatRecordRepository) GetUser(ctx context.Context, startId, endId, chatType int64, startTime, endTime string) (chatRecordPos []*model.ChatRecord, err error) {

	rowBuilder := l.svcCtx.ChatRecordModel.RowBuilder()

	if startTime != "" {
		rowBuilder = rowBuilder.Where("created_at >= ?", startTime)
	}

	if endTime != "" {
		rowBuilder = rowBuilder.Where("created_at < ?", endTime)
	}

	if startId > 0 {
		rowBuilder = rowBuilder.Where("id >= ?", startId)
	}

	if endId > 0 {
		rowBuilder = rowBuilder.Where("created_at <= ?", endId)
	}

	if chatType > 0 {
		rowBuilder = rowBuilder.Where("chat_type == ?", chatType)
	}

	return l.svcCtx.ChatRecordModel.FindAll(ctx, rowBuilder)

}

func (l *ChatRecordRepository) GetAllKfRecord(openKfId, user, startTime, endTime string, state []int64, stateId int64) (chatRecordPos []*model.ChatRecord, err error) {
	l.Logger.Info(" GetCustomerChatRecord ChatRecordRepository All start ")

	rowBuilder := l.svcCtx.ChatRecordModel.RowBuilder()
	if user != "" {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"user": user})
	}

	if openKfId != "" {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"open_kf_id": openKfId})
	}
	if len(state) > 0 {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"state": state})
	}
	if stateId != 0 {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"state_id": stateId})
	}
	if startTime != "" {
		rowBuilder = rowBuilder.Where("created_at >= ?", startTime)
	}

	if endTime != "" {
		rowBuilder = rowBuilder.Where("created_at < ?", endTime)
	}

	chatRecordPos, err = l.svcCtx.ChatRecordModel.FindAll(l.ctx, rowBuilder)
	if err != nil {
		return
	}
	l.Logger.Info(" GetCustomerChatRecord ChatRecordRepository end ")

	return
}

func (l *ChatRecordRepository) GetOne(relationId, aq int64) (chatRecordPo *model.ChatRecord, err error) {

	rowBuilder := l.svcCtx.ChatRecordModel.RowBuilder()

	if relationId != 0 {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"id": relationId})
	}

	if aq != 0 {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"answer_or_question": aq})
	}

	return l.svcCtx.ChatRecordModel.FindOneByQuery(l.ctx, rowBuilder)
}
