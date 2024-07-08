package service

import (
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

type CustomerConfigService struct {
	logx.Logger
	ctx                      context.Context
	svcCtx                   *svc.ServiceContext
	CustomerConfigRepository *repository.CustomerConfigRepository
	StateRepository          *repository.StateRepository
	PromptRepository         *repository.PromptRepository
}

func NewCustomerConfigService(ctx context.Context, svcCtx *svc.ServiceContext) *CustomerConfigService {
	return &CustomerConfigService{
		Logger:                   logx.WithContext(ctx),
		ctx:                      ctx,
		svcCtx:                   svcCtx,
		CustomerConfigRepository: repository.NewCustomerConfigRepository(ctx, svcCtx),
		StateRepository:          repository.NewStateRepository(ctx, svcCtx),
		PromptRepository:         repository.NewPromptRepository(ctx, svcCtx),
	}
}

func (l *CustomerConfigService) GetPrompt(kfId, user string) (CustomerConfigPos *model.CustomerConfig, err error) {
	CustomerConfigPos, err = l.CustomerConfigRepository.GetByKfIdUseCache(kfId)
	if nil != err {
		return
	}
	if CustomerConfigPos == nil || CustomerConfigPos.Id <= 0 || CustomerConfigPos.Multiple == 0 {
		return
	}
	if CustomerConfigPos.PromptStates != "" {
		var promptStates map[int32]model.PromptStates
		err = json.Unmarshal([]byte(CustomerConfigPos.PromptStates), &promptStates)
		if err != nil {
			err = nil
			CustomerConfigPos.Prompt = ""
			return
		}
		//获取当前所在的节点
		stateInfo, stateErr := l.StateRepository.GetByUserAndKfId(user, kfId)
		if stateErr != nil {
			err = stateErr
			CustomerConfigPos.Prompt = ""
			return
		}
		var stateId int32
		if stateInfo == nil || stateInfo.Id <= 0 {
			stateId = promptStates[0].NextNode[0]
		} else {
			if states, ok := promptStates[int32(stateInfo.Id)]; ok {
				stateId = states.NextNode[0]
			}
		}
		if stateId == 0 {
			CustomerConfigPos.Prompt = ""
			return
		}
		//获取状态对应多prompt
		promptInfo, promptErr := l.PromptRepository.GetByKfIdAndState(kfId, stateId)
		if promptErr != nil {
			err = promptErr
			CustomerConfigPos.Prompt = ""
			return
		}
		if promptInfo != nil && promptInfo.Id >= 0 && promptInfo.Prompt != "" {
			CustomerConfigPos.Prompt = promptInfo.Prompt
		}
	}
	return
}

func (l *CustomerConfigService) GetByKfId(kfId string) (CustomerConfigPos *model.CustomerConfig, err error) {
	CustomerConfigPos, err = l.CustomerConfigRepository.GetByKfIdUseCache(kfId)
	if nil != err {
		return
	}
	if CustomerConfigPos == nil || CustomerConfigPos.Id <= 0 || CustomerConfigPos.Multiple == 0 {
		return
	}
	return
}

func (l *CustomerConfigService) GetByTypes(CustomerType []int64, name string) (CustomerConfigPos []*model.CustomerConfig, err error) {
	return l.CustomerConfigRepository.GetByTypes(CustomerType, name)
}
