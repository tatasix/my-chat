package customer

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/cache"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/model"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"time"
)

const (
	StatusGenerateTarot   = 1
	StatusGenerateProblem = 2
	StatusGenerateAnswer  = 3
	StatusNotContinue     = 4
)

type FortuneHandler struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func (f *FortuneHandler) HandleGet(m int32, r *types.GetEvaluationRequest) (*types.GetEvaluationResponse, error) { // 处理方法
	switch m {
	case 1:
		return f.RecordGenerateTarot(r)
	case 2:
		return f.RecordGenerateProblem(r)
	case 3:
		return f.RecordGenerateAnswer(r)
	case 4:
		return f.RecordNotContinue(r)
	default:
		return nil, util.ReturnError(xerr.SystemError)
	}
}

// GenerateTarot 生成塔罗牌
func (f *FortuneHandler) GenerateTarot(param *types.EvaluationSocketRequest, customerConfig *model.CustomerConfig, channel chan string, baseChannel chan any) (res *types.EvaluationSocketResponse, err error) {

	//tarot := util.GenerateRandomNumber(0, 77)
	tarot := param.Tarot
	baseInfo := types.EvaluationSocketResponse{
		Tarot: tarot,
		//TarotImage: util.GetTarotImage(tarot),
		TarotName: param.TarotName,
	}
	baseChannel <- baseInfo
	//获取prompt
	prompt, err := service.NewPromptService(f.ctx, f.svcCtx).GetByKfIdAndStateUseCache(param.OpenKfID, StatusGenerateTarot)
	if err != nil {
		return
	}

	if prompt == nil || prompt.Id <= 0 || prompt.Prompt == "" {
		err = util.ReturnError(xerr.ChatPromptEmpty)
		return
	}

	f.logger.Infof(" GenerateTarotInfo prompt:%s", prompt.Prompt)

	//调用ai接口获取问题的答案

	aiChatService := service.NewAiChatService(f.ctx, f.svcCtx)
	summary := aiChatService.SetSystemAndUserPrompt(prompt.Prompt, util.GetTarotName(tarot))
	////chatMessage, err := aiChatService.CustomerChat(customerConfig, summary)
	//chatMessage, err := aiChatService.CustomerChatStream(customerConfig, summary, channel, param.User)
	//
	//if err != nil {
	//	f.logger.Error(" GenerateTarot:CustomerChat:error::" + err.Error())
	//	err = util.ReturnError(xerr.SystemError)
	//	return
	//}
	chatMessage := param.Message

	f.logger.Infof(" GenerateTarotInfo chatMessage:%s", chatMessage)
	stateModel := &model.State{
		User:  param.User,
		KfId:  param.OpenKfID,
		State: StatusInit,
		Date:  util.TimeToSql(time.Now()),
	}
	_ = service.NewStateService(f.ctx, f.svcCtx).Insert(stateModel)

	questionModel := &model.ChatRecord{
		User:      param.User,
		MessageId: param.MessageID,
		OpenKfId:  param.OpenKfID,
		//Content:          util.GetTarotName(tarot),
		Content:          param.TarotName,
		StateId:          stateModel.Id,
		State:            StatusGenerateTarot,
		ChatType:         repository.ChatTypeCustomer,
		AnswerOrQuestion: repository.AnswerOrQuestionQuestion,
		MessageType:      repository.MessageTypeText,
	}
	_ = service.NewChatService(f.ctx, f.svcCtx).Insert(questionModel)
	paramJson, _ := json.Marshal(param)
	_ = service.NewChatService(f.ctx, f.svcCtx).Insert(&model.ChatRecord{
		RelationId: questionModel.Id,
		User:       param.User,
		MessageId:  param.MessageID,
		OpenKfId:   param.OpenKfID,
		//Content:          chatMessage,
		Content:          string(paramJson),
		StateId:          stateModel.Id,
		State:            StatusGenerateTarot,
		ChatType:         repository.ChatTypeApplication,
		AnswerOrQuestion: repository.AnswerOrQuestionAnswer,
		MessageType:      repository.MessageTypeText,
	})
	_ = service.NewStateService(f.ctx, f.svcCtx).UpdateStatus(param.User, param.OpenKfID, StatusGenerateTarot)

	res = &types.EvaluationSocketResponse{
		Tarot: tarot,
		//TarotImage:  util.GetTarotImage(tarot),
		//TarotName:   util.GetTarotName(tarot),
		TarotName:   param.TarotName,
		TarotStatus: param.TarotStatus,
		TarotInfo:   chatMessage,
		Status:      StatusGenerateTarot,
		IsEnd:       true,
		StatusIsEnd: false,
	}

	cache.SetCache(f.ctx, aiChatService.SetAssistantPrompt(summary, chatMessage), "Fortune", param.OpenKfID, param.User)
	cache.AddCache(f.ctx, "assistant", "塔罗牌牌面", "Fortune", param.OpenKfID, param.User)
	cache.AddCache(f.ctx, "user", param.TarotName, "Fortune", param.OpenKfID, param.User)
	cache.AddCache(f.ctx, "assistant", "塔罗牌正反位", "Fortune", param.OpenKfID, param.User)
	cache.AddCache(f.ctx, "user", param.TarotStatus, "Fortune", param.OpenKfID, param.User)
	cache.AddCache(f.ctx, "assistant", "幸运数字", "Fortune", param.OpenKfID, param.User)
	cache.AddCache(f.ctx, "user", param.LuckyNum, "Fortune", param.OpenKfID, param.User)
	cache.AddCache(f.ctx, "assistant", "幸运颜色", "Fortune", param.OpenKfID, param.User)
	cache.AddCache(f.ctx, "user", param.LuckyColor, "Fortune", param.OpenKfID, param.User)
	cache.AddCache(f.ctx, "assistant", "幸运食物", "Fortune", param.OpenKfID, param.User)
	cache.AddCache(f.ctx, "user", param.LuckyFood, "Fortune", param.OpenKfID, param.User)
	cache.AddCache(f.ctx, "assistant", "今天适宜做", "Fortune", param.OpenKfID, param.User)
	cache.AddCache(f.ctx, "user", param.Suitable, "Fortune", param.OpenKfID, param.User)
	cache.AddCache(f.ctx, "assistant", "今天不适宜做", "Fortune", param.OpenKfID, param.User)
	cache.AddCache(f.ctx, "user", param.Unsuitable, "Fortune", param.OpenKfID, param.User)

	return
}

