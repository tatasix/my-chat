package frontend

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/vars/myerror"
	"chat/service/chat/model"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	*service.WechatUserService
}

func NewUpdateUserPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPasswordLogic {
	return &UpdateUserPasswordLogic{
		Logger:            logx.WithContext(ctx),
		ctx:               ctx,
		svcCtx:            svcCtx,
		WechatUserService: service.NewWechatUserService(ctx, svcCtx),
	}
}

func (l *UpdateUserPasswordLogic) UpdateUserPassword(req *types.UpdateUserPasswordRequest) (resp *types.Response, err error) {
	if !util.CheckMobile(req.Mobile) {
		err = myerror.LoginMobileError
		return
	}

	//check verify code
	if !l.WechatUserService.VerifyCode(req.Mobile, req.Code) {
		err = util.ReturnError(xerr.LoginVerifyCodeError)
		return
	}

	origin, err := l.WechatUserService.GetByUser(req.User)
	if nil != err {
		return
	}

	if origin == nil || origin.Id <= 0 {
		err = util.ReturnError(xerr.LoginAccountNotExist)
		return
	}
	return assembler.Return(l.WechatUserService.Update(origin, &model.WechatUser{
		Password: util.Sha1En(req.Password + origin.Salt),
	}))
}
