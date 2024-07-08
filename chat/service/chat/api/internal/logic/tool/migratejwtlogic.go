package tool

import (
	"chat/common/jwt"
	"context"
	"time"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MigrateJWTLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMigrateJWTLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MigrateJWTLogic {
	return &MigrateJWTLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MigrateJWTLogic) MigrateJWT(req *types.MigrateJWTRequest) (resp *types.MigrateJWTResponse, err error) {
	//// 生成 token 并进行响应
	jwtObj := jwt.NewJwt(l.ctx, l.svcCtx.Config.FrontendAuth.AccessSecret)

	// 生成 token
	newToken, err := jwtObj.GetJwtToken(req.User, time.Now().Unix(), l.svcCtx.Config.FrontendAuth.AccessExpire)
	if nil != err {
		return
	}
	resp = &types.MigrateJWTResponse{Token: newToken}
	return
}
