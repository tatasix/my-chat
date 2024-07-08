package questionnaire

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/service/assembler"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"math"
)

type QuestionClassifyHandler struct {
	logger               logx.Logger
	ctx                  context.Context
	svcCtx               *svc.ServiceContext
	responseRepository   *repository.QuestionnaireResponseRepository
	resultRepository     *repository.QuestionnaireResultRepository
	resultMbtiRepository *repository.QuestionnaireResultMbtiRepository
	wechatUserRepository *repository.WechatUserRepository
}

func (s *QuestionClassifyHandler) initRepository() {
	s.responseRepository = repository.NewQuestionnaireResponseRepository(s.ctx, s.svcCtx)
	s.resultRepository = repository.NewQuestionnaireResultRepository(s.ctx, s.svcCtx)
	s.resultMbtiRepository = repository.NewQuestionnaireResultMbtiRepository(s.ctx, s.svcCtx)
	s.wechatUserRepository = repository.NewWechatUserRepository(s.ctx, s.svcCtx)
	return
}

func (s *QuestionClassifyHandler) Handle(customerConfig *model.CustomerConfig, config model.CustomerConfigConfig, relationId int64, user, openKfID string, channel chan string, baseInfo chan any) (resp *types.GenerateResultResponse, err error) {
	s.initRepository()
	resp = &types.GenerateResultResponse{}
	if len(config.Result) > 0 {
		for i := 0; i < len(config.Result)-1; i++ {
			v := config.Result[i]
			if v.Type%2 == 0 && v.Type != repository.QuestionnaireResultSummary {
				continue
			}
			err = s.deal(config, v, relationId, customerConfig.KfId, user)
			if err != nil {
				return resp, err
			}
		}
		resp, err = s.summary(customerConfig, config, relationId, user, openKfID, channel, baseInfo)
		if err != nil {
			return resp, err
		}
	}
	return
}

func (s *QuestionClassifyHandler) deal(config model.CustomerConfigConfig, configResult model.CustomerConfigConfigResultType, relationId int64, kfId, user string) (err error) {
	total := config.Total
	allResponse, err := s.responseRepository.GetAll("", "", relationId)
	if err != nil {
		return
	}
	if len(allResponse) != int(total) {
		err = util.ReturnError(xerr.QuestionnaireResponseError)
		return
	}
	baseType := configResult.Type
	evenType := configResult.Type + 1
	//处理基数
	sumBase, err := s.responseRepository.GetSum(relationId, baseType)
	if err != nil {
		s.logger.Errorf("QuestionClassifyHandler Handle GetSum err %+v", err.Error())
		return
	}
	//获取汇率
	rate := s.GetRate(config)
	scoreBase1 := float64(rate * float32(sumBase))
	//处理偶数
	sumEven, err := s.responseRepository.GetSum(relationId, evenType)
	if err != nil {
		s.logger.Errorf("QuestionClassifyHandler Handle GetSum err %+v", err.Error())
		return
	}
	//获取汇率
	scoreEven1 := float64(rate * float32(sumEven))

	scoreBase, scoreEven := s.getScore(configResult, scoreBase1, scoreEven1)
	scoreBaseFinal := util.Float64ToSql(math.Floor(scoreBase*100+0.5) / 100)
	scoreEvenFinal := util.Float64ToSql(math.Floor(scoreEven*100+0.5) / 100)
	//处理基数
	resultBaseFinal := s.getResultByType(config, baseType)
	result, err := s.resultRepository.GetByRelationIdAndType(relationId, baseType)
	if err != nil {
		s.logger.Errorf("QuestionClassifyHandler Handle GetByRelationIdAndType err %+v", err.Error())
		return
	}
	if result == nil || result.Id <= 0 {
		// 新增
		resultModel := &model.QuestionnaireResult{
			RelationId: relationId,
			User:       user,
			OpenKfId:   kfId,
			ScoreType:  baseType,
			Status:     repository.QuestionnaireResultStatusEnd,
			Result:     resultBaseFinal,
			Score:      scoreBaseFinal,
		}
		if err = s.resultRepository.Insert(resultModel); err != nil {
			return
		}

	} else {
		result.Status = repository.QuestionnaireResultStatusEnd
		result.Result = resultBaseFinal
		result.Score = scoreBaseFinal
		//修改
		err = s.resultRepository.Update(result)
	}

	//处理偶数
	resultEvenFinal := s.getResultByType(config, evenType)
	resultEven, err := s.resultRepository.GetByRelationIdAndType(relationId, evenType)
	if err != nil {
		s.logger.Errorf("QuestionClassifyHandler Handle GetByRelationIdAndType err %+v", err.Error())
		return
	}
	if resultEven == nil || resultEven.Id <= 0 {
		// 新增
		resultModel := &model.QuestionnaireResult{
			RelationId: relationId,
			User:       user,
			OpenKfId:   kfId,
			ScoreType:  evenType,
			Status:     repository.QuestionnaireResultStatusEnd,
			Result:     resultEvenFinal,
			Score:      scoreEvenFinal,
		}
		if err = s.resultRepository.Insert(resultModel); err != nil {
			return
		}

	} else {
		resultEven.Status = repository.QuestionnaireResultStatusEnd
		resultEven.Result = resultEvenFinal
		resultEven.Score = scoreEvenFinal
		//修改
		err = s.resultRepository.Update(resultEven)
	}
	return
}

