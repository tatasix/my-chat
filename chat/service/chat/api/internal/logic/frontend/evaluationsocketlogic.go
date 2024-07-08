package frontend

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/service/customer"
	"chat/service/chat/api/internal/util/sensitive"
	"chat/service/chat/model"
	"context"
	"encoding/json"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EvaluationSocketLogic struct {
	logx.Logger
	ctx                   context.Context
	svcCtx                *svc.ServiceContext
	customerConfigService *service.CustomerConfigService
	stateService          *service.StateService
}

func NewEvaluationSocketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EvaluationSocketLogic {
	return &EvaluationSocketLogic{
		Logger:                logx.WithContext(ctx),
		ctx:                   ctx,
		svcCtx:                svcCtx,
		customerConfigService: service.NewCustomerConfigService(ctx, svcCtx),
		stateService:          service.NewStateService(ctx, svcCtx),
	}
}

func (l *EvaluationSocketLogic) EvaluationSocket(req *types.EvaluationSocketRequest, channel chan string, baseChannel chan any) (resp *types.EvaluationSocketResponse, err error) {
	resp = &types.EvaluationSocketResponse{}
	if req.Message != "" {
		if !sensitive.NewSensitive(l.ctx, l.svcCtx).Check(req.Message) {
			response := &types.EvaluationSocketResponse{
				Message:      sensitive.ErrorMessage,
				MessageId:    req.MessageID,
				RelationId:   req.RelationId,
				IsEnd:        true,
				ErrorMessage: sensitive.ErrorMessage,
			}
			l.Logger.Errorf("SensitiveError user:%s message:%s", req.User, req.Message)
			return response, nil
		}
	}
	//获取客服相关信息
	customerInfo, err := l.customerConfigService.GetByKfId(req.OpenKfID)
	if nil != err {
		return
	}
	if customerInfo == nil || customerInfo.Id <= 0 {
		err = util.ReturnError(xerr.ChatCustomerNotExist)
		return
	}
	enable, _, err := service.NewRiskService(l.ctx, l.svcCtx).Check(req.User, req.OpenKfID)
	if err != nil {
		return
	}
	if enable == 2 {
		resp.Code = xerr.RightsNotHaveTimesError
		return
	} else if enable == 3 {
		resp.Code = xerr.RightsNotVip
		return
	}
	handler := customer.NewCustomerHandler(l.ctx, l.svcCtx, customerInfo.KfId)
	if handler == nil {
		err = util.ReturnError(xerr.SystemError)
		return
	}
	var promptStates map[int32]model.PromptStates
	if customerInfo.PromptStates == "" {
		err = util.ReturnError(xerr.ChatNodeNotExist)
		return
	}
	err = json.Unmarshal([]byte(customerInfo.PromptStates), &promptStates)
	if err != nil {
		err = util.ReturnError(xerr.SystemError)
		return
	}
	//获取当前的用户的状态
	stateInfo, stateErr := l.stateService.GetByUserAndKfId(req.User, req.OpenKfID)
	if stateErr != nil {
		err = stateErr
		return
	}
	var stateId int32
	if stateInfo == nil || stateInfo.State <= 0 {
		stateId = promptStates[0].NextNode[0]
	} else {
		if states, ok := promptStates[int32(stateInfo.State)]; ok {
			if len(states.NextNode) > 1 {
				for _, num := range states.NextNode {
					if num == int32(req.Status) {
						stateId = num
						break
					}
				}

			} else {
				stateId = states.NextNode[0]
			}
		}
	}
	if stateId == -1 {
		err = util.ReturnError(xerr.ChatTimesEmpty)
		return
	}
	if stateId == 0 {
		err = util.ReturnError(xerr.SystemError)
		return
	}

	return handler.Handle(stateId, req, customerInfo, stateInfo, channel, baseChannel)
}
