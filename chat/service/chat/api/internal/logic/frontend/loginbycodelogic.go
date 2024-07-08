package frontend

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	*service.WechatUserService
}

func NewLoginByCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByCodeLogic {
	return &LoginByCodeLogic{
		Logger:            logx.WithContext(ctx),
		ctx:               ctx,
		svcCtx:            svcCtx,
		WechatUserService: service.NewWechatUserService(ctx, svcCtx),
	}
}

func (l *LoginByCodeLogic) LoginByCode(req *types.LoginByCodeRequest) (resp *types.LoginResponse, err error) {
	if !util.CheckMobile(req.Mobile) {
		err = util.ReturnError(xerr.LoginMobileError)
		return
	}
	model, err := l.WechatUserService.GetByMobile(req.Mobile)
	if nil != err {
		return
	}
	//check verify code
	if !l.WechatUserService.VerifyCode(req.Mobile, req.Code) {
		err = util.ReturnError(xerr.LoginVerifyCodeError)
		return
	}
	var status int64
	var token string
	user := req.User
	if model == nil || model.Id <= 0 {
		//注册+登录了
		status = repository.UserStatusMobile
		newModel, errRegister := l.WechatUserService.DoRegister(req.User, req.Mobile, "", "", "", "", status, int64(req.IsLoginFree))
		if errRegister != nil {
			err = errRegister
			return
		}
		model = newModel
	} else {
		if model.Status == repository.UserStatusSuccess {
			token, err = l.WechatUserService.DoLogin(model, int64(req.IsLoginFree))
			if err != nil {
				return nil, err
			}
		}
		if model.User != req.User {
			// 旧用户重新登录
			l.Logger.Infof(" warring LoginByCode the request user not same origin user origin:%s request user:%s mobile:%s", model.User, req.User, req.Mobile)
		}
		user = model.User
		status = model.Status
	}

	return &types.LoginResponse{
		User:   user,
		Token:  token,
		Status: status,
	}, nil
}
