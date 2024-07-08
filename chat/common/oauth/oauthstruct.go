package oauth

// 临时票据结果
type CodeResult struct {
	Code int `json:"code"`
}

// 令牌结果
type TokenResult struct {
	AccessToken  string `json:"accessToken"`
	ExpireIn     string `json:"expireIn"`
	RefreshToken string `json:"refreshToken"`
	Uid          string `json:"uid"`
	OpenId       string `json:"openId"`
	AccessCode   string `json:"accessCode"`
	UnionId      string `json:"unionId"`

	// google
	Scope     string `json:"scope"`
	TokenType string `json:"tokenType"`
	IdToken   string `json:"idToken"`

	// mi
	MacAlgorithm string `json:"macAlgorithm"`
	MacKey       string `json:"macKey"`

	// wechat
	Code       string `json:"code"`
	SessionKey string `json:"sessionKey"`

	// twitter
	OauthToken             string `json:"oauthToken"`
	OauthTokenSecret       string `json:"oauthTokenSecret"`
	UserId                 string `json:"userId"`
	ScreenName             string `json:"screenName"`
	OauthCallbackConfirmed bool   `json:"oauthCallbackConfirmed"`
}

// 用户信息
type UserResult struct {
	UUID      string       `json:"uuid"`
	OpenId    string       `json:"open_id"`
	UnionId   string       `json:"union_id"`
	UserName  string       `json:"user_name"`
	NickName  string       `json:"nick_name"`
	AvatarUrl string       `json:"avatar"`
	Company   string       `json:"company"`
	Language  string       `json:"language"`
	City      string       `json:"city"`
	Province  string       `json:"province"`
	Country   string       `json:"country"`
	Blog      string       `json:"blog"`
	Location  string       `json:"location"`
	Mobile    string       `json:"mobile"`
	Email     string       `json:"email"`
	Remark    string       `json:"remark"`
	Url       string       `json:"url"`
	Gender    int32        `json:"gender"`
	Source    string       `json:"source"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
	Token     *TokenResult `json:"token"`
}

// 微信绑定手机
type WxMobileResult struct {
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     string `json:"countryCode"`
}
