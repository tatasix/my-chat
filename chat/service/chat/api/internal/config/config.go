package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	Mysql struct {
		DataSource string
	}

	RedisCache cache.CacheConf

	SystemVersion string `json:",optional,default=v0.6.2.2"`

	// jwt 配置
	Auth struct {
		AccessSecret string `json:",optional,default=13450cd8841c0f0"`
		AccessExpire int64  `json:",optional,default=25920000"`
	}

	// jwt 配置
	FrontendAuth struct {
		AccessSecret string `json:",optional,default=13safhasfuawefc0f0"`
		AccessExpire int64  `json:",optional,default=25920000"`
	}

	// 企业微信，配置信息
	WeCom struct {
		Port                  int `json:",optional,default=8887"`
		CorpID                string
		DefaultAgentSecret    string `json:",optional"`
		CustomerServiceSecret string `json:",optional"`
		CorpSecret            string `json:",optional"`
		Model                 string `json:",optional,default=gpt-3.5-turbo"`
		BasePrompt            string `json:",optional,default=你是ChatGPT，一个由OpenAI训练的大型语言模型，你旨在回答并解决人们的任何问题，并且可以使用多种语言与人交流。"`
		Welcome               string `json:",optional,default=您好！我是ChatGPT，一个由OpenAI训练的大型语言模型，我可以回答您的问题和进行交流。请告诉我您需要了解些什么，我会尽力为您提供答案。\n\n发送#help查看更多功能"`
		Token                 string `json:",optional"`
		EncodingAESKey        string `json:",optional"`
		MultipleApplication   []struct {
			AgentID         int64
			AgentSecret     string
			Model           string `json:",optional,default=gpt-3.5-turbo"`
			BasePrompt      string `json:",optional,default=你是ChatGPT，一个由OpenAI训练的大型语言模型，你旨在回答并解决人们的任何问题，并且可以使用多种语言与人交流。"`
			Welcome         string `json:",optional,default=您好！我是ChatGPT，一个由OpenAI训练的大型语言模型，我可以回答您的问题和进行交流。请告诉我您需要了解些什么，我会尽力为您提供答案。\n\n发送#help查看更多功能"`
			GroupEnable     bool   `json:",optional,default=false"`
			GroupName       string `json:",optional,default=ChatGPT应用内部交流群"`
			GroupChatID     string `json:",optional,default=ChatGPT202304021958"`
			EmbeddingEnable bool   `json:",optional,default=false"`
			EmbeddingMode   string `json:",optional,default=QA"`
			TopK            int    `json:",optional,default=1"`
			Score           string `json:",optional,default=0.3"`
		} `json:",optional"`
		MultipleCustomer []struct {
			EmbeddingEnable bool   `json:",optional,default=true"`
			EmbeddingMode   string `json:",optional,default=ARTICLE"`
			TopK            int    `json:",optional,default=1"`
			Score           string `json:",optional,default=0.3"`
			OpenKfId        string `json:",optional"`
		} `json:",optional"`
	}

	// openai 配置
	OpenAi struct {
		Key    []string
		Host   string `json:"host,optional,default=https://api.openai.com"`
		Origin string `json:"origin,optional,default=open_ai"`
		Engine string `json:"engine,optional,default="`
	}

	// http proxy 设置
	Proxy struct {
		Enable  bool   `json:",optional,default=false"`
		Socket5 string `json:",optional,default=127.0.0.1:1080"`
		Http    string `json:",optional,default="`
	} `json:",optional"`

	// ocr 配置
	OCR struct {
		Company string `json:",optional"`
		AliYun  struct {
			AccessKeyId     string
			AccessKeySecret string
		} `json:",optional"`
	} `json:",optional"`

	// embeddings 配置
	Embeddings struct {
		Enable bool   `json:",optional,default=false"`
		Mode   string `json:",optional,default=QA"`
		Milvus struct {
			Host     string   `json:",optional,default=127.0.0.1:19530"`
			Username string   `json:",optional"`
			Password string   `json:",optional"`
			Keywords []string `json:",optional"`
		} `json:",optional"`
	}

	// 流式传输 加快响应速度
	Response struct {
		Stream bool `json:",optional,default=true"`
	}

	// Plugins 配置
	Plugins struct {
		Enable bool `json:",optional,default=false"`
		List   []struct {
			Enable       bool   `json:",optional,default=true"`
			NameForHuman string `json:",optional"`
			NameForModel string `json:",optional"`
			DescForHuman string `json:",optional"`
			DescModel    string `json:",optional"`
			Auth         struct {
				Type string `json:",optional"`
			} `json:",optional"`
			API struct {
				URL string `json:",optional"`
			}
		} `json:",optional"`
	}

	// 语音解析与合成 后续可能会使用
	Speaker struct {
		Company string `json:",optional,default=openai"`
		AliYun  struct {
			AccessKeyId     string
			AccessKeySecret string
			AppKey          string
		} `json:",optional"`
	} `json:",optional"`

	// 作图相关配置，目前只支持 StableDiffusion
	Draw struct {
		Enable          bool `json:",optional,default=false"`
		StableDiffusion struct {
			Host string `json:",optional,default=http://localhost:7890"`
			Auth struct {
				Username string
				Password string
			}
		}
	} `json:",optional"`

	Domain  string
	H5Uri   string
	OssType string `json:"OssType,optional,default=local"`
	Env     string `json:",optional,default=dev"`
	Wechat  struct {
		Appid          string
		Secret         string
		RedirectUrl    string
		EncodingAESKey string
		Token          string
		Scope          string
		RedirectUri    string
	}
	Local struct {
		Path      string `json:",optional,default=false"`
		StorePath string `json:",optional,default=false"`
	} `json:",optional"`

	QiNiu struct {
		Zone          string // 存储区域
		Bucket        string // 空间名称
		ImgPath       string // CDN加速域名
		UseHTTPS      bool   // 是否使用https
		AccessKey     string // 秘钥AK
		SecretKey     string // 秘钥SK
		UseCdnDomains bool   // 上传是否使用CDN上传加速
	}

	Captcha struct {
		Open               bool `json:",optional,true"`         // 是否开启
		KeyLong            int  `json:",optional,default=4"`    // 验证码长度
		ImgWidth           int  `json:",optional,default=240"`  // 验证码宽度
		ImgHeight          int  `json:",optional,default=80"`   // 验证码高度
		OpenCaptcha        int  `json:",optional,default=0"`    // 防爆破验证码开启此数，0代表每次登录都需要验证码，其他数字代表错误密码此数，如3代表错误三次后出现验证码
		OpenCaptchaTimeOut int  `json:",optional,default=3600"` // 防爆破验证码超时时间，单位：s(秒)
	}
	WechatPay struct {
		Appid        string `json:"Appid"`
		MchId        string `json:"MchId"`
		ApiKey       string `json:"ApiKey"`
		ApiV3Key     string `json:"ApiV3Key"`
		MchSerialNo  string `json:"MchSerialNo"`
		PrivateKey   string `json:"PrivateKey"`
		NotifyUrl    string `json:"NotifyUrl"`
		WapReturnUrl string `json:"WapReturnUrl"`
	}

	Alipay struct {
		Appid                   string `json:"Appid"`
		PrivateKey              string `json:"PrivateKey"`
		PublicKey               string `json:"PublicKey"`
		AppPublicCertContent    string `json:"AppPublicCertContent"`
		AlipayRootCertContent   string `json:"AlipayRootCertContent"`
		AlipayPublicCertContent string `json:"AlipayPublicCertContent"`
		NotifyUrl               string `json:"NotifyUrl"`
		ReturnUrl               string `json:"ReturnUrl"`
		WapNotifyUrl            string `json:"WapNotifyUrl"`
		WapReturnUrl            string `json:"WapReturnUrl"`
		WapQuitUrl              string `json:"WapQuitUrl"`
	}

	AlipayApp struct {
		Appid      string `json:"Appid"`
		PrivateKey string `json:"PrivateKey"`
		PublicKey  string `json:"PublicKey"`
		NotifyUrl  string `json:"NotifyUrl"`
	}
	PayPrice bool `json:"PayPrice,optional,default=true"`

	ApplePay struct {
		Iss        string `json:"Iss"`
		Bid        string `json:"Bid"`
		Kid        string `json:"Kid"`
		PrivateKey string `json:"PrivateKey"`
		IsProd     bool   `json:"IsProd"`
	}

	Sensitive struct {
		Type      string `json:"Type"`
		BucketURL string `json:"BucketURL"`
		CIURL     string `json:"CIURL"`
		SecretId  string `json:"SecretId"`
		SecretKey string `json:"SecretKey"`
	}
}

type Redis struct {
	Host string `json:"host"`
	Pass string `json:"pass"`
}
