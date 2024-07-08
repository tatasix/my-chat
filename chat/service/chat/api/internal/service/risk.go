package service

import (
	"chat/common/redis"
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"time"
)

type RiskService struct {
	logx.Logger
	ctx                      context.Context
	svcCtx                   *svc.ServiceContext
	wechatUserRepository     *repository.WechatUserRepository
	customerConfigRepository *repository.CustomerConfigRepository
	rightsRecordRepository   *repository.RightsRecordRepository
	rightsTimesRepository    *repository.RightsTimesRepository
	rightsRepository         *repository.RightsRepository
}

func NewRiskService(ctx context.Context, svcCtx *svc.ServiceContext) *RiskService {
	return &RiskService{
		Logger:                   logx.WithContext(ctx),
		ctx:                      ctx,
		svcCtx:                   svcCtx,
		wechatUserRepository:     repository.NewWechatUserRepository(ctx, svcCtx),
		customerConfigRepository: repository.NewCustomerConfigRepository(ctx, svcCtx),
		rightsRepository:         repository.NewRightsRepository(ctx, svcCtx),
		rightsTimesRepository:    repository.NewRightsTimesRepository(ctx, svcCtx),
		rightsRecordRepository:   repository.NewRightsRecordRepository(ctx, svcCtx),
	}
}
func (l *RiskService) Check1(user, openKfId string) (bool, int64, error) {
	if openKfId == "" {
		return false, 0, util.ReturnError(xerr.ParamError)
	}
	riskError := util.ReturnError(xerr.RightsNotHaveTimesError)
	//先判断有没有次卡次数
	timesStr, errTimes := redis.Rdb.Get(l.ctx, fmt.Sprintf(redis.ChatRiskTimesKey, user, openKfId)).Result()
	times, _ := strconv.Atoi(timesStr)
	if errTimes == nil && times > 0 {
		l.Logger.Infof("RiskService Check true use times user:%s ,openKfId:%s", user, openKfId)
		return true, int64(times), nil
	}
	//再判断月卡年卡

	timesStr, err := redis.Rdb.Get(l.ctx, fmt.Sprintf(redis.ChatRiskVipKey, user, openKfId)).Result()
	if err != nil {
		//第一次来，需要查表，然后存储起来
		userInfo, errUser := l.wechatUserRepository.GetByUser(user)
		if errUser != nil {
			l.Logger.Errorf("RiskService Check false GetByUser error:%+v ,user:%s", errUser, user)
			return false, 0, riskError
		}
		if userInfo != nil && userInfo.LevelExpire != "" && userInfo.Level != 0 {

			t1Time, _ := time.Parse(vars.TimeFormat, userInfo.LevelExpire)

			if t1Time.After(time.Now()) {
				//获取次数
				config, err1 := l.customerConfigRepository.GetPayConfig(openKfId, userInfo.Level)
				if err1 != nil {
					l.Logger.Errorf("RiskService Check false GetPayConfig error:%+v ,user:%s", err1.Error(), user)
					return false, 0, riskError
				}
				times = int(config.Times)
				//缓存时间现在是一天过期，后续可以改成从config中获取
				t, _ := time.Parse(vars.TimeFormat, time.Now().Add(24*time.Hour).Format("2006-01-02")+" 00:00:00")

				redis.Rdb.IncrBy(l.ctx, fmt.Sprintf(redis.ChatRiskVipKey, user, openKfId), int64(times))
				redis.Rdb.Expire(l.ctx, fmt.Sprintf(redis.ChatRiskVipKey, user, openKfId), t.Sub(time.Now()))
				return true, int64(times), nil

			}

		}
	}

	times, _ = strconv.Atoi(timesStr)
	if times > 0 {
		return true, int64(times), nil
	}
	return false, 0, riskError
}

