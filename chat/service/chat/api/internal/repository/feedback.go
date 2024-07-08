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
	StatusToDo int64 = 1 //待处理
	StatusDone int64 = 2 //已处理
)

type FeedbackRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedbackRepository(ctx context.Context, svcCtx *svc.ServiceContext) *FeedbackRepository {
	return &FeedbackRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedbackRepository) GetAll(openKfId, user, startTime, endTime, order string, page, limit uint64, status int32) (FeedbackPos []*model.Feedback, count int64, err error) {

	countBuilder := l.svcCtx.FeedbackModel.CountBuilder("id")
	rowBuilder := l.svcCtx.FeedbackModel.RowBuilder()
	if user != "" {
		countBuilder = countBuilder.Where(squirrel.Eq{"user": user})
		rowBuilder = rowBuilder.Where(squirrel.Eq{"user": user})
	}

	if openKfId != "" {
		countBuilder = countBuilder.Where(squirrel.Eq{"open_kf_id": openKfId})
		rowBuilder = rowBuilder.Where(squirrel.Eq{"open_kf_id": openKfId})
	}

	if startTime != "" {
		countBuilder = countBuilder.Where("created_at >= ?", startTime)
		rowBuilder = rowBuilder.Where("created_at >= ?", startTime)
	}

	if endTime != "" {
		countBuilder = countBuilder.Where("created_at < ?", endTime)
		rowBuilder = rowBuilder.Where("created_at < ?", endTime)
	}

	if status != 0 {
		countBuilder = countBuilder.Where(squirrel.Eq{"status": status})
		rowBuilder = rowBuilder.Where(squirrel.Eq{"status": status})
	}

	count, err = l.svcCtx.FeedbackModel.FindCount(l.ctx, countBuilder)
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
	FeedbackPos, err = l.svcCtx.FeedbackModel.FindAll(l.ctx, rowBuilder)
	if err != nil {
		return
	}
	return
}

func (l *FeedbackRepository) GetById(id int64) (*model.Feedback, error) {
	return l.svcCtx.FeedbackModel.FindOne(l.ctx, id)
}

func (l *FeedbackRepository) GetByMessageAndCustomer(messageId, user string) (*model.Feedback, error) {
	return l.svcCtx.FeedbackModel.FindOneByQuery(l.ctx,
		l.svcCtx.FeedbackModel.RowBuilder().Where(squirrel.Eq{"message_id": messageId}).Where(squirrel.Eq{"user": user}),
	)
}

func (l *FeedbackRepository) Insert(Feedback *model.Feedback) (sql.Result, error) {
	l.svcCtx.FeedbackModel.BuildFiled(nil, Feedback)
	return l.svcCtx.FeedbackModel.Insert(l.ctx, Feedback)
}