// GenerateProblem 生成问题
func (f *FortuneHandler) GenerateProblem(param *types.EvaluationSocketRequest, state *model.State, channel chan string, baseChannel chan any) (res *types.EvaluationSocketResponse, err error) {
	baseInfo := types.EvaluationSocketResponse{}
	baseChannel <- baseInfo
	//获取prompt
	problems, err := service.NewConfigService(f.ctx, f.svcCtx).GetAll(repository.ConfigTypeProblem)
	if err != nil {
		return
	}

	if problems == nil || len(problems) <= 0 {
		err = util.ReturnError(xerr.ConfigEmpty)
		return
	}
	rand := util.GenerateRandomNumber(0, len(problems)-1)
	channel <- problems[rand].Value
	questionModel := &model.ChatRecord{
		User:             param.User,
		MessageId:        param.MessageID,
		OpenKfId:         param.OpenKfID,
		Content:          problems[rand].Value,
		StateId:          state.Id,
		State:            StatusGenerateProblem,
		ChatType:         repository.ChatTypeCustomer,
		AnswerOrQuestion: repository.AnswerOrQuestionQuestion,
		MessageType:      repository.MessageTypeText,
	}
	_ = service.NewChatService(f.ctx, f.svcCtx).Insert(questionModel)

	_ = service.NewStateService(f.ctx, f.svcCtx).UpdateStatus(param.User, param.OpenKfID, StatusGenerateProblem)

	var stateParam string
	firstStateData, err := service.NewChatService(f.ctx, f.svcCtx).GetByMessageAndState(state.Id, 1)
	if err == nil && firstStateData != nil && firstStateData.Content != "" {
		stateParam = firstStateData.Content
	}
	res = &types.EvaluationSocketResponse{
		Message:     problems[rand].Value,
		Status:      StatusGenerateProblem,
		IsEnd:       true,
		StatusIsEnd: false,
		RelationId:  strconv.Itoa(int(questionModel.Id)),
		Param:       stateParam,
	}
	cache.AddCache(f.ctx, "assistant", problems[rand].Value, "Fortune", param.OpenKfID, param.User)
	return
}

