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

type RightsTimesRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

const (
	RightsTimesPeriodDay  = 1
	RightsTimesPeriodZero = 2
)

func NewRightsTimesRepository(ctx context.Context, svcCtx *svc.ServiceContext) *RightsTimesRepository {
	return &RightsTimesRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RightsTimesRepository) GetTimes(user, kfId string) (rightsTimesPo []*model.RightsTimes, err error) {

	// 先获取次卡数据
	rowBuilder := l.svcCtx.RightsTimesModel.RowBuilder().Where(squirrel.Eq{"user": user, "open_kf_id": kfId}).Where(squirrel.Eq{"pay_type": PayTypeTimes}).Where("remain > 0")
	rightsTimesPo, err = l.svcCtx.RightsTimesModel.FindAll(l.ctx, rowBuilder)
	if err != nil {
		return
	}
	if rightsTimesPo != nil && len(rightsTimesPo) > 0 {
		return
	}

	return
}

func (l *RightsTimesRepository) GetByUserAndKfId(user, kfId, now string) (rightsTimesPo []*model.RightsTimes, err error) {
	rowBuilder := l.svcCtx.RightsTimesModel.RowBuilder().Where(squirrel.Eq{"user": user, "open_kf_id": kfId}).
		Where("start <= ?", now).Where("end >= ?", now)

	return l.svcCtx.RightsTimesModel.FindAll(l.ctx, rowBuilder)

}

func (l *RightsTimesRepository) UpdateById(id int64, RightsTimes *model.RightsTimes) error {
	old, err := l.GetById(id)
	if err != nil {
		return err
	}
	if old == nil || old.Id <= 0 {
		return util.ReturnError(xerr.RecordNotFound)
	}
	l.svcCtx.RightsTimesModel.BuildFiled(old, RightsTimes)
	return l.svcCtx.RightsTimesModel.Update(context.Background(), RightsTimes)
}

func (l *RightsTimesRepository) GetById(id int64) (RightsTimes *model.RightsTimes, err error) {
	return l.svcCtx.RightsTimesModel.FindOne(context.Background(), id)
}

func (l *RightsTimesRepository) Insert(RightsTimes *model.RightsTimes) (lastId int64, err error) {
	re, err := l.svcCtx.RightsTimesModel.Insert(l.ctx, RightsTimes)
	if err != nil {
		return
	}
	lastId, err = re.LastInsertId()
	return
}

func (l *RightsTimesRepository) Update(old, RightsTimes *model.RightsTimes) error {
	l.svcCtx.RightsTimesModel.BuildFiled(old, RightsTimes)
	return l.svcCtx.RightsTimesModel.Update(l.ctx, RightsTimes)
}

func (l *RightsTimesRepository) UpdateTimesById(id, times int64) (err error) {
	return l.svcCtx.RightsTimesModel.UpdateTimesById(l.ctx, id, times)
}

func (l *RightsTimesRepository) GetByMessage(messageId string) (*model.RightsTimes, error) {
	return l.svcCtx.RightsTimesModel.FindOneByQuery(l.ctx,
		l.svcCtx.RightsTimesModel.RowBuilder().Where(squirrel.Eq{"message_id": messageId}),
	)
}

func (l *RightsTimesRepository) GetByOutTradeNo(OutTradeNo string) (*model.RightsTimes, error) {
	return l.svcCtx.RightsTimesModel.FindOneByQuery(l.ctx,
		l.svcCtx.RightsTimesModel.RowBuilder().Where(squirrel.Eq{"out_trade_no": OutTradeNo}),
	)
}

func (l *RightsTimesRepository) GetAll(user, kfId, now, updateAt string, remain bool, payType int64) (rightsTimesPo []*model.RightsTimes, err error) {

	rowBuilder := l.svcCtx.RightsTimesModel.RowBuilder()
	if user != "" {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"user": user})
	}
	if kfId != "" {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"open_kf_id": kfId})
	}

	if now != "" {
		rowBuilder = rowBuilder.Where("start <= ?", now).Where("end >= ?", now)
	}

	if remain {
		rowBuilder = rowBuilder.Where("remain > 0")
	} else {
		rowBuilder = rowBuilder.Where("remain = 0")
	}

	if payType != 0 {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"pay_type": payType})
	}

	if updateAt != "" {
		rowBuilder = rowBuilder.Where("updated_at <= ?", updateAt+" 23:59:59").Where("updated_at >= ?", updateAt+" 00:00:00")
	}
	return l.svcCtx.RightsTimesModel.FindAll(l.ctx, rowBuilder)

}

func (l *RightsTimesRepository) AddTimesByUser(user, kfId, now string, times int64) (err error) {
	return l.svcCtx.RightsTimesModel.AddTimesByUser(l.ctx, user, kfId, now, times)
}

func (l *RightsTimesRepository) DeleteByUser(user string) error {
	return l.svcCtx.RightsTimesModel.DeleteByUser(l.ctx, user, time.Now().Format(vars.TimeFormat))
}
