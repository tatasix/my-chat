package questionnaire

import (
	"chat/service/chat/api/internal/service/questionnaire"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateResultLogic struct {
	logx.Logger
	ctx                  context.Context
	svcCtx               *svc.ServiceContext
	QuestionnaireService *questionnaire.QuestionnaireService
}

func NewGenerateResultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateResultLogic {
	return &GenerateResultLogic{
		Logger:               logx.WithContext(ctx),
		ctx:                  ctx,
		svcCtx:               svcCtx,
		QuestionnaireService: questionnaire.NewQuestionnaireService(ctx, svcCtx),
	}
}

func (l *GenerateResultLogic) GenerateResult(req *types.GenerateResultRequest, channel chan string, baseInfo chan any) (resp *types.GenerateResultResponse, err error) {
	return l.QuestionnaireService.GenerateResult(req.RelationId, req.User, req.OpenKfID, channel, baseInfo)
}
