package service

import (
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/api/internal/vars"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type UserLikeService struct {
	logx.Logger
	ctx                  context.Context
	svcCtx               *svc.ServiceContext
	UserLikeRepository   *repository.UserLikeRepository
	ChatRecordRepository *repository.ChatRecordRepository
}

func NewUserLikeService(ctx context.Context, svcCtx *svc.ServiceContext) *UserLikeService {
	return &UserLikeService{
		Logger:               logx.WithContext(ctx),
		ctx:                  ctx,
		svcCtx:               svcCtx,
		UserLikeRepository:   repository.NewUserLikeRepository(ctx, svcCtx),
		ChatRecordRepository: repository.NewChatRecordRepository(ctx, svcCtx),
	}
}

func (l *UserLikeService) GetUserLikeList(chatRecordUser, startCreatedAt, endCreatedAt string, order string, page, pageSize int) (resp *types.GetUserLikeListPageResult, err error) {

	userLikePos, count, err := l.UserLikeRepository.GetAll(chatRecordUser, startCreatedAt, endCreatedAt, order, uint64(page), uint64(pageSize))
	if err != nil {
		fmt.Printf("get userLike error: %v", err)
		return
	}
	if count <= 0 || len(userLikePos) <= 0 {
		resp = &types.GetUserLikeListPageResult{
			List:     nil,
			Total:    0,
			Page:     page,
			PageSize: pageSize,
		}
		return
	}
	var list []types.UserLikeResponse
	var content string
	if userLikePos != nil {
		for _, v := range userLikePos {
			chatRecordId, _ := strconv.ParseInt(v.ChatRecordId, 10, 64)
			chatRecordPos, err := l.ChatRecordRepository.GetById(chatRecordId)
			if chatRecordPos != nil && err == nil {
				content = chatRecordPos.Content
			}
			list = append(list, types.UserLikeResponse{
				Id:             v.Id,
				ChatRecordUser: v.User,
				ChatRecordId:   v.ChatRecordId,
				CreatedAt:      v.CreatedAt.Format(vars.TimeFormat),
				UpdatedAt:      v.UpdatedAt.Format(vars.TimeFormat),
				Content:        content,
			})
		}

	}

	return &types.GetUserLikeListPageResult{
		List:     list,
		Total:    count,
		Page:     page,
		PageSize: pageSize,
	}, nil
}