func (s *QuestionClassifyHandler) summary(customerConfig *model.CustomerConfig, config model.CustomerConfigConfig, relationId int64, user, openKfID string, channel chan string, baseInfo chan any) (resp *types.GenerateResultResponse, err error) {

	//要改为最终状态

	scoreType := []int64{0, repository.QuestionnaireResultStatusEnd}
	exist, err := s.resultRepository.GetByRelationIdAndTypes(relationId, scoreType)
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
	allResult, err := s.resultRepository.GetByRelationId(relationId)
	if err != nil {
		s.logger.Errorf("QuestionClassifyHandler summary GetByRelationId err %+v", err.Error())
		return
	}
	if allResult == nil || len(allResult) <= 0 {
		s.logger.Errorf("QuestionClassifyHandler summary result not exist")
		return
	}
	if len(allResult) != 9 {
		s.logger.Errorf("QuestionClassifyHandler summary deal result error")
		return
	}

	modelQuestionnaireResultMbti, mbitScores := s.dealAllResult(allResult, config)
	//可以输出baseInfo

	existMbti, err := s.resultMbtiRepository.GetByRelationId(relationId)
	if err != nil {
		s.logger.Errorf("QuestionClassifyHandler summary resultMbtiRepository GetByRelationId err %+v", err.Error())
		return
	}
	if existMbti != nil && existMbti.Id > 0 {
		//存在
		modelQuestionnaireResultMbti.Id = existMbti.Id
		err = s.resultMbtiRepository.Update(modelQuestionnaireResultMbti)
	} else {
		lastId, _ := s.resultMbtiRepository.Insert(modelQuestionnaireResultMbti)
		modelQuestionnaireResultMbti.Id = lastId
	}
	if err != nil {
		s.logger.Errorf("QuestionClassifyHandler summary resultMbtiRepository Insert err %+v", err.Error())
		return
	}
	exist.Status = repository.QuestionnaireResultStatusEnd
	exist.ScoreType = repository.QuestionnaireResultSummary
	err = s.resultRepository.Update(exist)
	lastResult := assembler.POTODTOGenerateResultByMbti(modelQuestionnaireResultMbti, exist, mbitScores)
	baseInfo <- lastResult
	if openKfID == vars.ChatMbti48 {
		//48题 需要更新用户的mbti结果
		if err = s.wechatUserRepository.UpdateByUser(user, &model.WechatUser{
			Mbti: modelQuestionnaireResultMbti.NameType,
		}); err != nil {
			return
		}
	}
	// 调用chatgpt解析最终结果
	chatMessage, err := s.dealContentByAI(customerConfig, modelQuestionnaireResultMbti, user, channel, modelQuestionnaireResultMbti.NameType, modelQuestionnaireResultMbti.ScoreResult)
	if err != nil {
		return
	}
	if chatMessage != "" {
		modelQuestionnaireResultMbti.Content = chatMessage
		_ = s.resultMbtiRepository.Update(modelQuestionnaireResultMbti)
	}

	lastResult.Content = chatMessage
	lastResult.Message = chatMessage
	lastResult.IsEnd = true
	return &lastResult, nil
}

