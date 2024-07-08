package rights

import (
	"chat/service/chat/api/internal/service/rights"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRightsLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	service *rights.RightsService
}

func NewGetRightsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRightsLogic {
	return &GetRightsLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		service: rights.NewRightsService(ctx, svcCtx),
	}
}

func (l *GetRightsLogic) GetRights(req *types.GetRightsRequest) (resp *types.GetRightsResponse, err error) {
	resp = &types.GetRightsResponse{}
	res, err := l.service.GetRights(req.RightsId)
	if err != nil {
		return
	}
	if res == nil || res.Id <= 0 {
		return
	}
	resp.RightsId = res.Id
	resp.Status = res.Status
	return
}
