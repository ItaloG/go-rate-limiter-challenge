package handlers

import (
	"context"
	"errors"
	"time"
)

type RateLimitInterface interface {
	VerifyLimit(ctx context.Context, key string, limit int64, duration time.Duration) (bool, error)
}

var ErrFoo = errors.New("you have reached the maximum number of requests or actions allowed within a certain time frame")

func Handle(r RateLimitInterface, blockIp string, blockIpLimit int64, blockIpTime int, clientIp string, blockToken string, blockTokenLimit int64, blockTokenTime int, clientToken string) error {
	ctx := context.Background()

	if blockToken != "" {
		if clientToken != blockToken {
			return nil
		}

		tokenIsLimited, err := r.VerifyLimit(ctx, blockToken, blockTokenLimit, time.Second*time.Duration(blockTokenTime))

		if err != nil {
			return err
		}

		if tokenIsLimited {
			return ErrFoo
		}
	}

	if blockIp != "" {
		if blockIp != clientIp {
			return nil
		}

		ipIsLimited, err := r.VerifyLimit(ctx, blockIp, blockIpLimit, time.Second*time.Duration(blockIpTime))

		if err != nil {
			return err
		}

		if ipIsLimited {
			return ErrFoo
		}
	}

	return nil
}
