package handlers

import (
	"context"
	"errors"
	"time"
)

type RateLimitInterface interface {
	VerifyLimit(ctx context.Context, key string, limit int64, duration time.Duration) (bool, error)
}

func Handle(r RateLimitInterface, key string, duration time.Duration, limit int64) error {
	ctx := context.Background()

	isLimited, err := r.VerifyLimit(ctx, key, limit, duration)

	if err != nil {
		return err
	}

	if isLimited {
		return errors.New("você esta limitado")
	}

	return nil
}
