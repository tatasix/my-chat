package structure

// UsageSubscription 订阅信息
type UsageSubscription struct {
	Object             string  `json:"object"`
	HasPaymentMethod   bool    `json:"has_payment_method"`
	Canceled           bool    `json:"canceled"`
	CanceledAt         any     `json:"canceled_at"`
	Delinquent         any     `json:"delinquent"`
	AccessUntil        int     `json:"access_until"` // key到期时间
	SoftLimit          int     `json:"soft_limit"`
	HardLimit          int     `json:"hard_limit"`
	SystemHardLimit    int     `json:"system_hard_limit"`
	SoftLimitUsd       float64 `json:"soft_limit_usd"`
	HardLimitUsd       float64 `json:"hard_limit_usd"` // 总计可使用金额
	SystemHardLimitUsd float64 `json:"system_hard_limit_usd"`
	Plan               struct {
		Title string `json:"title"`
		Id    string `json:"id"` // free 免费
	} `json:"plan"`
	AccountName     string `json:"account_name"` // 账户名称
	PoNumber        any    `json:"po_number"`
	BillingEmail    any    `json:"billing_email"`
	TaxIds          any    `json:"tax_ids"`
	BillingAddress  any    `json:"billing_address"`
	BusinessAddress any    `json:"business_address"`
}

// UsageDailyList 每日使用情况列表
type UsageDailyList struct {
	Object     string `json:"object"`
	DailyCosts []struct {
		Timestamp float64 `json:"timestamp"`
		LineItems []struct {
			Name string  `json:"name"`
			Cost float64 `json:"cost"`
		} `json:"line_items"`
	} `json:"daily_costs"`
	TotalUsage float64 `json:"total_usage"` //  505.79316000000006
}

// UsageInfo 使用情况
type UsageInfo struct {
	AccessUntil        string  `json:"access_until"`         // key 到期时间
	HardLimitUsd       float64 `json:"hard_limit_usd"`       // 总计金额
	AccountName        string  `json:"account_name"`         // 账户名称
	UsedAmountUsd      float64 `json:"used_amount_usd"`      // 已使用金额
	RemainingAmountUsd float64 `json:"remaining_amount_usd"` // 剩余可用金额
}

type (
	EmbeddingResponse struct {
		Object string         `json:"object"`
		Data   []Embedding    `json:"data"`
		Model  string         `json:"model"`
		Usage  EmbeddingUsage `json:"usage"`
	}

	EmbeddingUsage struct {
		PromptTokens int `json:"prompt_tokens"`
		TotalTokens  int `json:"total_tokens"`
	}

	Embedding struct {
		Object    string    `json:"object"`
		Embedding []float64 `json:"embedding"`
		Index     int       `json:"index"`
	}
)

type ChatModelMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
