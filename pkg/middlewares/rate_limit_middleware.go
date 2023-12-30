package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/ItaloG/go-rate-limiter-challenge/pkg/ip"
	"github.com/ItaloG/go-rate-limiter-challenge/pkg/middlewares/handlers"
	"github.com/ItaloG/go-rate-limiter-challenge/pkg/redis"
)

var (
	err error

	blockIp      string
	blockIpLimit int
	blockIpTime  int

	blockToken      string
	blockTokenLimit int
	blockTokenTime  int
)

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		redisAddr := os.Getenv("REDIS_HOST")
		redisClient := redis.NewRedisClient(redisAddr)
		redisRateLimit := redis.NewRedisRateLimit(redisClient.Client)

		if os.Getenv("BLOCK_IP_LIMIT") == "" {
			blockIpLimit = 0
		} else {
			blockIpLimit, err = strconv.Atoi(os.Getenv("BLOCK_IP_LIMIT"))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		if os.Getenv("BLOCK_IP_TIME") == "" {
			blockIpTime = 0
		} else {
			blockIpTime, err = strconv.Atoi(os.Getenv("BLOCK_IP_TIME"))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		if os.Getenv("BLOCK_TOKEN_LIMIT") == "" {
			blockTokenLimit = 0
		} else {
			blockTokenLimit, err = strconv.Atoi(os.Getenv("BLOCK_TOKEN_LIMIT"))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		if os.Getenv("BLOCK_TOKEN_TIME") == "" {
			blockTokenTime = 0
		} else {
			blockTokenTime, err = strconv.Atoi(os.Getenv("BLOCK_TOKEN_TIME"))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		blockIp = os.Getenv("BLOCK_IP")
		blockToken = os.Getenv("BLOCK_TOKEN")

		clientIp := ip.GetIp(r.RemoteAddr)
		clientToken := r.Header.Get("API_KEY")

		err = handlers.Handle(redisRateLimit, blockIp, int64(blockIpLimit), blockIpTime, clientIp, blockToken, int64(blockTokenLimit), blockTokenTime, clientToken)
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(err.Error()))
			return
		}

		fmt.Println("Rate Limit")
		next.ServeHTTP(w, r)
	})
}
