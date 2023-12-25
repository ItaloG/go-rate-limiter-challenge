package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ItaloG/go-rate-limiter-challenge/pkg/middlewares/handlers"
	"github.com/ItaloG/go-rate-limiter-challenge/pkg/redis"
)

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		redisClient := redis.NewRedisClient("localhost")
		redisRateLimit := &redis.RedisRateLimit{Client: redisClient.Client}

		err := handlers.Handle(redisRateLimit, "key", time.Second*8, 10)
		if err != nil {
			w.Write([]byte(err.Error()))
		}

		fmt.Println("Rate Limit")
		next.ServeHTTP(w, r)
	})
}
