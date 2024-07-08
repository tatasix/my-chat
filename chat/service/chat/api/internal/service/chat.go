package service

import (
	"chat/common/util"
	"chat/service/chat/api/internal/logic/assembler"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
	"context"
	"errors"
	"fmt"
	"github.com/gammazero/workerpool"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

const MaxExportNumber = 3000

var (
	RecordNotFoundErr = errors.New("数据不存在")
)

type ChatService struct {
	logx.Logger
	ctx                     context.Context
	svcCtx                  *svc.ServiceContext
	chatRepository          *repository.ChatRepository
	chatRecordRepository    *repository.ChatRecordRepository
	applicationConfig       *repository.ApplicationConfigRepository
	customerConfig          *repository.CustomerConfigRepository
	wechatUser              *repository.WechatUserRepository
	resourceUsageRepository *repository.ResourceUsageRepository
	wechatUserService       *WechatUserService
}

func NewChatService(ctx context.Context, svcCtx *svc.ServiceContext) *ChatService {
	return &ChatService{
		Logger:                  logx.WithContext(ctx),
		ctx:                     ctx,
		svcCtx:                  svcCtx,
		applicationConfig:       repository.NewApplicationConfigRepository(ctx, svcCtx),
		customerConfig:          repository.NewCustomerConfigRepository(ctx, svcCtx),
		chatRepository:          repository.NewChatRepository(ctx, svcCtx),
		chatRecordRepository:    repository.NewChatRecordRepository(ctx, svcCtx),
		wechatUser:              repository.NewWechatUserRepository(ctx, svcCtx),
		resourceUsageRepository: repository.NewResourceUsageRepository(ctx, svcCtx),
		wechatUserService:       NewWechatUserService(ctx, svcCtx),
	}
}

func (l *ChatService) FormatCondition(userNickname, kfName string, chatType int32) (userId, kfId string, agentId int64, err error) {
	if userNickname == "" || kfName == "" {
		return "", "", 0, fmt.Errorf("缺少必传参数")
	}
	if chatType == 2 {
		//客服聊天
		if userNickname != "" {
			//get userID by UserNickName
			wechatUserPo, err := l.wechatUser.GetByName(userNickname)
			if err != nil {
				return "", "", 0, err
			}

			if wechatUserPo != nil && wechatUserPo.User != "" {
				userId = wechatUserPo.User
			}
		}
		if kfName != "" {
			//get openKfID by OpenKfName
			kfPo, err := l.customerConfig.GetByName(kfName)
			if err != nil {
				return "", "", 0, err
			}

			if kfPo != nil && kfPo.KfId != "" {
				kfId = kfPo.KfId
			}
		}

	} else {
		//机器人聊天
		userId = userNickname
		if kfName != "" {
			//get openKfID by OpenKfName
			applicationPo, err := l.applicationConfig.GetByName(kfName)
			if err != nil {
				return "", "", 0, err
			}

			if applicationPo != nil && applicationPo.AgentId != 0 {
				agentId = applicationPo.AgentId
			}
		}

	}
	return
}

func (l *ChatService) GetChatList(userName, chatRecordUser, kfName, agentName, startCreatedAt, endCreatedAt string, chatType int32, order string, page, pageSize int) (resp *types.GetChatListPageResult, err error) {

	userId, openKfId, agentId, err := l.FormatRequestCondition(userName, kfName, agentName, chatType)
	if err != nil {
		if err == RecordNotFoundErr {
			return &types.GetChatListPageResult{
				List:     nil,
				Total:    0,
				Page:     page,
				PageSize: pageSize,
			}, nil
		}
		return
	}
	chatPos, count, err := l.chatRecordRepository.GetAll(agentId, openKfId, userId, chatRecordUser, startCreatedAt, endCreatedAt, order, uint64(page), uint64(pageSize), chatType, 1)
	if err != nil {
		fmt.Printf("GetSystemConfig error: %v", err)
		return
	}
	if count <= 0 || len(chatPos) <= 0 {
		//resp.Page = page
		//resp.PageSize = pageSize
		//return
		resp = &types.GetChatListPageResult{
			List:     nil,
			Total:    0,
			Page:     page,
			PageSize: pageSize,
		}
		return
	}
	var users, customers []string
	var applications, chatIds []int64
	for _, v := range chatPos {
		if v.AgentId == 0 {
			customers = append(customers, v.OpenKfId)
			users = append(users, v.User)
		} else {
			applications = append(applications, v.AgentId)
		}
		chatIds = append(chatIds, v.Id)
	}
	answer, err := l.GetAnswerByQuestion(chatIds)
	if err != nil {
		return
	}
	wechatUserPos, customerPos, applicationPos, err := l.GetOtherPo(users, customers, applications)
	chatDtos := assembler.POTODTOGetChatList(chatPos, answer, wechatUserPos, applicationPos, customerPos)
	return &types.GetChatListPageResult{
		List:     chatDtos,
		Total:    count,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

func (l *ChatService) GetExportList(userName, chatRecordUser, kfName, agentName, startCreatedAt, endCreatedAt string, chatType int32, order string, aq int64, page, pageSize int) (chatRecord []*model.ExportChatRecord, count int64, err error) {

	userId, openKfId, agentId, err := l.FormatRequestCondition(userName, kfName, agentName, chatType)
	if err != nil {
		if err == RecordNotFoundErr {
			err = nil
			return
		}
		return
	}
	chatPos, count, err := l.chatRecordRepository.GetAll(agentId, openKfId, userId, chatRecordUser, startCreatedAt, endCreatedAt, order, uint64(page), uint64(pageSize), chatType, aq)
	if err != nil {
		return
	}

	if count <= 0 || len(chatPos) <= 0 {
		return
	}
	var users, customers []string
	var applications []int64
	for _, v := range chatPos {
		if v.AgentId == 0 {
			customers = append(customers, v.OpenKfId)
			users = append(users, v.User)
		} else {
			applications = append(applications, v.AgentId)
		}
	}

	wechatUserPos, customerPos, applicationPos, err := l.GetOtherPo(users, customers, applications)
	chatRecord = assembler.POTODTOGetChatExportList(chatPos, wechatUserPos, applicationPos, customerPos)

	return
}

func (l *ChatService) GetAllRecord(agentId int64, openKfId, user, chatRecordUser, startTime, endTime, order string, page, limit uint64, chatType int32) ([]*model.ChatRecord, int64, error) {
	return l.chatRecordRepository.GetAll(agentId, openKfId, user, chatRecordUser, startTime, endTime, order, page, limit, chatType, 0)
}

func (l *ChatService) BatchInsert(record []*model.ChatRecord) (err error) {
	if len(record) <= 0 {
		return
	}
	_, err = l.chatRecordRepository.BatchInsert(record)
	return
}

func (l *ChatService) GetAllOld(page, pageSize uint64, order string) ([]*model.Chat, int64, error) {
	return l.chatRepository.GetAll(0, "", "", "", "", order, page, pageSize, 0)
}

func (l *ChatService) GetByIds(ids []int64) ([]*model.ChatRecord, error) {
	return l.chatRecordRepository.GetByIds(ids)
}

func (l *ChatService) Insert(chatRecord *model.ChatRecord) (err error) {
	_, err = l.chatRecordRepository.Insert(chatRecord)
	return
}

func (l *ChatService) InsertReturnInsertId(chatRecord *model.ChatRecord) (insertId int64, err error) {
	res, err := l.chatRecordRepository.Insert(chatRecord)
	if err != nil {
		return
	}
	insertId, _ = res.LastInsertId()
	return
}

func (l *ChatService) InsertV2(chatRecord *model.ChatRecord) (lastId int64, err error) {
	res, err := l.chatRecordRepository.Insert(chatRecord)
	lastId, err = res.LastInsertId()
	return
}

func (l *ChatService) GetByMessageAndCustomer(messageId, customerId string) ([]*model.ChatRecord, error) {
	return l.chatRecordRepository.GetByMessageAndCustomer(messageId, customerId)
}

func (l *ChatService) GetByMessageAndState(relationId int64, state int) (*model.ChatRecord, error) {
	return l.chatRecordRepository.GetByMessageAndState(relationId, state)
}

func (l *ChatService) Idempotent(messageId, customerId string) (string, int64, error) {
	originInfo, err := l.GetByMessageAndCustomer(messageId, customerId)
	if err != nil {
		return "", 0, err
	}
	if originInfo != nil && len(originInfo) > 0 {
		if len(originInfo) > 1 {
			for _, v := range originInfo {
				if v.AnswerOrQuestion == repository.AnswerOrQuestionAnswer {
					return v.Content, v.Emoji, nil
				}
			}
		}
		return "系统繁忙，请稍后再试～", 0, nil
	}
	return "", 0, nil
}

func (l *ChatService) FormatRequestCondition(userName, kfName, agentName string, chatType int32) (userId, openKfId string, agentId int64, err error) {
	if agentName != "" {
		applicationPo, err1 := l.applicationConfig.GetByName(agentName)
		if nil != err1 {
			err = err1
			return
		}
		if applicationPo == nil || applicationPo.AgentId == 0 {
			err = RecordNotFoundErr
			return
		}
		agentId = applicationPo.AgentId
	}
	if chatType == 2 && userName != "" {
		wechatUserPo, err2 := l.wechatUser.GetByName(userName)
		if nil != err2 {
			err = err2
			return
		}
		if wechatUserPo == nil || wechatUserPo.User == "" {
			err = RecordNotFoundErr
			return
		}
		userId = wechatUserPo.User
	} else {
		userId = userName
	}
	if kfName != "" {
		applicationPo, err3 := l.customerConfig.GetByName(kfName)
		if nil != err3 {
			err = err3
			return
		}
		if applicationPo == nil || applicationPo.KfId == "" {
			err = RecordNotFoundErr
			return
		}
		openKfId = applicationPo.KfId
	}
	return
}

func (l *ChatService) GetOtherPo(users, customers []string, applications []int64) (wechatUserPos []*model.WechatUser, customerPos []*model.CustomerConfig, applicationPos []*model.ApplicationConfig, err error) {

	if len(users) > 0 {
		wechatUserPos, err = l.wechatUser.GetByUsers(util.Unique(users))
		if err != nil {
			fmt.Printf("GetSystemConfig error: %v", err)
			return
		}
	}
	if len(customers) > 0 {
		customerPos, err = l.customerConfig.GetByKfIds(util.Unique(customers))
		if err != nil {
			fmt.Printf("GetSystemConfig error: %v", err)
			return
		}
	}
	if len(applications) > 0 {
		applicationPos, err = l.applicationConfig.GetByIds(util.Unique(applications))
		if err != nil {
			fmt.Printf("GetSystemConfig error: %v", err)
			return
		}
	}
	return
}

func (l *ChatService) GetAll(userName, kfName, agentName, startCreatedAt, endCreatedAt string, chatType int32, order string, aq int64, page, pageSize int) (chatPos []*model.ChatRecord, count int64, err error) {

	userId, openKfId, agentId, err := l.FormatRequestCondition(userName, kfName, agentName, chatType)
	if err != nil {
		if err == RecordNotFoundErr {
			err = nil
			return
		}
		return
	}
	return l.chatRecordRepository.GetAll(agentId, openKfId, userId, "", startCreatedAt, endCreatedAt, order, uint64(page), uint64(pageSize), chatType, aq)

}

func (l *ChatService) GetAnswerByQuestion(chatIds []int64) (chatPos []*model.ChatRecord, err error) {
	return l.chatRecordRepository.GetAnswerByQuestion(chatIds)
}

func (l *ChatService) SummarizeChatLog(ctx context.Context, startId, endId int64, startTime, endTime string) (err error) {
	chat, err := l.chatRecordRepository.GetUser(ctx, startId, endId, repository.ChatTypeCustomer, startTime, endTime)
	if err != nil {
		return
	}
	if chat == nil || len(chat) <= 0 {
		return
	}
	chatFormat := l.FormatChat(chat)
	fmt.Println(chatFormat)

	wp := workerpool.New(2)

	for k, v := range chatFormat {
		r := v
		user := k
		wp.Submit(func() {
			l.Deal(user, r)
		})
	}

	wp.StopWait()

	return
}

func (l *ChatService) Deal(user string, content []string) {
	//var result string
	//for _, v := range content {
	//
	//	result, err := l.aiChatService.Chat(user, v, result)
	//
	//	fmt.Println(err)
	//	fmt.Println(result)
	//}
	//fmt.Println(result)
	//err := l.wechatUserService.UpdateByUser(user, &model.WechatUser{
	//	User: "",
	//})
	//fmt.Println(err)
	return
}

func (l *ChatService) FormatChat(chat []*model.ChatRecord) map[string][]string {
	res := make(map[string]string)
	AnswerOrQuestion := make(map[int64]string)
	for _, v := range chat {
		str := ""
		if value, ok := AnswerOrQuestion[v.RelationId]; ok {
			if v.AnswerOrQuestion == repository.AnswerOrQuestionQuestion {
				str = str + "q:" + v.Content + " a:" + value
			} else {
				str = str + "q:" + value + " a:" + v.Content
			}
		} else {
			AnswerOrQuestion[v.RelationId] = v.Content
			continue
		}
		res[v.User] = res[v.User] + str
	}
	result := make(map[string][]string)

	for kk, vv := range res {
		if len(vv) > vars.MaxToken {
			result[kk] = CutString(vv)
		} else {
			result[kk] = []string{vv}
		}
	}
	return result
}

func CutString(s string) []string {
	var results []string
	var start, end int
	for start < len(s) {
		end = findEndIndex(s, start)
		if end == -1 { // 没有找到结尾，只能全部加到一个返回结果中
			return []string{s}
		}
		results = append(results, s[start:end]) // 切割并加入结果中
		start = end
	}
	return results
}

// 找到s[start:]中下一个分隔符的位置
func findEndIndex(s string, start int) int {
	for i := start + vars.MaxToken; i > start; i-- { // 从start往后找，找到第一个q:后面的
		if i >= len(s) { // 如果已到达字符串结尾
			return -1
		} else if s[i:i+2] == "q:" { // 找到了分隔符
			return i + 2
		}
	}
	return start + vars.MaxToken // 没有找到分隔符，切割到最大值
}

func (l *ChatService) GetAllKfRecord(kfName, user, startCreatedAt, endCreatedAt string, state []int64, stateId int64) (resp []*model.ChatRecord, err error) {
	return l.chatRecordRepository.GetAllKfRecord(kfName, user, startCreatedAt, endCreatedAt, state, stateId)
}

func (l *ChatService) GetByRelationId(relation int64) (resp *model.ChatRecord, err error) {
	return l.chatRecordRepository.GetOne(relation, repository.AnswerOrQuestionQuestion)
}

func (l *ChatService) SaveTimesAndToken(user string, token int64) {

	err := l.resourceUsageRepository.SaveTimesAndToken(user, time.Now(), 1, token)
	if err != nil {
		l.Logger.Errorf("SaveTimesAndToken err:%+v", err)
	}
	return
}
