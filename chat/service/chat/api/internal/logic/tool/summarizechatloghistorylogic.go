package tool

import (
	"chat/service/chat/api/internal/service"
	"context"
	"fmt"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"github.com/gammazero/workerpool"
	"github.com/zeromicro/go-zero/core/logx"
)

type SummarizeChatLogHistoryLogic struct {
	logx.Logger
	ctx         context.Context
	svcCtx      *svc.ServiceContext
	chatService *service.ChatService
}

func NewSummarizeChatLogHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SummarizeChatLogHistoryLogic {
	return &SummarizeChatLogHistoryLogic{
		Logger:      logx.WithContext(ctx),
		ctx:         ctx,
		svcCtx:      svcCtx,
		chatService: service.NewChatService(ctx, svcCtx),
	}
}

func (l *SummarizeChatLogHistoryLogic) SummarizeChatLogHistory(req *types.SummarizeChatLogRequest) (resp *types.Response, err error) {
	wp := workerpool.New(2)
	requests := []string{"alpha", "beta", "gamma", "delta", "epsilon"}

	for _, r := range requests {
		r := r
		wp.Submit(func() {
			fmt.Println("Handling request:", r)
		})
	}

	wp.StopWait()
	return
}
