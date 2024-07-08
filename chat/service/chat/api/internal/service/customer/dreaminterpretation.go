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
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

const (
	StatusGetAnswer1 = 1
	StatusGetAnswer2 = 2
)

type DreamInterpretationHandler struct {
	logger      logx.Logger
	ctx         context.Context
	svcCtx      *svc.ServiceContext
	relationId1 int64
	relationId2 int64
}

func (d *DreamInterpretationHandler) GetAnswer1(param *types.EvaluationSocketRequest, customerConfig *model.CustomerConfig, channel chan string, baseChannel chan any) (res *types.EvaluationSocketResponse, err error) {
	baseInfo := types.EvaluationSocketResponse{}
	baseChannel <- baseInfo

	//获取prompt
	prompt, err := service.NewPromptService(d.ctx, d.svcCtx).GetByKfIdAndStateUseCache(param.OpenKfID, StatusGetAnswer1)
	if err != nil {
		return
	}

	if prompt == nil || prompt.Id <= 0 || prompt.Prompt == "" {
		err = util.ReturnError(xerr.ChatPromptEmpty)
		return
	}
	d.logger.Infof(" GetAnswer1 prompt:%s", prompt.Prompt)

	//调用ai接口获取问题的答案
	aiChatService := service.NewAiChatService(d.ctx, d.svcCtx)
	summary := aiChatService.SetSystemAndUserPrompt(prompt.Prompt, param.Message)

	chatMessage, err := aiChatService.CustomerChatStream(customerConfig, summary, channel, param.User)
	if err != nil {
		d.logger.Error(" GetAnswer1:CustomerChat:error::" + err.Error())
		err = util.ReturnError(xerr.SystemError)
		return
	}
	d.logger.Infof(" GetAnswer1 chatMessage:%s", chatMessage)
	stateModel := &model.State{
		User:  param.User,
		KfId:  param.OpenKfID,
		State: StatusInit,
		Date:  util.TimeToSql(time.Now()),
	}
	_ = service.NewStateService(d.ctx, d.svcCtx).Insert(stateModel)

	questionModel := &model.ChatRecord{
		User:             param.User,
		MessageId:        param.MessageID,
		OpenKfId:         param.OpenKfID,
		Content:          param.Message,
		StateId:          stateModel.Id,
		State:            StatusGetAnswer1,
		ChatType:         repository.ChatTypeCustomer,
		AnswerOrQuestion: repository.AnswerOrQuestionQuestion,
		MessageType:      repository.MessageTypeText,
	}
	_ = service.NewChatService(d.ctx, d.svcCtx).Insert(questionModel)

	_ = service.NewChatService(d.ctx, d.svcCtx).Insert(&model.ChatRecord{
		RelationId:       questionModel.Id,
		User:             param.User,
		MessageId:        param.MessageID,
		OpenKfId:         param.OpenKfID,
		Content:          chatMessage,
		StateId:          stateModel.Id,
		State:            StatusGetAnswer1,
		ChatType:         repository.ChatTypeApplication,
		AnswerOrQuestion: repository.AnswerOrQuestionAnswer,
		MessageType:      repository.MessageTypeText,
	})
	_ = service.NewStateService(d.ctx, d.svcCtx).Insert(&model.State{
		User:  param.User,
		KfId:  param.OpenKfID,
		State: StatusGetAnswer1,
		Date:  util.TimeToSql(time.Now()),
	})
	_ = service.NewStateService(d.ctx, d.svcCtx).UpdateStatus(param.User, param.OpenKfID, StatusGetAnswer1)

	res = &types.EvaluationSocketResponse{
		Message:     chatMessage,
		Status:      StatusGetAnswer1,
		IsEnd:       true,
		StatusIsEnd: false,
	}
	cache.SetCache(d.ctx, aiChatService.SetAssistantPrompt(summary, chatMessage), "DreamInterpretation", param.OpenKfID, param.User)

	return
}

