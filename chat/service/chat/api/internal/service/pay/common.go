package pay

import (
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/svc"
	"context"
	"net/http"
)

// PayClient 支付客户端接口
type PayClient interface {
	GetQrCode(charge *Charge) (string, error)
	wapPay(charge *Charge) (string, error)
	pcPay(charge *Charge) (string, error)
	// GetNotifyResult 处理结果
	GetNotifyResult(r *http.Request) (res *PayCallback, err error)
}

// GetPayClient 得到需要支付的客户端
func GetPayClient(ctx context.Context, svcCtx *svc.ServiceContext, payMethod int64) PayClient {
	//如果使用余额支付
	switch payMethod {
	case repository.PayMethodAlipay:
		return NewAlipayService(ctx, svcCtx)
	case repository.PayMethodWechat:
		return NewWechatPayService(ctx, svcCtx)
	case repository.PayMethodApple:
		return NewApplePayService(ctx, svcCtx)
	}

	return nil
}

// Charge 支付参数
type Charge struct {
	TradeNum    string `json:"tradeNum,omitempty"`
	Origin      string `json:"origin,omitempty"`
	UserID      string `json:"userId,omitempty"`
	PayMethod   int64  `json:"payMethod,omitempty"`
	MoneyFee    int64  `json:"MoneyFee,omitempty"`
	CallbackURL string `json:"callbackURL,omitempty"`
	ReturnURL   string `json:"returnURL,omitempty"`
	QuitUrl     string `json:"QuitUrl,omitempty"`
	ShowURL     string `json:"showURL,omitempty"`
	Describe    string `json:"describe,omitempty"`
	OpenID      string `json:"openid,omitempty"`
	Source      int64  `json:"source,omitempty"`
	PayType     int64  `json:"pay_type,omitempty"`
	RightsId    int64  `json:"rights_id,omitempty"`
	MessageId   string `json:"message_id,omitempty"`
}

// PayCallback 支付返回
type PayCallback struct {
	Result       string
	OutTradeNo   string
	OrderNum     string
	CallBackURL  string
	CallBackInfo string
	Status       int64
}

// CallbackReturn 回调业务代码时的参数
type CallbackReturn struct {
	IsSucceed     bool   `json:"isSucceed"`
	OrderNum      string `json:"orderNum"`
	TradeNum      string `json:"tradeNum"`
	UserID        string `json:"userID"`
	MoneyFee      int64  `json:"moneyFee"`
	Sign          string `json:"sign"`
	ThirdDiscount int64  `json:"thirdDiscount"`
}

// BaseResult 支付结果
type BaseResult struct {
	IsSucceed     bool   // 是否交易成功
	TradeNum      string // 交易流水号
	MoneyFee      int64  // 支付金额
	TradeTime     string // 交易时间
	ContractNum   string // 交易单号
	UserInfo      string // 支付账号信息(有可能有，有可能没有)
	ThirdDiscount int64  // 第三方优惠
}
