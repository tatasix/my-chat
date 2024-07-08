package frontend

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type SaveFeedbackLogic struct {
	logx.Logger
	ctx             context.Context
	svcCtx          *svc.ServiceContext
	feedbackService *service.FeedbackService
}

func NewSaveFeedbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveFeedbackLogic {
	return &SaveFeedbackLogic{
		Logger:          logx.WithContext(ctx),
		ctx:             ctx,
		svcCtx:          svcCtx,
		feedbackService: service.NewFeedbackService(ctx, svcCtx),
	}
}

func (l *SaveFeedbackLogic) SaveFeedback(req *types.SaveFeedbackRequest) (resp *types.Response, err error) {

	info, err := l.feedbackService.Idempotent(req.MessageId, req.User)
	if err != nil {
		return
	}
	if info != nil && info.Id > 0 {
		err = util.ReturnError(xerr.RepeatRequest)
		return
	}
	if err = l.feedbackService.Insert(&model.Feedback{
		User:      req.User,
		MessageId: req.MessageId,
		OpenKfId:  req.OpenKfId,
		Contact:   req.Contact,
		Images:    req.Images,
		Content:   util.StringToSql(req.Content),
		Status:    repository.StatusToDo,
	}); err != nil {
		return
	}
	return &types.Response{Message: "ok"}, nil
}
