package questionnaire

import (
	"chat/common/util"
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/service/questionnaire"
	"chat/service/chat/model"
	"context"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuestionnaireLogic struct {
	logx.Logger
	ctx                  context.Context
	svcCtx               *svc.ServiceContext
	QuestionnaireService *questionnaire.QuestionnaireService
}

func NewGetQuestionnaireLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionnaireLogic {
	return &GetQuestionnaireLogic{
		Logger:               logx.WithContext(ctx),
		ctx:                  ctx,
		svcCtx:               svcCtx,
		QuestionnaireService: questionnaire.NewQuestionnaireService(ctx, svcCtx),
	}
}

func (l *GetQuestionnaireLogic) GetQuestionnaire(req *types.GetQuestionnaireRequest) (resp *types.GetQuestionnaireResponse, err error) {
	resp = &types.GetQuestionnaireResponse{}
	//问题
	question, err := l.QuestionnaireService.GetQuestion(req.OpenKfId)
	if nil != err {
		return
	}
	//结果
	resultResponse, err := l.QuestionnaireService.GetResult(req.OpenKfId, req.User)
	if nil != err {
		return
	}
	var result []types.QuestionnaireResult
	var response []*model.QuestionnaireResponse
	if resultResponse != nil && len(resultResponse.Result) > 0 {
		result = resultResponse.Result
		relationId := util.ConvertToInt64(result[0].RelationId)
		//回答
		response, err = l.QuestionnaireService.GetResponse(req.OpenKfId, req.User, relationId)
		if nil != err {
			return
		}
	}

	return assembler.POTODTOGetQuestionnaire(question, result, response), nil
}
