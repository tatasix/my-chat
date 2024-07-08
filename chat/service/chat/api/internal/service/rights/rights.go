package rights

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/service/pay"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"time"
)

type RightsService struct {
	logx.Logger
	ctx                      context.Context
	svcCtx                   *svc.ServiceContext
	rightsRepository         *repository.RightsRepository
	customerConfigRepository *repository.CustomerConfigRepository
	wechatUserRepository     *repository.WechatUserRepository
	configRepository         *repository.ConfigRepository
	rightsTimesRepository    *repository.RightsTimesRepository
}

func NewRightsService(ctx context.Context, svcCtx *svc.ServiceContext) *RightsService {
	return &RightsService{
		Logger:                   logx.WithContext(ctx),
		ctx:                      ctx,
		svcCtx:                   svcCtx,
		rightsRepository:         repository.NewRightsRepository(ctx, svcCtx),
		customerConfigRepository: repository.NewCustomerConfigRepository(ctx, svcCtx),
		wechatUserRepository:     repository.NewWechatUserRepository(ctx, svcCtx),
		configRepository:         repository.NewConfigRepository(ctx, svcCtx),
		rightsTimesRepository:    repository.NewRightsTimesRepository(ctx, svcCtx),
	}
}

func (l *RightsService) Pay(user, openKfId, messageId, returnUrl, quitUrl string, payMethod, payType, source int64) (lastId int64, url string, err error) {
	//return l.Test()
	if payType != repository.PayTypeTimes && payType != repository.PayTypeMonth && payType != repository.PayTypeYear {
		err = util.ReturnError(xerr.ParamError)
		return
	}
	if payType == repository.PayTypeMonth || payType == repository.PayTypeYear {
		openKfId = ""
	} else {
		if openKfId == "" {
			return 0, "", util.ReturnError(xerr.ParamError)
		}
	}
	start := time.Now()

	add := l.GetAddTime(payType)

	//幂等
	exist, err := l.rightsRepository.GetByMessage(messageId)
	if err != nil {
		return
	}
	if exist != nil && exist.Id > 0 {
		err = util.ReturnError(xerr.SystemBusyError)
		return
	}

	//用户
	userInfo, err := l.wechatUserRepository.GetByUser(user)
	if err != nil {
		return
	}
	if userInfo == nil || userInfo.Id <= 0 {
		err = util.ReturnError(xerr.LoginAccountNotExist)
		return
	}
	t := time.Now()
	if userInfo.LevelExpire != "" {
		t, err = time.Parse(vars.TimeFormat, userInfo.LevelExpire)
		if err != nil {
			return
		}
	}

	if time.Now().Before(t) {
		start = t
	}

	payConfig, err := l.customerConfigRepository.GetPayConfig(openKfId, payType)
	if err != nil {
		return
	}
	if payConfig.Amount <= 0 {
		err = util.ReturnError(xerr.RightsAmountError)
		return
	}
	if !l.svcCtx.Config.PayPrice {
		payConfig.Amount = 1
	}
	var times, period int64
	if payType == repository.PayTypeTimes {
		//月卡，年卡的配置需要根据客服来取，所以这里不存
		times = payConfig.Times
		period = payConfig.Period
	}

	snapshot, _ := json.Marshal(payConfig)
	rights := &model.Rights{
		User:       user,
		OpenKfId:   openKfId,
		PayType:    payType,
		OutTradeNo: util.GenerateSnowflakeString(),
		PayMethod:  payMethod,
		MessageId:  messageId,
		Source:     source,
		Price:      payConfig.Amount,
		Times:      times,
		Period:     period,
		Status:     repository.StatusNotPay,
		Start:      sql.NullTime{Valid: true, Time: start},
		End:        sql.NullTime{Valid: true, Time: start.Add(add * 24 * time.Hour)},
		Snapshot:   string(snapshot),
	}
	lastId, err = l.rightsRepository.Insert(rights)
	if err != nil {
		return
	}
	//支付
	charge := &pay.Charge{
		TradeNum:  rights.OutTradeNo,
		UserID:    user,
		PayMethod: payMethod,
		MoneyFee:  payConfig.Amount,
		Source:    source,
		OpenID:    userInfo.Openid,
		Describe:  payConfig.Describe,
		ReturnURL: returnUrl,
		QuitUrl:   quitUrl,
		PayType:   payType,
		RightsId:  lastId,
		MessageId: messageId,
	}

	url, err = pay.GetPayClient(l.ctx, l.svcCtx, charge.PayMethod).GetQrCode(charge)
	return

}

func (l *RightsService) Notify(payMethod int64, r *http.Request) (callbackRes string, err error) {
	flag := true
	errMessage := ""
	//解析内容
	result, err := pay.GetPayClient(l.ctx, l.svcCtx, payMethod).GetNotifyResult(r)

	if result == nil {
		l.Logger.Errorf("pay Notify  error result nil")
		return "", nil
	}
	if result.OutTradeNo == "" {
		l.Logger.Errorf("pay Notify error OutTradeNo empty")
		return
	}
	if err != nil {
		l.Logger.Errorf("pay Notify request:%+v error: %+v", r, err)
		flag = false
		errMessage = err.Error()
	} else if result.Status != 1 {
		l.Logger.Errorf("pay Notify error status 0 result %s", result.Result)
		flag = false
		errMessage = result.Result

	}

	outTradeNo := result.OutTradeNo
	//幂等
	rightsInfo, err := l.rightsRepository.GetByOutTradeNo(outTradeNo)
	if err != nil {
		l.Logger.Errorf("Notify error err:%+v", err)
		return
	}
	if rightsInfo == nil || rightsInfo.Id <= 0 {
		l.Logger.Errorf("Notify error record not found, outTradeNo:%s", outTradeNo)
		return
	}

	if flag {
		//成功
		if rightsInfo.Status != repository.StatusNotPay {
			l.Logger.Errorf("Notify error Status not right, rightsInfo:%+v", rightsInfo)
			return
		}
		err = l.NotifySuccess(rightsInfo.Id, rightsInfo)
	} else {
		//失败
		err = l.NotifyFail(rightsInfo.Id, errMessage)
	}

	return result.CallBackInfo, nil
}

