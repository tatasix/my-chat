package repository

import (
	"chat/common/redis"
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CustomerConfigRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCustomerConfigRepository(ctx context.Context, svcCtx *svc.ServiceContext) *CustomerConfigRepository {
	return &CustomerConfigRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CustomerConfigRepository) GetAll(isPrompt int, agentName, model, startTime, endTime, order string, page, limit uint64) (CustomerConfigPo []*model.CustomerConfig, count int64, err error) {

	countBuilder := l.svcCtx.CustomerConfigModel.CountBuilder("id")
	rowBuilder := l.svcCtx.CustomerConfigModel.RowBuilder()
	if agentName != "" {
		countBuilder = countBuilder.Where(squirrel.Eq{"kf_name": agentName})
		rowBuilder = rowBuilder.Where(squirrel.Eq{"kf_name": agentName})
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

	if isPrompt == 1 {
		countBuilder = countBuilder.Where(squirrel.Or{squirrel.NotEq{"prompt": ""}, squirrel.Eq{"multiple": 1}})
		rowBuilder = rowBuilder.Where(squirrel.Or{squirrel.NotEq{"prompt": ""}, squirrel.Eq{"multiple": 1}})
	} else if isPrompt == 2 {
		countBuilder = countBuilder.Where(squirrel.And{squirrel.Eq{"prompt": ""}, squirrel.Eq{"multiple": 0}})
		rowBuilder = rowBuilder.Where(squirrel.And{squirrel.Eq{"prompt": ""}, squirrel.Eq{"multiple": 0}})
	}

	count, err = l.svcCtx.CustomerConfigModel.FindCount(l.ctx, countBuilder)
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
	CustomerConfigPo, err = l.svcCtx.CustomerConfigModel.FindAll(l.ctx, rowBuilder)
	if err != nil {
		return
	}
	return
}

func (l *CustomerConfigRepository) GetById(id int64) (CustomerConfigPo *model.CustomerConfig, err error) {
	return l.svcCtx.CustomerConfigModel.FindOne(l.ctx, id)
}

func (l *CustomerConfigRepository) GetByKfId(kfId string) (CustomerConfigPo *model.CustomerConfig, err error) {
	return l.svcCtx.CustomerConfigModel.FindOneByQuery(l.ctx,
		l.svcCtx.CustomerConfigModel.RowBuilder().Where(squirrel.Eq{"kf_id": kfId}),
	)
}

func (l *CustomerConfigRepository) Insert(CustomerConfigPo *model.CustomerConfig) (sql.Result, error) {
	l.svcCtx.CustomerConfigModel.BuildFiled(nil, CustomerConfigPo)
	return l.svcCtx.CustomerConfigModel.Insert(l.ctx, CustomerConfigPo)
}

func (l *CustomerConfigRepository) Update(id int64, CustomerConfigPo *model.CustomerConfig) error {
	old, err := l.GetById(id)
	if err != nil {
		return err
	}
	if old == nil || old.Id <= 0 {
		return errors.New("record not find")
	}
	l.svcCtx.CustomerConfigModel.BuildFiled(old, CustomerConfigPo)
	return l.svcCtx.CustomerConfigModel.Update(l.ctx, CustomerConfigPo)
}

func (l *CustomerConfigRepository) GetByName(kfName string) (customerPo *model.CustomerConfig, err error) {

	customerPo, err = l.svcCtx.CustomerConfigModel.FindOneByQuery(l.ctx,
		l.svcCtx.CustomerConfigModel.RowBuilder().Where(squirrel.Eq{"kf_name": kfName}),
	)
	return
}

func (l *CustomerConfigRepository) GetByKfIds(kfId []string) (CustomerConfigPo []*model.CustomerConfig, err error) {
	return l.svcCtx.CustomerConfigModel.FindAll(l.ctx,
		l.svcCtx.CustomerConfigModel.RowBuilder().Where(squirrel.Eq{"kf_id": kfId}),
	)
}

func (l *CustomerConfigRepository) Delete(id int64) error {
	return l.svcCtx.CustomerConfigModel.Delete(l.ctx, id)
}

func (l *CustomerConfigRepository) GetByKfIdUseCache(kfId string) (customerPo *model.CustomerConfig, err error) {

	customerPoCache, err1 := redis.Rdb.Get(l.ctx, fmt.Sprintf(redis.CustomerConfigGetByKfIdCacheKey, kfId)).Result()

	if err1 == nil {
		_ = json.Unmarshal([]byte(customerPoCache), &customerPo)
		return
	} else {
		customerPo, err = l.svcCtx.CustomerConfigModel.FindOneByQuery(l.ctx,
			l.svcCtx.CustomerConfigModel.RowBuilder().Where(squirrel.Eq{"kf_id": kfId}),
		)
		if err == nil && customerPo != nil && customerPo.Id > 0 {
			redisData, _ := json.Marshal(customerPo)
			redis.Rdb.Set(l.ctx, fmt.Sprintf(redis.CustomerConfigGetByKfIdCacheKey, kfId), string(redisData), redis.CustomerConfigGetByKfIdCacheExpireTime)
		}
	}

	return
}

func (l *CustomerConfigRepository) GetPayConfig(openKfId string, payType int64) (payConfig model.PayConfig, err error) {
	if openKfId == "" {
		openKfId = vars.ChatCat
	}
	customer, err := l.GetByKfIdUseCache(openKfId)
	if err != nil {
		return
	}
	if customer == nil || customer.Id <= 0 {
		err = util.ReturnError(xerr.ConfigEmpty)
		return
	}

	var payConfigs []model.PayConfig
	_ = json.Unmarshal([]byte(customer.Pay), &payConfigs)

	payConfig = payConfigs[payType]
	return
}

func (l *CustomerConfigRepository) GetAllPayConfig(kfId string) (CustomerConfigPo []*model.CustomerConfig, err error) {
	rowBuilder := l.svcCtx.CustomerConfigModel.RowBuilder()
	if kfId != "" {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"kf_id": kfId})
	}
	rowBuilder = rowBuilder.Where("pay <> ''")

	return l.svcCtx.CustomerConfigModel.FindAll(l.ctx, rowBuilder)
}

func (l *CustomerConfigRepository) GetByTypes(customerType []int64, name string) (customerPo []*model.CustomerConfig, err error) {
	rowBuilder := l.svcCtx.CustomerConfigModel.RowBuilder()
	if len(customerType) > 0 {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"type": customerType})
	}
	if name != "" {
		rowBuilder = rowBuilder.Where("kf_name LIKE ? ", fmt.Sprint("%", name, "%"))
	}
	return l.svcCtx.CustomerConfigModel.FindAll(l.ctx, rowBuilder)
}
