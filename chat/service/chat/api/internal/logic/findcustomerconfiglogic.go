package logic

import (
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"context"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCustomerConfigLogic struct {
	logx.Logger
	ctx                      context.Context
	svcCtx                   *svc.ServiceContext
	customerConfigRepository *repository.CustomerConfigRepository
	promptRepository         *repository.PromptRepository
}

func NewFindCustomerConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCustomerConfigLogic {
	return &FindCustomerConfigLogic{
		Logger:                   logx.WithContext(ctx),
		ctx:                      ctx,
		svcCtx:                   svcCtx,
		customerConfigRepository: repository.NewCustomerConfigRepository(ctx, svcCtx),
		promptRepository:         repository.NewPromptRepository(ctx, svcCtx),
	}
}

func (l *FindCustomerConfigLogic) FindCustomerConfig(req *types.FindCustomerConfigRequest) (resp *types.FindCustomerConfigResult, err error) {
	customerConfigPo, err := l.customerConfigRepository.GetById(req.Id)
	if err != nil {
		fmt.Printf("GetSystemConfig error: %v", err)
		return
	}

	if customerConfigPo == nil && customerConfigPo.Id <= 0 {
		return nil, errors.New("找不到对应数据")
	}
	subConfigPo, err := l.promptRepository.GetByKfId(customerConfigPo.KfId)

	dto := assembler.POTODTOFindCustomerConfig(customerConfigPo, subConfigPo)

	return &dto, err
}
