package frontend

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/service/customer"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetEvaluationLogic struct {
	logx.Logger
	ctx                   context.Context
	svcCtx                *svc.ServiceContext
	customerConfigService *service.CustomerConfigService
	stateService          *service.StateService
}

func NewGetEvaluationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEvaluationLogic {
	return &GetEvaluationLogic{
		Logger:                logx.WithContext(ctx),
		ctx:                   ctx,
		svcCtx:                svcCtx,
		customerConfigService: service.NewCustomerConfigService(ctx, svcCtx),
		stateService:          service.NewStateService(ctx, svcCtx),
	}
}

func (l *GetEvaluationLogic) GetEvaluation(req *types.GetEvaluationRequest) (resp *types.GetEvaluationResponse, err error) {
	//获取客服相关信息
	customerInfo, err := l.customerConfigService.GetByKfId(req.OpenKfID)
	if nil != err {
		return
	}
	if customerInfo == nil || customerInfo.Id <= 0 {
		err = util.ReturnError(xerr.ChatCustomerNotExist)
		return
	}
	handler := customer.NewCustomerHandler(l.ctx, l.svcCtx, customerInfo.KfId)
	if handler == nil {
		err = util.ReturnError(xerr.SystemError)
		return
	}

	//获取当前的用户的状态
	stateInfo, err := l.stateService.GetByUserAndKfId(req.User, req.OpenKfID)
	if err != nil {
		return
	}
	if stateInfo == nil || stateInfo.State <= 0 {
		resp = &types.GetEvaluationResponse{
			Status: 0,
			IsEnd:  false,
		}
		return
	}

	return handler.HandleGet(int32(stateInfo.State), req)
}
