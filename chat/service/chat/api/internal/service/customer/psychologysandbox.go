package customer

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/cache"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/service/assembler"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"time"
)

const (
	StatusInit                                  = 0
	StatusPsychologySandboxGenerateProblemOne   = 1
	StatusPsychologySandboxGenerateProblemTwo   = 2
	StatusPsychologySandboxGenerateProblemThree = 3
	StatusPsychologySandboxGenerateResult       = 4
)
const (
	BaseProblem = "你好，请问你今天想测什么？"
)

type PsychologySandboxHandler struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// psychologySandboxGenerateProblemOne 生成第一个问题
func (p *PsychologySandboxHandler) psychologySandboxGenerateProblemOne(param *types.EvaluationSocketRequest, customerConfig *model.CustomerConfig, channel chan string) (res *types.EvaluationSocketResponse, err error) {
	res = &types.EvaluationSocketResponse{}

	//获取prompt
	prompt, err := service.NewPromptService(p.ctx, p.svcCtx).GetByKfIdAndStateUseCache(param.OpenKfID, StatusPsychologySandboxGenerateProblemOne)
	if err != nil {
		return
	}

	if prompt == nil || prompt.Id <= 0 || prompt.Prompt == "" {
		err = util.ReturnError(xerr.ChatPromptEmpty)
		return
	}

	p.logger.Infof(" psychologySandboxType prompt:%s", prompt.Prompt)

	//调用ai接口获取问题的答案

	aiChatService := service.NewAiChatService(p.ctx, p.svcCtx)
	summary := aiChatService.SetSystemAndUserPrompt(prompt.Prompt, BaseProblem)
	summary = aiChatService.SetAssistantPrompt(summary, param.Message)

	chatMessage, err := aiChatService.CustomerChatStream(customerConfig, summary, channel, param.User)

	if err != nil {
		p.logger.Error(" psychologySandboxGenerateProblemOne:CustomerChatStream:error::" + err.Error())
		err = util.ReturnError(xerr.SystemError)
		return
	}
	p.logger.Infof(" psychologySandboxGenerateProblemOne chatMessage:%s", chatMessage)

	stateModel := &model.State{
		User:  param.User,
		KfId:  param.OpenKfID,
		State: StatusInit,
		Date:  util.TimeToSql(time.Now()),
	}
	err1 := service.NewStateService(p.ctx, p.svcCtx).Insert(stateModel)
	if err1 != nil {
		p.logger.Errorf("psychologySandboxGenerateProblemOne NewStateService Insert err %+v", err1)
	}

	// save question
	questionModel := &model.ChatRecord{
		User:             param.User,
		MessageId:        param.MessageID,
		OpenKfId:         param.OpenKfID,
		Content:          BaseProblem,
		State:            StatusPsychologySandboxGenerateProblemOne,
		StateId:          stateModel.Id,
		ChatType:         repository.ChatTypeCustomer,
		AnswerOrQuestion: repository.AnswerOrQuestionQuestion,
		MessageType:      repository.MessageTypeText,
	}
	err2 := service.NewChatService(p.ctx, p.svcCtx).Insert(questionModel)
	if err2 != nil {
		p.logger.Errorf("psychologySandboxGenerateProblemOne NewChatService Insert err %+v", err2)
	}

	// save answer
	err3 := service.NewChatService(p.ctx, p.svcCtx).Insert(&model.ChatRecord{
		RelationId:       questionModel.Id,
		User:             param.User,
		MessageId:        param.MessageID,
		OpenKfId:         param.OpenKfID,
		Content:          param.Message,
		State:            StatusPsychologySandboxGenerateProblemOne,
		StateId:          stateModel.Id,
		ChatType:         repository.ChatTypeCustomer,
		AnswerOrQuestion: repository.AnswerOrQuestionAnswer,
		MessageType:      repository.MessageTypeText,
	})
	if err3 != nil {
		p.logger.Errorf("psychologySandboxGenerateProblemOne NewChatService Insert err %+v", err3)
	}

	// save question
	questionOneModel := &model.ChatRecord{
		User:             param.User,
		MessageId:        param.MessageID,
		OpenKfId:         param.OpenKfID,
		Content:          chatMessage,
		State:            StatusPsychologySandboxGenerateProblemOne,
		StateId:          stateModel.Id,
		ChatType:         repository.ChatTypeCustomer,
		AnswerOrQuestion: repository.AnswerOrQuestionQuestion,
		MessageType:      repository.MessageTypeText,
	}

	err4 := service.NewChatService(p.ctx, p.svcCtx).Insert(questionOneModel)
	if nil != err4 {
		p.logger.Errorf("psychologySandboxGenerateProblemOne NewChatService Insert err %+v", err4)
	}

	res = &types.EvaluationSocketResponse{
		RelationId:       strconv.Itoa(int(questionOneModel.Id)),
		Message:          chatMessage,
		QuestionOrAnswer: int32(repository.AnswerOrQuestionQuestion),
		Status:           StatusPsychologySandboxGenerateProblemOne,
		IsEnd:            true,
		StatusIsEnd:      false,
		MessageId:        param.MessageID,
	}
	err5 := service.NewStateService(p.ctx, p.svcCtx).UpdateStatus(param.User, param.OpenKfID, StatusPsychologySandboxGenerateProblemOne)
	if err5 != nil {
		p.logger.Errorf("psychologySandboxGenerateProblemOne NewStateService UpdateStatus err %+v", err5)
	}

	cache.SetCache(p.ctx, aiChatService.SetUserPrompt(summary, chatMessage), "PsychologySandbox", param.OpenKfID, param.User)

	return
}

