package pay

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
	"context"
	"database/sql"
	_ "github.com/go-pay/gopay"
	"github.com/go-pay/gopay/apple"
	_ "github.com/go-pay/gopay/apple"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	sandBoxExpireDateAdd = 2 * time.Minute       //沙盒会员过期时间加 3 分钟
	prodExpireDateAdd    = 24 * 60 * time.Minute //正式会员过期时间加 1 天
)

type ApplePayService struct {
	ctx context.Context
	logx.Logger
	svcCtx *svc.ServiceContext

	rwMu      sync.RWMutex
	appleTest *apple.Client
	appleProd *apple.Client

	isProd                         bool
	rightsRepository               *repository.RightsRepository
	wechatUserRepository           *repository.WechatUserRepository
	rightsTimesRepository          *repository.RightsTimesRepository
	applePaySubscriptionRepository *repository.ApplePaySubscriptionRepository
}

func NewApplePayService(ctx context.Context, svcCtx *svc.ServiceContext) *ApplePayService {
	iss := svcCtx.Config.ApplePay.Iss
	bid := svcCtx.Config.ApplePay.Bid
	kid := svcCtx.Config.ApplePay.Kid
	privateKey := svcCtx.Config.ApplePay.PrivateKey
	//isProd := svcCtx.Config.ApplePay.IsProd
	appleCliTest, err := apple.NewClient(iss, bid, kid, privateKey, false)
	if err != nil {
		xlog.Color(xlog.RedBright).Info(err.Error())
		return nil
	}
	appleCliProd, err := apple.NewClient(iss, bid, kid, privateKey, true)
	if err != nil {
		xlog.Color(xlog.RedBright).Info(err.Error())
		return nil
	}

	return &ApplePayService{
		Logger:                         logx.WithContext(ctx),
		ctx:                            ctx,
		svcCtx:                         svcCtx,
		appleTest:                      appleCliTest,
		appleProd:                      appleCliProd,
		isProd:                         svcCtx.Config.ApplePay.IsProd,
		rightsRepository:               repository.NewRightsRepository(ctx, svcCtx),
		wechatUserRepository:           repository.NewWechatUserRepository(ctx, svcCtx),
		rightsTimesRepository:          repository.NewRightsTimesRepository(ctx, svcCtx),
		applePaySubscriptionRepository: repository.NewApplePaySubscriptionRepository(ctx, svcCtx),
	}
}

