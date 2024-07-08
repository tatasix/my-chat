package service

import (
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type StateService struct {
	logx.Logger
	ctx             context.Context
	svcCtx          *svc.ServiceContext
	StateRepository *repository.StateRepository
}

func NewStateService(ctx context.Context, svcCtx *svc.ServiceContext) *StateService {
	return &StateService{
		Logger:          logx.WithContext(ctx),
		ctx:             ctx,
		svcCtx:          svcCtx,
		StateRepository: repository.NewStateRepository(ctx, svcCtx),
	}
}

func (l *StateService) GetByUserAndKfId(user, kfId string) (res *model.State, err error) {
	//return l.StateRepository.GetByUserAndKfId
	res, err = l.StateRepository.GetByUserAndKfId(user, kfId)
	return
}

func (l *StateService) Insert(po *model.State) (err error) {
	return l.StateRepository.Insert(po)
}

func (l *StateService) UpdateStatus(user, kfId string, status int64) (err error) {
	return l.StateRepository.UpdateStatus(user, kfId, status)
}
