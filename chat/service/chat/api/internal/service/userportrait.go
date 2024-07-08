package service

import (
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserPortraitService struct {
	logx.Logger
	ctx                    context.Context
	svcCtx                 *svc.ServiceContext
	UserPortraitRepository *repository.UserPortraitRepository
}

func NewUserPortraitService(ctx context.Context, svcCtx *svc.ServiceContext) *UserPortraitService {
	return &UserPortraitService{
		Logger:                 logx.WithContext(ctx),
		ctx:                    ctx,
		svcCtx:                 svcCtx,
		UserPortraitRepository: repository.NewUserPortraitRepository(ctx, svcCtx),
	}
}

func (l *UserPortraitService) Update(id int64, model *model.UserPortrait) (err error) {
	return l.UserPortraitRepository.Update(id, model)
}

func (l *UserPortraitService) UpdateByUser(user string, model *model.UserPortrait) (err error) {
	return l.UserPortraitRepository.UpdateByUser(user, model)
}
