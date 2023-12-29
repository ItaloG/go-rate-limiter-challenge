package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ItaloG/go-rate-limiter-challenge/pkg/middlewares/handlers"
	"github.com/ItaloG/go-rate-limiter-challenge/pkg/redis"
)

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RemoteAddr)
		redisAddr := os.Getenv("REDIS_HOST")
		redisClient := redis.NewRedisClient(redisAddr)
		redisRateLimit := &redis.RedisRateLimit{Client: redisClient.Client}

		err := handlers.Handle(redisRateLimit, "teste", time.Second*3, 2)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		fmt.Println("Rate Limit")
		next.ServeHTTP(w, r)
	})
}
