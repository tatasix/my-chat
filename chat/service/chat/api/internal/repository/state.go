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

type StateRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStateRepository(ctx context.Context, svcCtx *svc.ServiceContext) *StateRepository {
	return &StateRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StateRepository) GetByUserAndKfId(user, kfId string) (StatePo *model.State, err error) {
	return l.svcCtx.StateModel.FindOneByQuery(l.ctx,
		l.svcCtx.StateModel.RowBuilder().Where(squirrel.Eq{"user": user}).Where(squirrel.Eq{"kf_id": kfId}).Where(squirrel.Eq{"date": time.Now().Format("2006-01-02")}),
	)
}

func (l *StateRepository) UpdateById(id int64, State *model.State) error {
	old, err := l.GetById(id)
	if err != nil {
		return err
	}
	if old == nil || old.Id <= 0 {
		return util.ReturnError(xerr.RecordNotFound)
	}
	l.svcCtx.StateModel.BuildFiled(old, State)
	return l.svcCtx.StateModel.Update(context.Background(), State)
}

func (l *StateRepository) GetById(id int64) (State *model.State, err error) {
	return l.svcCtx.StateModel.FindOne(context.Background(), id)
}

func (l *StateRepository) Insert(State *model.State) (err error) {
	l.svcCtx.StateModel.BuildFiled(nil, State)
	_, err = l.svcCtx.StateModel.Insert(l.ctx, State)
	return
}

func (l *StateRepository) Update(old, State *model.State) error {
	l.svcCtx.StateModel.BuildFiled(old, State)
	return l.svcCtx.StateModel.Update(l.ctx, State)
}

func (l *StateRepository) UpdateStatus(user, kfId string, status int64) (err error) {
	old, err := l.GetByUserAndKfId(user, kfId)
	if err != nil {
		return err
	}
	if old == nil || old.Id <= 0 {
		return util.ReturnError(xerr.RecordNotFound)
	}
	old.State = status
	old.UpdatedAt = time.Now()
	return l.svcCtx.StateModel.Update(context.Background(), old)
}
