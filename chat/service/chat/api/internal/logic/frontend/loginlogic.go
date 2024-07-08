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

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	*service.WechatUserService
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger:            logx.WithContext(ctx),
		ctx:               ctx,
		svcCtx:            svcCtx,
		WechatUserService: service.NewWechatUserService(ctx, svcCtx),
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	if !util.CheckMobile(req.Mobile) {
		err = util.ReturnError(xerr.LoginMobileError)
		return
	}
	model, err := l.WechatUserService.GetByMobile(req.Mobile)
	if nil != err {
		return
	}

	if model == nil || model.Id <= 0 {
		err = util.ReturnError(xerr.LoginAccountOrPasswordError)
		return
	}
	//登录
	if util.Sha1En(req.Password+model.Salt) != model.Password {
		err = util.ReturnError(xerr.LoginAccountOrPasswordError)
		return
	}
	token, err := l.WechatUserService.DoLogin(model, int64(req.IsLoginFree))
	if err != nil {
		return
	}
	return &types.LoginResponse{
		Token:  token,
		User:   model.User,
		Status: model.Status,
	}, nil
}
