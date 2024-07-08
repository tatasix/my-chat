package pay

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strconv"
	"sync"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/shopspring/decimal"
)

type AlipayService struct {
	ctx context.Context
	logx.Logger
	svcCtx *svc.ServiceContext

	rwMu sync.RWMutex

	alipay *alipay.Client
}

func NewAlipayService(ctx context.Context, svcCtx *svc.ServiceContext) *AlipayService {

	alipayCli, err := alipay.NewClient(svcCtx.Config.Alipay.Appid, svcCtx.Config.Alipay.PrivateKey, true)
	if err != nil {
		panic(err)
	}
	// Debug开关，输出/关闭日志
	alipayCli.DebugSwitch = gopay.DebugOff
	// 配置公共参数
	alipayCli.SetCharset(alipay.UTF8).
		SetSignType(alipay.RSA2)

	return &AlipayService{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		alipay: alipayCli,
	}
}

// GetQrCode 支付码获取(用户扫码支付)
func (a *AlipayService) GetQrCode(charge *Charge) (qrCode string, err error) {
	if charge.Source == repository.SourcePC {
		return a.pcPay(charge)
	} else if charge.Source == repository.SourceH5 {
		return a.wapPay(charge)
	} else if charge.Source == repository.SourceAndroid {
		return a.androidPay(charge)
	}
	return "", util.ReturnError(xerr.SystemBusyError)
}

// PcPay 支付码获取(用户扫码支付)
func (a *AlipayService) pcPay(charge *Charge) (qrCode string, err error) {
	a.alipay = a.alipay.SetReturnUrl(a.svcCtx.Config.Alipay.ReturnUrl + charge.ReturnURL).SetNotifyUrl(a.svcCtx.Config.Alipay.NotifyUrl)
	// 生成单号
	amount := decimal.NewFromInt(charge.MoneyFee).DivRound(decimal.NewFromInt(100), 2).String()
	// 构造参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", charge.Describe).
		Set("out_trade_no", charge.TradeNum).
		Set("total_amount", amount).
		Set("qr_pay_mode", 4).
		Set("qrcode_width", 120)
	// 发起支付
	qrCode, err = a.alipay.TradePagePay(a.ctx, bm)
	if err != nil {
		if bizError, ok := alipay.IsBizError(err); ok {
			a.Logger.Errorf("s.alipay.TradePagePay(%v), bizError:%v", bm, bizError)
			return
		}
		a.Logger.Errorf("s.alipay.TradePagePay(%v), err:%v", bm, err)
		return
	}
	return
}

// GetNotifyResult 支付码获取(用户扫码支付)
func (a *AlipayService) GetNotifyResult(r *http.Request) (res *PayCallback, err error) {
	// 解析异步通知的参数
	// req：*http.Request
	notifyReq, err := alipay.ParseNotifyToBodyMap(r)
	if err != nil {
		a.Logger.Errorf("alipay.ParseNotifyToBodyMap err :%+v", err)
		return
	}
	a.Logger.Infof("alipay.ParseNotifyToBodyMap notifyReq:%+v ", notifyReq)

	// 支付宝异步通知验签（公钥模式）
	_, err = alipay.VerifySign(a.svcCtx.Config.Alipay.PublicKey, notifyReq)
	if err != nil {
		a.Logger.Errorf("alipay.VerifySign err :%+v", err)
		return
	}
	a.Logger.Infof("VerifySign ok ")
	var result AliWebPayResult
	err = notifyReq.Unmarshal(&result)
	if err != nil {
		a.Logger.Errorf("alipay.Unmarshal err :%+v", err)
		return
	}

	a.Logger.Infof("alipay.notifyReq.Unmarshal result:%+v ", result)
	res = &PayCallback{
		OutTradeNo: result.OutTradeNo,
	}
	if result.TradeStatus != "TRADE_SUCCESS" {
		res.Result = "支付失败"
		return
	}
	// ====异步通知，返回支付宝平台的信息====
	// 文档：https://opendocs.alipay.com/open/203/105286
	// 程序执行完后必须打印输出“success”（不包含引号）。如果商户反馈给支付宝的字符不是success这7个字符，支付宝服务器会不断重发通知，直到超过24小时22分钟。一般情况下，25小时以内完成8次通知（通知的间隔频率一般是：4m,10m,10m,1h,2h,6h,15h）
	res.Status = 1
	res.CallBackInfo = "success"
	return
}

