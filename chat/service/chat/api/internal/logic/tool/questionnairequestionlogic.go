package tool

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/service/questionnaire"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

type QuestionnaireQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuestionnaireQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuestionnaireQuestionLogic {
	return &QuestionnaireQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Question struct {
	QuestionNumber int64    `json:"QuestionNumber"`
	Title          string   `json:"Title"`
	Type           int64    `json:"type"`
	ScoreType      int64    `json:"score_letter"`
	Answers        []Answer `json:"answers"`
}

type Answer struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Score   int64  `json:"score"`
}

func (l *QuestionnaireQuestionLogic) QuestionnaireQuestion() (resp *types.Response, err error) {
	if l.svcCtx.Config.Env != "dev" {
		err = util.ReturnError(xerr.SystemError)
		return
	}
	baseInfo := vars.GetTheBaseInfo()
	for _, vv := range baseInfo {
		if vv.Info == "" || vv.OpenKfId == "" {
			continue
		}
		jsonStr := vv.Info
		kfId := vv.OpenKfId
		questions := make([]Question, 0)
		_ = json.Unmarshal([]byte(jsonStr), &questions)
		for _, v := range questions {
			answer, _ := json.Marshal(v.Answers)
			_, _ = questionnaire.NewQuestionnaireService(l.ctx, l.svcCtx).InsertQuestionnaireQuestion(&model.QuestionnaireQuestion{
				OpenKfId:     kfId,
				QuestionType: v.Type,
				ScoreType:    v.ScoreType,
				Question:     v.Title,
				Sort:         v.QuestionNumber,
				Option:       string(answer),
			})
		}
	}

	return
}
