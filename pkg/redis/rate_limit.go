package redis

import (
	"context"
	"fmt"
	"time"

	redisV9 "github.com/redis/go-redis/v9"
)

type RedisRateLimit struct {
	Client *redisV9.Client
}

func (r *RedisRateLimit) VerifyLimit(ctx context.Context, key string, limit int64, duration time.Duration) (bool, error) {
	isLimited := false

	count, err := r.Client.Incr(ctx, key).Result()
	fmt.Println(count)
	if err != nil {
		isLimited = true
		return isLimited, err
	}

	if count > limit {
		isLimited = true
		return isLimited, nil
	}

	err = r.Client.Expire(ctx, key, duration).Err()
	if err != nil {
		isLimited = true
		return isLimited, err
	}

	return isLimited, nil
}