func (a *ApplePayService) ReceiptVerification(req types.ApplePayConfirmReq) (err error) {
	if req.RightsId <= 0 {
		err = util.ReturnError(xerr.ParamError)
		return
	}

	rightsPo, err := a.rightsRepository.GetById(req.RightsId)
	if err != nil || rightsPo == nil {
		err = util.ReturnError(xerr.RecordNotFound)
		return
	}

	originalTransactionId := req.TxnId
	//次卡
	if rightsPo.OpenKfId != "" && rightsPo.PayType == repository.PayTypeTimes {
		singleRsp, err := a.appleProd.GetTransactionInfo(a.ctx, originalTransactionId)
		a.isProd = true
		if err != nil {
			singleRsp, err = a.appleTest.GetTransactionInfo(a.ctx, originalTransactionId)
			a.isProd = false
			if err != nil {
				a.Logger.Errorf("client.Get single TransactionInfo(),err:%+v", err)
				return err
			}
		}
		singleTi, err := singleRsp.DecodeSignedTransaction()
		xlog.Color(xlog.GreenBright).Debugf("TransactionInfo:%+v", singleTi)
		if err != nil {
			a.Logger.Errorf("client.Get single TransactionInfo decode,err:%+v", err)
			return err
		}
		if singleTi.InAppOwnershipType == "PURCHASED" {
			err = a.NotifySuccess(rightsPo.Id, rightsPo, "")
			if err != nil {
				a.Logger.Errorf("client. single pay NotifySuccess,err:%+v", err)
				return err
			}
		}

		newRightsPo := rightsPo
		newRightsPo.AppleTxnId = req.TxnId
		newRightsPo.Status = repository.StatusSuccess
		newRightsPo.Reason = singleTi.InAppOwnershipType
		err = a.rightsRepository.Update(rightsPo, newRightsPo)
		return err
	}

	//subscription
	rsp, err := a.GetLastTransactions(originalTransactionId)
	if err != nil {
		a.Logger.Errorf("client.GetAllSubscriptionStatuses(),err:%+v", err)
		return
	}

	for _, dataItem := range rsp.Data {
		for _, lastTransaction := range dataItem.LastTransactions {
			//status
			//1
			//The auto-renewable subscription is active.
			//2
			//The auto-renewable subscription is expired.
			//3
			//The auto-renewable subscription is in a billing retry period.
			//4
			//The auto-renewable subscription is in a Billing Grace Period.
			//5
			//The auto-renewable subscription is revoked. The App Store refunded the transaction or revoked it from Family Sharing.

			if lastTransaction.Status != 1 {
				//xlog.Color(xlog.RedBright).Infof("apple subscription err : rights_id:%+v fail,apple pay subscription status:%+v", req.RightsId, lastTransaction.Status)
				a.Logger.Errorf("apple subscription err : rights_id:%+v fail,apple pay subscription status:%+v", req.RightsId, lastTransaction.Status)
				continue
			}

			ti, _ := lastTransaction.DecodeTransactionInfo()
			//xlog.Color(xlog.GreenBright).Debugf("TransactionInfo:%+v", ti)

			ri, _ := lastTransaction.DecodeRenewalInfo()
			//xlog.Color(xlog.GreenBright).Debugf("RenewalInfo:%+v", ri)

			subscriptionPo, err := a.applePaySubscriptionRepository.GetByOriTxnId(ti.OriginalTransactionId)
			//xlog.Color(xlog.Red).Debug(subPo)

			if err != nil {
				return err
			}

			subPo := &model.ApplePaySubscription{
				UserId:                rightsPo.User,
				OriginalTransactionId: ti.OriginalTransactionId,
				TransactionId:         ti.TransactionId,
				OrderId:               strconv.FormatInt(rightsPo.Id, 10),
				Status:                ri.AutoRenewStatus,
				ProductId:             ri.ProductId,
				OriginalProductId:     ri.AutoRenewProductId,
				ExpiresDate:           util.TimeToSql(time.Unix(0, ti.ExpiresDate)),
				PurchaseDate:          util.TimeToSql(time.Unix(0, ti.PurchaseDate*int64(time.Millisecond))),
				TransactionReason:     ti.TransactionReason,
				TransactionType:       ti.Type,
				InAppOwnershipType:    ti.InAppOwnershipType,
				Price:                 ti.Price,
				Currency:              ti.Currency,
			}
			if subscriptionPo == nil || subscriptionPo.Id == 0 {
				a.applePaySubscriptionRepository.Insert(subPo)
			} else {
				a.applePaySubscriptionRepository.Update(subscriptionPo.Id, subPo)
			}

			if ri.AutoRenewStatus != 0 {
				if a.isProd {
					//正式
					ti.ExpiresDate = ti.ExpiresDate*int64(time.Millisecond) + int64(prodExpireDateAdd)
				} else {
					//沙盒
					ti.ExpiresDate = ti.ExpiresDate*int64(time.Millisecond) + int64(sandBoxExpireDateAdd)
				}
				a.NotifySuccess(rightsPo.Id, rightsPo, util.TimeFormat(time.Unix(0, ti.ExpiresDate)))
			}

		}
	}
	return
}
func (a *ApplePayService) CheckApplePaySubscription() (err error) {
	pos, err := a.applePaySubscriptionRepository.GetByStatusAndExpiresDate(1, util.TimeFormat(time.Now()))
	if err != nil {
		a.Logger.Errorf("applePay.GetByStatus err :%+v", err)
		return
	}

	if pos != nil && len(pos) > 0 {
		for _, v := range pos {
			xlog.Color(xlog.Red).Debug(v)
			rsp, err := a.GetLastTransactions(v.OriginalTransactionId)
			if err != nil {
				a.Logger.Errorf("applePay.GetLastTransactions err :%+v", err)
				return err
			}
			for _, dataItem := range rsp.Data {
				for _, lastTransaction := range dataItem.LastTransactions {
					ti, _ := lastTransaction.DecodeTransactionInfo()
					xlog.Color(xlog.GreenBright).Debugf("TransactionInfo:%+v", ti)

					ri, _ := lastTransaction.DecodeRenewalInfo()
					xlog.Color(xlog.GreenBright).Debugf("RenewalInfo:%+v", ri)

					subPo := &model.ApplePaySubscription{
						UserId:                v.UserId,
						OriginalTransactionId: ti.OriginalTransactionId,
						TransactionId:         ti.TransactionId,
						OrderId:               v.OrderId,
						Status:                ri.AutoRenewStatus,
						ProductId:             ri.ProductId,
						OriginalProductId:     ri.AutoRenewProductId,
						ExpiresDate:           util.TimeToSql(time.Unix(0, ti.ExpiresDate*int64(time.Millisecond))),
						PurchaseDate:          util.TimeToSql(time.Unix(0, ti.PurchaseDate*int64(time.Millisecond))),
						TransactionReason:     ti.TransactionReason,
						TransactionType:       ti.Type,
						InAppOwnershipType:    ti.InAppOwnershipType,
						Price:                 ti.Price,
						Currency:              ti.Currency,
					}

					a.applePaySubscriptionRepository.Update(v.Id, subPo)

					if ri.AutoRenewStatus == 1 && v.TransactionId != subPo.TransactionId {
						//创建订单
						orderId, _ := strconv.ParseInt(v.OrderId, 10, 64)
						oldRightsPo, err := a.rightsRepository.GetById(orderId)

						if err != nil || oldRightsPo == nil || oldRightsPo.Id <= 0 {
							a.Logger.Errorf("applePay.CheckApplePaySubscription getRightsById err :%+v", err)
							return err
						}

						newRights := &model.Rights{
							MessageId:  util.GenerateSnowflakeString(),
							User:       oldRightsPo.User,
							OpenKfId:   oldRightsPo.OpenKfId,
							PayType:    oldRightsPo.PayType,
							OutTradeNo: util.GenerateSnowflakeString(),
							PayMethod:  oldRightsPo.PayMethod,
							Price:      oldRightsPo.Price,
							Period:     oldRightsPo.Period,
							Times:      oldRightsPo.Times,
							Source:     oldRightsPo.Source,
							Status:     repository.StatusDone,
							Reason:     ti.TransactionReason,
							Start:      util.TimeToSql(time.Unix(0, ti.PurchaseDate*int64(time.Millisecond))),
							End:        util.TimeToSql(time.Unix(0, ti.ExpiresDate*int64(time.Millisecond))),
							Snapshot:   oldRightsPo.Snapshot,
							AppleTxnId: ti.TransactionId,
						}

						newId, err := a.rightsRepository.Insert(newRights)
						if err != nil || newId <= 0 {
							a.Logger.Errorf("applePay.CheckApplePaySubscription Rights Insert err :%+v", err)
							return err
						}

						//	更新会员时间
						if a.isProd {
							//正式
							ti.ExpiresDate = ti.ExpiresDate*int64(time.Millisecond) + int64(prodExpireDateAdd)
						} else {
							//沙盒
							ti.ExpiresDate = ti.ExpiresDate*int64(time.Millisecond) + int64(sandBoxExpireDateAdd)
						}
						err = a.NotifySuccess(newId, newRights, util.TimeFormat(time.Unix(0, ti.ExpiresDate)))

						if err != nil {
							a.Logger.Errorf("applePay.CheckApplePaySubscription NotifySuccess err :%+v", err)
							return err
						}
					}
				}
			}

		}

	}

	return err
}

