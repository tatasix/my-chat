package captcha

import (
	"chat/common/redis"
	"context"
	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

func NewDefaultRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_",
	}
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
	logger     logx.Logger
}

func (rs *RedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	rs.logger = logx.WithContext(ctx)
	return rs
}

func (rs *RedisStore) Set(id string, value string) error {
	err := redis.Rdb.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		rs.logger.Errorf("RedisStoreSetError! ", err)
	}
	return err
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := redis.Rdb.Get(rs.Context, key).Result()
	if err != nil {
		rs.logger.Errorf("RedisStoreGetError! ", err)
		return ""
	}
	if clear {
		err := redis.Rdb.Del(rs.Context, key).Err()
		if err != nil {
			rs.logger.Errorf("RedisStoreClearError! ", err)
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}
