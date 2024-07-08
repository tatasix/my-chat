package logic

import (
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"context"
	"database/sql"
	"github.com/cockroachdb/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCustomerConfigLogic struct {
	logx.Logger
	ctx                      context.Context
	svcCtx                   *svc.ServiceContext
	customerConfigRepository *repository.CustomerConfigRepository
}

func NewUpdateCustomerConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCustomerConfigLogic {
	return &UpdateCustomerConfigLogic{
		Logger:                   logx.WithContext(ctx),
		ctx:                      ctx,
		svcCtx:                   svcCtx,
		customerConfigRepository: repository.NewCustomerConfigRepository(ctx, svcCtx),
	}
}

func (l *UpdateCustomerConfigLogic) UpdateCustomerConfig(req *types.CustomerConfig) (resp *types.Response, err error) {
	if req == nil || req.Id <= 0 {
		return nil, errors.New("缺少必传参数")
	}

	var score sql.NullFloat64
	if req.Score > 0 {
		score.Valid = true
		score.Float64 = req.Score
	}

	customerConfigPo, err := l.customerConfigRepository.GetById(req.Id)

	if err != nil || customerConfigPo == nil || customerConfigPo.Id <= 0 {
		return nil, errors.New("找不到对应数据")
	}

	customerConfigPo.KfName = req.KfName
	customerConfigPo.Prompt = req.Prompt
	customerConfigPo.PostModel = req.PostModel
	customerConfigPo.EmbeddingEnable = req.EmbeddingEnable
	customerConfigPo.EmbeddingMode = req.EmbeddingMode
	customerConfigPo.Score = score
	customerConfigPo.TopK = req.TopK
	customerConfigPo.ClearContextTime = req.ClearContextTime

	err = l.customerConfigRepository.Update(req.Id, customerConfigPo)
	if err != nil {
		return
	}
	return &types.Response{
		Message: "ok",
	}, nil

}