func (d *DreamInterpretationHandler) GetAnswer2(param *types.EvaluationSocketRequest, customerConfig *model.CustomerConfig, state *model.State, channel chan string, baseChannel chan any) (res *types.EvaluationSocketResponse, err error) {
	baseInfo := types.EvaluationSocketResponse{}
	baseChannel <- baseInfo

	//获取prompt
	prompt, err := service.NewPromptService(d.ctx, d.svcCtx).GetByKfIdAndStateUseCache(param.OpenKfID, StatusGetAnswer2)
	if err != nil {
		return
	}

	if prompt == nil || prompt.Id <= 0 || prompt.Prompt == "" {
		err = util.ReturnError(xerr.ChatPromptEmpty)
		return
	}
	d.logger.Infof(" GetAnswer2 prompt:%s", prompt.Prompt)

	//调用ai接口获取问题的答案
	aiChatService := service.NewAiChatService(d.ctx, d.svcCtx)
	summary := cache.GetCache(d.ctx, "DreamInterpretation", param.OpenKfID, param.User)
	summary = aiChatService.SetSystemPrompt(summary, prompt.Prompt)
	summary = aiChatService.SetUserPrompt(summary, param.Message)

	chatMessage, err := aiChatService.CustomerChatStream(customerConfig, summary, channel, param.User)
	if err != nil {
		d.logger.Error(" GetAnswer2:CustomerChat:error::" + err.Error())
		err = util.ReturnError(xerr.SystemError)
		return
	}
	d.logger.Infof(" GetAnswer2 chatMessage:%s", chatMessage)
	questionModel := &model.ChatRecord{
		User:             param.User,
		MessageId:        param.MessageID,
		OpenKfId:         param.OpenKfID,
		Content:          param.Message,
		StateId:          state.Id,
		State:            StatusGetAnswer2,
		ChatType:         repository.ChatTypeCustomer,
		AnswerOrQuestion: repository.AnswerOrQuestionQuestion,
		MessageType:      repository.MessageTypeText,
	}
	_ = service.NewChatService(d.ctx, d.svcCtx).Insert(questionModel)

	lastId, _ := service.NewChatService(d.ctx, d.svcCtx).InsertV2(&model.ChatRecord{
		RelationId:       questionModel.Id,
		User:             param.User,
		MessageId:        param.MessageID,
		OpenKfId:         param.OpenKfID,
		Content:          chatMessage,
		StateId:          state.Id,
		State:            StatusGetAnswer2,
		ChatType:         repository.ChatTypeApplication,
		AnswerOrQuestion: repository.AnswerOrQuestionAnswer,
		MessageType:      repository.MessageTypeText,
	})
	_ = service.NewStateService(d.ctx, d.svcCtx).UpdateStatus(param.User, param.OpenKfID, StatusGetAnswer2)
	res = &types.EvaluationSocketResponse{
		Message:     chatMessage,
		Status:      StatusGetAnswer2,
		IsEnd:       true,
		StatusIsEnd: true,
	}
	go service.NewRiskService(context.Background(), d.svcCtx).Reduce(param.User, param.OpenKfID, lastId)
	return
}

func (d *DreamInterpretationHandler) HandleGet(f int32, param *types.GetEvaluationRequest) (*types.GetEvaluationResponse, error) { // 处理方法
	switch f {
	case StatusGetAnswer1:
		return d.GetAnswer1Record(param)
	case StatusGetAnswer2:
		return d.GetAnswer2Record(param)

	default:
		return nil, util.ReturnError(xerr.SystemError)
	}
}

func (d *DreamInterpretationHandler) GetAnswer2Record(param *types.GetEvaluationRequest) (res *types.GetEvaluationResponse, err error) {
	now := util.NowTimeFormat()
	chatRecords, err := service.NewChatService(d.ctx, d.svcCtx).GetAllKfRecord(param.OpenKfID, param.User, now+" 00:00:00", now+" 23:59:59", []int64{StatusGetAnswer2}, 0)
	if err != nil {
		return
	}
	if chatRecords == nil || len(chatRecords) < 1 {
		err = util.ReturnError(xerr.SystemError)
		return
	}
	var answer string
	for _, v := range chatRecords {
		if v.AnswerOrQuestion == repository.AnswerOrQuestionAnswer {
			answer = v.Content
		}
	}
	res = &types.GetEvaluationResponse{
		Message: answer,
		Status:  StatusGetAnswer2,
		IsEnd:   true,
	}
	return
}

func (d *DreamInterpretationHandler) GetAnswer1Record(param *types.GetEvaluationRequest) (res *types.GetEvaluationResponse, err error) {
	now := util.NowTimeFormat()
	chatRecords, err := service.NewChatService(d.ctx, d.svcCtx).GetAllKfRecord(param.OpenKfID, param.User, now+" 00:00:00", now+" 23:59:59", []int64{StatusGetAnswer1}, 0)
	if err != nil {
		return
	}
	if chatRecords == nil || len(chatRecords) < 1 {
		err = util.ReturnError(xerr.SystemError)
		return
	}
	var answer string
	for _, v := range chatRecords {
		if v.AnswerOrQuestion == repository.AnswerOrQuestionAnswer {
			answer = v.Content
		}
	}
	res = &types.GetEvaluationResponse{
		Message: answer,
		Status:  StatusGetAnswer1,
		IsEnd:   false,
	}
	return
}

func (d *DreamInterpretationHandler) HandleGetSocket(state *model.State, r *types.GetEvaluationSocketRequest) (res *types.GetEvaluationSocketResponse, err error) {
	return
}

func (d *DreamInterpretationHandler) Handle(f int32, param *types.EvaluationSocketRequest, customerConfig *model.CustomerConfig, state *model.State, channel chan string, baseChannel chan any) (res *types.EvaluationSocketResponse, err error) {
	switch f {
	case StatusGetAnswer1:
		return d.GetAnswer1(param, customerConfig, channel, baseChannel)
	case StatusGetAnswer2:
		return d.GetAnswer2(param, customerConfig, state, channel, baseChannel)

	default:
		return nil, util.ReturnError(xerr.SystemError)
	}
}
