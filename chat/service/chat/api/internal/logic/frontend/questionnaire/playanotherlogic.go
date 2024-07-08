package questionnaire

import (
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/service/questionnaire"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayAnotherLogic struct {
	logx.Logger
	ctx                  context.Context
	svcCtx               *svc.ServiceContext
	QuestionnaireService *questionnaire.QuestionnaireService
}

func NewPlayAnotherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayAnotherLogic {
	return &PlayAnotherLogic{
		Logger:               logx.WithContext(ctx),
		ctx:                  ctx,
		svcCtx:               svcCtx,
		QuestionnaireService: questionnaire.NewQuestionnaireService(ctx, svcCtx),
	}
}

func (l *PlayAnotherLogic) PlayAnother(req *types.PlayAnotherRequest) (resp *types.Response, err error) {
	return assembler.Return(l.QuestionnaireService.PlayAnother(assembler.DTOTOPOQuestionnaireResult(req)))
}
