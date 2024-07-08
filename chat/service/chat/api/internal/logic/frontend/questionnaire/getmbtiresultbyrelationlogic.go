package questionnaire

import (
	"chat/common/util"
	"chat/service/chat/api/internal/service/questionnaire"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMbtiResultByRelationLogic struct {
	logx.Logger
	ctx                  context.Context
	svcCtx               *svc.ServiceContext
	QuestionnaireService *questionnaire.QuestionnaireService
}

func NewGetMbtiResultByRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMbtiResultByRelationLogic {
	return &GetMbtiResultByRelationLogic{
		Logger:               logx.WithContext(ctx),
		ctx:                  ctx,
		svcCtx:               svcCtx,
		QuestionnaireService: questionnaire.NewQuestionnaireService(ctx, svcCtx),
	}
}

func (l *GetMbtiResultByRelationLogic) GetMbtiResultByRelation(req *types.GetMbtiResultByRelationRequest) (resp *types.GetMbtiResultByRelationResponse, err error) {
	return l.QuestionnaireService.GetMbtiResultByRelationId(util.ConvertToInt64(req.RelationId))
}
