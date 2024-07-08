package im

import (
	"chat/common/chatroom"
	"chat/common/response"
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/service"
	chatroomService "chat/service/chat/api/internal/service/chatroom"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/util/sensitive"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type Store struct {
	Svc *svc.ServiceContext
	Ctx context.Context
}

// SendMessage 发送消息回调
func (d *Store) SendMessage(msg chatroom.Message) {
	//处理消息
	channel := make(chan string, 50)
	go func() {
		defer func() {
			close(channel)
			// 处理 goroutine 泄漏
			recover()
		}()

		res, errChat := d.handleMessage(msg, channel)
		if errChat != nil {
			logx.Error("SendMessage error :", errChat)
			res = msg
			res.ErrorMessage = response.GetErrorMessage(errChat)
			if errChat == util.ReturnError(xerr.SystemBusyError) { //重复请求，忽略处理
				return
			}
		}
		logx.Infof("SendMessage message:%+v", res)
		//returnMessage, _ := json.Marshal(res)
		//发送消息
		//time.Sleep(3 * time.Second)
		SendAllMessage(msg.User, res)

	}()

	var rs []rune
	// 加快初次响应的时间 后续可改为阶梯式（用户体验好）
	length := 1
	var users []interface{}
	var isKf bool
	for {
		s, ok := <-channel
		if !ok {
			// 数据接受完成
			if len(rs) > 0 {
				isKf = true
				d.SendSocketMessage(msg.User, msg.ToUserId, string(rs), msg.MessageId, msg.ChannelId, msg.DialogueId, msg.Status, msg.Type, true, false, false, users)
				rs = []rune{}
			}

			d.SendSocketMessage(msg.User, msg.ToUserId, "", msg.MessageId, msg.ChannelId, msg.DialogueId, msg.Status, msg.Type, isKf, false, true, users)

			break
		}
		rs = append(rs, []rune(s)...)

		if len(rs) > length {
			isKf = true
			d.SendSocketMessage(msg.User, msg.ToUserId, string(rs), msg.MessageId, msg.ChannelId, msg.DialogueId, msg.Status, msg.Type, true, false, false, users)
			rs = []rune{}
			if length < 4 {
				length++
			}
		}
	}

}

func (d *Store) handleMessage(msg chatroom.Message, channel chan string) (res chatroom.Message, err error) {
	res = chatroom.Message{}
	switch msg.Type {
	case chatroom.SendMessageToCustomer:
		return d.HandelToCustomer(msg, channel)

	case chatroom.SendMessageToUserWriting:
		return d.HandelUserWriting()

	case chatroom.SendMessageGenerateOne:
		return d.HandelSendMessageGenerateOne(msg, channel)

	case chatroom.SendMessageCreate:
		//创建聊天室，并加入
		return d.HandelCreate(msg)

	case chatroom.SendMessageJoin:
		//邀请加入聊天室
		return d.HandelJoin(msg)

	case chatroom.SendMessageOut:
		//退出聊天室
		return d.HandelOut(msg)

	case chatroom.SendMessageStart:
		//游戏开始
		return d.HandelStart()

	case chatroom.SendMessageReady:
		//B进入聊天
		return d.HandelReady()

	default:
		return chatroom.Message{}, nil
	}
}

func (d *Store) DelaySendMessage(channelId string, msg chatroom.Message, sent []string) {
	//fmt.Println("delay：", channelId, msg, sent)
	//var ids []int64
	//switch msg.Type {
	//case 101:
	//	users, _ := d.ctx.GroupUserModel.FindUsersByChannelId(channelId)
	//	for _, u := range users {
	//		ids = append(ids, u.UserId)
	//	}
	//case 100:
	//	//var users []hasusers.UserUsers
	//	if channelId == "" {
	//		ids = append(ids, int64(msg.ToUserId))
	//	} else {
	//		if len(sent) == 2 { // 单聊发送人数为2，则无离线消息 return
	//			return
	//		}
	//		//users, _ = d.ctx.UserUsersModel.AllChannelIdUsers(channelId)
	//		//for _, u := range users {
	//		//	ids = append(ids, u.UserId)
	//		//}
	//	}
	//}

	//sentMap := make(map[int64]bool)
	//for _, v := range sent {
	//	sentMap[int64(v)] = true
	//}
	//msgByte, err := json.Marshal(msg)
	//if err != nil {
	//	return
	//}
	//for _, uid := range ids {
	//	if _, ok := sentMap[uid]; !ok {
	//		d.ctx.SendQueueModel.Insert(context.Background(), &sendqueue.SendQueues{UserId: uid, Message: string(msgByte), SendUserId: int64(msg.UserId)})
	//	}
	//}
}

