package frontend

import (
	"chat/common/jwt"
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/service/wechat"
	"chat/service/chat/model"
	"context"
	"time"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoByCodeLogic struct {
	logx.Logger
	ctx               context.Context
	svcCtx            *svc.ServiceContext
	wechatService     *wechat.OfficialAccount
	wechatUserService *service.WechatUserService
}

func NewGetInfoByCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoByCodeLogic {
	return &GetInfoByCodeLogic{
		Logger:            logx.WithContext(ctx),
		ctx:               ctx,
		svcCtx:            svcCtx,
		wechatService:     wechat.NewOfficialAccount(ctx, svcCtx),
		wechatUserService: service.NewWechatUserService(ctx, svcCtx),
	}
}

func (l *GetInfoByCodeLogic) GetInfoByCode(req *types.GetInfoByCodeRequest) (resp *types.LoginResponse, err error) {
	l.Logger.Infof("GetInfoByCodeLogic GetInfoByCode req.Code %s req.State %s ", req.Code, req.State)

	openID, unionId, _, err := l.wechatService.GetAccessToken(req.Code)
	if err != nil {
		l.Logger.Errorf("GetInfoByCodeLogic GetInfoByCode error %s", err.Error())
		return
	}
	l.Logger.Infof("GetInfoByCodeLogic GetInfoByCode openID %s unionId %s ", openID, unionId)

	user, err := l.wechatUserService.GetByUser(req.State)
	if err != nil {
		l.Logger.Errorf("GetInfoByCodeLogic wechatUserService GetByUser error %s", err.Error())
		return
	}
	if user == nil || user.Id <= 0 || user.Status != repository.UserStatusSuccess {
		err = util.ReturnError(xerr.RecordNotFound)
		return
	}

	err = l.wechatUserService.Update(user, &model.WechatUser{
		User:    req.State,
		Unionid: unionId,
		Openid:  openID,
	})
	if err != nil {
		l.Logger.Errorf("GetInfoByCodeLogic wechatUserService Update error %s", err.Error())
		return
	}

	// 生成 token 并进行响应
	jwtObj := jwt.NewJwt(l.ctx, l.svcCtx.Config.FrontendAuth.AccessSecret)
	token, _ := jwtObj.GetJwtToken(
		user.User,
		time.Now().Unix(),
		l.svcCtx.Config.FrontendAuth.AccessExpire,
	)

	resp = &types.LoginResponse{
		Token:  token,
		User:   user.User,
		Status: user.Status,
	}
	return
}
