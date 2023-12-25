package redis

import (
	"context"
	"time"

	redisV9 "github.com/redis/go-redis/v9"
)

type RedisRateLimit struct {
	Client *redisV9.Client
}

func (r *RedisRateLimit) AddKey(ctx context.Context, key string, duration time.Duration) error {
	err := r.Client.Incr(ctx, key).Err()
	if err != nil {
		return err
	}
	err = r.Client.Expire(ctx, key, duration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRateLimit) GetKey(ctx context.Context, key string) string {
	result := r.Client.Get(ctx, key)
	if result == nil {
		return ""
	}
	return result.Name()
}

func (r *RedisRateLimit) VerifyKey(ctx context.Context, key string, limit int) (bool, error) {
	count, err := r.Client.Get(ctx, key).Int()

	if err != nil {
		return false, err
	}

	if count < limit {
		return false, nil
	}

	return true, nil
}
