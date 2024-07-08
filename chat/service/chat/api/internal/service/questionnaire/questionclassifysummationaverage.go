package questionnaire

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/service/assembler"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/model"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"math"
)

type QuestionClassifySummationAverageHandler struct {
	logger               logx.Logger
	ctx                  context.Context
	svcCtx               *svc.ServiceContext
	responseRepository   *repository.QuestionnaireResponseRepository
	resultRepository     *repository.QuestionnaireResultRepository
	resultMbtiRepository *repository.QuestionnaireResultMbtiRepository
}

func (s *QuestionClassifySummationAverageHandler) initRepository() {
	s.responseRepository = repository.NewQuestionnaireResponseRepository(s.ctx, s.svcCtx)
	s.resultRepository = repository.NewQuestionnaireResultRepository(s.ctx, s.svcCtx)
	s.resultMbtiRepository = repository.NewQuestionnaireResultMbtiRepository(s.ctx, s.svcCtx)

	return
}

func (s *QuestionClassifySummationAverageHandler) Handle(customerConfig *model.CustomerConfig, config model.CustomerConfigConfig, relationId int64, user, openKfID string, channel chan string, baseInfo chan any) (resp *types.GenerateResultResponse, err error) {
	s.initRepository()
	resp = &types.GenerateResultResponse{}
	exist, err := s.resultRepository.GetByRelationIdAndTypes(relationId, []int64{0, repository.QuestionnaireResultStatusEnd})
	if err != nil {
		s.logger.Errorf("QuestionClassifyHandler summary GetByRelationIdAndType err %+v", err.Error())
		return
	}
	if exist == nil || exist.Id <= 0 {
		s.logger.Errorf("QuestionClassifyHandler summary result not exist")
		return
	}

	if exist.OpenKfId != customerConfig.KfId {
		err = util.ReturnError(xerr.ParamError)
		return
	}
	var scores []model.QuestionnaireResultMbtiScore
	questionnaireResultMbti := &model.QuestionnaireResultMbti{
		RelationId: relationId,
		User:       user,
		OpenKfId:   openKfID,
	}
	if len(config.Result) > 0 {
		for _, v := range config.Result {
			score, err1 := s.getScore(config, v, relationId)
			if err1 != nil {
				return nil, err1
			}
			scores = append(scores, score)
		}
	}
	if len(scores) > 0 {
		scoreJson, _ := json.Marshal(scores)
		questionnaireResultMbti.Score = string(scoreJson)
	}

	existMbti, err := s.resultMbtiRepository.GetByRelationId(relationId)
	if err != nil {
		s.logger.Errorf("QuestionClassifyHandler summary resultMbtiRepository GetByRelationId err %+v", err.Error())
		return
	}
	if existMbti != nil && existMbti.Id > 0 {
		//存在
		questionnaireResultMbti.Id = existMbti.Id
		err = s.resultMbtiRepository.Update(questionnaireResultMbti)
	} else {
		lastId, _ := s.resultMbtiRepository.Insert(questionnaireResultMbti)
		questionnaireResultMbti.Id = lastId
		exist.Status = repository.QuestionnaireResultStatusEnd
		exist.ScoreType = repository.QuestionnaireResultSummary
		err = s.resultRepository.Update(exist)
	}
	questionnaireResultMbti.Result = assembler.POTODTOGetResultToString(scores)
	lastResult := assembler.POTODTOGenerateResultByMbti(questionnaireResultMbti, exist, scores)
	lastResult.IsEnd = true
	resp = &lastResult
	baseInfo <- lastResult

	return
}

func (s *QuestionClassifySummationAverageHandler) getScore(config model.CustomerConfigConfig, configResult model.CustomerConfigConfigResultType, relationId int64) (resp model.QuestionnaireResultMbtiScore, err error) {

	//汇总记录
	sum, err := s.responseRepository.GetSum(relationId, configResult.Type)
	if err != nil {
		s.logger.Errorf("SummationHandler Handle GetSum err %+v", err.Error())
		return
	}
	//获取汇率
	rate := s.GetRate(config)
	score1 := float64(rate*float32(sum)) / configResult.Total
	score := math.Floor(score1 + 0.5)

	resp.Name = configResult.Name
	resp.Description = s.GetResult(configResult, score)
	resp.Score = score
	resp.Type = configResult.Type

	return
}
func (s *QuestionClassifySummationAverageHandler) GetRate(config model.CustomerConfigConfig) (rate float32) {
	//获取汇率
	switch config.Mod.Type {
	case 1:
		return config.Mod.Rate
	}
	return
}

func (s *QuestionClassifySummationAverageHandler) GetResult(config model.CustomerConfigConfigResultType, score float64) (result string) {
	for _, v := range config.Result {
		if v.Min <= score && v.Max >= score {
			return v.Res
		}
	}
	return
}

func (s *QuestionClassifySummationAverageHandler) GetReturn(customerConfig *model.CustomerConfig, user string) (response *types.GetQuestionnaireResultResponse, err error) {
	s.initRepository()

	response = &types.GetQuestionnaireResultResponse{}
	var results []types.QuestionnaireResult
	//先查看是否有进行中的
	inProgress, err := s.resultRepository.GetOne(0, -1, customerConfig.KfId, user)
	if err != nil {
		return
	}

	if inProgress != nil && inProgress.Id > 0 {
		results = append(results, assembler.POTODTOGetQuestionnaireResult(inProgress))
		response.Result = results
		return
	}
	//没有进行中的，那就需要获取最近的一次成功的结果
	completed, err := s.resultRepository.GetLastResult(customerConfig.KfId, user)
	if err != nil {
		return
	}

	if completed != nil && completed.RelationId > 0 {
		//从mbti表获取结果
		mbtiResult, err1 := s.resultMbtiRepository.GetByRelationId(completed.RelationId)
		if err1 != nil {
			return
		}
		if mbtiResult != nil && mbtiResult.Id > 0 {
			var scoreMbti []model.QuestionnaireResultMbtiScore
			_ = json.Unmarshal([]byte(mbtiResult.Score), &scoreMbti)
			mbtiResult.Result = assembler.POTODTOGetResultToString(scoreMbti)
			results = append(results, assembler.POTODTOGetQuestionnaireResultUseMbti(completed, mbtiResult))
			response.Result = results
			return
		}
	}
	return
}
