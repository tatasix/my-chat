package questionnaire

import (
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/service/questionnaire"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveQuestionnaireLogic struct {
	logx.Logger
	ctx                  context.Context
	svcCtx               *svc.ServiceContext
	QuestionnaireService *questionnaire.QuestionnaireService
}

func NewSaveQuestionnaireLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveQuestionnaireLogic {
	return &SaveQuestionnaireLogic{
		Logger:               logx.WithContext(ctx),
		ctx:                  ctx,
		svcCtx:               svcCtx,
		QuestionnaireService: questionnaire.NewQuestionnaireService(ctx, svcCtx),
	}
}

func (l *SaveQuestionnaireLogic) SaveQuestionnaire(req *types.SaveQuestionnaireResponseRequest) (resp *types.Response, err error) {
	return assembler.Return(l.QuestionnaireService.SaveQuestionnaire(assembler.DTOTOPOQuestionnaireResponse(req), req.Status))
}