// GenerateAnswer 生成答案
func (f *FortuneHandler) GenerateAnswer(param *types.EvaluationSocketRequest, customerConfig *model.CustomerConfig, state *model.State, channel chan string, baseChannel chan any) (res *types.EvaluationSocketResponse, err error) {
	if param == nil || param.RelationId == "" {
		err = util.ReturnError(xerr.ParamMiss)
		return
	}
	relationId, _ := strconv.Atoi(param.RelationId)

	//获取prompt
	prompt, err := service.NewPromptService(f.ctx, f.svcCtx).GetByKfIdAndStateUseCache(param.OpenKfID, StatusGenerateAnswer)
	if err != nil {
		return
	}

	if prompt == nil || prompt.Id <= 0 || prompt.Prompt == "" {
		err = util.ReturnError(xerr.ChatPromptEmpty)
		return
	}

	//调用ai接口获取问题的答案
	aiChatService := service.NewAiChatService(f.ctx, f.svcCtx)
	summary := cache.GetCache(f.ctx, "Fortune", param.OpenKfID, param.User)
	summary = aiChatService.SetUserPrompt(summary, param.Message)
	summary = aiChatService.SetSystemPrompt(summary, prompt.Prompt)
	f.logger.Infof(" GenerateAnswer prompt:%+v", summary)

	now := util.NowTimeFormat()
	chatRecords, err := service.NewChatService(f.ctx, f.svcCtx).GetAllKfRecord(param.OpenKfID, param.User, now+" 00:00:00", now+" 23:59:59", []int64{StatusGenerateTarot}, 0)
	if err != nil {
		return
	}
	if chatRecords == nil || len(chatRecords) < 1 {
		err = util.ReturnError(xerr.SystemError)
		return
	}
	var tarotName string
	for _, v := range chatRecords {
		if v.AnswerOrQuestion == repository.AnswerOrQuestionQuestion && v.State == StatusGenerateTarot {
			tarotName = v.Content
		}
	}
	baseInfo := types.EvaluationSocketResponse{
		TarotName:  tarotName,
		TarotImage: util.GetTarotImage(util.GetKeyFromTarotMap(tarotName)),
	}
	baseChannel <- baseInfo
	chatMessage, err := aiChatService.CustomerChatStream(customerConfig, summary, channel, param.User)
	if err != nil {
		f.logger.Error(" GenerateAnswer:CustomerChat:error::" + err.Error())
		err = util.ReturnError(xerr.SystemError)
		return
	}
	_ = service.NewChatService(f.ctx, f.svcCtx).Insert(&model.ChatRecord{
		RelationId:       int64(relationId),
		User:             param.User,
		MessageId:        param.MessageID,
		OpenKfId:         param.OpenKfID,
		Content:          param.Message,
		State:            StatusGenerateAnswer,
		StateId:          state.Id,
		ChatType:         repository.ChatTypeCustomer,
		AnswerOrQuestion: repository.AnswerOrQuestionAnswer,
		MessageType:      repository.MessageTypeText,
	})

	lastId, _ := service.NewChatService(f.ctx, f.svcCtx).InsertV2(&model.ChatRecord{
		User:             param.User,
		MessageId:        param.MessageID,
		OpenKfId:         param.OpenKfID,
		Content:          chatMessage,
		State:            StatusGenerateAnswer,
		StateId:          state.Id,
		ChatType:         repository.ChatTypeCustomer,
		AnswerOrQuestion: repository.AnswerOrQuestionSummarize,
		MessageType:      repository.MessageTypeText,
	})
	_ = service.NewStateService(f.ctx, f.svcCtx).UpdateStatus(param.User, param.OpenKfID, StatusGenerateAnswer)

	var stateParam string
	firstStateData, err := service.NewChatService(f.ctx, f.svcCtx).GetByMessageAndState(state.Id, 1)
	if err == nil && firstStateData != nil && firstStateData.Content != "" {
		stateParam = firstStateData.Content
	}
	res = &types.EvaluationSocketResponse{
		TarotName:   tarotName,
		TarotImage:  util.GetTarotImage(util.GetKeyFromTarotMap(tarotName)),
		Message:     chatMessage,
		Status:      StatusGenerateAnswer,
		IsEnd:       true,
		StatusIsEnd: true,
		Param:       stateParam,
	}
	go service.NewRiskService(context.Background(), f.svcCtx).Reduce(param.User, param.OpenKfID, lastId)
	return
}

// NotContinue 生成答案
func (f *FortuneHandler) NotContinue(param *types.EvaluationSocketRequest, channel chan string, baseChannel chan any) (res *types.EvaluationSocketResponse, err error) {
	baseInfo := types.EvaluationSocketResponse{}
	baseChannel <- baseInfo

	_ = service.NewStateService(f.ctx, f.svcCtx).UpdateStatus(param.User, param.OpenKfID, StatusNotContinue)
	res = &types.EvaluationSocketResponse{
		Status:      StatusNotContinue,
		IsEnd:       true,
		StatusIsEnd: true,
	}
	go service.NewRiskService(context.Background(), f.svcCtx).Reduce(param.User, param.OpenKfID, 0)
	return
}

