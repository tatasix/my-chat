package cache

import (
	"chat/common/redis"
	"context"
	"fmt"
)

func SetDialogueCache(ctx context.Context, value, kfId, dialogueId string, status uint8) {

	redis.Rdb.Set(ctx, fmt.Sprintf(redis.ChatRoomDialogueCacheKey, kfId, dialogueId, status), value, redis.ChatRoomDialogueCacheKeyExpireTime)

}

func GetDialogueCache(ctx context.Context, kfId, dialogueId string, status uint8) (string, error) {
	return redis.Rdb.Get(ctx, fmt.Sprintf(redis.ChatRoomDialogueCacheKey, kfId, dialogueId, status)).Result()
}