// psychologySandboxGenerateProblemTwo 生成第二个问题
func (p *PsychologySandboxHandler) psychologySandboxGenerateProblemTwo(param *types.EvaluationSocketRequest, customerConfig *model.CustomerConfig, state *model.State, channel chan string) (res *types.EvaluationSocketResponse, err error) {
	res = &types.EvaluationSocketResponse{}
	if param == nil || param.RelationId == "" {
		err = util.ReturnError(xerr.ParamMiss)
		return
	}
	relationId, _ := strconv.Atoi(param.RelationId)

	//生成第二个问题
	//获取prompt
	prompt, err := service.NewPromptService(p.ctx, p.svcCtx).GetByKfIdAndStateUseCache(param.OpenKfID, StatusPsychologySandboxGenerateProblemTwo)
	if err != nil {
		return
	}

	if prompt == nil || prompt.Id <= 0 || prompt.Prompt == "" {
		err = util.ReturnError(xerr.ChatPromptEmpty)
		return
	}

	p.logger.Infof(" psychologySandboxGenerateProblemTwo prompt:%s", prompt.Prompt)

	//调用ai接口获取问题的答案

	aiChatService := service.NewAiChatService(p.ctx, p.svcCtx)

	summary := cache.GetCache(p.ctx, "PsychologySandbox", param.OpenKfID, param.User)
	summary = aiChatService.SetSystemPrompt(summary, prompt.Prompt)
	summary = aiChatService.SetAssistantPrompt(summary, param.Message)
	chatMessage, err := aiChatService.CustomerChatStream(customerConfig, summary, channel, param.User)

	if err != nil {
		p.logger.Error(" psychologySandboxGenerateProblemTwo:CustomerChatStream:error::" + err.Error())
		err = util.ReturnError(xerr.SystemError)
		return
	}
	p.logger.Infof(" psychologySandboxGenerateProblemTwo chatMessage:%s", chatMessage)

	// save answer
	err1 := service.NewChatService(p.ctx, p.svcCtx).Insert(&model.ChatRecord{
		RelationId:       int64(relationId),
		User:             param.User,
		MessageId:        param.MessageID,
		OpenKfId:         param.OpenKfID,
		Content:          param.Message,
		State:            StatusPsychologySandboxGenerateProblemTwo,
		StateId:          state.Id,
		ChatType:         repository.ChatTypeCustomer,
		AnswerOrQuestion: repository.AnswerOrQuestionAnswer,
		MessageType:      repository.MessageTypeText,
	})
	if err1 != nil {
		p.logger.Errorf("psychologySandboxGenerateProblemTwo NewChatService Insert err %+v", err1)
	}

	// save question
	questionOneModel := &model.ChatRecord{
		User:             param.User,
		MessageId:        param.MessageID,
		OpenKfId:         param.OpenKfID,
		Content:          chatMessage,
		State:            StatusPsychologySandboxGenerateProblemTwo,
		StateId:          state.Id,
		ChatType:         repository.ChatTypeCustomer,
		AnswerOrQuestion: repository.AnswerOrQuestionQuestion,
		MessageType:      repository.MessageTypeText,
	}

	err2 := service.NewChatService(p.ctx, p.svcCtx).Insert(questionOneModel)
	if err2 != nil {
		p.logger.Errorf("psychologySandboxGenerateProblemTwo NewChatService Insert err %+v", err2)
	}
	err3 := service.NewStateService(p.ctx, p.svcCtx).UpdateStatus(param.User, param.OpenKfID, StatusPsychologySandboxGenerateProblemTwo)
	if err3 != nil {
		p.logger.Errorf("psychologySandboxGenerateProblemTwo NewStateService UpdateStatus err %+v", err3)
	}
	res = &types.EvaluationSocketResponse{
		RelationId:       strconv.Itoa(int(questionOneModel.Id)),
		Message:          chatMessage,
		QuestionOrAnswer: int32(repository.AnswerOrQuestionQuestion),
		Status:           StatusPsychologySandboxGenerateProblemTwo,
		IsEnd:            true,
		StatusIsEnd:      false,
		MessageId:        param.MessageID,
	}

	cache.SetCache(p.ctx, aiChatService.SetUserPrompt(summary, chatMessage), "PsychologySandbox", param.OpenKfID, param.User)

	return
}

