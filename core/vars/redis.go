package vars

import "time"

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

const LoginCodeKey = "chat:login:code:%s"
const LoginCodeExpireTime = 30 * time.Minute

const LogoutKey = "chat:logout:black"
const LogoutExpireTime = 30 * time.Minute
