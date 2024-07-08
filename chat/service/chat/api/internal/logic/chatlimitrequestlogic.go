package logic

import (
	"chat/common/redis"
	"chat/service/chat/api/internal/logic/frontend"
	"chat/service/chat/api/internal/svc"
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"
)

const (
	ConstUserRequestTimes    int32  = 25
	ConstUserRequestTimesPre string = "user:requested:times:%s"
)

type ChatLimitRequestLogic struct {
	Ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatLimitRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatLimitRequestLogic {
	return &ChatLimitRequestLogic{
		Ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *ChatLimitRequestLogic) ResetRequestTimesByUserToken(token string) {
	userId, _ := l.GetUserIdByToken(token)
	redis.Rdb.Del(l.Ctx, fmt.Sprintf(ConstUserRequestTimesPre, userId))
}

// 获取用户每天能请求的次数，负数为无限
func (l *ChatLimitRequestLogic) GetUserPerDayRequestTimes(userId string) int32 {
	//
	//后期从数据库读取，并保存到 redis 中
	return ConstUserRequestTimes
}

func (l *ChatLimitRequestLogic) GetUserTodayRequestedTimes(userId string) int32 {
	//从 redis 读取
	num, _ := redis.Rdb.Get(l.Ctx, fmt.Sprintf(ConstUserRequestTimesPre, userId)).Result()
	//
	if num == "" {
		num = l.SetUserTodayRequestedTimes(userId)
	}

	int64num, _ := strconv.ParseInt(num, 10, 32)

	return int32(int64num)
}

// 请求次数+1
func (l *ChatLimitRequestLogic) AddUserTodayRequestedTimes(userId string) error {
	//userId, err := l.GetUserIdByToken(token)
	//if err != nil {
	//	return err
	//}
	redis.Rdb.Incr(l.Ctx, fmt.Sprintf(ConstUserRequestTimesPre, userId))
	////test
	//times := l.GetUserTodayRequestedTimes(userId)
	//fmt.Println("requested times", times)
	return nil
}

// 晚上 12 点过期
func (l *ChatLimitRequestLogic) SetUserTodayRequestedTimes(userId string) string {
	now := time.Now() // 获取当前时间
	// 得到今晚0点时间(即明天0点时间)
	endDay := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	// 计算当前时间到今晚0点的时间差
	expire := endDay.Sub(now)
	redis.Rdb.Set(l.Ctx, fmt.Sprintf(ConstUserRequestTimesPre, userId), 0, expire).Result()
	return "0"
}

func (l *ChatLimitRequestLogic) GetUserIdByToken(token string) (string, error) {
	userLogic := frontend.NewGetUserLogic(l.Ctx, l.svcCtx)
	user, err := userLogic.WechatUserService.GetByToken(token)
	if err != nil {
		return "", err
	}
	userId := fmt.Sprintf("%d", user.Id)
	return userId, nil
}

func (l *ChatLimitRequestLogic) CheckUserRequestedTimes(token string, user string) (userId string, err error) {
	if token != "" {
		userId, err = l.GetUserIdByToken(token)
		if err != nil {
			return userId, errors.New("token 错误")
		}
	} else if user != "" {
		userId = user
	} else {
		return userId, errors.New("参数错误")
	}
	//fmt.Println(userId, "userid")
	maxTimes := l.GetUserPerDayRequestTimes(userId)
	requestedTimes := l.GetUserTodayRequestedTimes(userId)
	//fmt.Println("maxTimes:", maxTimes, "requestedTimes", requestedTimes)
	if requestedTimes >= maxTimes {
		return userId, errors.New("25次数已用完，请明天再来哟")
	}
	return userId, nil
}