// psychologySandboxGenerateProblemThree 生成第三个问题
func (p *PsychologySandboxHandler) psychologySandboxGenerateProblemThree(param *types.EvaluationSocketRequest, customerConfig *model.CustomerConfig, state *model.State, channel chan string) (res *types.EvaluationSocketResponse, err error) {
	res = &types.EvaluationSocketResponse{}
	if param == nil || param.RelationId == "" {
		err = util.ReturnError(xerr.ParamMiss)
		return
	}
	relationId, _ := strconv.Atoi(param.RelationId)

	//生成第三个问题
	//获取prompt
	prompt, err := service.NewPromptService(p.ctx, p.svcCtx).GetByKfIdAndStateUseCache(param.OpenKfID, StatusPsychologySandboxGenerateProblemThree)
	if err != nil {
		return
	}

	if prompt == nil || prompt.Id <= 0 || prompt.Prompt == "" {
		err = util.ReturnError(xerr.ChatPromptEmpty)
		return
	}

	p.logger.Infof(" psychologySandboxGenerateProblemThree prompt:%s", prompt.Prompt)

	//调用ai接口获取问题的答案

	aiChatService := service.NewAiChatService(p.ctx, p.svcCtx)
	summary := cache.GetCache(p.ctx, "PsychologySandbox", param.OpenKfID, param.User)
	summary = aiChatService.SetSystemPrompt(summary, prompt.Prompt)
	summary = aiChatService.SetAssistantPrompt(summary, param.Message)
	chatMessage, err := aiChatService.CustomerChatStream(customerConfig, summary, channel, param.User)
	if err != nil {
		p.logger.Error(" psychologySandboxGenerateProblemThree:CustomerChatStream:error::" + err.Error())
		err = util.ReturnError(xerr.SystemError)
		return
	}
	p.logger.Infof(" psychologySandboxGenerateProblemThree chatMessage:%s", chatMessage)

	// save answer
	err1 := service.NewChatService(p.ctx, p.svcCtx).Insert(&model.ChatRecord{
		RelationId:       int64(relationId),
		User:             param.User,
		MessageId:        param.MessageID,
		OpenKfId:         param.OpenKfID,
		Content:          param.Message,
		StateId:          state.Id,
		State:            StatusPsychologySandboxGenerateProblemThree,
		ChatType:         repository.ChatTypeCustomer,
		AnswerOrQuestion: repository.AnswerOrQuestionAnswer,
		MessageType:      repository.MessageTypeText,
	})
	if err1 != nil {
		p.logger.Errorf("psychologySandboxGenerateProblemThree NewChatService Insert err %+v", err1)
	}

	// save question
	questionOneModel := &model.ChatRecord{
		User:             param.User,
		MessageId:        param.MessageID,
		OpenKfId:         param.OpenKfID,
		Content:          chatMessage,
		StateId:          state.Id,
		State:            StatusPsychologySandboxGenerateProblemThree,
		ChatType:         repository.ChatTypeCustomer,
		AnswerOrQuestion: repository.AnswerOrQuestionQuestion,
		MessageType:      repository.MessageTypeText,
	}

	err2 := service.NewChatService(p.ctx, p.svcCtx).Insert(questionOneModel)
	if err2 != nil {
		p.logger.Errorf("psychologySandboxGenerateProblemThree NewChatService Insert err %+v", err2)
	}
	err3 := service.NewStateService(p.ctx, p.svcCtx).UpdateStatus(param.User, param.OpenKfID, StatusPsychologySandboxGenerateProblemThree)
	if err3 != nil {
		p.logger.Errorf("psychologySandboxGenerateProblemThree NewStateService UpdateStatus err %+v", err3)
	}
	res = &types.EvaluationSocketResponse{
		RelationId:       strconv.Itoa(int(questionOneModel.Id)),
		Message:          chatMessage,
		QuestionOrAnswer: int32(repository.AnswerOrQuestionQuestion),
		Status:           StatusPsychologySandboxGenerateProblemThree,
		IsEnd:            true,
		StatusIsEnd:      false,
		MessageId:        param.MessageID,
	}

	cache.SetCache(p.ctx, aiChatService.SetUserPrompt(summary, chatMessage), "PsychologySandbox", param.OpenKfID, param.User)

	return
}

