package repository

import (
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"database/sql"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type ApplicationConfigRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplicationConfigRepository(ctx context.Context, svcCtx *svc.ServiceContext) *ApplicationConfigRepository {
	return &ApplicationConfigRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplicationConfigRepository) GetAll(agentName, model, startTime, endTime, order string, page, limit uint64) (applicationConfigPo []*model.ApplicationConfig, count int64, err error) {

	countBuilder := l.svcCtx.ApplicationConfigModel.CountBuilder("id")
	rowBuilder := l.svcCtx.ApplicationConfigModel.RowBuilder()
	if agentName != "" {
		countBuilder = countBuilder.Where(squirrel.Eq{"agent_name": agentName})
		rowBuilder = rowBuilder.Where(squirrel.Eq{"agent_name": agentName})
	}
	if model != "" {
		countBuilder = countBuilder.Where(squirrel.Eq{"post_model": model})
		rowBuilder = rowBuilder.Where(squirrel.Eq{"post_model": model})
	}

	if startTime != "" {
		countBuilder = countBuilder.Where("created_at >= ?", startTime)
		rowBuilder = rowBuilder.Where("created_at >= ?", startTime)
	}

	if endTime != "" {
		countBuilder = countBuilder.Where("created_at < ?", endTime)
		rowBuilder = rowBuilder.Where("created_at < ?", endTime)
	}

	count, err = l.svcCtx.ApplicationConfigModel.FindCount(l.ctx, countBuilder)
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
	applicationConfigPo, err = l.svcCtx.ApplicationConfigModel.FindAll(l.ctx, rowBuilder)
	if err != nil {
		return
	}
	return
}

func (l *ApplicationConfigRepository) GetById(id int64) (applicationConfigPo *model.ApplicationConfig, err error) {
	return l.svcCtx.ApplicationConfigModel.FindOne(l.ctx, id)
}

func (l *ApplicationConfigRepository) Insert(applicationConfigPo *model.ApplicationConfig) (sql.Result, error) {
	l.svcCtx.ApplicationConfigModel.BuildFiled(nil, applicationConfigPo)
	return l.svcCtx.ApplicationConfigModel.Insert(l.ctx, applicationConfigPo)
}

func (l *ApplicationConfigRepository) Update(id int64, applicationConfigPo *model.ApplicationConfig) error {
	old, err := l.GetById(id)
	if err != nil {
		return err
	}
	if old == nil || old.Id <= 0 {
		return errors.New("record not find")
	}
	l.svcCtx.ApplicationConfigModel.BuildFiled(old, applicationConfigPo)
	return l.svcCtx.ApplicationConfigModel.Update(l.ctx, applicationConfigPo)
}

func (l *ApplicationConfigRepository) GetByName(agentName string) (customerPo *model.ApplicationConfig, err error) {

	customerPo, err = l.svcCtx.ApplicationConfigModel.FindOneByQuery(l.ctx,
		l.svcCtx.ApplicationConfigModel.RowBuilder().Where(squirrel.Eq{"agent_name": agentName}),
	)
	return
}

func (l *ApplicationConfigRepository) GetByIds(kfId []int64) (CustomerConfigPo []*model.ApplicationConfig, err error) {
	return l.svcCtx.ApplicationConfigModel.FindAll(l.ctx,
		l.svcCtx.ApplicationConfigModel.RowBuilder().Where(squirrel.Eq{"agent_id": kfId}),
	)
}

func (l *ApplicationConfigRepository) Delete(id int64) error {
	return l.svcCtx.ApplicationConfigModel.Delete(l.ctx, id)
}
