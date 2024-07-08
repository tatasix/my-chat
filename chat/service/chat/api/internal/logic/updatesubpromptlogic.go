package logic

import (
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/model"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSubPromptLogic struct {
	logx.Logger
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	promptRepository *repository.PromptRepository
}

func NewUpdateSubPromptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubPromptLogic {
	return &UpdateSubPromptLogic{
		Logger:           logx.WithContext(ctx),
		ctx:              ctx,
		svcCtx:           svcCtx,
		promptRepository: repository.NewPromptRepository(ctx, svcCtx),
	}
}

func (l *UpdateSubPromptLogic) UpdateSubPrompt(req *types.UpdateSubPromptRequest) (err error) {
	prompt, err := l.promptRepository.GetById(req.SubPromptId)

	if err != nil || prompt == nil || prompt.Id <= 0 {
		return errors.New("prompt 不存在")
	}

	l.promptRepository.Update(
		req.SubPromptId,
		&model.Prompt{
			StateId:   prompt.StateId,
			KfId:      prompt.KfId,
			Title:     req.Title,
			Prompt:    req.Prompt,
			UpdatedBy: prompt.UpdatedBy,
			CreatedBy: prompt.CreatedBy,
		})
	return
}
