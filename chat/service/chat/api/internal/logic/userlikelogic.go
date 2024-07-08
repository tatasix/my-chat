package logic

import (
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/model"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLikeLogic struct {
	logx.Logger
	ctx             context.Context
	svcCtx          *svc.ServiceContext
	userLikeService *service.UserLikeService
}

func NewUserLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLikeLogic {
	return &UserLikeLogic{
		Logger:          logx.WithContext(ctx),
		ctx:             ctx,
		svcCtx:          svcCtx,
		userLikeService: service.NewUserLikeService(ctx, svcCtx),
	}
}

func (l *UserLikeLogic) UserLike(req *types.UserLikeReq) (*types.UserLikeReply, error) {

	switch req.Type {
	case 1:
		row, _ := l.userLikeService.UserLikeRepository.GetByUserAndChatRecordId(req.User, req.ChatRecordId)

		if row != nil && row.Id > 0 {
			return nil, nil
		}
		_, err := l.userLikeService.UserLikeRepository.Insert(&model.UserLike{
			User:         req.User,
			ChatRecordId: req.ChatRecordId,
		})
		if err != nil {
			return nil, errors.New("收藏失败")
		}
		return nil, nil

	case 2:
		err := l.userLikeService.UserLikeRepository.DeleteByUserAndRecordId(req.User, req.ChatRecordId)
		return nil, err
	default:
		return nil, errors.New("非法参数")
	}

	return &types.UserLikeReply{Message: "success"}, nil
}
