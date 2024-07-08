package customer

import (
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type Handler interface {
	Handle(f int32, r *types.EvaluationSocketRequest, customerConfig *model.CustomerConfig, state *model.State, channel chan string, baseChannel chan any) (*types.EvaluationSocketResponse, error) // 处理方法
	HandleGet(f int32, r *types.GetEvaluationRequest) (*types.GetEvaluationResponse, error)                                                                                                         // 处理方法
	HandleGetSocket(state *model.State, r *types.GetEvaluationSocketRequest) (*types.GetEvaluationSocketResponse, error)                                                                            // 处理方法
}

// NewCustomerHandler Handler 工厂函数，用于创建不同类型的处理程序
func NewCustomerHandler(ctx context.Context, svcCtx *svc.ServiceContext, customer string) Handler {
	switch customer {
	case "oLet5ixVLgOqflofOJqjXqSJg0zYlF7U":
		return &FortuneHandler{
			logger: logx.WithContext(ctx),
			ctx:    ctx,
			svcCtx: svcCtx,
		}
	case "uIcMlmqSXJQ6259n6I3QMfSODVeFOwk5":
		return &DreamInterpretationHandler{
			logger: logx.WithContext(ctx),
			ctx:    ctx,
			svcCtx: svcCtx,
		}
	case "lGVmDRmK6W30ROG5eSN8wOlXlroy0D6t":
		return &PsychologySandboxHandler{
			logger: logx.WithContext(ctx),
			ctx:    ctx,
			svcCtx: svcCtx,
		}
	default:
		return nil
	}
}
