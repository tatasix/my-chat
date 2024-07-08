package tool

import (
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type SummarizeChatLogLogic struct {
	logx.Logger
	ctx         context.Context
	svcCtx      *svc.ServiceContext
	chatService *service.ChatService
}

func NewSummarizeChatLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SummarizeChatLogLogic {
	return &SummarizeChatLogLogic{
		Logger:      logx.WithContext(ctx),
		ctx:         ctx,
		svcCtx:      svcCtx,
		chatService: service.NewChatService(ctx, svcCtx),
	}
}

func (l *SummarizeChatLogLogic) SummarizeChatLog(r *types.SummarizeChatLogRequest) (resp *types.Response, err error) {
	// 获取昨天的聊天记录
	var startTime, endTime string
	yesterday := time.Now().Add(-24 * time.Hour).Format("2006-01-02")

	if r.StartTime != "" {
		startTime = r.StartTime
	} else {
		startTime = yesterday + " 00:00:00"
	}
	if r.EndTime != "" {
		endTime = r.EndTime
	} else {
		endTime = yesterday + " 23:59:59"
	}

	err = l.chatService.SummarizeChatLog(l.ctx, r.StartId, r.EndId, startTime, endTime)
	if err != nil {
		return
	}
	return &types.Response{Message: "ok"}, nil

}
