package model

import (
	"chat/common/util"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ StateModel = (*customStateModel)(nil)

type (
	// StateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStateModel.
	StateModel interface {
		stateModel
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*State, error)
		RowBuilder() squirrel.SelectBuilder
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*State, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		CountBuilder(field string) squirrel.SelectBuilder
		BuildFiled(old, new *State)
	}

	customStateModel struct {
		*defaultStateModel
	}
)

// NewStateModel returns a model for the database table.
func NewStateModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) StateModel {
	return &customStateModel{
		defaultStateModel: newStateModel(conn, c, opts...),
	}
}

func (m *defaultStateModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*State, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp State
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
func (m *defaultStateModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(stateRows).From(m.table).Where(squirrel.Eq{"is_deleted": IsNotDeleted})
}

func (m *defaultStateModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*State, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*State
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultStateModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

	query, values, err := countBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return 0, nil
		}
		return 0, err
	}
	return resp, nil
}

func (m *defaultStateModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultStateModel) BuildFiled(old, new *State) {
	if new == nil {
		return
	}
	if old == nil && new.Id == 0 {
		new.Id = util.GenerateSnowflakeInt64()
		return
	}
	new.Id = old.Id
	//还需要处理
	if new.User == "" {
		new.User = old.User
	}

	if new.KfId == "" {
		new.KfId = old.KfId
	}
	if new.State == 0 {
		new.State = old.State
	}

	new.Date = old.Date

	new.UpdatedAt = time.Now()

	return
}
