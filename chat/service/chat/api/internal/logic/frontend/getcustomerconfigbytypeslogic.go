package frontend

import (
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/service"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCustomerConfigByTypesLogic struct {
	logx.Logger
	ctx                   context.Context
	svcCtx                *svc.ServiceContext
	customerConfigService *service.CustomerConfigService
}

func NewGetCustomerConfigByTypesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCustomerConfigByTypesLogic {
	return &GetCustomerConfigByTypesLogic{
		Logger:                logx.WithContext(ctx),
		ctx:                   ctx,
		svcCtx:                svcCtx,
		customerConfigService: service.NewCustomerConfigService(ctx, svcCtx),
	}
}

func (l *GetCustomerConfigByTypesLogic) GetCustomerConfigByTypes(req *types.GetCustomerConfigByTypesRequest) (resp *types.GetCustomerConfigByTypeResponse, err error) {
	resp = &types.GetCustomerConfigByTypeResponse{}
	customer, err := l.customerConfigService.GetByTypes(req.Type, req.Name)
	if nil != err {
		return
	}
	resp.List = assembler.POTODTOGetCustomerList(customer)

	return
}