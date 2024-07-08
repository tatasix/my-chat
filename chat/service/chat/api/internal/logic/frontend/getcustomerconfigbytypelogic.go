package frontend

import (
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/service"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCustomerConfigByTypeLogic struct {
	logx.Logger
	ctx                   context.Context
	svcCtx                *svc.ServiceContext
	customerConfigService *service.CustomerConfigService
}

func NewGetCustomerConfigByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCustomerConfigByTypeLogic {
	return &GetCustomerConfigByTypeLogic{
		Logger:                logx.WithContext(ctx),
		ctx:                   ctx,
		svcCtx:                svcCtx,
		customerConfigService: service.NewCustomerConfigService(ctx, svcCtx),
	}
}

func (l *GetCustomerConfigByTypeLogic) GetCustomerConfigByType(req *types.GetCustomerConfigByTypeRequest) (resp *types.GetCustomerConfigByTypeResponse, err error) {
	resp = &types.GetCustomerConfigByTypeResponse{}
	customer, err := l.customerConfigService.GetByTypes([]int64{req.Type}, req.Name)
	if nil != err {
		return
	}
	resp.List = assembler.POTODTOGetCustomerList(customer)

	return
}
