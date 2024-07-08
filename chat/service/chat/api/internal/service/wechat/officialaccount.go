package wechat

import (
	"chat/service/chat/api/internal/svc"
	"context"
	wechat1 "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/zeromicro/go-zero/core/logx"
)

type OfficialAccount struct {
	logger          logx.Logger
	ctx             context.Context
	svcCtx          *svc.ServiceContext
	wc              *wechat1.Wechat
	officialAccount *officialaccount.OfficialAccount
}

// NewOfficialAccount 初始化
func NewOfficialAccount(ctx context.Context, svcCtx *svc.ServiceContext) *OfficialAccount {
	cfg := svcCtx.Config
	wc := wechat1.NewWechat()
	redisOpts := &cache.RedisOpts{
		Host:     cfg.RedisCache[0].Host,
		Password: cfg.RedisCache[0].Pass,
	}
	redisCache := cache.NewRedis(ctx, redisOpts)
	wc.SetCache(redisCache)

	offCfg := &offConfig.Config{
		AppID:          cfg.Wechat.Appid,
		AppSecret:      cfg.Wechat.Secret,
		Token:          cfg.Wechat.Token,
		EncodingAESKey: cfg.Wechat.EncodingAESKey,
	}

	officialAccount := wc.GetOfficialAccount(offCfg)

	return &OfficialAccount{
		logger:          logx.WithContext(ctx),
		ctx:             ctx,
		svcCtx:          svcCtx,
		wc:              wc,
		officialAccount: officialAccount,
	}
}

func (o *OfficialAccount) GetRedirectURL(state, returnUrl string) (url string, err error) {
	url, err = o.officialAccount.GetOauth().GetRedirectURL(o.svcCtx.Config.Wechat.RedirectUri+returnUrl, o.svcCtx.Config.Wechat.Scope, state)
	return
}

func (o *OfficialAccount) GetAccessToken(code string) (openId, unionID, accessToken string, err error) {
	access, err := o.officialAccount.GetOauth().GetUserAccessToken(code)
	openId = access.OpenID
	unionID = access.UnionID
	accessToken = access.AccessToken
	return
}

func (o *OfficialAccount) GetUserInfo(accessToken, openId, lang string) (ak string, err error) {
	access, err := o.officialAccount.GetOauth().GetUserInfo(accessToken, openId, lang)
	ak = access.OpenID
	return
}

func (o *OfficialAccount) GetConfig(uri string) (appId, nonceStr, signature string, timestamp int64, err error) {
	o.logger.Infof("OfficialAccount GetConfig %s", uri)
	info, err := o.officialAccount.GetJs().GetConfig(uri)
	if err != nil {
		return
	}
	appId = info.AppID
	nonceStr = info.NonceStr
	signature = info.Signature
	timestamp = info.Timestamp
	var accessToken string
	accessToken, err = o.officialAccount.GetJs().GetAccessToken()
	if err != nil {
		return
	}
	var ticketStr string
	ticketStr, err = o.officialAccount.GetJs().GetTicket(accessToken)
	if err != nil {
		return
	}
	o.logger.Infof("OfficialAccount GetConfig accessToken:%s  ticketStr:%s", accessToken, ticketStr)
	return
}
