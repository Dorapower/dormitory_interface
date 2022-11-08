package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "47.92.123.159:6379",
		Password: "",
		DB:       0,
	})
}
