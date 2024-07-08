package logic

import (
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/service"
	"context"
	"fmt"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFeedbackLogic struct {
	logx.Logger
	ctx             context.Context
	svcCtx          *svc.ServiceContext
	feedbackService *service.FeedbackService
}

func NewGetFeedbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFeedbackLogic {
	return &GetFeedbackLogic{
		Logger:          logx.WithContext(ctx),
		ctx:             ctx,
		svcCtx:          svcCtx,
		feedbackService: service.NewFeedbackService(ctx, svcCtx),
	}
}

func (l *GetFeedbackLogic) GetFeedback(req *types.GetFeedbackRequest) (resp *types.GetFeedbackResponse, err error) {
	feedbackPos, count, err := l.feedbackService.GetFeedbackList(req.OpenKfId, req.User, req.StartCreatedAt, req.EndCreatedAt, "id desc", uint64(req.Page), uint64(req.PageSize), int32(req.Status))
	if err != nil {
		fmt.Printf("GetFeedbackList error: %v", err)
		return
	}

	if count <= 0 || len(feedbackPos) <= 0 {
		return &types.GetFeedbackResponse{
			List:     nil,
			Total:    0,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, nil
	}

	feedbackDto := assembler.POTODTOGetFeedbackList(feedbackPos)
	return &types.GetFeedbackResponse{
		List:     feedbackDto,
		Total:    count,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
