package repository

import (
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

const ()

type StatisticsRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatisticsRepository(ctx context.Context, svcCtx *svc.ServiceContext) *StatisticsRepository {
	return &StatisticsRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatisticsRepository) GetByName(nickname string) (customerPo *model.Statistics, err error) {

	customerPo, err = l.svcCtx.StatisticsModel.FindOneByQuery(l.ctx,
		l.svcCtx.StatisticsModel.RowBuilder().Where(squirrel.Eq{"nickname": nickname}),
	)
	return
}

func (l *StatisticsRepository) GetByUsers(users []string) (StatisticsPo []*model.Statistics, err error) {
	return l.svcCtx.StatisticsModel.FindAll(l.ctx,
		l.svcCtx.StatisticsModel.RowBuilder().Where(squirrel.Eq{"user": users}),
	)
}

func (l *StatisticsRepository) GetByUser(user string) (StatisticsPo *model.Statistics, err error) {
	return l.svcCtx.StatisticsModel.FindOneByQuery(l.ctx,
		l.svcCtx.StatisticsModel.RowBuilder().Where(squirrel.Eq{"user": user}),
	)
}

func (l *StatisticsRepository) GetById(id int64) (Statistics *model.Statistics, err error) {
	return l.svcCtx.StatisticsModel.FindOne(context.Background(), id)
}

func (l *StatisticsRepository) Insert(Statistics *model.Statistics) (err error) {
	l.svcCtx.StatisticsModel.BuildFiled(nil, Statistics)
	_, err = l.svcCtx.StatisticsModel.Insert(l.ctx, Statistics)
	return
}

func (l *StatisticsRepository) Update(old, Statistics *model.Statistics) error {
	l.svcCtx.StatisticsModel.BuildFiled(old, Statistics)
	return l.svcCtx.StatisticsModel.Update(l.ctx, Statistics)
}

func (l *StatisticsRepository) GetAll(startTime, endTime, order string, page, limit uint64) (pos []*model.Statistics, count int64, err error) {

	countBuilder := l.svcCtx.StatisticsModel.CountBuilder("id")
	rowBuilder := l.svcCtx.StatisticsModel.RowBuilder()

	if startTime != "" {
		countBuilder = countBuilder.Where("date >= ?", startTime)
		rowBuilder = rowBuilder.Where("date >= ?", startTime)
	}

	if endTime != "" {
		countBuilder = countBuilder.Where("date <= ?", endTime)
		rowBuilder = rowBuilder.Where("date <= ?", endTime)
	}

	count, err = l.svcCtx.StatisticsModel.FindCount(l.ctx, countBuilder)
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
	pos, err = l.svcCtx.StatisticsModel.FindAll(l.ctx, rowBuilder)
	if err != nil {
		return
	}
	l.Logger.Info(" GetCustomerChatRecord ChatRecordRepository end ")

	return
}
