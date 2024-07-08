package model

import (
	"chat/common/page"
	"chat/common/util"
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ChatRoomModel = (*customChatRoomModel)(nil)

type (
	// ChatRoomModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatRoomModel.
	ChatRoomModel interface {
		chatRoomModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*ChatRoom, error)
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*ChatRoom, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, p, s int) ([]*ChatRoom, error)
		BuildFiled(old, new *ChatRoom)
		SoftDelete(ctx context.Context, channelId int64) error
	}

	customChatRoomModel struct {
		*defaultChatRoomModel
	}
)

// NewChatRoomModel returns a model for the database table.
func NewChatRoomModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ChatRoomModel {
	return &customChatRoomModel{
		defaultChatRoomModel: newChatRoomModel(conn, c, opts...),
	}
}

func (m *defaultChatRoomModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*ChatRoom, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp ChatRoom
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &resp, nil
}

func (m *defaultChatRoomModel) FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error) {

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

func (m *defaultChatRoomModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultChatRoomModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*ChatRoom, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ChatRoom
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultChatRoomModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, p, s int) ([]*ChatRoom, error) {

	pg := page.NewPage(p, s)

	query, values, err := rowBuilder.Limit(pg.Limit()).Offset(pg.Offset()).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ChatRoom
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// export logic
func (m *defaultChatRoomModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

// export logic
func (m *defaultChatRoomModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(chatRoomRows).From(m.table)
}

// export logic
func (m *defaultChatRoomModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

// export logic
func (m *defaultChatRoomModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

func (m *defaultChatRoomModel) BuildFiled(old, new *ChatRoom) {
	if new == nil {
		return
	}
	if old == nil {
		if new.ChannelId == 0 {
			new.ChannelId = util.GenerateSnowflakeInt64()
		}
		return
	}
	new.Id = old.Id
	return
}

func (m *defaultChatRoomModel) SoftDelete(ctx context.Context, channelId int64) error {
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set is_deleted=1 where `channel_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, channelId)
	})
	return err
}
