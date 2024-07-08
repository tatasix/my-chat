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
	feedbackFieldNames          = builder.RawFieldNames(&Feedback{})
	feedbackRows                = strings.Join(feedbackFieldNames, ",")
	feedbackRowsExpectAutoSet   = strings.Join(stringx.Remove(feedbackFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	feedbackRowsWithPlaceHolder = strings.Join(stringx.Remove(feedbackFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheFeedbackIdPrefix = "cache:feedback:id:"
)

type (
	feedbackModel interface {
		Insert(ctx context.Context, data *Feedback) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Feedback, error)
		Update(ctx context.Context, data *Feedback) error
		Delete(ctx context.Context, id int64) error
	}

	defaultFeedbackModel struct {
		sqlc.CachedConn
		table string
	}

	Feedback struct {
		Id        int64          `db:"id"`
		User      string         `db:"user"`       // weCom用户标识/customer用户标识
		MessageId string         `db:"message_id"` // message_id customer消息唯一ID
		OpenKfId  string         `db:"open_kf_id"` // 客服标识
		Title     string         `db:"title"`      // 用户反馈标题
		Content   sql.NullString `db:"content"`    // 用户反馈内容
		Reply     sql.NullString `db:"reply"`      // 回复内容
		Contact   string         `db:"contact"`    // 联系方式
		Images    string         `db:"images"`     // 图片
		Status    int64          `db:"status"`     // 状态:1待处理，2已处理
		CreatedAt time.Time      `db:"created_at"` // 创建时间
		UpdatedAt time.Time      `db:"updated_at"` // 更新时间
		IsDeleted int64          `db:"is_deleted"` // 是否删除，0否，1是
	}
)

func newFeedbackModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultFeedbackModel {
	return &defaultFeedbackModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`feedback`",
	}
}

func (m *defaultFeedbackModel) Delete(ctx context.Context, id int64) error {
	feedbackIdKey := fmt.Sprintf("%s%v", cacheFeedbackIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, feedbackIdKey)
	return err
}

func (m *defaultFeedbackModel) FindOne(ctx context.Context, id int64) (*Feedback, error) {
	feedbackIdKey := fmt.Sprintf("%s%v", cacheFeedbackIdPrefix, id)
	var resp Feedback
	err := m.QueryRowCtx(ctx, &resp, feedbackIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", feedbackRows, m.table)
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

func (m *defaultFeedbackModel) Insert(ctx context.Context, data *Feedback) (sql.Result, error) {
	feedbackIdKey := fmt.Sprintf("%s%v", cacheFeedbackIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, feedbackRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.User, data.MessageId, data.OpenKfId, data.Title, data.Content, data.Reply, data.Contact, data.Images, data.Status, data.IsDeleted)
	}, feedbackIdKey)
	return ret, err
}

func (m *defaultFeedbackModel) Update(ctx context.Context, data *Feedback) error {
	feedbackIdKey := fmt.Sprintf("%s%v", cacheFeedbackIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, feedbackRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.User, data.MessageId, data.OpenKfId, data.Title, data.Content, data.Reply, data.Contact, data.Images, data.Status, data.IsDeleted, data.Id)
	}, feedbackIdKey)
	return err
}

func (m *defaultFeedbackModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheFeedbackIdPrefix, primary)
}

func (m *defaultFeedbackModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", feedbackRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultFeedbackModel) tableName() string {
	return m.table
}