func (s *QuestionClassifyHandler) GetRate(config model.CustomerConfigConfig) (rate float32) {
	//获取汇率
	switch config.Mod.Type {
	case 1:
		return config.Mod.Rate
	}
	return
}

// getScore
// 例如，对于E/I倾向：
// 1. 将所有正面E的得分加起来，然后除以正面E题目的数量，得到E的平均分。
// 2. 将所有反面I的得分加起来，然后除以反面I题目的数量，得到I的平均分。
// 3. 比较E和I的平均分，得分较高的方向确定了该维度的倾向。
//
// 计算每种倾向的百分比
// **计算百分比**：
//
//	以E/I为例，E得分占总得分的百分比为：E平均分/（E平均分+I平均分） \times 100% \)
//
// 对于其他维度，使用相同的计算方法。
// 最终得出结论如：E 51% N-63% T-54% J-60%
func (s *QuestionClassifyHandler) getScore(config model.CustomerConfigConfigResultType, scoreBase, scoreEven float64) (baseFinal, evenFinal float64) {
	base := scoreBase / config.Total
	even := scoreEven / config.Total

	baseFinal = base / (base + even)
	evenFinal = even / (base + even)
	if baseFinal == evenFinal {
		baseFinal = baseFinal - 0.01
		evenFinal = evenFinal + 0.01
	}
	return
}

func (s *QuestionClassifyHandler) getResultByType(config model.CustomerConfigConfig, configResultType int64) string {
	for _, v := range config.Result {
		if v.Type == configResultType {
			return v.Name
		}
	}
	return ""
}

func (s *QuestionClassifyHandler) dealAllResult(questionnaireResult []*model.QuestionnaireResult, configResult model.CustomerConfigConfig) (questionnaireResultMbti *model.QuestionnaireResultMbti, mbitScores []model.QuestionnaireResultMbtiScore) {
	questionnaireResultMbti = &model.QuestionnaireResultMbti{
		RelationId: questionnaireResult[0].RelationId,
		User:       questionnaireResult[0].User,
		OpenKfId:   questionnaireResult[0].OpenKfId,
	}
	var customerConfigConfigResult []model.CustomerConfigConfigResult
	for _, vv := range configResult.Result {
		for _, v := range questionnaireResult {
			if v.ScoreType == vv.Type && vv.Type != repository.QuestionnaireResultSummary {
				mbitScores = append(mbitScores, model.QuestionnaireResultMbtiScore{
					Type:        vv.Type,
					Score:       util.SqlToFloat64(v.Score),
					Name:        vv.Name,
					Description: vv.Description,
				})
				break
			}
		}
		if vv.Type == 100 {
			customerConfigConfigResult = vv.Result
		}
	}
	scoreJson, _ := json.Marshal(mbitScores)
	questionnaireResultMbti.Score = string(scoreJson)
	s.getNameType(questionnaireResult, customerConfigConfigResult, questionnaireResultMbti)
	s.getLastResult(questionnaireResultMbti, mbitScores)
	return
}