func (l *RiskService) Reduce1(user, openKfId string) {
	//先判断有没有次卡次数
	timesStr, err := redis.Rdb.Get(l.ctx, fmt.Sprintf(redis.ChatRiskTimesKey, user, openKfId)).Result()
	times, _ := strconv.Atoi(timesStr)
	if err == nil && times > 0 {
		redis.Rdb.Decr(l.ctx, fmt.Sprintf(redis.ChatRiskTimesKey, user, openKfId))
		return
	}
	//再判断月卡年卡
	timesStr, err = redis.Rdb.Get(l.ctx, fmt.Sprintf(redis.ChatRiskVipKey, user, openKfId)).Result()
	if err == nil {
		times, _ = strconv.Atoi(timesStr)
		if times > 0 {
			redis.Rdb.Decr(l.ctx, fmt.Sprintf(redis.ChatRiskVipKey, user, openKfId))
			return
		}
	}
	l.Logger.Errorf("RiskService Reduce error user:%s  openKfId:%s", user, openKfId)
	return

}
func (l *RiskService) Check2(user, openKfId string) (enable, times int64, err error) {
	l.Logger.Infof("RiskService Check base data openKfId:%s ,user:%s", openKfId, user)
	if user == "121773743880028160" || user == "122588303228682240" || user == "nP1J2Xp4TzK5NrvGQs" {
		return 1, 25, nil
	}
	if openKfId == "" || user == "" {
		err = util.ReturnError(xerr.ParamError)
		return
	}
	//获取次功能的配置信息
	customerConfig, err1 := l.customerConfigRepository.GetByKfIdUseCache(openKfId)
	if err1 != nil {
		l.Logger.Errorf("RiskService Check false GetByKfIdUseCache error:%+v ,user:%s", err1.Error(), user)
		err = err1
		return
	}

	if customerConfig == nil || customerConfig.Id <= 0 {
		err = util.ReturnError(xerr.ConfigEmpty)
		return
	}
	var payConfigs []model.PayConfig
	_ = json.Unmarshal([]byte(customerConfig.Pay), &payConfigs)
	if payConfigs == nil || len(payConfigs) <= 0 {
		err = util.ReturnError(xerr.ConfigEmpty)
		return
	}
	//还需要查下，是不是次卡已经使用过了
	rightsTimes, err := l.rightsTimesRepository.GetAll(user, openKfId, "", time.Now().Format(vars.TimeFormat2), false, repository.PayTypeTimes)
	if nil != err {
		return
	}
	payConfigTimes := payConfigs[repository.PayTypeTimes]
	if int(payConfigTimes.Times) <= len(rightsTimes) {
		//当天已经玩过了，就也不可以再玩了
		enable = 2
		times = 0
		return
	}
	//先获取次卡数据
	rightsTimes, err = l.rightsTimesRepository.GetTimes(user, openKfId)
	if nil != err {
		return
	}
	if len(rightsTimes) > 0 {
		l.Logger.Infof("RiskService Check true 次卡 rightsTimes 大于 0")
		enable = 1
		times = 1
		return
	}

	now := time.Now().Format(vars.TimeFormat2)
	now1 := time.Now().Format(vars.TimeFormat)
	//获取会员卡数据或者免费的
	rightsTimes, err = l.rightsTimesRepository.GetByUserAndKfId(user, openKfId, now1)
	if nil != err {
		return
	}
	if len(rightsTimes) > 0 && rightsTimes[0].Id > 0 {
		if remain := rightsTimes[0].Remain; remain <= 0 {
			//当天次数已用完
			enable = 2
			times = 0
			l.Logger.Infof("RiskService Check false rightsTimes remain = 0")
		} else {
			enable = 1
			times = remain
			l.Logger.Infof("RiskService Check true rightsTimes remain  > 0")
		}
		return
	}
	//如果都没有数据，那很有可能是第一次来系统

	var period, payType, total, rightsId int64

	//看是否购买次卡月卡权益
	rights, err := l.rightsRepository.GetAll(user, "", []int64{repository.StatusSuccess}, []int64{repository.PayTypeYear, repository.PayTypeMonth}, true)
	if nil != err {
		return
	}
	if rights == nil || len(rights) <= 0 {
		//没有购买权益，需要再看此功能是否需要付费
		payConfig := payConfigs[repository.PayTypeFree]

		if !payConfig.Enable {
			if l.CheckCustomerIsFree(openKfId) {
				enable = 2
			} else {
				enable = 3
			}
			//此功能不可免费使用
			times = 0
			return
		}
		period = payConfig.Period
		payType = repository.PayTypeFree
		total = payConfig.Times
	} else {
		//月卡，或者年卡
		payConfig := payConfigs[rights[0].PayType]

		period = payConfig.Period
		payType = rights[0].PayType
		total = payConfig.Times
		rightsId = rights[0].Id
	}
	err = l.SaveRightsTimes(period, payType, total, rightsId, user, openKfId, now)
	if err != nil {
		return
	}
	enable = 1
	times = total
	return
}

