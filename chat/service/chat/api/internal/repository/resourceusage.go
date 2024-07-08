package repository

import (
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type ResourceUsageRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResourceUsageRepository(ctx context.Context, svcCtx *svc.ServiceContext) *ResourceUsageRepository {
	return &ResourceUsageRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResourceUsageRepository) GetByName(nickname string) (customerPo *model.ResourceUsage, err error) {

	customerPo, err = l.svcCtx.ResourceUsageModel.FindOneByQuery(l.ctx,
		l.svcCtx.ResourceUsageModel.RowBuilder().Where(squirrel.Eq{"nickname": nickname}),
	)
	return
}

func (l *ResourceUsageRepository) GetByUsers(users []string) (ResourceUsagePo []*model.ResourceUsage, err error) {
	return l.svcCtx.ResourceUsageModel.FindAll(l.ctx,
		l.svcCtx.ResourceUsageModel.RowBuilder().Where(squirrel.Eq{"user": users}),
	)
}

func (l *ResourceUsageRepository) GetByUserAndDate(user, date string, hour int) (ResourceUsagePo *model.ResourceUsage, err error) {
	return l.svcCtx.ResourceUsageModel.FindOneByQuery(l.ctx,
		l.svcCtx.ResourceUsageModel.RowBuilder().Where(squirrel.Eq{"user": user}).Where(squirrel.Eq{"date": date}).Where(squirrel.Eq{"hour": hour}),
	)
}

func (l *ResourceUsageRepository) GetById(id int64) (ResourceUsage *model.ResourceUsage, err error) {
	return l.svcCtx.ResourceUsageModel.FindOne(context.Background(), id)
}

func (l *ResourceUsageRepository) Insert(ResourceUsage *model.ResourceUsage) (err error) {
	l.svcCtx.ResourceUsageModel.BuildFiled(nil, ResourceUsage)
	_, err = l.svcCtx.ResourceUsageModel.Insert(l.ctx, ResourceUsage)
	return
}

func (l *ResourceUsageRepository) Update(old, ResourceUsage *model.ResourceUsage) error {
	l.svcCtx.ResourceUsageModel.BuildFiled(old, ResourceUsage)
	return l.svcCtx.ResourceUsageModel.Update(l.ctx, ResourceUsage)
}

func (l *ResourceUsageRepository) SaveTimesAndToken(user string, now time.Time, times, token int64) (err error) {
	date := now.Format(vars.TimeFormat2)
	hour := now.Hour()
	old, err := l.GetByUserAndDate(user, date, hour)
	if nil != err {
		return
	}
	ResourceUsage := &model.ResourceUsage{
		User:      user,
		Date:      now,
		Hour:      int64(hour),
		Times:     times,
		Token:     token,
		UpdatedAt: now,
	}
	if old != nil && old.Id > 0 {
		//存在
		return l.svcCtx.ResourceUsageModel.UpdateTimesAndToken(l.ctx, old.Id, times, token)
	}
	//不存在
	ResourceUsage.CreatedAt = now
	_, err = l.svcCtx.ResourceUsageModel.Insert(l.ctx, ResourceUsage)
	return
}
