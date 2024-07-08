package service

import (
	"chat/common/jwt"
	"chat/common/oauth"
	"chat/common/redis"
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/model"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type WechatUserService struct {
	logx.Logger
	ctx                  context.Context
	svcCtx               *svc.ServiceContext
	WechatUserRepository *repository.WechatUserRepository
}

func NewWechatUserService(ctx context.Context, svcCtx *svc.ServiceContext) *WechatUserService {
	return &WechatUserService{
		Logger:               logx.WithContext(ctx),
		ctx:                  ctx,
		svcCtx:               svcCtx,
		WechatUserRepository: repository.NewWechatUserRepository(ctx, svcCtx),
	}
}

func (l *WechatUserService) UpdateById(id int64, model *model.WechatUser) (err error) {
	return l.WechatUserRepository.UpdateById(id, model)
}

func (l *WechatUserService) Update(old, model *model.WechatUser) (err error) {
	return l.WechatUserRepository.Update(old, model)
}

func (l *WechatUserService) UpdateByUser(user string, model *model.WechatUser) (err error) {
	return l.WechatUserRepository.UpdateByUser(user, model)
}

func (l *WechatUserService) WechatAuthorizeCallBack(state, code string) (err1 error) {
	wxConf := &oauth.AuthConfig{
		ClientId:     l.svcCtx.Config.Wechat.Appid,
		ClientSecret: l.svcCtx.Config.Wechat.Secret,
		RedirectUrl:  l.svcCtx.Config.Wechat.RedirectUrl,
	}

	wxAuth := oauth.NewAuthWxWechat(wxConf)

	userInfo, err := wxAuth.GetUser(code)
	if nil != err {
		l.Logger.Errorf("WechatAuthorizeCallBack err %+v", err)
		return
	}
	l.Logger.Infof("wxAuth.GetUser %+v", userInfo)
	if userInfo == nil {
		l.Logger.Errorf("WechatAuthorizeCallBack get userInfo error")
		return
	}
	update := &model.WechatUser{
		User:     state,
		Nickname: userInfo.NickName,
		Avatar:   userInfo.AvatarUrl,
		Gender:   int64(userInfo.Gender),
		Unionid:  userInfo.UnionId,
		Openid:   userInfo.OpenId,
		City:     userInfo.City,
		Province: userInfo.Province,
		Country:  userInfo.Country,
	}
	//保存用户信息
	return l.WechatUserRepository.UpdateByUser(state, update)
}

func (l *WechatUserService) GetRedirectUrl(user string) string {
	wxConf := &oauth.AuthConfig{
		ClientId:     l.svcCtx.Config.Wechat.Appid,
		ClientSecret: l.svcCtx.Config.Wechat.Secret,
		RedirectUrl:  l.svcCtx.Config.Wechat.RedirectUrl,
	}

	wxAuth := oauth.NewAuthWxWechat(wxConf)

	return wxAuth.GetRedirectUrl(user)
}

func (l *WechatUserService) GetByMobile(mobile string) (res *model.WechatUser, err error) {
	return l.WechatUserRepository.GetByMobile(mobile)
}
func (l *WechatUserService) GetByToken(token string) (res *model.WechatUser, err error) {
	// 生成 token 并进行响应
	jwtObj := jwt.NewJwt(l.ctx, l.svcCtx.Config.FrontendAuth.AccessSecret)
	user, err := jwtObj.ParseToken(token)
	if nil != err {
		return
	}
	if user == "" {
		return &model.WechatUser{}, nil
	}
	return l.WechatUserRepository.GetByUser(user)
}

func (l *WechatUserService) VerifyCode(mobile, code string) bool {
	cacheCode, err := redis.Rdb.Get(l.ctx, fmt.Sprintf(redis.LoginCodeKey, mobile)).Result()
	if err == nil && cacheCode == code {
		redis.Rdb.Del(l.ctx, fmt.Sprintf(redis.LoginCodeKey, mobile))
		return true
	}
	return false
}

