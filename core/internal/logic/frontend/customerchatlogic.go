package frontend

import (
	"context"
	"core/internal/service/ai"
	"core/internal/service/structure"
	"fmt"

	"core/internal/svc"
	"core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CustomerChatLogic struct {
	logx.Logger
	ctx         context.Context
	svcCtx      *svc.ServiceContext
	chatService *ai.ChatService
}

func NewCustomerChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustomerChatLogic {
	return &CustomerChatLogic{
		Logger:      logx.WithContext(ctx),
		ctx:         ctx,
		svcCtx:      svcCtx,
		chatService: ai.NewChatService(),
	}
}

func (l *CustomerChatLogic) CustomerChat(req *types.CustomerChatRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	open := ai.NewOpenAi(l.ctx, l.svcCtx.Config)
	l.chatService.Set(open)
	l.chatService.SetConfig("gpt-3.5-turbo-0613", 0.8, 5000)
	info, err := l.chatService.Chat([]structure.ChatModelMessage{
		{
			Role:    "user",
			Content: "aaa",
		},
	})
	fmt.Println(info)
	fmt.Println(err)
	return
}
