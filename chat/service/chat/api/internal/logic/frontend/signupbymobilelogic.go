package frontend

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignupByMobileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	*service.WechatUserService
}

func NewSignupByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignupByMobileLogic {
	return &SignupByMobileLogic{
		Logger:            logx.WithContext(ctx),
		ctx:               ctx,
		svcCtx:            svcCtx,
		WechatUserService: service.NewWechatUserService(ctx, svcCtx),
	}
}

func (l *SignupByMobileLogic) SignupByMobile(req *types.SignupByMobileRequest) (resp *types.Response, err error) {
	if !util.CheckMobile(req.Mobile) {
		err = util.ReturnError(xerr.LoginMobileError)
		return
	}

	//check verify code
	if !l.WechatUserService.VerifyCode(req.Mobile, req.Code) {
		err = util.ReturnError(xerr.LoginVerifyCodeError)
		return
	}
	model, err := l.WechatUserService.GetByMobile(req.Mobile)
	if nil != err {
		return
	}

	if model != nil && model.Id > 0 {
		err = util.ReturnError(xerr.LoginAccountExist)
		return
	}
	_, err = l.WechatUserService.DoRegister(req.User, req.Mobile, req.Password, req.Name, req.Birthday, req.Constellation, repository.UserStatusSuccess, 0)
	return assembler.Return(err)
}
