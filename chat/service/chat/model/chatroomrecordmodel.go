package model

import (
	"chat/common/page"
	"chat/common/util"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ChatRoomRecordModel = (*customChatRoomRecordModel)(nil)

type (
	// ChatRoomRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatRoomRecordModel.
	ChatRoomRecordModel interface {
		chatRoomRecordModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*ChatRoomRecord, error)
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*ChatRoomRecord, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, p, s int) ([]*ChatRoomRecord, error)
		BuildFiled(old, new *ChatRoomRecord)
	}

	customChatRoomRecordModel struct {
		*defaultChatRoomRecordModel
	}
)

// NewChatRoomRecordModel returns a model for the database table.
func NewChatRoomRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ChatRoomRecordModel {
	return &customChatRoomRecordModel{
		defaultChatRoomRecordModel: newChatRoomRecordModel(conn, c, opts...),
	}
}

func (m *defaultChatRoomRecordModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*ChatRoomRecord, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp ChatRoomRecord
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &resp, nil
}

func (m *defaultChatRoomRecordModel) FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error) {

	query, values, err := sumBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultChatRoomRecordModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultChatRoomRecordModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*ChatRoomRecord, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ChatRoomRecord
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultChatRoomRecordModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, p, s int) ([]*ChatRoomRecord, error) {

	pg := page.NewPage(p, s)

	query, values, err := rowBuilder.Limit(pg.Limit()).Offset(pg.Offset()).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ChatRoomRecord
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// export logic
func (m *defaultChatRoomRecordModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

// export logic
func (m *defaultChatRoomRecordModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(chatRoomRecordRows).From(m.table)
}

// export logic
func (m *defaultChatRoomRecordModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

// export logic
func (m *defaultChatRoomRecordModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

func (m *defaultChatRoomRecordModel) BuildFiled(old, new *ChatRoomRecord) {
	if new == nil {
		return
	}
	if old == nil {
		newId := util.GenerateSnowflakeInt64()
		if new.Id == 0 {
			new.Id = newId
		}
		return
	}
	new.Id = old.Id
	return
}