func (s *QuestionClassifyHandler) getNameType(questionnaireResult []*model.QuestionnaireResult, customerConfigConfigResult []model.CustomerConfigConfigResult, questionnaireResultMbti *model.QuestionnaireResultMbti) {
	mbtiScore := make(map[int64]float64, len(questionnaireResult)-1)
	for _, v := range questionnaireResult {
		if v.ScoreType == repository.QuestionnaireResultSummary {
			continue
		}
		mbtiScore[v.ScoreType] = util.SqlToFloat64(v.Score)
	}
	var nameType string

	if mbtiScore[1] > mbtiScore[2] {
		nameType = nameType + "E"
	} else {
		nameType = nameType + "I"
	}

	if mbtiScore[3] > mbtiScore[4] {
		nameType = nameType + "S"
	} else {
		nameType = nameType + "N"
	}

	if mbtiScore[5] > mbtiScore[6] {
		nameType = nameType + "T"
	} else {
		nameType = nameType + "F"
	}

	if mbtiScore[7] > mbtiScore[8] {
		nameType = nameType + "J"
	} else {
		nameType = nameType + "P"
	}

	for _, vv := range customerConfigConfigResult {
		if vv.NameType == nameType {
			questionnaireResultMbti.Name = vv.Name
			questionnaireResultMbti.NameType = vv.NameType
			questionnaireResultMbti.Nickname = vv.Nickname
			questionnaireResultMbti.SimpleResult = vv.SimpleResult
			questionnaireResultMbti.Result = vv.Res
			questionnaireResultMbti.Label = vv.Label
			questionnaireResultMbti.Partner = vv.Partner
			questionnaireResultMbti.Mate = vv.Mate
			questionnaireResultMbti.Content = vv.Content
			questionnaireResultMbti.PcImage = vv.PcImage
			questionnaireResultMbti.H5Image = vv.H5Image
			questionnaireResultMbti.Color = vv.Color
		}
	}
	return
}
func (s *QuestionClassifyHandler) GetReturn(customerConfig *model.CustomerConfig, user string) (response *types.GetQuestionnaireResultResponse, err error) {
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
			results = append(results, assembler.POTODTOGetQuestionnaireResultUseMbti(completed, mbtiResult))
			response.Result = results
			return
		}
	}
	return
}

func (s *QuestionClassifyHandler) dealContentByAI(config *model.CustomerConfig, mbtiModel *model.QuestionnaireResultMbti, user string, channel chan string, nameType, result string) (chatMessage string, err error) {
	ctx := context.Background()
	userInfo, err := repository.NewWechatUserRepository(ctx, s.svcCtx).GetByUser(user)
	if err != nil {
		return
	}
	if userInfo == nil || userInfo.Id <= 0 {
		err = util.ReturnError(xerr.RecordNotFound)
		return
	}

	aiChatService := service.NewAiChatService(ctx, s.svcCtx)
	contextInfo := fmt.Sprintf("用户星座：%s 用户性别：%s 用户人格类型：%s 各人格类型占比：%s", userInfo.Constellation, util.GetGender(userInfo.Gender), nameType, result)
	summary := aiChatService.SetSystemAndUserPrompt(config.Prompt, contextInfo)

	chatMessage, err = aiChatService.CustomerChatStream(config, summary, channel, user)

	if err != nil {
		s.logger.Errorf("QuestionClassifyHandler dealContentByAI Chat err %+v", err.Error())
		return
	}
	s.logger.Infof("chat message %s", chatMessage)
	mbtiModel.Content = chatMessage
	if err = repository.NewQuestionnaireResultMbtiRepository(ctx, s.svcCtx).Update(mbtiModel); err != nil {
		s.logger.Errorf("QuestionClassifyHandler dealContentByAI MbtiUpdate err %+v", err.Error())
	}

	return
}

func (s *QuestionClassifyHandler) getLastResult(questionnaireResultMbti *model.QuestionnaireResultMbti, mbtiScores []model.QuestionnaireResultMbtiScore) {
	if len(mbtiScores) <= 0 {
		return
	}
	var result string
	for _, v := range []byte(questionnaireResultMbti.NameType) {
		for _, vv := range mbtiScores {
			if string(v) == vv.Name {
				result = result + vv.Name + " " + fmt.Sprintf("%.0f%%", vv.Score*100) + ","
			}
		}
	}

	questionnaireResultMbti.ScoreResult = result[:len(result)-1]
	return
}
