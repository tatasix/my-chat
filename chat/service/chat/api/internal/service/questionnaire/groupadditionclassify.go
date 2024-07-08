package questionnaire

import (
	"chat/common/util"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/service/assembler"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/model"
	"context"
	"encoding/json"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"sort"
)

type GroupAdditionClassifyHandler struct {
	logger             logx.Logger
	ctx                context.Context
	svcCtx             *svc.ServiceContext
	responseRepository *repository.QuestionnaireResponseRepository
	resultRepository   *repository.QuestionnaireResultRepository
	questionRepository *repository.QuestionnaireQuestionRepository
}

type cusConfig struct {
	Method int `json:"method"`
	Total  int `json:"total"`
	Mod    struct {
		Type int `json:"type"`
		Rate int `json:"rate"`
	} `json:"mod"`
	ResNeedGroupResult int           `json:"res_need_group_result"`
	ResNum             int           `json:"res_num"`
	GroupResult        []groutResult `json:"group_result"`
	Result             []result      `json:"result"`
}
type groutResult struct {
	Type        int      `json:"type"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Topics      []int64  `json:"topics"`
	Result      []result `json:"result"`
}

type result struct {
	Min int64  `json:"min"`
	Max int64  `json:"max"`
	Res string `json:"res"`
}

type groupRes struct {
	Name   string  `json:"name"`
	Desc   string  `json:"desc"`
	Score  int64   `json:"score"`
	Topics []int64 `json:"topics"`
	Res    string  `json:"res"`
}
type resp struct {
	Score  int64      `json:"score"`
	Res    string     `json:"res"`
	Detail []groupRes `json:"detail"`
}

func (s *GroupAdditionClassifyHandler) initRepository() {
	s.responseRepository = repository.NewQuestionnaireResponseRepository(s.ctx, s.svcCtx)
	s.resultRepository = repository.NewQuestionnaireResultRepository(s.ctx, s.svcCtx)
	s.questionRepository = repository.NewQuestionnaireQuestionRepository(s.ctx, s.svcCtx)
	return
}
func (s *GroupAdditionClassifyHandler) Handle(customerConfig *model.CustomerConfig, config model.CustomerConfigConfig, relationId int64, user, openKfID string, channel chan string, baseInfo chan any) (resp *types.GenerateResultResponse, err error) {
	s.initRepository()

	//判断是否答题完成
	//total := config.Total
	//allResponse, err := s.responseRepository.GetAll("", "", relationId)
	//if err != nil {
	//	return
	//}
	//if len(allResponse) != int(total) {
	//	err = util.ReturnError(xerr.QuestionnaireResponseError)
	//	return
	//}
	resp, err = s.summary(user, openKfID, customerConfig, relationId, baseInfo)
	return
}

func (s *GroupAdditionClassifyHandler) summary(user, openKfID string, config *model.CustomerConfig, relationId int64, baseInfo chan any) (re *types.GenerateResultResponse, err error) {
	if config.Config == "" {
		err = errors.New("config.Config配置不能为空")
	}
	var gRes cusConfig
	err = json.Unmarshal([]byte(config.Config), &gRes)
	if err != nil {
		return
	}
	var groupResult []groupRes
	var totalScore int64
	var totalRes string
	for _, v := range gRes.GroupResult {
		groupScore, err := s.getSumByTopics(openKfID, v.Topics, relationId, user)
		if err != nil {
			return nil, err
		}
		//获取分组结果
		groupScoreRes := ""
		for _, vv := range v.Result {
			if vv.Min <= groupScore && vv.Max >= groupScore {
				groupScoreRes = vv.Res
				totalRes += vv.Res + "，"
				break
			}
		}
		totalScore += groupScore

		groupResult = append(groupResult, groupRes{
			Name:   v.Name,
			Desc:   v.Description,
			Score:  groupScore,
			Res:    groupScoreRes,
			Topics: v.Topics,
		})
	}
	//总分的结果
	var resp resp
	if gRes.Result != nil && len(gRes.Result) > 0 {
		for _, vvv := range gRes.Result {
			if vvv.Min <= totalScore && vvv.Max >= totalScore {
				if gRes.ResNeedGroupResult == 1 {
					resp.Res = totalRes + vvv.Res
				} else {
					resp.Res = vvv.Res
				}
				break
			}
		}

	}

	resp.Score = totalScore
	//resp.Detail = groupResult

	//只取最高 gRes.ResNum 个结果
	groupResults := s.getSortDesc(groupResult)
	var detail []groupRes
	for x, y := range groupResults {
		if gRes.ResNum > 0 && gRes.ResNum <= x {
			break
		}
		detail = append(detail, y)
	}

	resp.Detail = detail

	exist, err := s.resultRepository.GetByRelationId(relationId)
	if err != nil {
		return
	}
	total := gRes.Total
	allResponse, err := s.responseRepository.GetAll("", "", relationId)
	if err != nil {
		return
	}

	var status int
	if len(allResponse) != int(total) {
		status = len(allResponse)
	} else {
		status = repository.QuestionnaireResultStatusEnd
	}

	resultModel := &model.QuestionnaireResult{
		RelationId: relationId,
		User:       user,
		OpenKfId:   openKfID,
		Score:      util.Float64ToSql(float64(totalScore)),
		Status:     int64(status),
	}
	byteData, _ := json.Marshal(resp)
	resultModel.Result = string(byteData)

	if exist != nil && len(exist) > 0 {
		resultModel.Id = exist[0].Id
		err = s.resultRepository.Update(resultModel)
	} else {
		err = s.resultRepository.Insert(resultModel)
	}
	po, err := s.resultRepository.GetByRelationId(relationId)
	lastResult := assembler.POTODTOGenerateResult(po[0])
	baseInfo <- lastResult
	return &lastResult, nil
}

// 获取分类总分
func (s *GroupAdditionClassifyHandler) getSumByTopics(openKfID string, topics []int64, relationId int64, user string) (total int64, err error) {
	quesPos, err := s.questionRepository.GetByKfId(openKfID)
	respPos, err := s.responseRepository.GetResponse(openKfID, user, relationId)

	if err != nil {
		return
	}

	if respPos == nil && len(respPos) <= 0 {
		return
	}
	var respTopics []int64
	for _, vv := range respPos {
		if vv.QuestionId >= 0 {
			respTopics = append(respTopics, vv.QuestionId)
		}
	}

	if len(respTopics) <= 0 {
		return
	}
	if quesPos != nil && len(quesPos) > 0 {
		var resp []int64
		for _, v := range quesPos {
			for _, vv := range topics {
				if v.Sort == vv {
					for _, vvv := range respTopics {
						if v.Id == vvv {
							resp = append(resp, v.Id)
						}
					}
				}
			}
		}
		total, err = s.responseRepository.GetSumByTopics(openKfID, resp, relationId)
	}
	return
}

// 根据分类降序
func (s *GroupAdditionClassifyHandler) getSortDesc(groups []groupRes) []groupRes {
	sort.Slice(groups, func(i, j int) bool {
		return groups[i].Score > groups[j].Score
	})

	return groups
}

func (s *GroupAdditionClassifyHandler) GetReturn(customerConfig *model.CustomerConfig, user string) (response *types.GetQuestionnaireResultResponse, err error) {
	s.initRepository()

	response = &types.GetQuestionnaireResultResponse{}
	var results []types.QuestionnaireResult
	//先查看是否有进行中的
	lastRes, err := s.resultRepository.GetLastResult(customerConfig.KfId, user)
	if err != nil {
		return
	}
	if lastRes != nil && lastRes.Id > 0 {
		results = append(results, assembler.POTODTOGetQuestionnaireResult(lastRes))
		response.Result = results
		return
	}

	return
}
