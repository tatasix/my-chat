package frontend

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type MobileIsExistsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	*service.WechatUserService
}

func NewMobileIsExistsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MobileIsExistsLogic {
	return &MobileIsExistsLogic{
		Logger:            logx.WithContext(ctx),
		ctx:               ctx,
		svcCtx:            svcCtx,
		WechatUserService: service.NewWechatUserService(ctx, svcCtx),
	}
}

func (l *MobileIsExistsLogic) MobileIsExists(req *types.MobileIsExistsRequest) (resp *types.MobileIsExistsResponse, err error) {
	if !util.CheckMobile(req.Mobile) {
		err = util.ReturnError(xerr.LoginMobileError)
		return
	}
	model, err := l.WechatUserService.GetByMobile(req.Mobile)
	if nil != err {
		return
	}
	resp = &types.MobileIsExistsResponse{IsExist: false}
	if model != nil && model.Id > 0 {
		resp.IsExist = true
	}
	return
}
