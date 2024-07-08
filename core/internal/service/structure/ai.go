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

type Message struct {
	Role    string  `json:"role"`
	Content string  `json:"content"`
	Name    *string `json:"name,omitempty"`
}

type GeneralOpenAIRequest struct {
	Model       string    `json:"model,omitempty"`
	Messages    []Message `json:"messages,omitempty"`
	Prompt      any       `json:"prompt,omitempty"`
	Stream      bool      `json:"stream,omitempty"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
	Temperature float64   `json:"temperature,omitempty"`
	TopP        float64   `json:"top_p,omitempty"`
	N           int       `json:"n,omitempty"`
	Input       any       `json:"input,omitempty"`
	Instruction string    `json:"instruction,omitempty"`
	Size        string    `json:"size,omitempty"`
}

type ChatRequest struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

type TextRequest struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	Prompt    string    `json:"prompt"`
	MaxTokens int       `json:"max_tokens"`
	//Stream   bool      `json:"stream"`
}

type ImageRequest struct {
	Prompt string `json:"prompt"`
	N      int    `json:"n"`
	Size   string `json:"size"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type OpenAIError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   string `json:"param"`
	Code    any    `json:"code"`
}

type OpenAIErrorWithStatusCode struct {
	OpenAIError
	StatusCode int `json:"status_code"`
}

type TextResponse struct {
	Choices []OpenAITextResponseChoice `json:"choices"`
	Usage   `json:"usage"`
	Error   OpenAIError `json:"error"`
}

type OpenAITextResponseChoice struct {
	Index        int `json:"index"`
	Message      `json:"message"`
	FinishReason string `json:"finish_reason"`
}

type OpenAITextResponse struct {
	Id      string                     `json:"id"`
	Object  string                     `json:"object"`
	Created int64                      `json:"created"`
	Choices []OpenAITextResponseChoice `json:"choices"`
	Usage   `json:"usage"`
}

type OpenAIEmbeddingResponseItem struct {
	Object    string    `json:"object"`
	Index     int       `json:"index"`
	Embedding []float64 `json:"embedding"`
}

type OpenAIEmbeddingResponse struct {
	Object string                        `json:"object"`
	Data   []OpenAIEmbeddingResponseItem `json:"data"`
	Model  string                        `json:"model"`
	Usage  `json:"usage"`
}

type ImageResponse struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	}
}

type ChatCompletionsStreamResponseChoice struct {
	Delta struct {
		Content string `json:"content"`
	} `json:"delta"`
	FinishReason string `json:"finish_reason,omitempty"`
}

type ChatCompletionsStreamResponse struct {
	Id      string                                `json:"id"`
	Object  string                                `json:"object"`
	Created int64                                 `json:"created"`
	Model   string                                `json:"model"`
	Choices []ChatCompletionsStreamResponseChoice `json:"choices"`
}

type CompletionsStreamResponse struct {
	Choices []struct {
		Text         string `json:"text"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
}
