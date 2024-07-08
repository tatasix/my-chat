package repository

import (
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type ApplePaySubscriptionRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplePaySubscriptionRepository(ctx context.Context, svcCtx *svc.ServiceContext) *ApplePaySubscriptionRepository {
	return &ApplePaySubscriptionRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplePaySubscriptionRepository) GetById(id int64) (applePaySubscription *model.ApplePaySubscription, err error) {
	return l.svcCtx.ApplePaySubscriptionModel.FindOne(l.ctx, id)
}

func (l *ApplePaySubscriptionRepository) GetByOriTxnId(txnId string) (applePaySubscription *model.ApplePaySubscription, err error) {
	return l.svcCtx.ApplePaySubscriptionModel.FindOneByQuery(l.ctx,
		l.svcCtx.ApplePaySubscriptionModel.RowBuilder().Where(
			squirrel.Eq{"original_transaction_id": txnId},
		),
	)
}

func (l *ApplePaySubscriptionRepository) GetByStatusAndExpiresDate(status int64, expiresDate string) (applePaySubscription []*model.ApplePaySubscription, err error) {
	return l.svcCtx.ApplePaySubscriptionModel.FindAll(l.ctx,
		l.svcCtx.ApplePaySubscriptionModel.RowBuilder().Where(
			squirrel.Eq{"status": status},
		).Where("expires_date <= ?", expiresDate),
	)
}

func (l *ApplePaySubscriptionRepository) Insert(applePaySubscription *model.ApplePaySubscription) (lastId int64, err error) {
	l.svcCtx.ApplePaySubscriptionModel.BuildFiled(nil, applePaySubscription)
	res, err := l.svcCtx.ApplePaySubscriptionModel.Insert(l.ctx, applePaySubscription)
	lastId, err = res.LastInsertId()
	return
}

func (l *ApplePaySubscriptionRepository) Update(id int64, applePaySubscription *model.ApplePaySubscription) error {
	old, err := l.GetById(id)
	if err != nil {
		return err
	}
	if old == nil || old.Id <= 0 {
		return errors.New("record not find")
	}
	l.svcCtx.ApplePaySubscriptionModel.BuildFiled(old, applePaySubscription)
	return l.svcCtx.ApplePaySubscriptionModel.Update(l.ctx, applePaySubscription)
}

func (l *ApplePaySubscriptionRepository) GetByIds(kfId []int64) (po []*model.ApplePaySubscription, err error) {
	return l.svcCtx.ApplePaySubscriptionModel.FindAll(l.ctx,
		l.svcCtx.ApplePaySubscriptionModel.RowBuilder().Where(squirrel.Eq{"agent_id": kfId}),
	)
}

func (l *ApplePaySubscriptionRepository) Delete(id int64) error {
	return l.svcCtx.ApplePaySubscriptionModel.Delete(l.ctx, id)
}