// LoginServer 登录成功后回调
func (d *Store) LoginServer(uid string) {
	//list, _ := d.ctx.UserModel.Friends(d.ctx.UserUsersModel, int64(uid))
	//list, _ := d.ctx.UserUsersModel.Friends(int64(uid))
	//var channelIds []string
	//for _, v := range list {
	//	channelIds = append(channelIds, GenChannelIdByFriend(int64(uid), v.HasUserId))
	//}
	//groups, _ := d.ctx.GroupUserModel.InGroups(context.Background(), int64(uid))
	//for _, v := range groups {
	//	channelIds = append(channelIds, v.ChannelId)
	//}
	//chatroom.JoinChannelIds(uid, channelIds...)
	//go func() {
	//	time.Sleep(time.Second * 1)
	//	queues, _ := d.ctx.SendQueueModel.FindByUserId(context.Background(), int64(uid))
	//	for _, queue := range queues {
	//		SendMessageToUid(uid, uid, queue.Message, 100)
	//		d.ctx.SendQueueModel.Delete(context.Background(), queue.Id)
	//	}
	//}()
}
func (d *Store) LogoutServer(uid string) {
	// 退出登陆回调
	//fmt.Println("logout ", uid)
}
func (d *Store) ErrorLogServer(err error) {
	// 错误消息回调
	logger := logx.WithContext(d.Ctx)
	logger.Errorf("Store err: %v", err)
}

func (d *Store) InfoLogServer(msg string) {
	// 错误消息回调
	logger := logx.WithContext(d.Ctx)
	logger.Infof("Store info: %v", msg)
}

func (d *Store) HandelToCustomer(msg chatroom.Message, channel chan string) (res chatroom.Message, err error) {
	//幂等性
	flag, err := service.NewChatRoomRecordService(d.Ctx, d.Svc).Idempotent(msg.DialogueId, msg.User, int8(msg.Status))
	if err != nil {
		return
	}
	if flag {
		err = util.ReturnError(xerr.SystemBusyError)
		return
	}
	if msg.Content != "" {
		if !sensitive.NewSensitive(d.Ctx, d.Svc).Check(msg.Content) {
			res = chatroom.Message{
				Content:      sensitive.ErrorMessage,
				ChannelId:    msg.ChannelId,
				ChannelTitle: msg.ChannelTitle,
				User:         msg.User,
				Users:        msg.Users,
				UserName:     msg.UserName,
				KfId:         msg.KfId,
				Type:         msg.Type,
				Status:       msg.Status,
				MessageId:    msg.MessageId,
				DialogueId:   msg.DialogueId,
				IsEnd:        true,
				ErrorMessage: sensitive.ErrorMessage,
			}
			return res, nil
		}
	}
	//获取客服相关信息
	customerInfo, err := service.NewCustomerConfigService(d.Ctx, d.Svc).GetByKfId(msg.KfId)
	if nil != err {
		return
	}
	if customerInfo == nil || customerInfo.Id <= 0 {
		err = util.ReturnError(xerr.ChatCustomerNotExist)
		return
	}
	handler := chatroomService.NewChatRoomHandler(d.Ctx, d.Svc, msg.KfId)
	if handler == nil {
		err = util.ReturnError(xerr.SystemError)
		return
	}

	res, err = handler.Handle(msg, customerInfo, channel)
	if err == util.ReturnError(xerr.ChatRoomOut) {
		// 退出
		err = d.UnJoinChannelIds(msg)
	}
	return
}

func (d *Store) HandelUserWriting() (res chatroom.Message, err error) {
	return
}

func (d *Store) HandelStart() (res chatroom.Message, err error) {
	return
}

func (d *Store) HandelReady() (res chatroom.Message, err error) {
	return
}

