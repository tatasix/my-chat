package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"

	"chat/service/chat/api/internal/config"
	"chat/service/chat/api/internal/middleware"
	"chat/service/chat/model"
)

type ServiceContext struct {
	Config                       config.Config
	UserModel                    model.UserModel
	ChatModel                    model.ChatModel
	ChatConfigModel              model.ChatConfigModel
	CustomerPromptModel          model.CustomerPromptModel
	PromptConfigModel            model.PromptConfigModel
	ApplicationConfigModel       model.ApplicationConfigModel
	CustomerConfigModel          model.CustomerConfigModel
	WechatUserModel              model.WechatUserModel
	ChatRecordModel              model.ChatRecordModel
	FeedbackModel                model.FeedbackModel
	UserPortraitModel            model.UserPortraitModel
	ConfigModel                  model.ConfigModel
	PromptModel                  model.PromptModel
	StateModel                   model.StateModel
	ChatRoomRecordModel          model.ChatRoomRecordModel
	ChatRoomModel                model.ChatRoomModel
	ChatRoomUsersModel           model.ChatRoomUsersModel
	ActivityLogModel             model.ActivityLogModel
	StatisticsModel              model.StatisticsModel
	ResourceUsageModel           model.ResourceUsageModel
	QuestionnaireResultModel     model.QuestionnaireResultModel
	QuestionnaireResultMbtiModel model.QuestionnaireResultMbtiModel
	QuestionnaireResponseModel   model.QuestionnaireResponseModel
	QuestionnaireQuestionModel   model.QuestionnaireQuestionModel
	UserLikeModel                model.UserLikeModel
	RightsModel                  model.RightsModel
	RightsRecordModel            model.RightsRecordModel
	RightsTimesModel             model.RightsTimesModel
	ApplePaySubscriptionModel    model.ApplePaySubscriptionModel

	AccessLog rest.Middleware
	Trace     rest.Middleware
	Login     rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:                       c,
		UserModel:                    model.NewUserModel(conn, c.RedisCache),
		ChatModel:                    model.NewChatModel(conn, c.RedisCache),
		ChatConfigModel:              model.NewChatConfigModel(conn, c.RedisCache),
		PromptConfigModel:            model.NewPromptConfigModel(conn, c.RedisCache),
		CustomerPromptModel:          model.NewCustomerPromptModel(conn, c.RedisCache),
		ApplicationConfigModel:       model.NewApplicationConfigModel(conn, c.RedisCache),
		CustomerConfigModel:          model.NewCustomerConfigModel(conn, c.RedisCache),
		WechatUserModel:              model.NewWechatUserModel(conn, c.RedisCache),
		ChatRecordModel:              model.NewChatRecordModel(conn, c.RedisCache),
		FeedbackModel:                model.NewFeedbackModel(conn, c.RedisCache),
		UserPortraitModel:            model.NewUserPortraitModel(conn, c.RedisCache),
		ConfigModel:                  model.NewConfigModel(conn, c.RedisCache),
		PromptModel:                  model.NewPromptModel(conn, c.RedisCache),
		ChatRoomRecordModel:          model.NewChatRoomRecordModel(conn, c.RedisCache),
		ChatRoomModel:                model.NewChatRoomModel(conn, c.RedisCache),
		StateModel:                   model.NewStateModel(conn, c.RedisCache),
		ChatRoomUsersModel:           model.NewChatRoomUsersModel(conn, c.RedisCache),
		ActivityLogModel:             model.NewActivityLogModel(conn, c.RedisCache),
		StatisticsModel:              model.NewStatisticsModel(conn, c.RedisCache),
		ResourceUsageModel:           model.NewResourceUsageModel(conn, c.RedisCache),
		QuestionnaireResultModel:     model.NewQuestionnaireResultModel(conn, c.RedisCache),
		QuestionnaireResultMbtiModel: model.NewQuestionnaireResultMbtiModel(conn, c.RedisCache),
		QuestionnaireResponseModel:   model.NewQuestionnaireResponseModel(conn, c.RedisCache),
		QuestionnaireQuestionModel:   model.NewQuestionnaireQuestionModel(conn, c.RedisCache),
		UserLikeModel:                model.NewUserLikeModel(conn, c.RedisCache),
		RightsModel:                  model.NewRightsModel(conn, c.RedisCache),
		RightsRecordModel:            model.NewRightsRecordModel(conn, c.RedisCache),
		RightsTimesModel:             model.NewRightsTimesModel(conn, c.RedisCache),
		ApplePaySubscriptionModel:    model.NewApplePaySubscriptionModel(conn, c.RedisCache),

		AccessLog: middleware.NewAccessLogMiddleware().Handle,
		Trace:     middleware.NewTraceMiddleware().Handle,
		Login:     middleware.NewLoginMiddleware(c).Handle,
	}
}
