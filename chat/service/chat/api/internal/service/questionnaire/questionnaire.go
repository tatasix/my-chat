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
)

type QuestionnaireService struct {
	logx.Logger
	ctx                             context.Context
	svcCtx                          *svc.ServiceContext
	QuestionnaireResponseRepository *repository.QuestionnaireResponseRepository
	QuestionnaireResultRepository   *repository.QuestionnaireResultRepository
	QuestionnaireQuestionRepository *repository.QuestionnaireQuestionRepository
	CustomerConfigRepository        *repository.CustomerConfigRepository
	ResultMbtiRepository            *repository.QuestionnaireResultMbtiRepository
	WechatUserRepository            *repository.WechatUserRepository
}

func NewQuestionnaireService(ctx context.Context, svcCtx *svc.ServiceContext) *QuestionnaireService {
	return &QuestionnaireService{
		Logger:                          logx.WithContext(ctx),
		ctx:                             ctx,
		svcCtx:                          svcCtx,
		QuestionnaireResponseRepository: repository.NewQuestionnaireResponseRepository(ctx, svcCtx),
		QuestionnaireResultRepository:   repository.NewQuestionnaireResultRepository(ctx, svcCtx),
		QuestionnaireQuestionRepository: repository.NewQuestionnaireQuestionRepository(ctx, svcCtx),
		CustomerConfigRepository:        repository.NewCustomerConfigRepository(ctx, svcCtx),
		ResultMbtiRepository:            repository.NewQuestionnaireResultMbtiRepository(ctx, svcCtx),
		WechatUserRepository:            repository.NewWechatUserRepository(ctx, svcCtx),
	}
}

func (l *QuestionnaireService) InsertQuestionnaireQuestion(questionnaireQuestion *model.QuestionnaireQuestion) (lastId int64, err error) {
	return l.QuestionnaireQuestionRepository.Insert(questionnaireQuestion)
}

func (l *QuestionnaireService) GetQuestion(kfId string) (question []*model.QuestionnaireQuestion, err error) {
	return l.QuestionnaireQuestionRepository.GetByKfId(kfId)
}

func (l *QuestionnaireService) GetResult1(kfId, user string) (question []*model.QuestionnaireResult, err error) {
	customerPo, err := l.QuestionnaireResultRepository.GetLastResult(kfId, user)
	if err != nil {
		return
	}
	if customerPo != nil && customerPo.RelationId > 0 {
		return l.QuestionnaireResultRepository.GetByRelationId(customerPo.RelationId)
	}
	return
}

func (l *QuestionnaireService) GetResult(kfId, user string) (results *types.GetQuestionnaireResultResponse, err error) {
	//获取客服信息
	customerConfig, err := l.CustomerConfigRepository.GetByKfIdUseCache(kfId)
	if err != nil {
		return
	}
	if customerConfig == nil || customerConfig.Id <= 0 {
		err = util.ReturnError(xerr.ChatCustomerNotExist)
		return
	}
	var config model.CustomerConfigConfig
	_ = json.Unmarshal([]byte(customerConfig.Config), &config)
	if config.Method == 0 {
		err = util.ReturnError(xerr.ConfigEmpty)
		return
	}

	return NewQuestionnaireHandler(context.Background(), l.svcCtx, config.Method).GetReturn(customerConfig, user)
}

func (l *QuestionnaireService) GetResponse(kfId, user string, relation int64) (question []*model.QuestionnaireResponse, err error) {
	if relation == 0 {
		return
	}
	return l.QuestionnaireResponseRepository.GetResponse(kfId, user, relation)
}

