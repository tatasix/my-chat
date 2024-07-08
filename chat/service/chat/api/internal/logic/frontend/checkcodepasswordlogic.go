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

type CheckCodePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	*service.WechatUserService
}

func NewCheckCodePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckCodePasswordLogic {
	return &CheckCodePasswordLogic{
		Logger:            logx.WithContext(ctx),
		ctx:               ctx,
		svcCtx:            svcCtx,
		WechatUserService: service.NewWechatUserService(ctx, svcCtx),
	}
}

func (l *CheckCodePasswordLogic) CheckCodePassword(req *types.CheckCodeRequest) (resp *types.CheckCodeReponse, err error) {
	if !util.CheckMobile(req.Mobile) {
		err = util.ReturnError(xerr.LoginMobileError)
		return
	}
	resp = &types.CheckCodeReponse{Data: false}
	//check verify code
	if l.WechatUserService.VerifyCode(req.Mobile, req.Code) {
		resp.Data = true
	}

	return
}
