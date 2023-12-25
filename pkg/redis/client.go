package redis

import (
	redisV9 "github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redisV9.Client
}

var singletonRedisClient *RedisClient

func NewRedisClient(address string) *RedisClient {
	if singletonRedisClient != nil {
		return singletonRedisClient
	}
	client := redisV9.NewClient(&redisV9.Options{
		Addr: address,
		DB:   0,
	})

	singletonRedisClient = &RedisClient{Client: client}

	return singletonRedisClient
}