func (l *RightsService) NotifyFail(id int64, reason string) (err error) {

	//更新状态
	return l.rightsRepository.UpdateStatusById(id, repository.StatusFail, reason)
}

func (l *RightsService) NotifySuccess(id int64, rights *model.Rights) (err error) {

	if rights.PayType != repository.PayTypeTimes {
		//更新用户
		userInfo, errUser := l.wechatUserRepository.GetByUser(rights.User)
		if err != nil {
			return errUser
		}
		if userInfo == nil || userInfo.Id <= 0 {
			return util.ReturnError(xerr.LoginAccountNotExist)
		}
		t := time.Now()
		if userInfo.LevelExpire != "" {
			t, err = time.Parse(vars.TimeFormat, userInfo.LevelExpire)
			if err != nil {
				return err
			}
		}

		add := l.GetAddTime(rights.PayType)
		start := time.Now()
		if time.Now().Before(t) {
			start = t
		}
		expire := start.Add(add * 24 * time.Hour).Format(vars.TimeFormat)
		err = l.wechatUserRepository.UpdateLevelById(userInfo.Id, rights.PayType, expire)

	} else {
		//次卡买了马上就得发放
		//把结果写进rights_times
		period := rights.Period
		if period == 0 {
			period = repository.RightsTimesPeriodDay
		}
		var start, end sql.NullTime
		if period == repository.RightsTimesPeriodDay {
			now := time.Now().Format(vars.TimeFormat2)
			start1, _ := time.Parse(vars.TimeFormat, now+" 00:00:00")
			end1, _ := time.Parse(vars.TimeFormat, now+" 23:59:59")
			start.Valid = true
			start.Time = start1
			end.Valid = true
			end.Time = end1
		}

		rightsTimesModel := &model.RightsTimes{
			RightsId: rights.Id,
			User:     rights.User,
			OpenKfId: rights.OpenKfId,
			Period:   rights.Period,
			PayType:  rights.PayType,
			Start:    start,
			End:      end,
			Total:    1,
			Remain:   1,
		}

		_, err = l.rightsTimesRepository.Insert(rightsTimesModel)
		if nil != err {
			return
		}
	}
	//更新状态
	err = l.rightsRepository.UpdateStatusById(id, repository.StatusSuccess, "success")
	if err != nil {
		return
	}
	//如果是非次卡 并且是疗愈喵，那次数也要加上去
	if rights.PayType != repository.PayTypeTimes {
		err = l.rightsTimesRepository.AddTimesByUser(rights.User, vars.ChatCat, time.Now().Format(vars.TimeFormat), 20)
	}
	return

}

func (l *RightsService) GetAddTime(payType int64) (add time.Duration) {
	switch payType {
	case repository.PayTypeTimes:
		add = 30
	case repository.PayTypeMonth:
		add = 30
	case repository.PayTypeYear:
		add = 365
	}
	return
}

//func (l *RightsService) Test() (callbackRes string, err error) {
//	outTradeNo := "1716729564903772160"
//	//幂等
//	rightsInfo, err := l.rightsRepository.GetByOutTradeNo(outTradeNo)
//	if err != nil {
//		l.Logger.Errorf("Notify error err:%+v", err)
//		return
//	}
//	if rightsInfo == nil || rightsInfo.Id <= 0 {
//		l.Logger.Errorf("Notify error record not found, outTradeNo:%s", outTradeNo)
//		return
//	}
//
//	//if flag {
//	//成功
//	if rightsInfo.Status != repository.StatusNotPay {
//		l.Logger.Errorf("Notify error Status not right, rightsInfo:%+v", rightsInfo)
//		return
//	}
//	err = l.NotifySuccess(rightsInfo.Id, rightsInfo)
//	//} else {
//	//	//失败
//	//	err = l.NotifyFail(rightsInfo.Id, "")
//	//}
//
//	return
//}

func (l *RightsService) GetRights(id int64) (Rights *model.Rights, err error) {
	return l.rightsRepository.GetById(id)
}

func (l *RightsService) ExpireRights() (err error) {
	//过期会员
	err = l.wechatUserRepository.UpdateLevel()
	if err != nil {
		return
	}

	//过期权益
	err = l.rightsRepository.UpdateStatus()
	return
}

func (l *RightsService) DeleteUserRights(user string) (err error) {
	//删会员
	err = l.wechatUserRepository.DeleteByUser(user)
	if err != nil {
		return
	}

	////过期权益
	//err = l.rightsRepository.DeleteByUser(user)
	//
	////过期权益
	//err = l.rightsTimesRepository.DeleteByUser(user)

	return
}