// Check
// enable 1 就是成功，2 次数用完了，3 要去付钱了
func (l *RiskService) Check(user, openKfId string) (enable, times int64, err error) {
	l.Logger.Infof("RiskService Check base data openKfId:%s ,user:%s", openKfId, user)
	if user == "121773743880028160" || user == "122588303228682240" || user == "nP1J2Xp4TzK5NrvGQs" {
		return 1, 25, nil
	}
	if openKfId == "" || user == "" {
		err = util.ReturnError(xerr.ParamError)
		return
	}
	//获取次功能的配置信息
	customerConfig, err1 := l.customerConfigRepository.GetByKfIdUseCache(openKfId)
	if err1 != nil {
		l.Logger.Errorf("RiskService Check false GetByKfIdUseCache error:%+v ,user:%s", err1.Error(), user)
		err = err1
		return
	}

	if customerConfig == nil || customerConfig.Id <= 0 {
		err = util.ReturnError(xerr.ConfigEmpty)
		return
	}
	var payConfigs []model.PayConfig
	_ = json.Unmarshal([]byte(customerConfig.Pay), &payConfigs)
	if payConfigs == nil || len(payConfigs) <= 0 {
		err = util.ReturnError(xerr.ConfigEmpty)
		return
	}
	payConfigTimes := payConfigs[repository.PayTypeTimes]
	used := 0
	var rightsTimes []*model.RightsTimes
	if payConfigTimes.Enable {
		//次卡可用才处理次卡逻辑
		//还需要查下，是不是次卡已经使用过了
		rightsTimes, err = l.rightsTimesRepository.GetAll(user, openKfId, "", time.Now().Format(vars.TimeFormat2), false, repository.PayTypeTimes)
		if nil != err {
			return
		}
		if payConfigTimes.Enable && int(payConfigTimes.Times) <= len(rightsTimes) {
			//当天已经玩过了，就也不可以再玩了
			enable = 2
			times = 0
			return enable, times, nil
		}
		used = len(rightsTimes)
		//先获取次卡数据
		rightsTimes, err = l.rightsTimesRepository.GetTimes(user, openKfId)
		if nil != err {
			return 0, 0, err
		}
		if len(rightsTimes) > 0 {
			l.Logger.Infof("RiskService Check true 次卡 rightsTimes 大于 0")
			enable = 1
			times = 1
			return enable, times, nil
		}
	}

	now := time.Now().Format(vars.TimeFormat2)
	now1 := time.Now().Format(vars.TimeFormat)
	//获取会员卡数据或者免费的
	rightsTimes, err = l.rightsTimesRepository.GetByUserAndKfId(user, openKfId, now1)
	if nil != err {
		return
	}
	if len(rightsTimes) > 0 && rightsTimes[0].Id > 0 {
		if remain := rightsTimes[0].Remain; remain <= 0 {
			//当天次数已用完
			enable = 2
			times = 0
			l.Logger.Infof("RiskService Check false rightsTimes remain = 0")
		} else {
			enable = 1
			times = remain
			l.Logger.Infof("RiskService Check true rightsTimes remain  > 0")
		}
		return
	}
	//如果都没有数据，那很有可能是第一次来系统

	var period, payType, total, rightsId int64

	//看是否购买次卡月卡权益
	rights, err := l.rightsRepository.GetAll(user, "", []int64{repository.StatusSuccess}, []int64{repository.PayTypeYear, repository.PayTypeMonth}, true)
	if nil != err {
		return
	}
	if rights == nil || len(rights) <= 0 {
		//没有购买权益，需要再看此功能是否需要付费
		payConfig := payConfigs[repository.PayTypeFree]

		if !payConfig.Enable {
			if l.CheckCustomerIsFree(openKfId) {
				enable = 2
			} else {
				enable = 3
			}
			//此功能不可免费使用
			times = 0
			return
		}
		period = payConfig.Period
		payType = repository.PayTypeFree
		total = payConfig.Times
	} else {
		//月卡，或者年卡
		payConfig := payConfigs[rights[0].PayType]

		period = payConfig.Period
		payType = rights[0].PayType
		total = payConfig.Times - int64(used)
		rightsId = rights[0].Id
	}
	err = l.SaveRightsTimes(period, payType, total, rightsId, user, openKfId, now)
	if err != nil {
		return
	}
	enable = 1
	times = total
	return
}

