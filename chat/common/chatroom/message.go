package chatroom

const (
	SendMessageToCustomer    = 1
	SendMessageGenerateOne   = 2
	SendMessageToUserWriting = 3
	SendMessageCreate        = 4
	SendMessageJoin          = 5
	SendMessageOut           = 6
	SendMessageDisconnect    = 7
	SendMessageStart         = 8
	SendMessageReady         = 9
	SendMessagePong          = 10
)

type Message struct {
	ChannelId    string        `json:"channel_id"`    // 管道ID
	ChannelTitle string        `json:"channel_title"` // 管道标题
	User         string        `json:"user"`
	Users        []interface{} `json:"users"`
	UserName     string        `json:"user_name"`
	Detail       interface{}   `json:"detail"`
	ToUserId     string        `json:"to_user_id"`
	KfId         string        `json:"kf_id"`
	Type         uint8         `json:"type"`       // 消息类型
	Status       uint8         `json:"status"`     // 消息类型
	Content      string        `json:"content"`    // 消息内容
	MessageId    string        `json:"message_id"` // 消息内容
	SendTime     string        `json:"send_time"`
	DialogueId   string        `json:"dialogue_id"`
	ErrorMessage string        `json:"error_message"`
	IsEnd        bool          `json:"is_end"`
	IsStreamEnd  bool          `json:"is_stream_end"`
	IsKf         bool          `json:"is_kf"`
	Heartbeat    bool          `json:"heartbeat"`
}

type messageInterface interface {
	SendMessage(msg Message)
	DelaySendMessage(channelId string, msg Message, uids []string)
	LoginServer(uid string)
	LogoutServer(uid string)
	ErrorLogServer(err error)
	InfoLogServer(msg string)
}

var msgStore messageInterface
