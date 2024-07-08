package model

import (
	"chat/common/util"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ ConfigModel = (*customConfigModel)(nil)

type (
	// ConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customConfigModel.
	ConfigModel interface {
		configModel
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*Config, error)
		RowBuilder() squirrel.SelectBuilder
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*Config, error)
		CountBuilder(field string) squirrel.SelectBuilder
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		BuildFiled(old, new *Config)
	}

	customConfigModel struct {
		*defaultConfigModel
	}
)

// NewConfigModel returns a model for the database table.
func NewConfigModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ConfigModel {
	return &customConfigModel{
		defaultConfigModel: newConfigModel(conn, c, opts...),
	}
}

func (m *defaultConfigModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*Config, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp Config
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &resp, nil
}

// export logic
func (m *defaultConfigModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(configRows).From(m.table).Where(squirrel.Eq{"is_deleted": IsNotDeleted})
}

func (m *defaultConfigModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*Config, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Config
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
func (m *defaultConfigModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table).Where(squirrel.Eq{"is_deleted": IsNotDeleted})
}

func (m *defaultConfigModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

	query, values, err := countBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultConfigModel) BuildFiled(old, new *Config) {
	if new == nil {
		return
	}
	if old == nil && new.Id == 0 {
		new.Id = util.GenerateSnowflakeInt64()
		return
	}
	new.Id = old.Id
	if new.ConfigType == 0 {
		new.ConfigType = old.ConfigType
	}

	if new.Name == "" {
		new.Name = old.Name
	}

	if new.Description == "" {
		new.Description = old.Description
	}

	if new.Value == "" {
		new.Value = old.Value
	}

	if new.CreatedBy == "" {
		new.CreatedBy = old.CreatedBy
	}

	if new.UpdatedBy == "" {
		new.UpdatedBy = old.UpdatedBy
	}

	if new.IsDeleted == 0 {
		new.IsDeleted = old.IsDeleted
	}

	new.UpdatedAt = time.Now()

	return
}
