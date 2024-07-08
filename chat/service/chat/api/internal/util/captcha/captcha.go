package captcha

import (
	"chat/service/chat/api/internal/svc"
	"context"
	"github.com/mojocn/base64Captcha"
)

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
var store = NewDefaultRedisStore()

//var store = base64Captcha.DefaultMemStore

type Captcha struct{}

// Captcha
// @Summary   生成验证码
func (b *Captcha) Captcha(ctx context.Context, svcCtx *svc.ServiceContext) (captchaId string, picPath string, captchaLength int, err error) {
	if !svcCtx.Config.Captcha.Open {
		return
	}

	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(svcCtx.Config.Captcha.ImgHeight, svcCtx.Config.Captcha.ImgWidth, svcCtx.Config.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(ctx))
	//cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		return
	}
	return id, b64s, svcCtx.Config.Captcha.KeyLong, nil
}

func (b *Captcha) CaptchaVerify(ctx context.Context, captchaId, captcha string) bool {
	if captchaId == "" || captcha == "" {
		return false
	}
	return store.UseWithCtx(ctx).Verify(captchaId, captcha, true)
}
