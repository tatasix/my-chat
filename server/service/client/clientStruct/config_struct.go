package clientStruct

type GetConfigRequest struct {
	ConfigType int64 `json:"config_type,optional"`
}

type AddConfigRequest struct {
	ConfigType  int64  `json:"config_type"` // 配置类型：1，openai key
	Description string `json:"description"` // 配置描述
	Value       string `json:"value"`       // 配置内容
	Name        string `json:"name"`        // 配置内容
	CreatedBy   string `json:"created_by"`  // 创建人
}

type GetConfigResponse struct {
	List []Config `json:"list"`
}

type Config struct {
	Id          int64  `json:"id"`
	ConfigType  int64  `json:"config_type"` // 配置类型：1，openai key
	Name        string `json:"name"`        // 配置名
	Description string `json:"description"` // 配置描述
	Value       string `json:"value"`       // 配置内容
	CreatedAt   string `json:"created_at"`  // 创建时间
	UpdatedAt   string `json:"updated_at"`  // 更新时间
	UpdatedBy   string `json:"updated_by"`  // 更新人
	CreatedBy   string `json:"created_by"`  // 创建人
}

type UpdateConfigRequest struct {
	ConfigType  int64  `json:"config_type,optional"` // 配置类型：1，openai key
	Id          int64  `json:"id"`
	Description string `json:"description,optional"` // 配置描述
	Value       string `json:"value,optional"`       // 配置内容
	Name        string `json:"name,optional"`        // 配置内容
	UpdateBy    string `json:"updated_by"`           // 创建人
}