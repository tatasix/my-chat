package redis

import (
	"github.com/go-redis/redis/v8"
	"time"
)

var Rdb *redis.Client

const WelcomeCacheKey = "chat:wecome:%d:%s"
const CursorCacheKey = "chat:cursor:%s"
const EmbeddingsCacheKey = "chat:embeddings:%s"
const UserSessionAgentDefaultKey = "session_agent:default:%s:%s"
const UserSessionListKey = "user:session:list:%s"
const SessionKey = "session:%s"

const OpenAIUsedKey = "chat:open_ai:used_key"

const ChatHistoryExportKey = "chat:history:export:%s"

const CustomerConfigGetByKfIdCacheKey = "chat:repository:customer_config:%s"
const CustomerConfigGetByKfIdCacheExpireTime = 10 * time.Minute

const CustomerConfigGetByTypeCacheKey = "chat:repository:customer_config:type:%d"
const CustomerConfigGetByTypeCacheExpireTime = 10 * time.Minute

const LoginCodeKey = "chat:login:code:%s"
const LoginCodeExpireTime = 30 * time.Minute

const LogoutKey = "chat:logout:black"
const LogoutExpireTime = 30 * time.Minute

const ChatContextCacheKey = "chat:context:customer_chat:%s:%s:%s"

const CacheCaptchaKey = "chat:cache:captcha:%s"

const PromptGetByKfIdAndStateCacheKey = "chat:repository:prompt:%s:%d"
const PromptGetByKfIdAndStateCacheExpireTime = 10 * time.Minute

const GetByChannelIdCacheKey = "chat:repository:chatroom:%d"
const GetByChannelIdCacheExpireTime = 10 * time.Minute

const ChatRoomDialogueCacheKey = "chat:chat_room:%s:%s:%d"
const ChatRoomDialogueCacheKeyExpireTime = 60 * time.Minute

const ChatRiskTimesKey = "chat:risk:times:%s:%s"
const ChatRiskVipKey = "chat:risk:vip:%s:%s"

func Init(Host, Pass string) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     Host,
		Password: Pass,
		DB:       1,
	})
}

func Close() {
	err := Rdb.Close()
	if err != nil {
		return
	}
}