func (l *RiskService) SaveRightsTimes(period, payType, total, rightsId int64, user, openKfId, now string) (err error) {
	if period == 0 {
		period = repository.RightsTimesPeriodDay
	}
	var start, end sql.NullTime
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("无法加载时区：", err)
		return
	}

	if period == repository.RightsTimesPeriodDay {
		start1, _ := time.ParseInLocation(vars.TimeFormat, now+" 00:00:00", loc)
		end1, _ := time.ParseInLocation(vars.TimeFormat, now+" 23:59:59", loc)
		start.Valid = true
		start.Time = start1
		end.Valid = true
		end.Time = end1
	}

	//把结果写进rights_times
	rightsTimesModel := &model.RightsTimes{
		RightsId: rightsId,
		User:     user,
		OpenKfId: openKfId,
		PayType:  payType,
		Period:   period,
		Start:    start,
		End:      end,
		Total:    total,
		Remain:   total,
	}

	_, err = l.rightsTimesRepository.Insert(rightsTimesModel)
	return
}

func (l *RiskService) Reduce(user, openKfId string, relationId int64) {

	payType := 100
	rightsTimes, err := l.rightsTimesRepository.GetTimes(user, openKfId)
	if nil != err {
		return
	}
	if len(rightsTimes) > 0 && rightsTimes[0].Id > 0 {
		//先处理次卡数据
		payType = repository.PayTypeTimes
	} else {
		//处理会员卡数据
		rightsTimes, err = l.rightsTimesRepository.GetAll(user, openKfId, time.Now().Format(vars.TimeFormat), "", true, 0)
		if nil != err {
			return
		}
		if len(rightsTimes) > 0 && rightsTimes[0].Id > 0 {
			payType = int(rightsTimes[0].PayType)
		}
	}

	if payType == 100 {
		l.Logger.Errorf("RiskService Reduce not have access")
		return
	}

	rightsTimesInfo := rightsTimes[0]
	//添加记录
	rightsRecordModel := &model.RightsRecord{
		RightsId:      rightsTimesInfo.RightsId,
		RightsTimesId: rightsTimesInfo.Id,
		RelationId:    relationId,
		User:          user,
		OpenKfId:      openKfId,
	}

	//var conn sqlx.SqlConn
	//err = conn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
	//
	//	_, err = l.rightsRecordRepository.Insert(rightsRecordModel)
	//	if nil != err {
	//		return err
	//	}
	//	//更新rights_times表的次数
	//	err = l.rightsTimesRepository.UpdateTimesById(rightsTimesInfo.Id, 1)
	//	if nil != err {
	//		return err
	//	}
	//	if payType == repository.PayTypeTimes {
	//		//更新rights表的权益使用状态，次卡才需要更新
	//		err = l.rightsRepository.UpdateStatusById(rightsTimesInfo.RightsId, repository.StatusUsed, "")
	//		if nil != err {
	//			return err
	//		}
	//	}
	//	return nil
	//})
	//if nil != err {
	//	return
	//}

	_, err = l.rightsRecordRepository.Insert(rightsRecordModel)
	if nil != err {
		return
	}
	//更新rights_times表的次数
	err = l.rightsTimesRepository.UpdateTimesById(rightsTimesInfo.Id, 1)
	if nil != err {
		return
	}
	if payType == repository.PayTypeTimes {
		//更新rights表的权益使用状态，次卡才需要更新
		err = l.rightsRepository.UpdateStatusById(rightsTimesInfo.RightsId, repository.StatusUsed, "")
		if nil != err {
			return
		}
	}
	return
}

func (l *RiskService) CheckCustomerIsFree(kfId string) bool {
	haystack := []string{
		vars.ChatFortune,
		vars.ChatDream,
		//vars.ChatSandbox,
	}
	for _, item := range haystack {
		if item == kfId {
			return true
		}
	}
	return false
}