func (d *Store) HandelCreate(msg chatroom.Message) (res chatroom.Message, err error) {
	err = JoinChannelIds(msg.User, msg.ChannelId)
	if err != nil {
		d.ErrorLogServer(err)
	}
	users := GetAllUser(msg.User, msg.ChannelId)
	d.SendSocketMessage(msg.User, msg.User, msg.UserName+" join the channel", msg.MessageId, msg.ChannelId, msg.DialogueId, msg.Status, msg.Type, false, true, false, users)
	return
}
func (d *Store) HandelJoin(msg chatroom.Message) (res chatroom.Message, err error) {

	err = JoinChannelIds(msg.User, msg.ChannelId)
	if err != nil {
		d.ErrorLogServer(err)
		return
	}
	chatRoomService := service.NewChatRoomService(d.Ctx, d.Svc)
	info, err := chatRoomService.GetChatRoomByChannelId(util.ConvertToInt64(msg.ChannelId))
	if err != nil {
		d.ErrorLogServer(err)
		return
	}
	if info == nil {
		return
	}
	users := GetAllUser(msg.User, msg.ChannelId)

	d.SendSocketMessage(msg.User, "", msg.UserName+" join the channel", msg.MessageId, msg.ChannelId, msg.DialogueId, msg.Status, msg.Type, false, true, false, users)
	return
}

func (d *Store) HandelOut(msg chatroom.Message) (res chatroom.Message, err error) {
	//退出聊天室，关闭聊天室，数据表也要处理
	err = d.UnJoinChannelIds(msg)
	if err != nil {
		d.ErrorLogServer(err)
	}
	return
}

func (d *Store) SendSocketMessage(user, toUser, message, messageId, channelId, dialogueId string, status, tp uint8, isKf, isEnd, isStreamEnd bool, users []interface{}) {
	if tp == chatroom.SendMessageGenerateOne {
		tp = chatroom.SendMessageToCustomer
	}
	SendAllMessage(user, chatroom.Message{
		User:        user,
		ToUserId:    toUser,
		Type:        tp,
		Status:      status,
		Content:     message,
		MessageId:   messageId,
		ChannelId:   channelId,
		IsKf:        isKf,
		IsEnd:       isEnd,
		DialogueId:  dialogueId,
		IsStreamEnd: isStreamEnd,
		Users:       users,
	})
}

func (d *Store) UnJoinChannelIds(msg chatroom.Message) (err error) {
	//推送一条要退出的消息
	//var users []interface{}
	//d.SendSocketMessage(msg.User, "", msg.UserName+"断开链接", msg.MessageId, msg.ChannelId, msg.DialogueId, 0, chatroom.SendMessageDisconnect, true, true, false, users)

	chatRoomService := service.NewChatRoomService(d.Ctx, d.Svc)
	channelId := util.ConvertToInt64(msg.ChannelId)
	//退出chat room user
	err = chatRoomService.DeleteChatRoomUsersByChannelId(channelId)
	if nil != err {
		return
	}
	//退出chat room
	//err = chatRoomService.Delete(channelId)
	//if nil != err {
	//	return
	//}
	//退出channel
	err = UnJoinChannelIds(msg.User, msg.ChannelId)
	return
}

func (d *Store) HandelSendMessageGenerateOne(msg chatroom.Message, channel chan string) (res chatroom.Message, err error) {
	//幂等性
	msg.Status = 0
	flag, err := service.NewChatRoomRecordService(d.Ctx, d.Svc).Idempotent(msg.DialogueId, msg.User, -1)
	if err != nil {
		return
	}
	if flag {
		err = util.ReturnError(xerr.SystemBusyError)
		return
	}

	//获取客服相关信息
	customerInfo, err := service.NewCustomerConfigService(d.Ctx, d.Svc).GetByKfId(msg.KfId)
	if nil != err {
		return
	}
	if customerInfo == nil || customerInfo.Id <= 0 {
		err = util.ReturnError(xerr.ChatCustomerNotExist)
		return
	}
	handler := chatroomService.NewChatRoomHandler(d.Ctx, d.Svc, msg.KfId)
	if handler == nil {
		err = util.ReturnError(xerr.SystemError)
		return
	}

	res, err = handler.Handle(msg, customerInfo, channel)
	if err == util.ReturnError(xerr.ChatRoomOut) {
		// 退出
		err = d.UnJoinChannelIds(msg)
	}
	return
}
