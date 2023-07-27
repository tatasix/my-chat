package redis

import (
	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

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
