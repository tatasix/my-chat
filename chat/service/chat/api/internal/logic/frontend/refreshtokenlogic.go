package frontend

import (
	"chat/common/jwt"
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/service"
	"context"
	"time"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	*service.WechatUserService
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger:            logx.WithContext(ctx),
		ctx:               ctx,
		svcCtx:            svcCtx,
		WechatUserService: service.NewWechatUserService(ctx, svcCtx),
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenRequest) (resp *types.RefreshTokenResponse, err error) {

	//// 生成 token 并进行响应
	jwtObj := jwt.NewJwt(l.ctx, l.svcCtx.Config.FrontendAuth.AccessSecret)

	user, err := jwtObj.ParseToken(req.Authorization)
	if err != nil || user == "" {
		return
	}

	userPo, err := l.WechatUserRepository.GetByUser(user)
	if err != nil || userPo == nil || userPo.Id == 0 {
		err = util.ReturnError(xerr.RecordNotFound)
		return
	}

	// 生成 token
	newToken, err := jwtObj.GetJwtToken(req.User, time.Now().Unix(), l.svcCtx.Config.FrontendAuth.AccessExpire)
	if nil != err {
		return
	}

	//旧token过期
	//err = jwtObj.AddBlack(req.Authorization)
	//if nil != err {
	//	return
	//}
	resp = &types.RefreshTokenResponse{Token: newToken}
	return
}
