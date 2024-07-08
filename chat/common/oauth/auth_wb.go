package oauth

import (
	"chat/common/util"
	"errors"
	"strconv"
)

// 微博授权登录
type AuthWb struct {
	BaseRequest
}

func NewAuthWb(conf *AuthConfig) *AuthWb {
	authRequest := &AuthWb{}
	authRequest.Set("weibo", conf)

	authRequest.authorizeUrl = "https://api.weibo.com/oauth2/authorize"
	authRequest.TokenUrl = "https://api.weibo.com/oauth2/access_token"
	authRequest.userInfoUrl = "https://api.weibo.com/2/users/show.json"

	return authRequest
}

// 获取登录地址
func (a *AuthWb) GetRedirectUrl(state string) (*CodeResult, error) {
	url := util.NewUrlBuilder(a.authorizeUrl).
		AddParam("response_type", "code").
		AddParam("client_id", a.config.ClientId).
		AddParam("redirect_uri", a.config.RedirectUrl).
		AddParam("state", a.GetState(state)).
		Build()

	_, err := util.PostV2(url)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// 获取token
func (a *AuthWb) GetToken(code string) (*TokenResult, error) {
	url := util.NewUrlBuilder(a.TokenUrl).
		AddParam("grant_type", "authorization_code").
		AddParam("code", code).
		AddParam("client_id", a.config.ClientId).
		AddParam("client_secret", a.config.ClientSecret).
		AddParam("redirect_uri", a.config.RedirectUrl).
		Build()

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
		Scope:        m["scope"],
		TokenType:    m["token_type"],
	}
	return token, nil
}

// 获取第三方用户信息
func (a *AuthWb) GetUserInfo(openId string, accessToken string) (*UserResult, error) {
	url := util.NewUrlBuilder(a.TokenUrl).
		AddParam("uid", openId).
		AddParam("access_token", accessToken).
		Build()

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
		UUID:      m["id"],
		UserName:  m["login"],
		NickName:  m["name"],
		AvatarUrl: m["avatar_url"],
		Company:   m["company"],
		Blog:      m["blog"],
		Location:  m["location"],
		Email:     m["email"],
		Remark:    m["bio"],
		Url:       m["html_url"],
		CreatedAt: m["created_at"],
		UpdatedAt: m["updated_at"],
		Source:    a.registerSource,
		Gender:    sex,
	}
	return user, nil
}
