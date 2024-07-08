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

type CreateSubPromptLogic struct {
	logx.Logger
	ctx                      context.Context
	svcCtx                   *svc.ServiceContext
	promptRepository         *repository.PromptRepository
	customerConfigRepository *repository.CustomerConfigRepository
}

func NewCreateSubPromptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSubPromptLogic {
	return &CreateSubPromptLogic{
		Logger:                   logx.WithContext(ctx),
		ctx:                      ctx,
		svcCtx:                   svcCtx,
		promptRepository:         repository.NewPromptRepository(ctx, svcCtx),
		customerConfigRepository: repository.NewCustomerConfigRepository(ctx, svcCtx),
	}
}

func (l *CreateSubPromptLogic) CreateSubPrompt(req *types.CreateSubPromptRequest) (resp *types.Response, err error) {

	customerPo, err := l.customerConfigRepository.GetByKfId(req.KfId)

	if err != nil {
		return
	}

	if customerPo == nil || customerPo.Id <= 0 {
		return nil, errors.New("客服不存在")
	}

	count, err := l.promptRepository.GetCountByKfId(req.KfId)

	l.promptRepository.Insert(&model.Prompt{
		StateId: count + 1,
		KfId:    req.KfId,
		Title:   req.Title,
		Prompt:  req.Prompt,
	})

	//更新 multiple =1
	customerPo.Multiple = 1
	l.customerConfigRepository.Update(customerPo.Id, customerPo)

	return
}
