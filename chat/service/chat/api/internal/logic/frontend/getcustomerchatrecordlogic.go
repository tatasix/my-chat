package frontend

import (
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCustomerChatRecordLogic struct {
	logx.Logger
	ctx             context.Context
	svcCtx          *svc.ServiceContext
	chatService     *service.ChatService
	userLikeService *service.UserLikeService
}

func NewGetCustomerChatRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCustomerChatRecordLogic {
	return &GetCustomerChatRecordLogic{
		Logger:          logx.WithContext(ctx),
		ctx:             ctx,
		svcCtx:          svcCtx,
		chatService:     service.NewChatService(ctx, svcCtx),
		userLikeService: service.NewUserLikeService(ctx, svcCtx),
	}
}

func (l *GetCustomerChatRecordLogic) GetCustomerChatRecord(req *types.GetCustomerChatRecordRequest) (resp *types.GetCustomerChatRecordResponse, err error) {
	l.Logger.Info(" GetCustomerChatRecord start ")

	chatRecordPos, count, err := l.chatService.GetAllRecord(0, req.OpenKfID, req.UserId, "", req.StartCreatedAt, req.EndCreatedAt, "created_at desc", uint64(req.Page), uint64(req.PageSize), 2)
	if err != nil {
		fmt.Printf("chatRecordRepository.GetAll error: %v", err)
		return
	}

	if count <= 0 || len(chatRecordPos) <= 0 {
		return &types.GetCustomerChatRecordResponse{
			List:     nil,
			Total:    0,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, nil
	}

	chatRecords := assembler.POTODTOGetChatRecordList(chatRecordPos)
	l.Logger.Info(" GetCustomerChatRecord end ")
	if req.UserId != "" {
		userLike, err := l.userLikeService.UserLikeRepository.GetByUser(req.UserId)

		if err == nil && userLike != nil && len(userLike) > 0 {
			arr := []string{}
			for _, v := range userLike {
				arr = append(arr, v.ChatRecordId)
			}

			for kk, vv := range chatRecords {
				if vv.ChatRecordId != "" {
					for _, vvv := range arr {
						if vvv == vv.ChatRecordId {
							chatRecords[kk].IsLike = 1
						}
					}
				}
			}
		}

	}
	return &types.GetCustomerChatRecordResponse{
		List:     chatRecords,
		Total:    count,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
