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

var _ RightsRecordModel = (*customRightsRecordModel)(nil)

type (
	// RightsRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRightsRecordModel.
	RightsRecordModel interface {
		rightsRecordModel
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*RightsRecord, error)
		RowBuilder() squirrel.SelectBuilder
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*RightsRecord, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		CountBuilder(field string) squirrel.SelectBuilder
		UpdateStatusById(ctx context.Context, id, status int64, reason string) (err error)
		BuildFiled(old, new *RightsRecord)
	}

	customRightsRecordModel struct {
		*defaultRightsRecordModel
	}
)

// NewRightsRecordModel returns a model for the database table.
func NewRightsRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RightsRecordModel {
	return &customRightsRecordModel{
		defaultRightsRecordModel: newRightsRecordModel(conn, c, opts...),
	}
}

func (m *defaultRightsRecordModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*RightsRecord, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp RightsRecord
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
func (m *defaultRightsRecordModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(rightsRecordRows).From(m.table).Where(squirrel.Eq{"is_deleted": IsNotDeleted})
}

func (m *defaultRightsRecordModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*RightsRecord, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*RightsRecord
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultRightsRecordModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultRightsRecordModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultRightsRecordModel) BuildFiled(old, new *RightsRecord) {
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

func (m *defaultRightsRecordModel) UpdateStatusById(ctx context.Context, id, status int64, reason string) (err error) {

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `status`= ?,`reason`= ? where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, status, reason, id)
	})

	return
}
