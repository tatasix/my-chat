package frontend

import (
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/service"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	*service.WechatUserService
}

func NewUpdateLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLoginLogic {
	return &UpdateLoginLogic{
		Logger:            logx.WithContext(ctx),
		ctx:               ctx,
		svcCtx:            svcCtx,
		WechatUserService: service.NewWechatUserService(ctx, svcCtx),
	}
}

func (l *UpdateLoginLogic) UpdateLogin(req *types.UpdateLoginRequest) (resp *types.LoginResponse, err error) {

	wechatUser, err := l.WechatUserService.UpdateLogin(req.User, req.Password, req.Name, req.Birthday, req.Constellation, req.Status)
	if err != nil {
		return
	}
	if wechatUser.Status != repository.UserStatusSuccess {
		return &types.LoginResponse{
			User:   req.User,
			Status: req.Status,
		}, nil
	}
	token, err := l.WechatUserService.DoLogin(wechatUser, wechatUser.IsLoginFree)
	if err != nil {
		return
	}
	return &types.LoginResponse{
		User:   req.User,
		Status: req.Status,
		Token:  token,
	}, nil
}
