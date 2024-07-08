package repository

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

const ()

type ActivityLogRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityLogRepository(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityLogRepository {
	return &ActivityLogRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityLogRepository) GetByName(nickname string) (customerPo *model.ActivityLog, err error) {

	customerPo, err = l.svcCtx.ActivityLogModel.FindOneByQuery(l.ctx,
		l.svcCtx.ActivityLogModel.RowBuilder().Where(squirrel.Eq{"nickname": nickname}),
	)
	return
}

func (l *ActivityLogRepository) GetByUsers(users []string) (ActivityLogPo []*model.ActivityLog, err error) {
	return l.svcCtx.ActivityLogModel.FindAll(l.ctx,
		l.svcCtx.ActivityLogModel.RowBuilder().Where(squirrel.Eq{"user": users}),
	)
}

func (l *ActivityLogRepository) GetByUser(user string) (ActivityLogPo *model.ActivityLog, err error) {
	return l.svcCtx.ActivityLogModel.FindOneByQuery(l.ctx,
		l.svcCtx.ActivityLogModel.RowBuilder().Where(squirrel.Eq{"user": user}),
	)
}

func (l *ActivityLogRepository) UpdateById(id int64, ActivityLog *model.ActivityLog) error {
	old, err := l.GetById(id)
	if err != nil {
		return err
	}
	if old == nil || old.Id <= 0 {
		return util.ReturnError(xerr.RecordNotFound)
	}
	l.svcCtx.ActivityLogModel.BuildFiled(old, ActivityLog)
	return l.svcCtx.ActivityLogModel.Update(context.Background(), ActivityLog)
}

func (l *ActivityLogRepository) GetById(id int64) (ActivityLog *model.ActivityLog, err error) {
	return l.svcCtx.ActivityLogModel.FindOne(context.Background(), id)
}

func (l *ActivityLogRepository) Insert(ActivityLog *model.ActivityLog) (err error) {
	l.svcCtx.ActivityLogModel.BuildFiled(nil, ActivityLog)
	_, err = l.svcCtx.ActivityLogModel.Insert(l.ctx, ActivityLog)
	return
}

func (l *ActivityLogRepository) Update(old, ActivityLog *model.ActivityLog) error {
	l.svcCtx.ActivityLogModel.BuildFiled(old, ActivityLog)
	return l.svcCtx.ActivityLogModel.Update(l.ctx, ActivityLog)
}

func (l *ActivityLogRepository) UpdateUser(new, old string) error {
	return l.svcCtx.ActivityLogModel.UpdateUser(l.ctx, new, old)
}

func (l *ActivityLogRepository) GetDaily(startTime, endTime string) (count int64, err error) {
	countBuilder := l.svcCtx.ActivityLogModel.CountBuilder("distinct user").
		Where("begin_time >= ?", startTime).Where("begin_time <= ?", endTime)

	return l.svcCtx.ActivityLogModel.FindCount(l.ctx, countBuilder)
}