// psychologySandboxGenerateResult 生成答案
func (p *PsychologySandboxHandler) psychologySandboxGenerateResult(param *types.EvaluationSocketRequest, customerConfig *model.CustomerConfig, state *model.State, channel chan string) (res *types.EvaluationSocketResponse, err error) {

	res = &types.EvaluationSocketResponse{}
	if param == nil || param.RelationId == "" {
		err = util.ReturnError(xerr.ParamMiss)
		return
	}
	relationId, _ := strconv.Atoi(param.RelationId)

	//获取prompt
	prompt, err := service.NewPromptService(p.ctx, p.svcCtx).GetByKfIdAndStateUseCache(param.OpenKfID, StatusPsychologySandboxGenerateResult)
	if err != nil {
		return
	}

	if prompt == nil || prompt.Id <= 0 || prompt.Prompt == "" {
		err = util.ReturnError(xerr.ChatPromptEmpty)
		return
	}

	p.logger.Infof(" psychologySandboxGenerateResult prompt:%s", prompt.Prompt)

	//调用ai接口获取问题的答案

	aiChatService := service.NewAiChatService(p.ctx, p.svcCtx)
	summary := cache.GetCache(p.ctx, "PsychologySandbox", param.OpenKfID, param.User)
	summary = aiChatService.SetSystemPrompt(summary, prompt.Prompt)
	summary = aiChatService.SetAssistantPrompt(summary, param.Message)
	chatMessage, err := aiChatService.CustomerChatStream(customerConfig, summary, channel, param.User)

	if err != nil {
		p.logger.Error(" psychologySandboxGenerateResult:CustomerChatStream:error::" + err.Error())
		err = util.ReturnError(xerr.SystemError)
		return
	}
	p.logger.Infof(" psychologySandboxGenerateResult chatMessage:%s", chatMessage)

	// save answer
	err1 := service.NewChatService(p.ctx, p.svcCtx).Insert(&model.ChatRecord{
		RelationId:       int64(relationId),
		User:             param.User,
		MessageId:        param.MessageID,
		OpenKfId:         param.OpenKfID,
		Content:          param.Message,
		StateId:          state.Id,
		State:            StatusPsychologySandboxGenerateResult,
		ChatType:         repository.ChatTypeCustomer,
		AnswerOrQuestion: repository.AnswerOrQuestionAnswer,
		MessageType:      repository.MessageTypeText,
	})
	if err1 != nil {
		p.logger.Errorf("psychologySandboxGenerateResult NewChatService Insert err %+v", err1)
	}

	// save question
	questionOneModel := &model.ChatRecord{
		User:             param.User,
		MessageId:        param.MessageID,
		OpenKfId:         param.OpenKfID,
		Content:          chatMessage,
		StateId:          state.Id,
		State:            StatusPsychologySandboxGenerateResult,
		ChatType:         repository.ChatTypeCustomer,
		AnswerOrQuestion: repository.AnswerOrQuestionSummarize,
		MessageType:      repository.MessageTypeText,
	}

	err2 := service.NewChatService(p.ctx, p.svcCtx).Insert(questionOneModel)
	if err2 != nil {
		p.logger.Errorf("psychologySandboxGenerateResult NewChatService Insert err %+v", err2)
	}
	err3 := service.NewStateService(p.ctx, p.svcCtx).UpdateStatus(param.User, param.OpenKfID, StatusPsychologySandboxGenerateResult)
	if err3 != nil {
		p.logger.Errorf("psychologySandboxGenerateResult NewStateService UpdateStatus err %+v", err3)
	}
	res = &types.EvaluationSocketResponse{
		Message:          chatMessage,
		QuestionOrAnswer: int32(repository.AnswerOrQuestionSummarize),
		Status:           StatusPsychologySandboxGenerateResult,
		IsEnd:            true,
		StatusIsEnd:      true,
		MessageId:        param.MessageID,
	}

	cache.SetCache(p.ctx, aiChatService.SetUserPrompt(summary, chatMessage), "PsychologySandbox", param.OpenKfID, param.User)
	go service.NewRiskService(context.Background(), p.svcCtx).Reduce(param.User, param.OpenKfID, state.Id)

	return
}

