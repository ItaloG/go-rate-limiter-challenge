package redis

import (
	"context"
	"time"

	"github.com/alicebob/miniredis/v2"
	redisV9 "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
)

type RedisRateLimitTestSuite struct {
	suite.Suite
	Server *miniredis.Miniredis
	Client *redisV9.Client
}

func (suite *RedisRateLimitTestSuite) SetupSuite() {
	server, err := miniredis.Run()
	suite.NoError(err)
	suite.Server = server
	suite.Client = redisV9.NewClient(&redisV9.Options{
		Addr: server.Addr(),
	})
}

func (suite *RedisRateLimitTestSuite) TearDownTest() {
	suite.Client.Close()
	suite.Server.Close()
}

func (suite *RedisRateLimitTestSuite) TestCallingVerifyLimitLessThanLimitAndDuration_ShouldReturnFalse() {
	rateLimiter := NewRedisRateLimit(suite.Client)

	key := "test_key"
	limit := int64(5)
	duration := time.Second

	isLimited, err := rateLimiter.VerifyLimit(context.Background(), key, limit, duration)
	suite.NoError(err)
	suite.False(isLimited)
}

func (suite *RedisRateLimitTestSuite) TestCallingVerifyLimitMoreThanLimitAndDuration_ShouldReturnTrue() {
	rateLimiter := NewRedisRateLimit(suite.Client)

	key := "test_key"
	limit := 5
	duration := time.Second

	for i := 0; i < limit; i++ {
		_, _ = rateLimiter.VerifyLimit(context.Background(), key, int64(limit), duration)
	}
	isLimited, err := rateLimiter.VerifyLimit(context.Background(), key, int64(limit), duration)
	suite.NoError(err)
	suite.True(isLimited)
}
