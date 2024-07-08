package frontend

import (
	"chat/common/redis"
	"chat/common/sms"
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/api/internal/util/captcha"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
)

type SendSmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsLogic {
	return &SendSmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendSmsLogic) SendSms(req *types.SendSmsRequest) (resp *types.Response, err error) {
	if !util.CheckMobile(req.Mobile) {
		err = util.ReturnError(xerr.LoginMobileError)
		return
	}
	if l.svcCtx.Config.Captcha.Open {
		captchaObj := captcha.Captcha{}
		if !captchaObj.CaptchaVerify(l.ctx, req.CaptchaId, req.Captcha) {
			err = util.ReturnError(xerr.LoginCaptchaError)
			return
		}
	}

	//生成随机数
	code := util.GetRandomNum(6)
	redis.Rdb.Set(l.ctx, fmt.Sprintf(redis.LoginCodeKey, req.Mobile), code, redis.LoginCodeExpireTime)
	msg := strings.Replace(sms.SMSTemplate, "{code}", code, 1)
	err = sms.SendSms(l.ctx, req.Mobile, msg)
	if err != nil {
		return
	}

	return &types.Response{Message: "ok"}, nil
}
