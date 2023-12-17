package redis

import (
	"context"
	"lightyear/core/global"
	"time"
)

var ctx = context.Background()

const (
	DEFAULT_EXPIRATION_TIME = time.Hour * 24 * 7
)

func SetWithDefaultExpiration(key string, value any) error {
	return Set(key, value, DEFAULT_EXPIRATION_TIME)
}

func Set(key string, value any, expiration time.Duration) error {
	return global.Redis.Set(ctx, key, value, expiration).Err()
}

func Get(key string) (any, error) {
	return global.Redis.Get(ctx, key).Result()
}

func Del(key string) error {
	return global.Redis.Del(ctx, key).Err()
}
