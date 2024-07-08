package lock

import (
	"chat/common/redis"
	"context"
	"fmt"
	"time"
)

func Lock(key string, timeout time.Duration) (bool, error) {
	return redis.Rdb.SetNX(context.Background(), key, "", timeout).Result()
}

func Release(key string) error {
	res, err := redis.Rdb.Del(context.Background(), key).Result()
	fmt.Println(res)
	return err
}