func (l *WechatUserService) DoLogin(model *model.WechatUser, isLoginFree int64) (token string, err error) {
	var expire int64
	expire = 86400
	if isLoginFree != 0 {
		expire = l.svcCtx.Config.FrontendAuth.AccessExpire
	}
	// 生成 token 并进行响应
	jwtObj := jwt.NewJwt(l.ctx, l.svcCtx.Config.FrontendAuth.AccessSecret)
	return jwtObj.GetJwtToken(
		model.User,
		time.Now().Unix(),
		expire,
	)
}

func (l *WechatUserService) Logout(token string) (err error) {

	// 生成 token 并进行响应
	jwtObj := jwt.NewJwt(l.ctx, l.svcCtx.Config.FrontendAuth.AccessSecret)
	return jwtObj.AddBlack(token)
}

func (l *WechatUserService) Signout(token string) (err error) {
	jwtObj := jwt.NewJwt(l.ctx, l.svcCtx.Config.FrontendAuth.AccessSecret)

	//delete user
	user, err := jwtObj.ParseToken(token)
	if err != nil {
		return
	}
	//xlog.Color(xlog.Green).Debug(user, err)
	userPo, err := l.WechatUserRepository.GetByUser(user)
	if err != nil || userPo == nil || userPo.Id == 0 {
		err = util.ReturnError(xerr.RecordNotFound)
		return
	}
	userPo.IsDeleted = 1
	err = l.WechatUserRepository.UpdateById(userPo.Id, userPo)
	if err != nil {
		return
	}
	err = jwtObj.AddBlack(token)
	return
}

func (l *WechatUserService) DoRegister(user, mobile, password, name, birthday, constellation string, status, isLoginFree int64) (new *model.WechatUser, err error) {
	wechatUserPo, err := l.WechatUserRepository.GetByUser(user)
	if err != nil {
		l.Logger.Errorf("DoRegister GetByUser err:%+v", err)
		return
	}
	var salt, pass string
	if password != "" {
		salt = util.GetRandomBoth(4)
		pass = util.Sha1En(password + salt)
	}
	new = &model.WechatUser{
		User:          user,
		UserType:      repository.UserTypePc,
		Mobile:        mobile,
		Password:      pass,
		Name:          name,
		Birthday:      birthday,
		Constellation: constellation,
		Salt:          salt,
		Status:        status,
		IsLoginFree:   isLoginFree,
		UpdatedAt:     time.Now(),
	}
	//已存在就只需要更新了
	if wechatUserPo != nil && wechatUserPo.Id >= 0 {
		updateErr := l.WechatUserRepository.Update(wechatUserPo, new)
		if updateErr != nil {
			err = updateErr
			l.Logger.Errorf("DoRegister GetByUser err:%+v", err)
			return
		}
		return
	}
	insertErr := l.WechatUserRepository.Insert(new)
	if insertErr != nil {
		err = insertErr
		l.Logger.Errorf("DoRegister GetByUser err:%+v", err)
		return
	}
	return
}

func (l *WechatUserService) UpdateLogin(user, password, name, birthday, constellation string, status int64) (newWechatUser *model.WechatUser, err error) {
	wechatUserPo, err := l.WechatUserRepository.GetByUser(user)
	if err != nil {
		l.Logger.Errorf("DoRegister GetByUser err:%+v", err)
		return
	}
	if wechatUserPo == nil || wechatUserPo.Id <= 0 {
		err = util.ReturnError(xerr.RecordNotFound)
		return
	}
	var salt, pass string

	if password != "" {
		salt = util.GetRandomBoth(4)
		pass = util.Sha1En(password + salt)
	}
	newWechatUser = &model.WechatUser{
		User:          user,
		Password:      pass,
		Name:          name,
		Salt:          salt,
		Birthday:      birthday,
		Constellation: constellation,
		Status:        status,
		UpdatedAt:     time.Now(),
	}
	//已存在就只需要更新了
	err = l.WechatUserRepository.Update(wechatUserPo, newWechatUser)
	if err != nil {
		l.Logger.Errorf("DoRegister GetByUser err:%+v", err)
		return
	}

	return
}

func (l *WechatUserService) GetByUser(user string) (res *model.WechatUser, err error) {
	return l.WechatUserRepository.GetByUser(user)
}
