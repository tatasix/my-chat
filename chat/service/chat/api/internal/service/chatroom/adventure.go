package chatroom

import (
	"chat/common/chatroom"
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/cache"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/service"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

const (
	StatusZero  = 0
	StatusOne   = 1
	StatusTwo   = 2
	StatusThree = 3
	StatusFour  = 4
	StatusFive  = 5
	StatusSix   = 6
)

var StatusMap = map[uint8]uint8{
	StatusZero:  StatusOne,
	StatusOne:   StatusTwo,
	StatusTwo:   StatusThree,
	StatusThree: StatusFour,
	StatusFour:  StatusFive,
	StatusFive:  StatusSix,
}

type AdventureHandler struct {
	logger logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func (a *AdventureHandler) StatusZero(msg chatroom.Message, customerInfo *model.CustomerConfig, channel chan string) (res chatroom.Message, err error) {

	//获取prompt
	prompt, err := service.NewPromptService(a.ctx, a.svcCtx).GetByKfIdAndStateUseCache(msg.KfId, int32(StatusMap[msg.Status]))
	if err != nil {
		return
	}

	if prompt == nil || prompt.Id <= 0 || prompt.Prompt == "" {
		err = util.ReturnError(xerr.ChatPromptEmpty)
		return
	}

	a.logger.Infof(" AdventureHandler StatusOne prompt:%s", prompt.Prompt)

	//调用ai接口获取问题的答案
	user, err := a.GetUserByChatRoom(msg.ChannelId)
	if err != nil {
		err = util.ReturnError(xerr.SystemError)
		return
	}
	aiChatService := service.NewAiChatService(a.ctx, a.svcCtx)
	summary := aiChatService.SetSystemAndUserPrompt(prompt.Prompt, "")
	//chatMessage, err := aiChatService.CustomerChat(customerConfig, summary)
	chatMessage, err := aiChatService.CustomerChatStream(customerInfo, summary, channel, user)

	if err != nil {
		a.logger.Error(" GenerateTarot:CustomerChat:error::" + err.Error())
		err = util.ReturnError(xerr.SystemError)
		return
	}
	a.logger.Infof("  AdventureHandler StatusOne  chatMessage:%s", chatMessage)
	_, _ = service.NewChatRoomRecordService(a.ctx, a.svcCtx).Insert(&model.ChatRoomRecord{
		ChannelId:   util.ConvertToInt64(msg.ChannelId),
		SendUserId:  msg.KfId,
		Message:     chatMessage,
		MessageId:   msg.MessageId,
		Status:      int64(StatusMap[msg.Status]),
		DialogueId:  msg.DialogueId,
		MessageType: repository.MessageTypeCustomer,
	})

	cache.SetCache(a.ctx, aiChatService.SetAssistantPrompt(summary, chatMessage), "adventure", msg.DialogueId, msg.KfId)

	now := time.Now()
	//timeout := time.After(65 * time.Second)

	res = msg
	res.SendTime = now.Format(vars.TimeFormat)
	res.IsEnd = true
	res.IsKf = true
	res.Content = chatMessage
	res.Type = chatroom.SendMessageToCustomer
	res.Status = StatusMap[msg.Status]

	return
	// 等待倒计时结束或超时
	//select {
	//case <-timeout:
	//	a.logger.Infof(" AdventureHandler StatusOne 倒计时结束，触发其他逻辑")
	//	// 在这里添加你希望在倒计时结束后触发的逻辑
	//	return chatroom.Message{}, a.DealTimeOut(msg) // 设置一个超时错误返回码
	//default:
	//	return
	//}
}

// StatusSix 总结
func (a *AdventureHandler) StatusSix(msg chatroom.Message, customerInfo *model.CustomerConfig, channel chan string) (res chatroom.Message, err error) {
	//获取prompt
	prompt, err := service.NewPromptService(a.ctx, a.svcCtx).GetByKfIdAndStateUseCache(msg.KfId, int32(msg.Status))
	if err != nil {
		return
	}

	if prompt == nil || prompt.Id <= 0 || prompt.Prompt == "" {
		err = util.ReturnError(xerr.ChatPromptEmpty)
		return
	}
	//调用ai接口获取问题的答案
	user, err := a.GetUserByChatRoom(msg.ChannelId)
	if err != nil {
		err = util.ReturnError(xerr.SystemError)
		return
	}
	//调用ai接口获取问题的答案
	aiChatService := service.NewAiChatService(a.ctx, a.svcCtx)
	summary := cache.GetCache(a.ctx, "adventure", msg.DialogueId, msg.KfId)
	summary = aiChatService.SetUserPrompt(summary, "")
	summary = aiChatService.SetSystemPrompt(summary, prompt.Prompt)
	a.logger.Infof(" AdventureHandler prompt:%+v", summary)
	chatMessage, err := aiChatService.CustomerChatStream(customerInfo, summary, channel, user)

	if err != nil {
		a.logger.Error(" AdventureHandler:CustomerChatStream:error::" + err.Error())
		err = util.ReturnError(xerr.SystemError)
		return
	}

	_, _ = service.NewChatRoomRecordService(a.ctx, a.svcCtx).Insert(&model.ChatRoomRecord{
		ChannelId:   util.ConvertToInt64(msg.ChannelId),
		SendUserId:  msg.KfId,
		Message:     chatMessage,
		MessageId:   msg.MessageId,
		Status:      int64(msg.Status),
		MessageType: repository.MessageTypeSummarize,
		DialogueId:  msg.DialogueId,
	})
	cache.DeleteCache(a.ctx, "adventure", msg.DialogueId, msg.KfId)
	now := time.Now()

	res = msg
	res.SendTime = now.Format(vars.TimeFormat)
	res.IsEnd = true
	res.IsKf = true
	res.Content = chatMessage
	return
}

func (a *AdventureHandler) Handle(msg chatroom.Message, customerInfo *model.CustomerConfig, channel chan string) (res chatroom.Message, err error) {
	switch msg.Status {
	case StatusZero:
		return a.StatusZero(msg, customerInfo, channel)

	case StatusOne, StatusTwo, StatusThree, StatusFour, StatusFive:
		return a.DealMessage(msg, customerInfo, channel)

	case StatusSix:
		return a.StatusSix(msg, customerInfo, channel)

	default:
		return chatroom.Message{}, util.ReturnError(xerr.RequestError)
	}
}

func (a *AdventureHandler) DealMessage(msg chatroom.Message, customerInfo *model.CustomerConfig, channel chan string) (res chatroom.Message, err error) {
	//判断消息是第一条还是第二条
	userOne, err1 := cache.GetDialogueCache(a.ctx, msg.KfId, msg.DialogueId, msg.Status)
	if err1 != nil || userOne == "" {

		cache.SetDialogueCache(a.ctx, a.formatMessage(msg.UserName, msg.Content), msg.KfId, msg.DialogueId, msg.Status)

		//只用保存缓存，并存储消息日志

		_, _ = service.NewChatRoomRecordService(a.ctx, a.svcCtx).Insert(&model.ChatRoomRecord{
			ChannelId:   util.ConvertToInt64(msg.ChannelId),
			SendUserId:  msg.User,
			Message:     msg.Content,
			MessageId:   msg.MessageId,
			Status:      int64(msg.Status),
			MessageType: repository.MessageTypeUser,
			DialogueId:  msg.DialogueId,
		})

		return

	}

	prompt, err := service.NewPromptService(a.ctx, a.svcCtx).GetByKfIdAndStateUseCache(msg.KfId, int32(StatusMap[msg.Status]))
	if err != nil {
		return
	}

	if prompt == nil || prompt.Id <= 0 || prompt.Prompt == "" {
		err = util.ReturnError(xerr.ChatPromptEmpty)
		return
	}

	a.logger.Infof(" AdventureHandler DealMessage prompt:%s", prompt.Prompt)

	//调用ai接口获取问题的答案
	user, err := a.GetUserByChatRoom(msg.ChannelId)
	if err != nil {
		err = util.ReturnError(xerr.SystemError)
		return
	}

	//调用ai接口获取问题的答案
	aiChatService := service.NewAiChatService(a.ctx, a.svcCtx)
	summary := cache.GetCache(a.ctx, "adventure", msg.DialogueId, msg.KfId)

	content := userOne + a.formatMessage(msg.UserName, msg.Content)
	summary = aiChatService.SetUserPrompt(summary, content)
	summary = aiChatService.SetSystemPrompt(summary, prompt.Prompt)
	a.logger.Infof(" AdventureHandler prompt:%+v", summary)
	chatMessage, err := aiChatService.CustomerChatStream(customerInfo, summary, channel, user)

	if err != nil {
		a.logger.Error(" AdventureHandler:CustomerChatStream:error::" + err.Error())
		err = util.ReturnError(xerr.SystemError)
		return
	}
	a.logger.Infof("  AdventureHandler DealMessage  chatMessage:%s", chatMessage)
	_, _ = service.NewChatRoomRecordService(a.ctx, a.svcCtx).Insert(&model.ChatRoomRecord{
		ChannelId:   util.ConvertToInt64(msg.ChannelId),
		SendUserId:  msg.User,
		Message:     msg.Content,
		MessageId:   msg.MessageId,
		Status:      int64(msg.Status),
		MessageType: repository.MessageTypeUser,
		DialogueId:  msg.DialogueId,
	})

	lastId, _ := service.NewChatRoomRecordService(a.ctx, a.svcCtx).Insert(&model.ChatRoomRecord{
		ChannelId:   util.ConvertToInt64(msg.ChannelId),
		SendUserId:  msg.KfId,
		Message:     chatMessage,
		MessageId:   msg.MessageId,
		Status:      int64(StatusMap[msg.Status]),
		MessageType: repository.MessageTypeCustomer,
		DialogueId:  msg.DialogueId,
	})
	cache.SetCache(a.ctx, aiChatService.SetAssistantPrompt(summary, chatMessage), "adventure", msg.DialogueId, msg.KfId)
	now := time.Now()
	//timeout := time.After(80 * time.Second)
	if msg.Status == StatusFive {
		go service.NewRiskService(context.Background(), a.svcCtx).Reduce(user, msg.KfId, lastId)
	}

	res = msg
	res.Status = StatusMap[msg.Status]
	res.SendTime = now.Format(vars.TimeFormat)
	res.IsEnd = true
	res.IsKf = true
	res.Content = chatMessage
	return
	// 开始倒计时
	// 等待倒计时结束或超时
	//select {
	//case <-timeout:
	//	a.logger.Infof(" AdventureHandler DealMessage 倒计时结束，触发其他逻辑")
	//	// 在这里添加你希望在倒计时结束后触发的逻辑
	//	return chatroom.Message{}, a.DealTimeOut(msg)
	//default:
	//	return
	//}
}

func (a *AdventureHandler) DealTimeOut(msg chatroom.Message) (err error) {
	//倒计时结束，
	//如果当前状态总共只有一条记录，或者没有记录，那就说明需要结束这次聊天
	//如果当前状态有两条记录，需要主动提交（前端控制）
	//如果当前状态有三条记录，那就不用处理，已经有别的流程处理过了
	info, err := service.NewChatRoomRecordService(a.ctx, a.svcCtx).GetByDialogueIdAndStatus(msg.DialogueId, msg.Status)
	if err != nil {
		a.logger.Error(" AdventureHandler:DealTimeOut:error::" + err.Error())
		return
	}
	if info == nil || len(info) < 3 {
		err = util.ReturnError(xerr.ChatRoomOut)
	}
	return
}

func (a *AdventureHandler) formatMessage(userName, content string) string {
	return fmt.Sprintf("%s:%s", userName, content)
}

func (a *AdventureHandler) GetUserByChatRoom(id string) (res string, err error) {
	chatroomService := service.NewChatRoomService(a.ctx, a.svcCtx)
	chatroomInfo, err := chatroomService.GetChatRoomByChannelId(util.ConvertToInt64(id))
	if err != nil {
		a.logger.Error(" GenerateTarot:CustomerChat:error::" + err.Error())
		err = util.ReturnError(xerr.SystemError)
		return
	}
	if chatroomInfo == nil || chatroomInfo.Id <= 0 {
		return
	}
	res = chatroomInfo.User
	return
}
