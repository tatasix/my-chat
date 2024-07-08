package tool

import (
	"chat/common/util"
	"chat/service/chat/api/internal/repository"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EncryptMobileLogic struct {
	logx.Logger
	ctx                  context.Context
	svcCtx               *svc.ServiceContext
	wechatUserRepository *repository.WechatUserRepository
}

func NewEncryptMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EncryptMobileLogic {
	return &EncryptMobileLogic{
		Logger:               logx.WithContext(ctx),
		ctx:                  ctx,
		svcCtx:               svcCtx,
		wechatUserRepository: repository.NewWechatUserRepository(ctx, svcCtx),
	}
}

func (l *EncryptMobileLogic) EncryptMobile() (resp *types.Response, err error) {
	all, err := l.wechatUserRepository.GetAllHaveMobile()
	if err != nil {
		return
	}
	if len(all) > 0 {
		for _, v := range all {
			mobile := util.Encrypt(v.Mobile)
			l.Logger.Infof("encrypt mobile id:%d old:%s new:%s", v.Id, v.Mobile, mobile)
			v.Mobile = mobile
			err = l.wechatUserRepository.UpdateV2(v)
			if err != nil {
				return
			}
		}
	}
	return
}
