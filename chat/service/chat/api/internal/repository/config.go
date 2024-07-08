package repository

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type ConfigRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

const (
	ConfigTypeOpenAi  = 1 // open ai keys
	ConfigTypeProblem = 2 //每日运势问题
	ConfigTypeType    = 3 //心理沙盘类型
	ConfigTypePay     = 4 //付款配置
)

func NewConfigRepository(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigRepository {
	return &ConfigRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigRepository) GetAll(configType int64) (ConfigPo []*model.Config, err error) {

	rowBuilder := l.svcCtx.ConfigModel.RowBuilder()

	if configType != 0 {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"config_type": configType})
	}

	rowBuilder = rowBuilder.OrderBy("id asc")

	ConfigPo, err = l.svcCtx.ConfigModel.FindAll(l.ctx, rowBuilder)
	if err != nil {
		return
	}
	return
}

func (l *ConfigRepository) GetById(id int64) (ConfigPo *model.Config, err error) {
	return l.svcCtx.ConfigModel.FindOne(l.ctx, id)
}

func (l *ConfigRepository) Insert(ConfigPo *model.Config) (sql.Result, error) {
	l.svcCtx.ConfigModel.BuildFiled(nil, ConfigPo)
	return l.svcCtx.ConfigModel.Insert(l.ctx, ConfigPo)
}

func (l *ConfigRepository) Update(id int64, ConfigPo *model.Config) error {
	old, err := l.GetById(id)
	if err != nil {
		return err
	}
	if old == nil || old.Id <= 0 {
		return errors.New("record not find")
	}
	l.svcCtx.ConfigModel.BuildFiled(old, ConfigPo)
	return l.svcCtx.ConfigModel.Update(l.ctx, ConfigPo)
}

func (l *ConfigRepository) GetByName(value string, configType int64) (customerPo *model.Config, err error) {
	builder := l.svcCtx.ConfigModel.RowBuilder()
	if configType != 0 {
		builder = builder.Where(squirrel.Eq{"config_type": configType})
	}
	builder = builder.Where(squirrel.Eq{"value": value})
	customerPo, err = l.svcCtx.ConfigModel.FindOneByQuery(l.ctx, builder)
	return
}

func (l *ConfigRepository) GetByIds(kfId []int64) (CustomerConfigPo []*model.Config, err error) {
	return l.svcCtx.ConfigModel.FindAll(l.ctx,
		l.svcCtx.ConfigModel.RowBuilder().Where(squirrel.Eq{"agent_id": kfId}),
	)
}

func (l *ConfigRepository) Delete(id int64) error {
	return l.svcCtx.ConfigModel.Delete(l.ctx, id)
}

func (l *ConfigRepository) GetOne(name string, configType int64) (customerPo *model.Config, err error) {
	builder := l.svcCtx.ConfigModel.RowBuilder()
	if configType != 0 {
		builder = builder.Where(squirrel.Eq{"config_type": configType})
	}
	builder = builder.Where(squirrel.Eq{"name": name})
	customerPo, err = l.svcCtx.ConfigModel.FindOneByQuery(l.ctx, builder)
	return
}

func (l *ConfigRepository) GetPayConfig(payType, configTypePay int64) (payConfig *model.PayConfig, err error) {
	config, err1 := l.GetOne(strconv.Itoa(int(payType)), configTypePay)
	if err1 != nil {
		return payConfig, err1
	}
	if config == nil || config.Id <= 0 {
		return payConfig, util.ReturnError(xerr.ConfigEmpty)
	}
	_ = json.Unmarshal([]byte(config.Value), &payConfig)
	return
}