func (f *FortuneHandler) RecordGenerateTarot(param *types.GetEvaluationRequest) (res *types.GetEvaluationResponse, err error) {
	now := util.NowTimeFormat()
	chatRecords, err := service.NewChatService(f.ctx, f.svcCtx).GetAllKfRecord(param.OpenKfID, param.User, now+" 00:00:00", now+" 23:59:59", []int64{StatusGenerateTarot}, 0)
	if err != nil {
		return
	}
	if chatRecords == nil || len(chatRecords) < 1 {
		err = util.ReturnError(xerr.SystemError)
		return
	}
	var question, answer string
	for _, v := range chatRecords {
		if v.AnswerOrQuestion == repository.AnswerOrQuestionQuestion {
			question = v.Content
		}
		if v.AnswerOrQuestion == repository.AnswerOrQuestionAnswer {
			answer = v.Content
		}
	}
	var stateParam string
	firstStateData, err := service.NewChatService(f.ctx, f.svcCtx).GetByMessageAndState(chatRecords[0].StateId, 1)
	if err == nil && firstStateData != nil && firstStateData.Content != "" {
		stateParam = firstStateData.Content
	}
	res = &types.GetEvaluationResponse{
		TarotName:  question,
		TarotImage: util.GetTarotImage(util.GetKeyFromTarotMap(question)),
		Status:     StatusGenerateTarot,
		Message:    answer,
		IsEnd:      false,
		Param:      stateParam,
	}
	return
}

func (f *FortuneHandler) RecordGenerateProblem(param *types.GetEvaluationRequest) (res *types.GetEvaluationResponse, err error) {
	now := util.NowTimeFormat()
	chatRecords, err := service.NewChatService(f.ctx, f.svcCtx).GetAllKfRecord(param.OpenKfID, param.User, now+" 00:00:00", now+" 23:59:59", []int64{StatusGenerateProblem}, 0)
	if err != nil {
		return
	}
	if chatRecords == nil || len(chatRecords) < 1 {
		err = util.ReturnError(xerr.SystemError)
		return
	}

	var stateParam string
	firstStateData, err := service.NewChatService(f.ctx, f.svcCtx).GetByMessageAndState(chatRecords[0].StateId, 1)
	if err == nil && firstStateData != nil && firstStateData.Content != "" {
		stateParam = firstStateData.Content
	}
	res = &types.GetEvaluationResponse{
		Message:    chatRecords[0].Content,
		RelationId: strconv.Itoa(int(chatRecords[0].RelationId)),
		Status:     StatusGenerateProblem,
		IsEnd:      false,
		Param:      stateParam,
	}
	return
}

func (f *FortuneHandler) RecordGenerateAnswer(param *types.GetEvaluationRequest) (res *types.GetEvaluationResponse, err error) {

	now := util.NowTimeFormat()
	chatRecords, err := service.NewChatService(f.ctx, f.svcCtx).GetAllKfRecord(param.OpenKfID, param.User, now+" 00:00:00", now+" 23:59:59", []int64{StatusGenerateAnswer, StatusGenerateTarot}, 0)
	if err != nil {
		return
	}
	if chatRecords == nil || len(chatRecords) < 1 {
		err = util.ReturnError(xerr.SystemError)
		return
	}
	var answer, tarotName string
	for _, v := range chatRecords {
		if v.AnswerOrQuestion == repository.AnswerOrQuestionQuestion && v.State == StatusGenerateTarot {
			tarotName = v.Content
		}
		if v.AnswerOrQuestion == repository.AnswerOrQuestionSummarize && v.State == StatusGenerateAnswer {
			answer = v.Content
		}
	}

	var stateParam string
	firstStateData, err := service.NewChatService(f.ctx, f.svcCtx).GetByMessageAndState(chatRecords[0].StateId, 1)
	if err == nil && firstStateData != nil && firstStateData.Content != "" {
		stateParam = firstStateData.Content
	}
	res = &types.GetEvaluationResponse{
		TarotName:  tarotName,
		TarotImage: util.GetTarotImage(util.GetKeyFromTarotMap(tarotName)),
		Message:    answer,
		Status:     StatusGenerateAnswer,
		IsEnd:      true,
		Param:      stateParam,
	}
	return
}

func (f *FortuneHandler) RecordNotContinue(param *types.GetEvaluationRequest) (res *types.GetEvaluationResponse, err error) {
	res = &types.GetEvaluationResponse{
		Status: StatusNotContinue,
		IsEnd:  true,
	}
	return
}

func (f *FortuneHandler) HandleGetSocket(state *model.State, r *types.GetEvaluationSocketRequest) (res *types.GetEvaluationSocketResponse, err error) {
	return
}

func (f *FortuneHandler) Handle(m int32, r *types.EvaluationSocketRequest, customerConfig *model.CustomerConfig, state *model.State, channel chan string, baseChannel chan any) (res *types.EvaluationSocketResponse, err error) {
	switch m {
	case 1:
		return f.GenerateTarot(r, customerConfig, channel, baseChannel)
	case 2:
		return f.GenerateProblem(r, state, channel, baseChannel)
	case 3:
		return f.GenerateAnswer(r, customerConfig, state, channel, baseChannel)
	case 4:
		return f.NotContinue(r, channel, baseChannel)
	default:
		err = util.ReturnError(xerr.SystemError)
		return
	}
}
