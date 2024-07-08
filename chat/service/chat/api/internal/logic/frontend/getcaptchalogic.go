package frontend

import (
	"chat/service/chat/api/internal/util/captcha"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaLogic) GetCaptcha() (resp *types.GetCaptchaResponse, err error) {
	resp = &types.GetCaptchaResponse{}
	if l.svcCtx.Config.Captcha.Open {
		captchaObj := captcha.Captcha{}

		id, image, length, err1 := captchaObj.Captcha(l.ctx, l.svcCtx)
		if err1 != nil {
			err = err1
			return
		}
		resp.CaptchaId = id
		resp.Captcha = image
		resp.CaptchaLength = length
	}
	return
}
