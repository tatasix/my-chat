// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	chatRoomUsersFieldNames          = builder.RawFieldNames(&ChatRoomUsers{})
	chatRoomUsersRows                = strings.Join(chatRoomUsersFieldNames, ",")
	chatRoomUsersRowsExpectAutoSet   = strings.Join(stringx.Remove(chatRoomUsersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	chatRoomUsersRowsWithPlaceHolder = strings.Join(stringx.Remove(chatRoomUsersFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheChatRoomUsersIdPrefix = "cache:chatRoomUsers:id:"
)

type (
	chatRoomUsersModel interface {
		Insert(ctx context.Context, data *ChatRoomUsers) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ChatRoomUsers, error)
		Update(ctx context.Context, data *ChatRoomUsers) error
		Delete(ctx context.Context, id int64) error
	}

	defaultChatRoomUsersModel struct {
		sqlc.CachedConn
		table string
	}

	ChatRoomUsers struct {
		Id        int64     `db:"id"`
		GroupId   int64     `db:"group_id"`
		User      string    `db:"user"`       // 用户
		ChannelId int64     `db:"channel_id"` // 频道id
		IsManager int64     `db:"is_manager"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		IsDeleted int64     `db:"is_deleted"` // 是否删除
	}
)

func newChatRoomUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultChatRoomUsersModel {
	return &defaultChatRoomUsersModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`chat_room_users`",
	}
}

func (m *defaultChatRoomUsersModel) Delete(ctx context.Context, id int64) error {
	chatRoomUsersIdKey := fmt.Sprintf("%s%v", cacheChatRoomUsersIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, chatRoomUsersIdKey)
	return err
}

func (m *defaultChatRoomUsersModel) FindOne(ctx context.Context, id int64) (*ChatRoomUsers, error) {
	chatRoomUsersIdKey := fmt.Sprintf("%s%v", cacheChatRoomUsersIdPrefix, id)
	var resp ChatRoomUsers
	err := m.QueryRowCtx(ctx, &resp, chatRoomUsersIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", chatRoomUsersRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultChatRoomUsersModel) Insert(ctx context.Context, data *ChatRoomUsers) (sql.Result, error) {
	chatRoomUsersIdKey := fmt.Sprintf("%s%v", cacheChatRoomUsersIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, chatRoomUsersRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.GroupId, data.User, data.ChannelId, data.IsManager, data.IsDeleted)
	}, chatRoomUsersIdKey)
	return ret, err
}

func (m *defaultChatRoomUsersModel) Update(ctx context.Context, data *ChatRoomUsers) error {
	chatRoomUsersIdKey := fmt.Sprintf("%s%v", cacheChatRoomUsersIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, chatRoomUsersRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.GroupId, data.User, data.ChannelId, data.IsManager, data.IsDeleted, data.Id)
	}, chatRoomUsersIdKey)
	return err
}

func (m *defaultChatRoomUsersModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheChatRoomUsersIdPrefix, primary)
}

func (m *defaultChatRoomUsersModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", chatRoomUsersRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultChatRoomUsersModel) tableName() string {
	return m.table
}