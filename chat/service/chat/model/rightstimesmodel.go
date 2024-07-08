package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ RightsTimesModel = (*customRightsTimesModel)(nil)

type (
	// RightsTimesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRightsTimesModel.
	RightsTimesModel interface {
		rightsTimesModel
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*RightsTimes, error)
		RowBuilder() squirrel.SelectBuilder
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*RightsTimes, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		CountBuilder(field string) squirrel.SelectBuilder
		UpdateTimesById(ctx context.Context, id, times int64) (err error)
		AddTimesByUser(ctx context.Context, user, kfId, now string, times int64) (err error)
		BuildFiled(old, new *RightsTimes)
		DeleteByUser(ctx context.Context, user, now string) error
	}

	customRightsTimesModel struct {
		*defaultRightsTimesModel
	}
)

// NewRightsTimesModel returns a model for the database table.
func NewRightsTimesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RightsTimesModel {
	return &customRightsTimesModel{
		defaultRightsTimesModel: newRightsTimesModel(conn, c, opts...),
	}
}

func (m *defaultRightsTimesModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*RightsTimes, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp RightsTimes
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
func (m *defaultRightsTimesModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(rightsTimesRows).From(m.table).Where(squirrel.Eq{"is_deleted": IsNotDeleted})
}

func (m *defaultRightsTimesModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*RightsTimes, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*RightsTimes
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultRightsTimesModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultRightsTimesModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultRightsTimesModel) BuildFiled(old, new *RightsTimes) {
	if new == nil {
		return
	}
	if old == nil && new.Id == 0 {
		//new.OutTradeNo = util.GenerateSnowflakeString()
		return
	}
	new.Id = old.Id

	new.UpdatedAt = time.Now()

	return
}

func (m *defaultRightsTimesModel) UpdateTimesById(ctx context.Context, id, times int64) (err error) {

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `used`= `used`+ ?,`remain`= `remain` - ? where `id` = ? and `remain`>= ?", m.table)
		return conn.ExecCtx(ctx, query, times, times, id, times)
	})

	return
}
func (m *defaultRightsTimesModel) AddTimesByUser(ctx context.Context, user, kfId, now string, times int64) (err error) {
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `total`= `total`+ ?,`remain`= `remain` + ? where `open_kf_id` = ? and `user` = ? and `start`<= ? and `end`>= ? ", m.table)
		return conn.ExecCtx(ctx, query, times, times, kfId, user, now, now)
	})
	return
}

func (m *defaultRightsTimesModel) DeleteByUser(ctx context.Context, user, now string) (err error) {
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user` = ?  and `start`<= ? and `end`>= ? ", m.table)
		return conn.ExecCtx(ctx, query, user, now, now)
	})
	return
}
