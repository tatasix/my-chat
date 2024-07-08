package rights

import (
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/repository"
	"context"
	"fmt"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPayInfoLogic struct {
	logx.Logger
	ctx                      context.Context
	svcCtx                   *svc.ServiceContext
	customerConfigRepository *repository.CustomerConfigRepository
}

func NewGetPayInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPayInfoLogic {
	return &GetPayInfoLogic{
		Logger:                   logx.WithContext(ctx),
		ctx:                      ctx,
		svcCtx:                   svcCtx,
		customerConfigRepository: repository.NewCustomerConfigRepository(ctx, svcCtx),
	}
}

func (l *GetPayInfoLogic) GetPayInfo(req *types.GetPayInfoRequest) (resp *types.GetPayInfoResponse, err error) {

	customerConfigPos, err := l.customerConfigRepository.GetAllPayConfig(req.OpenKfID)
	if err != nil {
		fmt.Printf("GetSystemConfig error: %v", err)
		return
	}
	if len(customerConfigPos) <= 0 {
		return &types.GetPayInfoResponse{}, nil
	}

	return &types.GetPayInfoResponse{
		List: assembler.POTODTOGetPayConfig(customerConfigPos),
	}, nil
}
