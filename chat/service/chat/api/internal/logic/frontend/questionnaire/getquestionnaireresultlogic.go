package questionnaire

import (
	"chat/service/chat/api/internal/service/questionnaire"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuestionnaireResultLogic struct {
	logx.Logger
	ctx                  context.Context
	svcCtx               *svc.ServiceContext
	QuestionnaireService *questionnaire.QuestionnaireService
}

func NewGetQuestionnaireResultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionnaireResultLogic {
	return &GetQuestionnaireResultLogic{
		Logger:               logx.WithContext(ctx),
		ctx:                  ctx,
		svcCtx:               svcCtx,
		QuestionnaireService: questionnaire.NewQuestionnaireService(ctx, svcCtx),
	}
}

func (l *GetQuestionnaireResultLogic) GetQuestionnaireResult(req *types.GetQuestionnaireRequest) (resp *types.GetQuestionnaireResultResponse, err error) {
	return l.QuestionnaireService.GetResult(req.OpenKfId, req.User)
}