func (a *ApplePayService) GetLastTransactions(originalTransactionId string) (rsp *apple.AllSubscriptionStatusesRsp, err error) {

	rsp, err = a.appleProd.GetAllSubscriptionStatuses(a.ctx, originalTransactionId)
	a.isProd = true
	if err != nil {
		rsp, err = a.appleTest.GetAllSubscriptionStatuses(a.ctx, originalTransactionId)
		a.isProd = false
		if err != nil {
			a.Logger.Errorf("client.GetAllSubscriptionStatuses(),err:%+v", err)
			return
		}
	}

	if len(rsp.Data) <= 0 {
		a.Logger.Errorf(" GetLastTransactions error rsp:%+v", rsp)
		err = util.ReturnError(xerr.RecordNotFound)
		return
	}

	return
}

func (a *ApplePayService) Notify(r *http.Request) (err error) {
	//换成主动定时获取
	return
}

func (a *ApplePayService) NotifySuccess(id int64, rights *model.Rights, expireDate string) (err error) {

	if rights.PayType != repository.PayTypeTimes {
		//更新用户
		userInfo, errUser := a.wechatUserRepository.GetByUser(rights.User)
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
		var expire string
		if expireDate == "" {
			add := a.GetAddTime(rights.PayType)
			start := time.Now()
			if time.Now().Before(t) {
				start = t
			}
			expire = start.Add(add * 24 * time.Hour).Format(vars.TimeFormat)
		} else {
			expire = expireDate
		}

		err = a.wechatUserRepository.UpdateLevelById(userInfo.Id, rights.PayType, expire)

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

		_, err = a.rightsTimesRepository.Insert(rightsTimesModel)
		if nil != err {
			return
		}
	}
	//更新状态
	err = a.rightsRepository.UpdateStatusById(id, repository.StatusSuccess, "success")
	if err != nil {
		return
	}
	//如果是非次卡 并且是疗愈喵，那次数也要加上去
	if rights.PayType != repository.PayTypeTimes {
		err = a.rightsTimesRepository.AddTimesByUser(rights.User, vars.ChatCat, time.Now().Format(vars.TimeFormat), 20)
	}
	return

}

