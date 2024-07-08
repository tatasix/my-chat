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

type FindFeedbackLogic struct {
	logx.Logger
	ctx             context.Context
	svcCtx          *svc.ServiceContext
	feedbackService *service.FeedbackService
}

func NewFindFeedbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindFeedbackLogic {
	return &FindFeedbackLogic{
		Logger:          logx.WithContext(ctx),
		ctx:             ctx,
		svcCtx:          svcCtx,
		feedbackService: service.NewFeedbackService(ctx, svcCtx),
	}
}

func (l *FindFeedbackLogic) FindFeedback(req *types.FindFeedbackRequest) (resp *types.FeedbackResponse, err error) {
	feedbackPo, err := l.feedbackService.FindFeedback(req.Id)
	if err != nil {
		fmt.Printf("GetFeedbackList error: %v", err)
		return
	}

	feedbackDto := assembler.POTODTOFindFeedback(feedbackPo)
	return &feedbackDto, nil
}
