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
	applePaySubscriptionFieldNames          = builder.RawFieldNames(&ApplePaySubscription{})
	applePaySubscriptionRows                = strings.Join(applePaySubscriptionFieldNames, ",")
	applePaySubscriptionRowsExpectAutoSet   = strings.Join(stringx.Remove(applePaySubscriptionFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	applePaySubscriptionRowsWithPlaceHolder = strings.Join(stringx.Remove(applePaySubscriptionFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheApplePaySubscriptionIdPrefix = "cache:applePaySubscription:id:"
)

type (
	applePaySubscriptionModel interface {
		Insert(ctx context.Context, data *ApplePaySubscription) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ApplePaySubscription, error)
		Update(ctx context.Context, data *ApplePaySubscription) error
		Delete(ctx context.Context, id int64) error
	}

	defaultApplePaySubscriptionModel struct {
		sqlc.CachedConn
		table string
	}

	ApplePaySubscription struct {
		Id                    int64        `db:"id"`
		UserId                string       `db:"user_id"`                 // user id
		OriginalTransactionId string       `db:"original_transaction_id"` // 原始订阅唯一标识
		OriginalReceipt       string       `db:"original_receipt"`        // 原始票据
		Receipt               string       `db:"receipt"`                 // 上一次的票据
		TransactionId         string       `db:"transaction_id"`          // 每一期的唯一标识
		OrderId               string       `db:"order_id"`                // 订单 id 对应 rights 表 rights_id
		Status                int64        `db:"status"`                  // 订阅中,订阅关闭
		RenewStatus           int64        `db:"renew_status"`            // 续费:扣费成功,等待扣费,扣费失败
		OriginalProductId     string       `db:"original_product_id"`     // 一开始订阅的商品 id
		ProductId             string       `db:"product_id"`              // 现在订阅的商品 id
		ExpiresDate           sql.NullTime `db:"expires_date"`            // 订阅周期过期时间
		PurchaseDate          sql.NullTime `db:"purchase_date"`           // 订单周期开始时间
		TransactionReason     string       `db:"transaction_reason"`
		TransactionType       string       `db:"transaction_type"`
		InAppOwnershipType    string       `db:"in_app_ownership_type"`
		Price                 int64        `db:"price"`
		Currency              string       `db:"currency"`
		CreatedAt             time.Time    `db:"created_at"`
		UpdatedAt             time.Time    `db:"updated_at"`
		IsDeleted             int64        `db:"is_deleted"`
	}
)

func newApplePaySubscriptionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultApplePaySubscriptionModel {
	return &defaultApplePaySubscriptionModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`apple_pay_subscription`",
	}
}

func (m *defaultApplePaySubscriptionModel) Delete(ctx context.Context, id int64) error {
	applePaySubscriptionIdKey := fmt.Sprintf("%s%v", cacheApplePaySubscriptionIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, applePaySubscriptionIdKey)
	return err
}

func (m *defaultApplePaySubscriptionModel) FindOne(ctx context.Context, id int64) (*ApplePaySubscription, error) {
	applePaySubscriptionIdKey := fmt.Sprintf("%s%v", cacheApplePaySubscriptionIdPrefix, id)
	var resp ApplePaySubscription
	err := m.QueryRowCtx(ctx, &resp, applePaySubscriptionIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", applePaySubscriptionRows, m.table)
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

func (m *defaultApplePaySubscriptionModel) Insert(ctx context.Context, data *ApplePaySubscription) (sql.Result, error) {
	applePaySubscriptionIdKey := fmt.Sprintf("%s%v", cacheApplePaySubscriptionIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, applePaySubscriptionRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.OriginalTransactionId, data.OriginalReceipt, data.Receipt, data.TransactionId, data.OrderId, data.Status, data.RenewStatus, data.OriginalProductId, data.ProductId, data.ExpiresDate, data.PurchaseDate, data.TransactionReason, data.TransactionType, data.InAppOwnershipType, data.Price, data.Currency, data.IsDeleted)
	}, applePaySubscriptionIdKey)
	return ret, err
}

func (m *defaultApplePaySubscriptionModel) Update(ctx context.Context, data *ApplePaySubscription) error {
	applePaySubscriptionIdKey := fmt.Sprintf("%s%v", cacheApplePaySubscriptionIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, applePaySubscriptionRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.OriginalTransactionId, data.OriginalReceipt, data.Receipt, data.TransactionId, data.OrderId, data.Status, data.RenewStatus, data.OriginalProductId, data.ProductId, data.ExpiresDate, data.PurchaseDate, data.TransactionReason, data.TransactionType, data.InAppOwnershipType, data.Price, data.Currency, data.IsDeleted, data.Id)
	}, applePaySubscriptionIdKey)
	return err
}

func (m *defaultApplePaySubscriptionModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheApplePaySubscriptionIdPrefix, primary)
}

func (m *defaultApplePaySubscriptionModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", applePaySubscriptionRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultApplePaySubscriptionModel) tableName() string {
	return m.table
}