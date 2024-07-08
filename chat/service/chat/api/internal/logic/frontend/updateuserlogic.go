package frontend

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/util/sensitive"
	"chat/service/chat/model"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	*service.WechatUserService
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger:            logx.WithContext(ctx),
		ctx:               ctx,
		svcCtx:            svcCtx,
		WechatUserService: service.NewWechatUserService(ctx, svcCtx),
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserRequest) (resp *types.Response, err error) {
	if "" == req.User {
		err = util.ReturnError(xerr.ParamMiss)
		return
	}
	if req.Name != "" {
		if !sensitive.NewSensitive(l.ctx, l.svcCtx).Check(req.Name) {
			l.Logger.Errorf("SensitiveError user:%s name:%s", req.User, req.Name)
			err = util.ReturnError(xerr.SensitiveError)
			return
		}
	}
	return assembler.Return(l.WechatUserService.UpdateByUser(req.User, &model.WechatUser{
		Name:          req.Name,
		Birthday:      req.Birthday,
		Gender:        int64(req.Sex),
		Constellation: req.Constellation,
		Mbti:          req.Mbti,
	}))
}
