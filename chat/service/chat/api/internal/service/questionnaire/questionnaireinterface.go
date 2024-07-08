package questionnaire

import (
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type Handler interface {
	Handle(customerConfig *model.CustomerConfig, config model.CustomerConfigConfig, relationId int64, user, openKfID string, channel chan string, baseInfo chan any) (resp *types.GenerateResultResponse, err error) // 处理方法
	GetReturn(customerConfig *model.CustomerConfig, user string) (results *types.GetQuestionnaireResultResponse, err error)                                                                                          // 处理方法
}

// NewQuestionnaireHandler Handler 工厂函数，用于创建不同类型的处理程序
func NewQuestionnaireHandler(ctx context.Context, svcCtx *svc.ServiceContext, method int64) Handler {
	switch method {
	case 5:
		//分类加和求平均值
		return &QuestionClassifySummationAverageHandler{
			logger: logx.WithContext(ctx),
			ctx:    ctx,
			svcCtx: svcCtx,
		}
	case 1:
		//普通加和
		return &SummationHandler{
			logger: logx.WithContext(ctx),
			ctx:    ctx,
			svcCtx: svcCtx,
		}
	case 2:
		// 48题
		return &QuestionClassifyHandler{
			logger: logx.WithContext(ctx),
			ctx:    ctx,
			svcCtx: svcCtx,
		}
	case 3:
		//选项分类（分类加和）93题
		return &OptionClassifyHandler{
			logger: logx.WithContext(ctx),
			ctx:    ctx,
			svcCtx: svcCtx,
		}
	case 4:
		//选项分类（分类加和）
		return &GroupAdditionClassifyHandler{
			logger: logx.WithContext(ctx),
			ctx:    ctx,
			svcCtx: svcCtx,
		}
	default:
		return nil
	}
}
