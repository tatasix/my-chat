package service

import (
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type PromptService struct {
	logx.Logger
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	PromptRepository *repository.PromptRepository
}

func NewPromptService(ctx context.Context, svcCtx *svc.ServiceContext) *PromptService {
	return &PromptService{
		Logger:           logx.WithContext(ctx),
		ctx:              ctx,
		svcCtx:           svcCtx,
		PromptRepository: repository.NewPromptRepository(ctx, svcCtx),
	}
}

func (l *PromptService) GetByKfIdAndState(kfId string, state int32) (res *model.Prompt, err error) {
	return l.PromptRepository.GetByKfIdAndState(kfId, state)
}

func (l *PromptService) GetByKfIdAndStateUseCache(kfId string, state int32) (res *model.Prompt, err error) {
	return l.PromptRepository.GetByKfIdAndStateUseCache(kfId, state)
}
