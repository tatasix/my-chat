package oauth

import (
	"chat/common/util"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

// AuthWxWechat 微信授权登录（第三方应用）
type AuthWxWechat struct {
	BaseRequest
}

func NewAuthWxWechat(conf *AuthConfig) *AuthWxWechat {
	authRequest := &AuthWxWechat{}
	authRequest.Set("weixin", conf)

	authRequest.authorizeUrl = "https://open.weixin.qq.com/connect/qrconnect"
	authRequest.TokenUrl = "https://api.weixin.qq.com/sns/oauth2/access_token"
	authRequest.userInfoUrl = "https://api.weixin.qq.com/sns/userinfo"

	return authRequest
}

// GetRedirectUrl 获取登录地址
func (a *AuthWxWechat) GetRedirectUrl(state string) string {
	url := util.NewUrlBuilder(a.authorizeUrl).
		AddParam("response_type", "code").
		AddParam("appid", a.config.ClientId).
		AddParam("redirect_uri", a.config.RedirectUrl).
		AddParam("scope", "snsapi_login").
		AddParam("state", a.GetState(state)).
		Build()

	return url
}

// GetWebAccessToken 获取token
func (a *AuthWxWechat) GetWebAccessToken(code string) (*TokenResult, error) {
	url := util.NewUrlBuilder(a.TokenUrl).
		AddParam("grant_type", "authorization_code").
		AddParam("code", code).
		AddParam("appid", a.config.ClientId).
		AddParam("secret", a.config.ClientSecret).
		AddParam("redirect_uri", a.config.RedirectUrl).
		Build()
	logx.Infof("GetWebAccessToken url", url)

	body, err := util.PostV2(url)
	if err != nil {
		return nil, err
	}
	m := util.JsonToMSS(body)
	if _, ok := m["error"]; ok {
		return nil, errors.New(m["error_description"])
	}
	token := &TokenResult{
		AccessToken:  m["access_token"],
		RefreshToken: m["refresh_token"],
		ExpireIn:     m["expires_in"],
		OpenId:       m["openid"],
		UnionId:      m["unionid"],
		Scope:        m["scope"],
		TokenType:    m["token_type"],
	}
	if token.AccessToken == "" {
		return nil, errors.New("获取AccessToken数据为空！")
	}
	return token, nil
}

// GetAppAccessToken 通过移动应用获取AccessToken
func (a *AuthWxWechat) GetAppAccessToken(code string) (*TokenResult, error) {
	url := util.NewUrlBuilder(a.TokenUrl).
		AddParam("grant_type", "authorization_code").
		AddParam("code", code).
		AddParam("appid", a.config.ClientId).
		AddParam("secret", a.config.ClientSecret).
		Build()

	logx.Infof("GetAppAccessToken url", url)
	body, err := util.PostV2(url)
	if err != nil {
		return nil, err
	}
	m := util.JsonToMSS(body)
	if _, ok := m["error"]; ok {
		return nil, errors.New(m["error_description"])
	}
	token := &TokenResult{
		AccessToken:  m["access_token"],
		RefreshToken: m["refresh_token"],
		ExpireIn:     m["expires_in"],
		OpenId:       m["openid"],
		UnionId:      m["unionid"],
		Scope:        m["scope"],
		TokenType:    m["token_type"],
	}
	if token.AccessToken == "" {
		return nil, errors.New("获取AccessToken数据为空！")
	}
	return token, nil
}

// GetUserInfo 获取第三方用户信息
// https://github.com/geiqin/thirdparty/blob/main/oauth/auth_wx_wechat.go
func (a *AuthWxWechat) GetUserInfo(accessToken string, openId string) (*UserResult, error) {
	url := util.NewUrlBuilder(a.userInfoUrl).
		AddParam("openid", openId).
		AddParam("access_token", accessToken).
		Build()

	logx.Infof("GetUserInfo url", url)
	body, err := util.GetV2(url)

	if err != nil {
		return nil, err
	}
	m := util.JsonToMSS(body)
	if _, ok := m["error"]; ok {
		return nil, errors.New(m["error_description"])
	}
	var sex int32
	if v, ok := m["sex"]; ok {
		sexInt, _ := strconv.Atoi(v)
		sex = int32(sexInt)
	}

	user := &UserResult{
		OpenId:    m["openid"],
		UnionId:   m["unionid"],
		UserName:  m["nickname"],
		NickName:  m["nickname"],
		AvatarUrl: m["headimgurl"],
		City:      m["city"],
		Province:  m["province"],
		Country:   m["country"],
		Language:  m["language"],
		Source:    a.registerSource,
		Gender:    sex,
	}
	if user.OpenId == "" {
		return nil, errors.New("获取用户信息为空！")
	}
	return user, nil
}

// GetUser 直接获取第三方用户信息
func (a *AuthWxWechat) GetUser(code string) (user *UserResult, err error) {
	tokenInfo, err := a.GetWebAccessToken(code)
	if err != nil {
		return
	}

	return a.GetUserInfo(tokenInfo.AccessToken, tokenInfo.OpenId)
}