//record

func (p *PsychologySandboxHandler) HandleGet(m int32, r *types.GetEvaluationRequest) (res *types.GetEvaluationResponse, err error) { // 处理方法
	return
}

func (p *PsychologySandboxHandler) HandleGetSocket(state *model.State, r *types.GetEvaluationSocketRequest) (*types.GetEvaluationSocketResponse, error) { // 处理方法
	switch state.State {
	case StatusPsychologySandboxGenerateProblemOne:
		return p.recordPsychologySandboxGenerateProblemOne(r, state)
	case StatusPsychologySandboxGenerateProblemTwo:
		return p.recordPsychologySandboxGenerateProblemTwo(r, state)
	case StatusPsychologySandboxGenerateProblemThree:
		return p.recordPsychologySandboxGenerateProblemThree(r, state)
	case StatusPsychologySandboxGenerateResult:
		return p.recordPsychologySandboxGenerateResult(r, state)
	default:
		return nil, util.ReturnError(xerr.SystemError)
	}
}

func (p *PsychologySandboxHandler) getChatRecordByStateId(param *types.GetEvaluationSocketRequest, state *model.State, isEnd bool) (res *types.GetEvaluationSocketResponse, err error) {
	now := util.NowTimeFormat()
	chatRecords, err := service.NewChatService(p.ctx, p.svcCtx).GetAllKfRecord(param.OpenKfID, param.User, now+" 00:00:00", now+" 23:59:59", []int64{}, state.Id)
	if err != nil {
		return
	}
	if chatRecords == nil || len(chatRecords) < 1 {
		err = util.ReturnError(xerr.SystemError)
		return
	}

	res = &types.GetEvaluationSocketResponse{
		List:   assembler.POTODTOGetEvaluationSocket(chatRecords),
		Status: state.State,
		IsEnd:  isEnd,
	}
	return
}

func (p *PsychologySandboxHandler) recordPsychologySandboxGenerateProblemOne(param *types.GetEvaluationSocketRequest, state *model.State) (res *types.GetEvaluationSocketResponse, err error) {
	return p.getChatRecordByStateId(param, state, false)
}

func (p *PsychologySandboxHandler) recordPsychologySandboxGenerateProblemTwo(param *types.GetEvaluationSocketRequest, state *model.State) (res *types.GetEvaluationSocketResponse, err error) {
	return p.getChatRecordByStateId(param, state, false)
}

func (p *PsychologySandboxHandler) recordPsychologySandboxGenerateProblemThree(param *types.GetEvaluationSocketRequest, state *model.State) (res *types.GetEvaluationSocketResponse, err error) {
	return p.getChatRecordByStateId(param, state, false)
}

func (p *PsychologySandboxHandler) recordPsychologySandboxGenerateResult(param *types.GetEvaluationSocketRequest, state *model.State) (res *types.GetEvaluationSocketResponse, err error) {
	return p.getChatRecordByStateId(param, state, true)
}

func (p *PsychologySandboxHandler) Handle(m int32, r *types.EvaluationSocketRequest, customerConfig *model.CustomerConfig, state *model.State, channel chan string, baseChannel chan any) (res *types.EvaluationSocketResponse, err error) {
	baseInfo := types.EvaluationSocketResponse{}
	baseChannel <- baseInfo
	switch m {
	case StatusPsychologySandboxGenerateProblemOne:
		return p.psychologySandboxGenerateProblemOne(r, customerConfig, channel)
	case StatusPsychologySandboxGenerateProblemTwo:
		return p.psychologySandboxGenerateProblemTwo(r, customerConfig, state, channel)
	case StatusPsychologySandboxGenerateProblemThree:
		return p.psychologySandboxGenerateProblemThree(r, customerConfig, state, channel)
	case StatusPsychologySandboxGenerateResult:
		return p.psychologySandboxGenerateResult(r, customerConfig, state, channel)
	default:
		return nil, util.ReturnError(xerr.SystemError)
	}
}
