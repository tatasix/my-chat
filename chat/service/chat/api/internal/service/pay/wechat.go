package pay

import (
	"chat/common/util"
	"chat/common/xerr"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/svc"
	"context"
	"encoding/json"
	"github.com/go-pay/gopay/wechat/v3"
	"net/http"
	"net/url"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"sync"

	"github.com/go-pay/gopay"
	"github.com/shopspring/decimal"
)

type WechatPayService struct {
	ctx context.Context
	logx.Logger
	svcCtx *svc.ServiceContext

	rwMu sync.RWMutex

	wechatPay *wechat.ClientV3
}

func NewWechatPayService(ctx context.Context, svcCtx *svc.ServiceContext) *WechatPayService {

	// NewClientV3 初始化微信客户端 v3
	// mchid：商户ID 或者服务商模式的 sp_mchid
	// serialNo：商户证书的证书序列号
	// apiV3Key：apiV3Key，商户平台获取
	// privateKey：私钥 apiclient_key.pem 读取后的内容
	client, err := wechat.NewClientV3(svcCtx.Config.WechatPay.MchId, svcCtx.Config.WechatPay.MchSerialNo, svcCtx.Config.WechatPay.ApiV3Key, svcCtx.Config.WechatPay.PrivateKey)
	if err != nil {
		logx.WithContext(ctx).Error(err)
		return nil
	}

	// 设置微信平台API证书和序列号（推荐开启自动验签，无需手动设置证书公钥等信息）
	//client.SetPlatformCert([]byte(""), "")

	// 启用自动同步返回验签，并定时更新微信平台API证书（开启自动验签时，无需单独设置微信平台API证书和序列号）
	err = client.AutoVerifySign()
	if err != nil {
		logx.WithContext(ctx).Error(err)
		return nil
	}

	// 自定义配置http请求接收返回结果body大小，默认 10MB
	//client.SetBodySize() // 没有特殊需求，可忽略此配置

	// 打开Debug开关，输出日志，默认是关闭的
	client.DebugSwitch = gopay.DebugOn

	return &WechatPayService{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		wechatPay: client,
	}
}

// GetQrCode 支付码获取(用户扫码支付)
func (w *WechatPayService) GetQrCode(charge *Charge) (qrCode string, err error) {
	if charge.Source == repository.SourcePC {
		return w.pcPay(charge)
	} else if charge.Source == repository.SourceH5 {
		return w.wapPay(charge)
	} else if charge.Source == repository.SourceMiniProgram {
		return w.miniProgramPay(charge)
	}
	return "", util.ReturnError(xerr.SystemBusyError)
}

// GetNotifyResult 支付码获取(用户扫码支付)
func (w *WechatPayService) GetNotifyResult(r *http.Request) (res *PayCallback, err error) {
	res = &PayCallback{}
	// 解析异步通知的参数
	notifyReq, err := wechat.V3ParseNotify(r)
	if err != nil {
		w.Logger.Errorf("wechatPay.V3ParseNotify err :%+v", err)
		return
	}
	if notifyReq == nil || notifyReq.Summary != "支付成功" {
		res.Result = notifyReq.Summary
		return
	}
	// 获取微信平台证书
	certMap := w.wechatPay.WxPublicKeyMap()
	// 验证异步通知的签名
	err = notifyReq.VerifySignByPKMap(certMap)
	if err != nil {
		w.Logger.Errorf("wechatPay.VerifySignByPKMap err :%+v", err)
		return
	}
	// 敏感信息解密
	//w.wechatPay.V3DecryptText()

	// ====↓↓↓====异步通知参数解密====↓↓↓====

	// 普通支付通知解密
	result, err := notifyReq.DecryptCipherText(w.svcCtx.Config.WechatPay.ApiV3Key)
	if err != nil {
		w.Logger.Error(err)
		return
	}
	//result
	w.Logger.Infof("wechatPay.DecryptCipherText notifyReq:%+v ", result)

	res.Status = 1
	res.OutTradeNo = result.OutTradeNo
	back, _ := json.Marshal(&wechat.V3NotifyRsp{Code: gopay.SUCCESS, Message: "成功"})
	res.CallBackInfo = string(back)
	return
}

func (w *WechatPayService) wapPay(charge *Charge) (qrCode string, err error) {
	amount := decimal.NewFromInt(charge.MoneyFee).DivRound(decimal.NewFromInt(1), 2).IntPart()

	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("appid", w.svcCtx.Config.WechatPay.Appid).
		Set("mchid", w.svcCtx.Config.WechatPay.MchId).
		Set("description", charge.Describe).
		Set("out_trade_no", charge.TradeNum).
		Set("time_expire", expire).
		Set("notify_url", w.svcCtx.Config.WechatPay.NotifyUrl).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", amount).
				Set("currency", "CNY")
		}).
		SetBodyMap("scene_info", func(bm gopay.BodyMap) {
			bm.Set("payer_client_ip", util.GetLocalIp()).
				SetBodyMap("h5_info", func(bm gopay.BodyMap) {
					bm.Set("type", "Wap")
				})
		})

	// 发起支付
	rsp, err := w.wechatPay.V3TransactionH5(w.ctx, bm)
	if err != nil {
		return
	}
	// return
	qrCode = rsp.Response.H5Url + "&redirect_url=" + url.QueryEscape(w.svcCtx.Config.WechatPay.WapReturnUrl+charge.ReturnURL)

	return
}

func (w *WechatPayService) pcPay(charge *Charge) (qrCode string, err error) {
	// 生成单号
	amount := decimal.NewFromInt(charge.MoneyFee).DivRound(decimal.NewFromInt(1), 2).IntPart()

	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("appid", w.svcCtx.Config.WechatPay.Appid).
		Set("mchid", w.svcCtx.Config.WechatPay.MchId).
		Set("description", charge.Describe).
		Set("out_trade_no", charge.TradeNum).
		Set("time_expire", expire).
		Set("notify_url", w.svcCtx.Config.WechatPay.NotifyUrl).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", amount).
				Set("currency", "CNY")
		})

	// 发起支付
	rsp, err := w.wechatPay.V3TransactionNative(w.ctx, bm)
	if err != nil {
		return
	}
	// return
	qrCode = rsp.Response.CodeUrl

	return
}

// GetQrCode 支付码获取(用户扫码支付)
func (w *WechatPayService) miniProgramPay(charge *Charge) (qrCode string, err error) {
	if charge.OpenID == "" {
		err = util.ReturnError(xerr.ParamMiss)
		return
	}
	// 生成单号
	amount := decimal.NewFromInt(charge.MoneyFee).DivRound(decimal.NewFromInt(1), 2).IntPart()

	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("appid", w.svcCtx.Config.WechatPay.Appid).
		Set("mchid", w.svcCtx.Config.WechatPay.MchId).
		Set("description", charge.Describe).
		Set("out_trade_no", charge.TradeNum).
		Set("time_expire", expire).
		Set("notify_url", w.svcCtx.Config.WechatPay.NotifyUrl).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", amount).
				Set("currency", "CNY")
		}).
		SetBodyMap("payer", func(bm gopay.BodyMap) {
			bm.Set("openid", charge.OpenID)
		})

	// 发起支付
	rsp, err := w.wechatPay.V3TransactionJsapi(w.ctx, bm)
	if err != nil {
		return
	}
	// return
	//qrCode = rsp.Response.PrepayId

	jsapiInfo, err := w.wechatPay.PaySignOfJSAPI(w.svcCtx.Config.WechatPay.Appid, rsp.Response.PrepayId)
	str, _ := json.Marshal(jsapiInfo)
	qrCode = string(str)
	return
}
