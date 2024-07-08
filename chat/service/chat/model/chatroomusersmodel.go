package model

import (
	"chat/common/page"
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ChatRoomUsersModel = (*customChatRoomUsersModel)(nil)

type (
	// ChatRoomUsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatRoomUsersModel.
	ChatRoomUsersModel interface {
		chatRoomUsersModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*ChatRoomUsers, error)
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*ChatRoomUsers, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, p, s int) ([]*ChatRoomUsers, error)
		BuildFiled(old, new *ChatRoomUsers)
		SoftDelete(ctx context.Context, channelId int64) error
	}

	customChatRoomUsersModel struct {
		*defaultChatRoomUsersModel
	}
)

// NewChatRoomUsersModel returns a model for the database table.
func NewChatRoomUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ChatRoomUsersModel {
	return &customChatRoomUsersModel{
		defaultChatRoomUsersModel: newChatRoomUsersModel(conn, c, opts...),
	}
}

func (m *defaultChatRoomUsersModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*ChatRoomUsers, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp ChatRoomUsers
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &resp, nil
}

func (m *defaultChatRoomUsersModel) FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error) {

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

func (m *defaultChatRoomUsersModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultChatRoomUsersModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*ChatRoomUsers, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ChatRoomUsers
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultChatRoomUsersModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, p, s int) ([]*ChatRoomUsers, error) {

	pg := page.NewPage(p, s)

	query, values, err := rowBuilder.Limit(pg.Limit()).Offset(pg.Offset()).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ChatRoomUsers
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// export logic
func (m *defaultChatRoomUsersModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

// export logic
func (m *defaultChatRoomUsersModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(chatRoomUsersRows).From(m.table)
}

// export logic
func (m *defaultChatRoomUsersModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table).Where(squirrel.Eq{"is_deleted": IsNotDeleted})
}

// export logic
func (m *defaultChatRoomUsersModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

func (m *defaultChatRoomUsersModel) BuildFiled(old, new *ChatRoomUsers) {
	if new == nil {
		return
	}
	if old == nil {
		return
	}
	new.Id = old.Id
	return
}

func (m *defaultChatRoomUsersModel) SoftDelete(ctx context.Context, channelId int64) error {
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set is_deleted=1 where `channel_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, channelId)
	})
	return err
}
