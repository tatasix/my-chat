package questionnaire

import (
	"chat/common/util"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/service/assembler"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"math"
)

type SummationHandler struct {
	logger             logx.Logger
	ctx                context.Context
	svcCtx             *svc.ServiceContext
	responseRepository *repository.QuestionnaireResponseRepository
	resultRepository   *repository.QuestionnaireResultRepository
}

func (s *SummationHandler) initRepository() {
	s.responseRepository = repository.NewQuestionnaireResponseRepository(s.ctx, s.svcCtx)
	s.resultRepository = repository.NewQuestionnaireResultRepository(s.ctx, s.svcCtx)
	return
}

func (s *SummationHandler) Handle(customerConfig *model.CustomerConfig, config model.CustomerConfigConfig, relationId int64, user, openKfID string, channel chan string, baseInfo chan any) (resp *types.GenerateResultResponse, err error) {
	s.initRepository()
	resp = &types.GenerateResultResponse{}
	if len(config.Result) > 0 {
		for _, v := range config.Result {
			//if v.Type == repository.QuestionnaireResultSummary {
			//	s.summary(config, v, relationId, customerConfig.KfId, user)
			//} else {
			resp, err = s.Deal(config, v, relationId, baseInfo)
			//}
		}
	}
	return
}

func (s *SummationHandler) Deal(config model.CustomerConfigConfig, configResult model.CustomerConfigConfigResultType, relationId int64, baseInfo chan any) (resp *types.GenerateResultResponse, err error) {
	//获取所有的数据

	results, err := s.resultRepository.GetByRelationId(relationId)
	if err != nil {
		s.logger.Errorf("SummationHandler Handle GetByRelationId err %+v", err.Error())
		return
	}
	if results == nil || len(results) <= 0 {
		s.logger.Errorf("SummationHandler Handle GetByRelationId result is empty")
		return
	}
	result := results[0]

	//汇总记录

	sum, err := s.responseRepository.GetSum(relationId, configResult.Type)
	if err != nil {
		s.logger.Errorf("SummationHandler Handle GetSum err %+v", err.Error())
		return
	}
	//获取汇率
	rate := s.GetRate(config)
	score1 := float64(rate * float32(sum))
	score := math.Round(score1*1000) / 1000
	result.Status = repository.QuestionnaireResultStatusEnd
	result.Result = s.GetResult(configResult, score)
	result.Score = util.Float64ToSql(score)
	result.ScoreType = configResult.Type
	err = s.resultRepository.Update(result)
	lastResult := assembler.POTODTOGenerateResult(result)
	lastResult.IsEnd = true
	baseInfo <- lastResult
	return &lastResult, nil
}

func (s *SummationHandler) summary(config model.CustomerConfigConfig, configResult model.CustomerConfigConfigResultType, relationId int64, kfId, user string) {

	sum, err := s.responseRepository.GetSum(relationId, 0)
	if err != nil {
		s.logger.Errorf("SummationHandler Handle GetSum err %+v", err.Error())
		return
	}
	//获取汇率
	rate := s.GetRate(config)
	score1 := float64(rate * float32(sum))
	score := math.Round(score1*1000) / 1000
	// 获取result表看看有咩有数据
	resultModel := &model.QuestionnaireResult{
		RelationId: relationId,
		User:       user,
		OpenKfId:   kfId,
		ScoreType:  configResult.Type,
	}

	resultModel.Status = repository.QuestionnaireResultStatusEnd
	resultModel.Result = s.GetResult(configResult, score)
	resultModel.Score = util.Float64ToSql(score)
	err = s.resultRepository.Insert(resultModel)

	return
}

func (s *SummationHandler) GetRate(config model.CustomerConfigConfig) (rate float32) {
	//获取汇率
	switch config.Mod.Type {
	case 1:
		return config.Mod.Rate
	}
	return
}

func (s *SummationHandler) GetResult(config model.CustomerConfigConfigResultType, score float64) (result string) {
	for _, v := range config.Result {
		if v.Min <= score && v.Max >= score {
			return v.Res
		}
	}
	return
}

func (s *SummationHandler) GetReturn(customerConfig *model.CustomerConfig, user string) (results *types.GetQuestionnaireResultResponse, err error) {
	s.initRepository()
	results = &types.GetQuestionnaireResultResponse{}
	customerPo, err := s.resultRepository.GetLastResult(customerConfig.KfId, user)
	if err != nil {
		return
	}
	if customerPo != nil && customerPo.RelationId > 0 {
		result, err1 := s.resultRepository.GetByRelationId(customerPo.RelationId)
		if err1 != nil {
			return
		}
		if len(result) >= 0 {
			var resultDto []types.QuestionnaireResult
			for _, v := range result {
				resultDto = append(resultDto, assembler.POTODTOGetQuestionnaireResult(v))
			}
			results.Result = resultDto
		}
	}
	return
}
