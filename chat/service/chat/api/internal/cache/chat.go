package cache

import (
	"chat/common/openai"
	"chat/common/redis"
	"chat/service/chat/api/internal/vars"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func SetCache(ctx context.Context, summary []openai.ChatModelMessage, f, kfId, user string) {
	byteData, _ := json.Marshal(summary)
	t, _ := time.Parse(vars.TimeFormat, time.Now().Add(24*time.Hour).Format("2006-01-02")+" 00:00:00")

	redis.Rdb.Set(ctx, fmt.Sprintf(redis.ChatContextCacheKey, f, kfId, user), string(byteData), t.Sub(time.Now()))

}

func AddCache(ctx context.Context, role, content, f, kfId, user string) {
	var old []openai.ChatModelMessage
	data, _ := redis.Rdb.Get(ctx, fmt.Sprintf(redis.ChatContextCacheKey, f, kfId, user)).Result()
	_ = json.Unmarshal([]byte(data), &old)

	newSummary := append(old, openai.ChatModelMessage{
		Role:    role,
		Content: content,
	})

	byteData, _ := json.Marshal(newSummary)
	t, _ := time.Parse(vars.TimeFormat, time.Now().Add(24*time.Hour).Format("2006-01-02")+" 00:00:00")

	redis.Rdb.Set(ctx, fmt.Sprintf(redis.ChatContextCacheKey, f, kfId, user), string(byteData), t.Sub(time.Now()))

}

func GetCache(ctx context.Context, f, kfId, user string) []openai.ChatModelMessage {
	var m []openai.ChatModelMessage
	data, _ := redis.Rdb.Get(ctx, fmt.Sprintf(redis.ChatContextCacheKey, f, kfId, user)).Result()
	_ = json.Unmarshal([]byte(data), &m)
	return m
}

func DeleteCache(ctx context.Context, f, kfId, user string) {
	redis.Rdb.Del(ctx, fmt.Sprintf(redis.ChatContextCacheKey, f, kfId, user))
}
