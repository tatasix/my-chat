package service

import (
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type FeedbackService struct {
	logx.Logger
	ctx                context.Context
	svcCtx             *svc.ServiceContext
	FeedbackRepository *repository.FeedbackRepository
}

func NewFeedbackService(ctx context.Context, svcCtx *svc.ServiceContext) *FeedbackService {
	return &FeedbackService{
		Logger:             logx.WithContext(ctx),
		ctx:                ctx,
		svcCtx:             svcCtx,
		FeedbackRepository: repository.NewFeedbackRepository(ctx, svcCtx),
	}
}

func (l *FeedbackService) GetFeedbackList(openKfId, user, startTime, endTime, order string, page, limit uint64, status int32) (FeedbackPos []*model.Feedback, count int64, err error) {
	return l.FeedbackRepository.GetAll(openKfId, user, startTime, endTime, order, page, limit, status)
}

func (l *FeedbackService) Insert(dto *model.Feedback) (err error) {
	_, err = l.FeedbackRepository.Insert(dto)
	return
}

func (l *FeedbackService) Idempotent(messageId, user string) (FeedbackPos *model.Feedback, err error) {
	return l.FeedbackRepository.GetByMessageAndCustomer(messageId, user)
}

func (l *FeedbackService) FindFeedback(id int64) (FeedbackPo *model.Feedback, err error) {
	return l.FeedbackRepository.GetById(id)
}
