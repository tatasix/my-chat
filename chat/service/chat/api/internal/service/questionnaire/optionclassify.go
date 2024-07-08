package questionnaire

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/service/assembler"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"math"
	"reflect"
)

type score struct {
	E float64 `json:"e"`
	I float64 `json:"i"`
	S float64 `json:"s"`
	N float64 `json:"n"`
	F float64 `json:"f"`
	T float64 `json:"t"`
	J float64 `json:"j"`
	P float64 `json:"p"`
}
type mbtiResultScore struct {
	Type        int     `json:"type"`
	Score       float64 `json:"score"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}
type OptionClassifyHandler struct {
	logger               logx.Logger
	ctx                  context.Context
	svcCtx               *svc.ServiceContext
	responseRepository   *repository.QuestionnaireResponseRepository
	resultRepository     *repository.QuestionnaireResultRepository
	resultMbtiRepository *repository.QuestionnaireResultMbtiRepository
}

func (s *OptionClassifyHandler) initRepository() {
	s.responseRepository = repository.NewQuestionnaireResponseRepository(s.ctx, s.svcCtx)
	s.resultRepository = repository.NewQuestionnaireResultRepository(s.ctx, s.svcCtx)
	s.resultMbtiRepository = repository.NewQuestionnaireResultMbtiRepository(s.ctx, s.svcCtx)
	return
}

func (s *OptionClassifyHandler) Handle(customerConfig *model.CustomerConfig, config model.CustomerConfigConfig, relationId int64, user, openKfID string, channel chan string, baseInfo chan any) (resp *types.GenerateResultResponse, err error) {
	s.initRepository()
	resp = &types.GenerateResultResponse{}
	result, err := s.resultRepository.GetByRelationIdAndType(relationId, 0)
	if err != nil {
		s.logger.Errorf("OptionclassifyHandler Handle GetByRelationId err %+v", err.Error())
		return
	}
	if result == nil || result.Id <= 0 {
		s.logger.Errorf("OptionclassifyHandler Handle GetByRelationId result is empty")
		return
	}

	total := config.Total
	allResponse, err := s.responseRepository.GetAll("", "", relationId)
	if err != nil {
		return
	}
	if len(allResponse) != int(total) {
		err = util.ReturnError(xerr.QuestionnaireResponseError)
		return
	}

	mbti, mbtiPercent, scoreRes, err := s.summary(relationId)
	if err != nil {
		s.logger.Errorf("OptionclassifyHandler Handle GetMBTI err %+v", err.Error())
		return
	}
	//percent, err := s.getMbtiPercent(mbti)
	//if err != nil {
	//	return
	//}
	var content string
	questionnaireResultMbti := &model.QuestionnaireResultMbti{}
	for _, vv := range config.Mbti93Result {
		if mbti == vv.NameType {
			contentByte, _ := json.Marshal(vv)
			content = string(contentByte)
			questionnaireResultMbti.Name = vv.Name
			questionnaireResultMbti.NameType = vv.NameType
			questionnaireResultMbti.Nickname = vv.Nickname
			questionnaireResultMbti.SimpleResult = vv.Res1
			questionnaireResultMbti.Result = vv.Res
			questionnaireResultMbti.Label = vv.Label
			questionnaireResultMbti.Partner = vv.Partner
			questionnaireResultMbti.Mate = vv.Mate
			questionnaireResultMbti.PcImage = vv.PcImage
			questionnaireResultMbti.H5Image = vv.H5Image
			questionnaireResultMbti.Color = vv.Color
			break
		}
	}
	result.Status = repository.QuestionnaireResultStatusEnd
	result.Result = content
	result.Mbti = mbti
	err = s.resultRepository.Update(result)

	//insert into table resultMbti
	questionnaireResultMbti.Content = ""
	questionnaireResultMbti.RelationId = relationId
	questionnaireResultMbti.User = user
	questionnaireResultMbti.OpenKfId = openKfID
	questionnaireResultMbti.Score = mbtiPercent
	questionnaireResultMbti.ScoreResult = scoreRes

	existMbti, err := s.resultMbtiRepository.GetByRelationId(relationId)
	if err != nil {
		s.logger.Errorf("QuestionClassifyHandler summary resultMbtiRepository GetByRelationId err %+v", err.Error())
		return
	}

	if existMbti != nil && existMbti.Id > 0 {
		questionnaireResultMbti.Id = existMbti.Id
		err = s.resultMbtiRepository.Update(questionnaireResultMbti)
	} else {
		lastId, _ := s.resultMbtiRepository.Insert(questionnaireResultMbti)
		questionnaireResultMbti.Id = lastId
	}

	var mbtiScore []model.QuestionnaireResultMbtiScore
	_ = json.Unmarshal([]byte(mbtiPercent), &mbtiScore)
	lastResult := assembler.POTODTOGenerateResultByMbti(questionnaireResultMbti, result, mbtiScore)
	baseInfo <- lastResult

	// 调用chatgpt解析最终结果
	chatMessage, err := s.dealContentByAI(customerConfig, questionnaireResultMbti, user, channel, mbti, scoreRes)
	if err != nil {
		return
	}
	if chatMessage != "" {
		questionnaireResultMbti.Content = chatMessage
		lastResult.Content = chatMessage
	}

	err = s.resultMbtiRepository.Update(questionnaireResultMbti)

	return &lastResult, nil
}

func (s *OptionClassifyHandler) summary(relationId int64) (mbti string, mbtiString, scoreRes string, err error) {
	questionnaireResponsePo, err := s.responseRepository.GetAll("", "", relationId)

	if err != nil {
		s.logger.Errorf("OptionclassifyHandler Handle summary err %+v", err.Error())
		return
	}
	var E, I, S, N, T, F, J, P int8
	var MBTI, EI, SN, TF, JP string
	if len(questionnaireResponsePo) > 0 {
		for _, v := range questionnaireResponsePo {
			if v.Mbti == "" {
				continue
			}
			switch v.Mbti {
			case "E":
				E++
			case "I":
				I++
			case "S":
				S++
			case "N":
				N++
			case "T":
				T++
			case "F":
				F++
			case "J":
				J++
			case "P":
				P++
			default:
				s.logger.Errorf("OptionclassifyHandler Handle summary err %+v", "非法 mbti 类型")
				err = errors.New("非法 mbti 类型")
				return
			}
		}
	}

	EI = s.getMaxValue("E", "I", E, I)
	SN = s.getMaxValue("S", "N", S, N)
	TF = s.getMaxValue("T", "F", T, F)
	JP = s.getMaxValue("J", "P", J, P)

	MBTI = EI + SN + TF + JP

	EPercent, IPercent := s.calculatePercentages(E, I)
	SPercent, NPercent := s.calculatePercentages(S, N)
	TPercent, FPercent := s.calculatePercentages(T, F)
	JPercent, PPercent := s.calculatePercentages(J, P)

	scoreMap := make(map[string]string)
	scoreMap["E"] = "外向"
	scoreMap["I"] = "内向"
	scoreMap["S"] = "实感"
	scoreMap["N"] = "直觉"
	scoreMap["F"] = "思考"
	scoreMap["T"] = "情感"
	scoreMap["J"] = "判断"
	scoreMap["P"] = "直觉"

	var mbtiScore []mbtiResultScore
	p := score{EPercent, IPercent, SPercent, NPercent, FPercent, TPercent, JPercent, PPercent}

	//score result
	v := reflect.ValueOf(p)
	t := reflect.TypeOf(p)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		key := t.Field(i).Name
		value := field.Float()
		scoreRes += key + " " + fmt.Sprintf("%.0f%%", value) + ","
		//table mbitResult score
		mbtiScore = append(mbtiScore, mbtiResultScore{
			Type:        i,
			Score:       value / 100,
			Name:        key,
			Description: scoreMap[key],
		})
	}
	if scoreRes != "" {
		scoreRes = scoreRes[:len(scoreRes)-1]
	}
	jsonBytes, err := json.Marshal(mbtiScore)
	if err != nil {
		return
	}
	mbtiScoreString := string(jsonBytes)
	return MBTI, mbtiScoreString, scoreRes, nil
}

func (s *OptionClassifyHandler) getMaxValue(first, second string, firstNum, sercondNum int8) string {
	if firstNum <= sercondNum {
		return second
	} else {
		return first
	}
}

func (s *OptionClassifyHandler) calculatePercentages(a, b int8) (float64, float64) {
	if (a == 0 && b == 0) || a == b {
		return 49, 51
	}
	total := float64(a + b)
	aPercentage := math.Round((float64(a) / total) * 100)
	bPercentage := math.Round((float64(b) / total) * 100)
	return aPercentage, bPercentage
}

// 占人群百分比
func (s *OptionClassifyHandler) getMbtiPercent(mbti string) (percent float64, err error) {
	num, total, err := s.resultRepository.CountByMbti(mbti)

	if err != nil {
		return
	}
	num += 1
	percent = math.Round(float64(num) / float64(total) * 100)
	return
}

func (s *OptionClassifyHandler) GetReturn(customerConfig *model.CustomerConfig, user string) (response *types.GetQuestionnaireResultResponse, err error) {
	s.initRepository()

	response = &types.GetQuestionnaireResultResponse{}
	var results []types.QuestionnaireResult
	//先查看是否有进行中的
	lastRes, err := s.resultRepository.GetLastResult(customerConfig.KfId, user)
	if err != nil {
		return
	}
	if lastRes != nil && lastRes.Id > 0 && lastRes.Status != repository.QuestionnaireResultStatusEnd {
		results = append(results, assembler.POTODTOGetQuestionnaireResult(lastRes))
		response.Result = results
		return
	}
	//没有进行中的，那就需要获取最近的一次成功的结果

	if lastRes != nil && lastRes.RelationId > 0 {
		//从mbti表获取结果
		mbtiResult, err1 := s.resultMbtiRepository.GetByRelationId(lastRes.RelationId)
		if err1 != nil {
			return
		}
		if mbtiResult != nil && mbtiResult.Id > 0 {
			results = append(results, assembler.POTODTOGetQuestionnaireResultUseMbti(lastRes, mbtiResult))
			response.Result = results
			return
		}
	}
	return
}

func (s *OptionClassifyHandler) dealContentByAI(config *model.CustomerConfig, mbtiModel *model.QuestionnaireResultMbti, user string, channel chan string, nameType, result string) (chatMessage string, err error) {
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
