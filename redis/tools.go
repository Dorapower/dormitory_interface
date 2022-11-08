package redis

import (
	"github.com/go-redis/redis/v8"
)

func QueryAvaliableCountviaRedis(building string) (cnt int, err error) {
	cnt, err = rdb.Get(ctx, building).Int()
	if err == redis.Nil {
		return -1, redis.Nil
	} else if err != nil {
		panic(err)
	}
	return cnt, nil
}

func UpdateAvailableCount(building string, cnt int) {
	err := rdb.Set(ctx, building, cnt, 0).Err()
	if err != nil {
		panic(err)
	}
}
