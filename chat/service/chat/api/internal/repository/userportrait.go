package repository

import (
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	UserTypeWechat1  = 1
	UserTypeVisitor1 = 2
)

type UserPortraitRepository struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserPortraitRepository(ctx context.Context, svcCtx *svc.ServiceContext) *UserPortraitRepository {
	return &UserPortraitRepository{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserPortraitRepository) GetByName(nickname string) (customerPo *model.UserPortrait, err error) {

	customerPo, err = l.svcCtx.UserPortraitModel.FindOneByQuery(context.Background(),
		l.svcCtx.UserPortraitModel.RowBuilder().Where(squirrel.Eq{"nickname": nickname}),
	)
	return
}

func (l *UserPortraitRepository) GetByUsers(users []string) (CustomerConfigPo []*model.UserPortrait, err error) {
	return l.svcCtx.UserPortraitModel.FindAll(context.Background(),
		l.svcCtx.UserPortraitModel.RowBuilder().Where(squirrel.Eq{"user": users}),
	)
}

func (l *UserPortraitRepository) GetByUser(user string) (CustomerConfigPo *model.UserPortrait, err error) {
	return l.svcCtx.UserPortraitModel.FindOneByQuery(context.Background(),
		l.svcCtx.UserPortraitModel.RowBuilder().Where(squirrel.Eq{"user": user}),
	)
}

func (l *UserPortraitRepository) Update(id int64, UserPortrait *model.UserPortrait) error {
	old, err := l.GetById(id)
	if err != nil {
		return err
	}
	if old == nil || old.Id <= 0 {
		return errors.New("record not find")
	}
	l.svcCtx.UserPortraitModel.BuildFiled(old, UserPortrait)
	return l.svcCtx.UserPortraitModel.Update(context.Background(), UserPortrait)
}

func (l *UserPortraitRepository) UpdateByUser(user string, UserPortrait *model.UserPortrait) error {
	old, err := l.GetByUser(user)
	if err != nil {
		return err
	}
	if old == nil || old.Id <= 0 {
		return errors.New("record not find")
	}
	l.svcCtx.UserPortraitModel.BuildFiled(old, UserPortrait)
	return l.svcCtx.UserPortraitModel.Update(context.Background(), UserPortrait)
}

func (l *UserPortraitRepository) GetById(id int64) (UserPortrait *model.UserPortrait, err error) {
	return l.svcCtx.UserPortraitModel.FindOne(context.Background(), id)
}