func (l *QuestionnaireService) SaveQuestionnaire(questionnaireQuestion *model.QuestionnaireResponse, status int64) (err error) {
	//幂等性校验
	exist, err := l.QuestionnaireResponseRepository.GetByRelationIdAndQuestionId(questionnaireQuestion.RelationId, questionnaireQuestion.QuestionId)
	if err != nil {
		return
	}
	if exist != nil && exist.Id > 0 {
		if exist.OptionId == questionnaireQuestion.OptionId {
			//重复请求
			err = util.ReturnError(xerr.SystemBusyError)
			return
		}
	}

	allResponse, err := l.QuestionnaireResponseRepository.GetAll("", "", questionnaireQuestion.RelationId)
	if err != nil {
		return
	}
	//if len(allResponse) != int(status-1) {
	l.Logger.Infof("SaveQuestionnaire err allResponse:%d status:%d", len(allResponse), status)
	//	err = util.ReturnError(xerr.QuestionnaireStatusError)
	//	return
	//}

	//获取客服信息
	customerConfig, err := l.CustomerConfigRepository.GetByKfIdUseCache(questionnaireQuestion.OpenKfId)
	if err != nil {
		return
	}
	if customerConfig == nil || customerConfig.Id <= 0 {
		err = util.ReturnError(xerr.ChatCustomerNotExist)
		return
	}
	var config model.CustomerConfigConfig
	_ = json.Unmarshal([]byte(customerConfig.Config), &config)
	if config.Method == 0 {
		err = util.ReturnError(xerr.ConfigEmpty)
		return
	}
	total := config.Total
	if status <= 0 || status > total {
		err = util.ReturnError(xerr.QuestionnaireStatusError)
		return
	}
	// 获取问题信息
	question, err := l.QuestionnaireQuestionRepository.GetById(questionnaireQuestion.QuestionId)
	if err != nil {
		return
	}
	if question == nil || question.Id <= 0 {
		err = util.ReturnError(xerr.ConfigEmpty)
		return
	}
	var options []model.QuestionnaireQuestionOption
	_ = json.Unmarshal([]byte(question.Option), &options)
	option, err := l.QuestionnaireQuestionRepository.GetOption(options, questionnaireQuestion.OptionId)
	if err != nil {
		return
	}
	if option.Id == 0 {
		err = util.ReturnError(xerr.ConfigEmpty)
		return
	}
	optionStr, _ := json.Marshal(option)
	questionnaireQuestion.Score = option.Score
	questionnaireQuestion.ScoreType = question.ScoreType
	questionnaireQuestion.SnapshotOption = string(optionStr)
	questionnaireQuestion.Mbti = option.Mbti
	if exist != nil && exist.Id > 0 && exist.OptionId != questionnaireQuestion.OptionId {
		//修改问题的答案
		exist.OptionId = questionnaireQuestion.OptionId
		exist.Score = questionnaireQuestion.Score
		exist.ScoreType = questionnaireQuestion.ScoreType
		exist.SnapshotOption = questionnaireQuestion.SnapshotOption
		exist.MessageId = questionnaireQuestion.MessageId
		exist.Question = questionnaireQuestion.Question
		exist.Answer = questionnaireQuestion.Answer
		exist.Mbti = questionnaireQuestion.Mbti

		if err = l.QuestionnaireResponseRepository.Update(exist); err != nil {
			return err
		}

	} else {
		// 获取result表看看有咩有数据
		resultModel := &model.QuestionnaireResult{
			RelationId: questionnaireQuestion.RelationId,
			User:       questionnaireQuestion.User,
			OpenKfId:   questionnaireQuestion.OpenKfId,
			//ScoreType:  question.ScoreType,
			Status: status,
		}
		result, err1 := l.QuestionnaireResultRepository.GetByRelationIdAndType(questionnaireQuestion.RelationId, 0)
		if err1 != nil {
			return err1
		}
		if result != nil && result.Id >= 0 {
			if int(result.Status) == repository.QuestionnaireResultStatusEnd {
				err = util.ReturnError(xerr.QuestionnaireHaveDone)
				return
			}
			if result.OpenKfId != questionnaireQuestion.OpenKfId || result.User != questionnaireQuestion.User {
				err = util.ReturnError(xerr.QuestionnaireRelationError)
				return
			}
			//更新
			resultModel.Id = result.Id
			resultModel.CreatedAt = result.CreatedAt
			err = l.QuestionnaireResultRepository.Update(resultModel)

		} else {
			err = l.QuestionnaireResultRepository.Insert(resultModel)
		}
		if err != nil {
			return err
		}
		//保存问题的答案
		_, err = l.QuestionnaireResponseRepository.Insert(questionnaireQuestion)
		if err != nil {
			return err
		}
	}
	//<<<<<<< HEAD
	//	if len(allResponse) == int(total-1) {
	//		go NewQuestionnaireHandler(context.Background(), l.svcCtx, config.Method).Handle(customerConfig, config, questionnaireQuestion.RelationId)
	//	}
	//=======

	//if len(allResponse) == int(total-1) {
	//
	//	go NewQuestionnaireHandler(context.Background(), l.svcCtx, config.Method).Handle(customerConfig, config, questionnaireQuestion.RelationId)
	//
	//}

	return
}

func (l *QuestionnaireService) GenerateResult(relationId, user, openKfID string, channel chan string, baseInfo chan any) (resp *types.GenerateResultResponse, err error) {
	// TODO 幂等性校验
	resp = &types.GenerateResultResponse{}
	//获取客服信息
	customerConfig, err := l.CustomerConfigRepository.GetByKfIdUseCache(openKfID)
	if err != nil {
		return
	}
	if customerConfig == nil || customerConfig.Id <= 0 {
		err = util.ReturnError(xerr.ChatCustomerNotExist)
		return
	}
	var config model.CustomerConfigConfig
	_ = json.Unmarshal([]byte(customerConfig.Config), &config)
	if config.Method == 0 {
		err = util.ReturnError(xerr.ConfigEmpty)
		return
	}

	return NewQuestionnaireHandler(context.Background(), l.svcCtx, config.Method).Handle(customerConfig, config, util.ConvertToInt64(relationId), user, openKfID, channel, baseInfo)
}

func (l *QuestionnaireService) PlayAnother(questionnaireQuestion *model.QuestionnaireResult) (err error) {

	result, err := l.QuestionnaireResultRepository.GetByRelationId(questionnaireQuestion.RelationId)
	if err != nil {
		return
	}
	if result != nil && len(result) > 0 {
		err = util.ReturnError(xerr.SystemBusyError)
		return
	} else {
		err = l.QuestionnaireResultRepository.Insert(questionnaireQuestion)
	}

	return
}

func (l *QuestionnaireService) GetMbtiResultByRelationId(relationId int64) (result *types.GetMbtiResultByRelationResponse, err error) {
	result = &types.GetMbtiResultByRelationResponse{}
	mbtiResult, err := l.ResultMbtiRepository.GetByRelationId(relationId)
	if err != nil {
		return
	}
	if mbtiResult == nil || mbtiResult.Id <= 0 {
		return
	}
	wechatUser, err := l.WechatUserRepository.GetByUser(mbtiResult.User)
	if err != nil {
		return
	}
	if wechatUser == nil || wechatUser.Id <= 0 {
		err = util.ReturnError(xerr.LoginAccountNotExist)
		return
	}
	return assembler.POTODTOGetMbtiResultByRelationId(wechatUser, mbtiResult), nil
}
