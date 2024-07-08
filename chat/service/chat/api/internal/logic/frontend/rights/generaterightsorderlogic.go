package rights

import (
	"chat/service/chat/api/internal/service/rights"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateRightsOrderLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	service *rights.RightsService
}

func NewGenerateRightsOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateRightsOrderLogic {
	return &GenerateRightsOrderLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		service: rights.NewRightsService(ctx, svcCtx),
	}
}

func (l *GenerateRightsOrderLogic) GenerateRightsOrder(req *types.GenerateRightsOrderRequest) (resp *types.GenerateRightsOrderResponse, err error) {
	lastId, res, err := l.service.Pay(req.User, req.OpenKfID, req.MessageID, req.ReturnUrl, req.QuitUrl, req.PayMethod, req.PayType, req.Source)
	if nil != err {
		return
	}
	resp = &types.GenerateRightsOrderResponse{
		Url:       res,
		PayType:   req.PayType,
		PayMethod: req.PayMethod,
		Source:    req.Source,
		RightsId:  lastId,
	}
	return
}
