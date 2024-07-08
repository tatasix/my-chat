package service

import (
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityLogService struct {
	logx.Logger
	ctx                   context.Context
	svcCtx                *svc.ServiceContext
	ActivityLogRepository *repository.ActivityLogRepository
}

func NewActivityLogService(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityLogService {
	return &ActivityLogService{
		Logger:                logx.WithContext(ctx),
		ctx:                   ctx,
		svcCtx:                svcCtx,
		ActivityLogRepository: repository.NewActivityLogRepository(ctx, svcCtx),
	}
}

func (l *ActivityLogService) UpdateById(id int64, model *model.ActivityLog) (err error) {
	return l.ActivityLogRepository.UpdateById(id, model)
}

func (l *ActivityLogService) Update(old, model *model.ActivityLog) (err error) {
	return l.ActivityLogRepository.Update(old, model)
}

func (l *ActivityLogService) Insert(dto *model.ActivityLog) (err error) {

	return l.ActivityLogRepository.Insert(dto)
}

func (l *ActivityLogService) UpdateUser(new, old string) (err error) {
	return l.ActivityLogRepository.UpdateUser(new, old)
}
