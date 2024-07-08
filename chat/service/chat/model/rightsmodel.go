package model

import (
	"chat/common/util"
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ RightsModel = (*customRightsModel)(nil)

type (
	// RightsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRightsModel.
	RightsModel interface {
		rightsModel
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*Rights, error)
		RowBuilder() squirrel.SelectBuilder
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*Rights, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		CountBuilder(field string) squirrel.SelectBuilder
		UpdateStatusById(ctx context.Context, id, status int64, reason string) (err error)
		UpdateStatus(ctx context.Context, expire string, status int64) (err error)
		BuildFiled(old, new *Rights)
		DeleteByUser(ctx context.Context, user string) error
	}

	customRightsModel struct {
		*defaultRightsModel
	}
)

// NewRightsModel returns a model for the database table.
func NewRightsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RightsModel {
	return &customRightsModel{
		defaultRightsModel: newRightsModel(conn, c, opts...),
	}
}

func (m *defaultRightsModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*Rights, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp Rights
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
func (m *defaultRightsModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(rightsRows).From(m.table).Where(squirrel.Eq{"is_deleted": IsNotDeleted})
}

func (m *defaultRightsModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*Rights, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Rights
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultRightsModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultRightsModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultRightsModel) BuildFiled(old, new *Rights) {
	if new == nil {
		return
	}
	if old == nil && new.Id == 0 {
		new.OutTradeNo = util.GenerateSnowflakeString()
		return
	}
	new.Id = old.Id

	new.UpdatedAt = time.Now()

	return
}

func (m *defaultRightsModel) UpdateStatusById(ctx context.Context, id, status int64, reason string) (err error) {

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `status`= ?,`reason`= ? where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, status, reason, id)
	})

	return
}

func (m *defaultRightsModel) UpdateStatus(ctx context.Context, expire string, status int64) (err error) {

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `status`= ? where `pay_type` != 1 and `status` = 2 and `end` < ? ", m.table)
		return conn.ExecCtx(ctx, query, status, expire)
	})

	return
}

func (m *defaultRightsModel) DeleteByUser(ctx context.Context, user string) (err error) {
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user` = ? ", m.table)
		return conn.ExecCtx(ctx, query, user)
	})
	return
}
