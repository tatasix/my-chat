package sms

import (
	"chat/common/util"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	url1 "net/url"
)

const (
	//SMSTemplate 模版
	SMSTemplate = "【白小喵】您的验证码为{code}，该验证码30分钟内有效，请勿泄漏于他人！"
	// Url 账号 https://api.smsbao.com/sms?u=brchen&p=5e32e280ed7c4fe6b827a093dd01c7d3&m={mobile}&c={content}
	Url = "https://api.smsbao.com/sms"
	//ApiKey 密码
	ApiKey = "5e32e280ed7c4fe6b827a093dd01c7d3"
	//Account 发送url，
	Account = "brchen"
)

//0     短信发送成功
//-1    参数不全
//-2    服务器空间不支持,请确认支持curl或者fsocket，联系您的空间商解决或者更换空间
//30	错误密码
//40	账号不存在
//41	余额不足
//43	IP地址限制
//50	内容含有敏感词
//51	手机号码不正确

// SendSms 发送短信
func SendSms(ctx context.Context, mobile, message string) error {
	if mobile == "" {
		return errors.New("mobile is not null")
	}
	if !util.CheckMobile(mobile) {
		return errors.New("mobile is irregular")
	}

	url := util.NewUrlBuilder(Url).
		AddParam("u", Account).
		AddParam("p", ApiKey).
		AddParam("m", mobile).
		AddParam("c", message).
		Build()
	//不同信道参数可能不同，具体查看其开发文档
	body, err := util.GetV2(url)
	if err != nil {
		return err
	}
	decodedStr, _ := url1.QueryUnescape(url)
	logx.WithContext(ctx).Infof("SendSms url:%s decode-url:%s response:%+v err:%+v", url, decodedStr, body, err)
	if body != "0" {
		return errors.New("message send error")
	}
	return nil
}
