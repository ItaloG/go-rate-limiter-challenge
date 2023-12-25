package handlers

import (
	"context"
	"errors"
	"time"
)

type RateLimitInterface interface {
	AddKey(ctx context.Context, key string, duration time.Duration) error
	GetKey(ctx context.Context, key string) string
	VerifyKey(ctx context.Context, key string, limit int) (bool, error)
}

func Handle(r RateLimitInterface, key string, duration time.Duration, limit int) error {
	ctx := context.Background()
	defer ctx.Done()
	result := r.GetKey(ctx, key)
	if result == "" {
		r.AddKey(ctx, key, duration)
	}

	isLimited, err := r.VerifyKey(ctx, key, limit)

	if err != nil {
		return err
	}

	if isLimited {
		return errors.New("você esta limitado")
	}

	return nil
}
