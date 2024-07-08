package repository

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

const (
	// 1 次卡 2 月卡 3 年卡

	PayTypeFree  = 0
	PayTypeTimes = 1
	PayTypeMonth = 2
	PayTypeYear  = 3

	// 1支付宝 2微信

	PayMethodAlipay = 1
	PayMethodWechat = 2
	PayMethodApple  = 3

	// 1pc 2H5

	SourcePC          = 1
	SourceH5          = 2
	SourceMiniProgram = 3
	SourceAndroid     = 4
	SourceIos         = 5

	// 权益状态：1未付款 2付款成功 3权益已使用 4付款失败 5过期

	StatusNotPay  = 1
	StatusSuccess = 2
	StatusUsed    = 3
	StatusFail    = 4
	StatusExpire  = 5
)

type RightsRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRightsRepository(ctx context.Context, svcCtx *svc.ServiceContext) *RightsRepository {
	return &RightsRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RightsRepository) GetByUserAndKfId(user, kfId string, status []int64) (RightsPo []*model.Rights, err error) {
	rowBuilder := l.svcCtx.RightsModel.RowBuilder().Where(squirrel.Eq{"user": user}).Where(squirrel.Eq{"open_kf_id": kfId}).
		Where(squirrel.Eq{"status": status}).Where("end 	>= ?", time.Now())

	RightsPo, err = l.svcCtx.RightsModel.FindAll(l.ctx, rowBuilder)
	if err != nil {
		return
	}
	l.Logger.Info(" GetCustomerChatRecord ChatRecordRepository end ")
	return
}

func (l *RightsRepository) UpdateById(id int64, Rights *model.Rights) error {
	old, err := l.GetById(id)
	if err != nil {
		return err
	}
	if old == nil || old.Id <= 0 {
		return util.ReturnError(xerr.RecordNotFound)
	}
	l.svcCtx.RightsModel.BuildFiled(old, Rights)
	return l.svcCtx.RightsModel.Update(context.Background(), Rights)
}

func (l *RightsRepository) GetById(id int64) (rights *model.Rights, err error) {
	return l.svcCtx.RightsModel.FindOneByQuery(l.ctx,
		l.svcCtx.RightsModel.RowBuilder().Where(squirrel.Eq{"id": id}),
	)
}

func (l *RightsRepository) Insert(Rights *model.Rights) (lastId int64, err error) {
	re, err := l.svcCtx.RightsModel.Insert(l.ctx, Rights)
	if err != nil {
		return
	}
	lastId, err = re.LastInsertId()
	return
}

func (l *RightsRepository) Update(old, Rights *model.Rights) error {
	l.svcCtx.RightsModel.BuildFiled(old, Rights)
	return l.svcCtx.RightsModel.Update(l.ctx, Rights)
}

func (l *RightsRepository) UpdateStatusById(id, status int64, reason string) (err error) {
	return l.svcCtx.RightsModel.UpdateStatusById(l.ctx, id, status, reason)
}

func (l *RightsRepository) GetByMessage(messageId string) (*model.Rights, error) {
	return l.svcCtx.RightsModel.FindOneByQuery(l.ctx,
		l.svcCtx.RightsModel.RowBuilder().Where(squirrel.Eq{"message_id": messageId}),
	)
}

func (l *RightsRepository) GetByOutTradeNo(OutTradeNo string) (*model.Rights, error) {
	return l.svcCtx.RightsModel.FindOneByQuery(l.ctx,
		l.svcCtx.RightsModel.RowBuilder().Where(squirrel.Eq{"out_trade_no": OutTradeNo}),
	)
}

func (l *RightsRepository) UpdateStatus() error {
	return l.svcCtx.RightsModel.UpdateStatus(l.ctx, time.Now().Format(vars.TimeFormat), StatusExpire)
}

func (l *RightsRepository) GetAll(user, kfId string, status, payType []int64, canUse bool) (RightsPo []*model.Rights, err error) {
	rowBuilder := l.svcCtx.RightsModel.RowBuilder()
	if user != "" {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"user": user})
	}

	if kfId != "" {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"open_kf_id": kfId})
	}

	if len(status) >= 0 {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"status": status})
	}

	if len(payType) >= 0 {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"pay_type": payType})
	}

	if canUse {
		rowBuilder = rowBuilder.Where("end >= ?", time.Now())
	}

	RightsPo, err = l.svcCtx.RightsModel.FindAll(l.ctx, rowBuilder)
	if err != nil {
		return
	}
	return
}

func (l *RightsRepository) DeleteByUser(user string) error {
	return l.svcCtx.RightsModel.DeleteByUser(l.ctx, user)
}

func (l *RightsRepository) GetLastByAppleTxnId(appleTxnId string) (rightsPo *model.Rights, err error) {
	pos, err := l.svcCtx.RightsModel.FindAll(l.ctx,
		l.svcCtx.RightsModel.RowBuilder().Where(squirrel.Eq{"apple_txn_id": appleTxnId}).OrderBy("id desc").Limit(1),
	)

	if err != nil {
		return
	}

	if pos != nil && len(pos) > 0 {
		return pos[0], nil
	}
	return
}
