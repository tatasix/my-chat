package logic

import (
	"chat/service/chat/api/internal/repository"
	"context"
	"errors"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCustomerConfigLogic struct {
	logx.Logger
	ctx                      context.Context
	svcCtx                   *svc.ServiceContext
	customerConfigRepository *repository.CustomerConfigRepository
}

func NewDeleteCustomerConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCustomerConfigLogic {
	return &DeleteCustomerConfigLogic{
		Logger:                   logx.WithContext(ctx),
		ctx:                      ctx,
		svcCtx:                   svcCtx,
		customerConfigRepository: repository.NewCustomerConfigRepository(ctx, svcCtx),
	}
}

func (l *DeleteCustomerConfigLogic) DeleteCustomerConfig(req *types.FindCustomerConfigRequest) (resp *types.Response, err error) {
	po, err := l.customerConfigRepository.GetById(req.Id)

	if err != nil || po == nil || po.Id <= 0 {
		return nil, errors.New("数据不存在")
	}
	po.IsDeleted = 1
	err = l.customerConfigRepository.Update(req.Id, po)
	if err != nil {
		return
	}
	return &types.Response{
		Message: "ok",
	}, nil
}
