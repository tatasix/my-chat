package repository

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type RightsRecordRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRightsRecordRepository(ctx context.Context, svcCtx *svc.ServiceContext) *RightsRecordRepository {
	return &RightsRecordRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RightsRecordRepository) GetByUserAndKfId(user, kfId string, status []int64) (RightsRecordPo []*model.RightsRecord, err error) {
	rowBuilder := l.svcCtx.RightsRecordModel.RowBuilder().Where(squirrel.Eq{"user": user}).Where(squirrel.Eq{"kf_id": kfId}).
		Where(squirrel.Eq{"status": status}).Where("end <= ?", time.Now())

	RightsRecordPo, err = l.svcCtx.RightsRecordModel.FindAll(l.ctx, rowBuilder)
	if err != nil {
		return
	}
	l.Logger.Info(" GetCustomerChatRecord ChatRecordRepository end ")
	return
}

func (l *RightsRecordRepository) UpdateById(id int64, RightsRecord *model.RightsRecord) error {
	old, err := l.GetById(id)
	if err != nil {
		return err
	}
	if old == nil || old.Id <= 0 {
		return util.ReturnError(xerr.RecordNotFound)
	}
	l.svcCtx.RightsRecordModel.BuildFiled(old, RightsRecord)
	return l.svcCtx.RightsRecordModel.Update(context.Background(), RightsRecord)
}

func (l *RightsRecordRepository) GetById(id int64) (RightsRecord *model.RightsRecord, err error) {
	return l.svcCtx.RightsRecordModel.FindOne(context.Background(), id)
}

func (l *RightsRecordRepository) Insert(RightsRecord *model.RightsRecord) (lastId int64, err error) {
	re, err := l.svcCtx.RightsRecordModel.Insert(l.ctx, RightsRecord)
	if err != nil {
		return
	}
	lastId, err = re.LastInsertId()
	return
}

func (l *RightsRecordRepository) Update(old, RightsRecord *model.RightsRecord) error {
	l.svcCtx.RightsRecordModel.BuildFiled(old, RightsRecord)
	return l.svcCtx.RightsRecordModel.Update(l.ctx, RightsRecord)
}

func (l *RightsRecordRepository) UpdateStatusById(id, status int64, reason string) (err error) {
	return l.svcCtx.RightsRecordModel.UpdateStatusById(l.ctx, id, status, reason)
}

func (l *RightsRecordRepository) GetByMessage(messageId string) (*model.RightsRecord, error) {
	return l.svcCtx.RightsRecordModel.FindOneByQuery(l.ctx,
		l.svcCtx.RightsRecordModel.RowBuilder().Where(squirrel.Eq{"message_id": messageId}),
	)
}

func (l *RightsRecordRepository) GetByOutTradeNo(OutTradeNo string) (*model.RightsRecord, error) {
	return l.svcCtx.RightsRecordModel.FindOneByQuery(l.ctx,
		l.svcCtx.RightsRecordModel.RowBuilder().Where(squirrel.Eq{"out_trade_no": OutTradeNo}),
	)
}