// AliWebPayResult 支付宝支付结果回调
type AliWebPayResult struct {
	NotifyTime       string `json:"notify_time"`
	NotifyType       string `json:"notify_type"`
	NotifyID         string `json:"notify_id"`
	SignType         string `json:"sign_type"`
	Sign             string `json:"sign"`
	OutTradeNo       string `json:"out_trade_no"`
	Subject          string `json:"subject"`
	PayMentType      string `json:"payment_type"`
	TradeNum         string `json:"trade_no"`
	TradeStatus      string `json:"trade_status"`
	GmtPayMent       string `json:"gmt_payment"`
	GmtClose         string `json:"gmt_close"`
	SellerEmail      string `json:"seller_email"`
	BuyerEmail       string `json:"buyer_email"`
	SellerID         string `json:"seller_id"`
	BuyerID          string `json:"buyer_id"`
	Price            string `json:"price"`
	TotalFee         string `json:"total_fee"`
	Quantity         string `json:"quantity"`
	Body             string `json:"body"`
	Discount         string `json:"discount"`
	IsTotalFeeAdjust string `json:"is_total_fee_adjust"`
	UseCoupon        string `json:"use_coupon"`
	RefundStatus     string `json:"refund_status"`
	GmtRefund        string `json:"gmt_refund"`
}

func (a *AlipayService) wapPay(charge *Charge) (qrCode string, err error) {
	returnUrl := a.svcCtx.Config.Alipay.WapReturnUrl + charge.ReturnURL + "?pay_type=" + strconv.Itoa(int(charge.PayType)) + "&pay_method=" + strconv.Itoa(int(charge.PayMethod)) + "&source=" + strconv.Itoa(int(charge.Source)) + "&rights_id=" + strconv.Itoa(int(charge.RightsId)) + "&message_id=" + charge.MessageId
	a.alipay = a.alipay.SetReturnUrl(returnUrl).SetNotifyUrl(a.svcCtx.Config.Alipay.WapNotifyUrl)
	// 生成单号
	amount := decimal.NewFromInt(charge.MoneyFee).DivRound(decimal.NewFromInt(100), 2).String()
	// 构造参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", charge.Describe).
		Set("out_trade_no", charge.TradeNum).
		Set("total_amount", amount).
		Set("quit_url", a.svcCtx.Config.Alipay.WapQuitUrl+charge.QuitUrl).
		Set("product_code", "FAST_INSTANT_TRADE_PAY")
	// 发起支付
	qrCode, err = a.alipay.TradeWapPay(a.ctx, bm)
	if err != nil {
		if bizError, ok := alipay.IsBizError(err); ok {
			a.Logger.Errorf("s.alipay.TradeWapPay(%v), bizError:%v", bm, bizError)
			return
		}
		a.Logger.Errorf("s.alipay.TradeWapPay(%v), err:%v", bm, err)
		return
	}
	return
}

func (a *AlipayService) androidPay(charge *Charge) (qrCode string, err error) {
	client, err := alipay.NewClient(a.svcCtx.Config.AlipayApp.Appid, a.svcCtx.Config.AlipayApp.PrivateKey, true)
	if err != nil {
		panic(err)
	}
	// Debug开关，输出/关闭日志
	client.DebugSwitch = gopay.DebugOff
	// 配置公共参数
	client.SetCharset(alipay.UTF8).
		SetSignType(alipay.RSA2).SetNotifyUrl(a.svcCtx.Config.AlipayApp.NotifyUrl)
	// 生成单号
	amount := decimal.NewFromInt(charge.MoneyFee).DivRound(decimal.NewFromInt(100), 2).String()
	// 构造参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", charge.Describe).
		Set("out_trade_no", charge.TradeNum).
		Set("total_amount", amount).
		Set("product_code", "FAST_INSTANT_TRADE_PAY")
	// 发起支付
	qrCode, err = client.TradeAppPay(a.ctx, bm)
	if err != nil {
		if bizError, ok := alipay.IsBizError(err); ok {
			a.Logger.Errorf("s.alipay.TradeWapPay(%v), bizError:%v", bm, bizError)
			return
		}
		a.Logger.Errorf("s.alipay.TradeWapPay(%v), err:%v", bm, err)
		return
	}
	return
}