func (a *ApplePayService) GetAddTime(payType int64) (add time.Duration) {
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

func (a *ApplePayService) GetQrCode(charge *Charge) (qrCode string, err error) {
	return
}

func (a *ApplePayService) wapPay(charge *Charge) (qrCode string, err error) {
	return
}

func (a *ApplePayService) pcPay(charge *Charge) (qrCode string, err error) {
	return
}
func (a *ApplePayService) GetNotifyResult(r *http.Request) (res *PayCallback, err error) {
	return
}

//func (a *ApplePayService) GetNotificationHistory(transactionId string) {
//
//	//transactionId := "2000000489292555"
//
//	currentTime := time.Now()
//	// 获取 24 小时前的时间
//	previousTime := currentTime.Add(-24 * 8 * time.Hour)
//	// 转换为毫秒级时间戳
//	currentTimestamp := currentTime.UnixNano() / int64(time.Millisecond)
//	previousTimestamp := previousTime.UnixNano() / int64(time.Millisecond)
//
//	bm := make(gopay.BodyMap)
//	bm.Set("startDate", previousTimestamp).
//		Set("endDate", currentTimestamp).
//		Set("transactionId", transactionId)
//
//	// 发起请求
//	rsp, err := a.apple.GetNotificationHistory(a.ctx, "", bm)
//	if err != nil {
//		xlog.Color(xlog.Red).Errorf("client.GetNotificationHistory(),err:%+v", err)
//		return
//	}
//	for _, v := range rsp.NotificationHistory {
//		payload, err := apple.DecodeSignedPayload(v.SignedPayload)
//		if err != nil {
//			xlog.Errorf("DecodeSignedPayload(),err:+v", err)
//			continue
//		}
//		xlog.Infof("payload: %+v", payload)
//		//renew info
//		ri, _ := payload.DecodeRenewalInfo()
//		xlog.Color(xlog.GreenBright).Debugf("RenewalInfo:%+v", ri)
//
//	}
//}
