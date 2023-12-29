package handlers

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockRateLimit struct {
	IsLimited bool
}

func (m *MockRateLimit) VerifyLimit(ctx context.Context, key string, limit int64, duration time.Duration) (bool, error) {
	return m.IsLimited, nil
}

func TestGivenAEmptyBlockToken_AndEmptyBlockIp_ShouldReturnNill(t *testing.T) {
	rateLimitMock := &MockRateLimit{}
	blockIp := ""
	clientIp := ""
	blockIpLimit := int64(0)
	blockIpTime := 0
	blockToken := ""
	clientToken := ""
	blockTokenLimit := int64(10)
	blockTokenTime := 60

	err := Handle(rateLimitMock, blockIp, blockIpLimit, blockIpTime, clientIp, blockToken, blockTokenLimit, blockTokenTime, clientToken)
	assert.Nil(t, err)
}

func TestGivenABlockToken_AndClientTokenIsDiffOfBlockToken_ShouldReturnNill(t *testing.T) {
	rateLimitMock := &MockRateLimit{}
	blockIp := ""
	clientIp := ""
	blockIpLimit := int64(0)
	blockIpTime := 0
	blockToken := "block_token"
	clientToken := "client_token"
	blockTokenLimit := int64(10)
	blockTokenTime := 60

	err := Handle(rateLimitMock, blockIp, blockIpLimit, blockIpTime, clientIp, blockToken, blockTokenLimit, blockTokenTime, clientToken)
	assert.Nil(t, err)
}

func TestGivenABlockIp_AndClientIpIsDiffOfBlockIp_ShouldReturnNil(t *testing.T) {
	rateLimitMock := &MockRateLimit{}
	blockIp := "block_ip"
	clientIp := "client_ip"
	blockIpLimit := int64(10)
	blockIpTime := 60
	blockToken := ""
	clientToken := ""
	blockTokenLimit := int64(0)
	blockTokenTime := 0

	err := Handle(rateLimitMock, blockIp, blockIpLimit, blockIpTime, clientIp, blockToken, blockTokenLimit, blockTokenTime, clientToken)
	assert.Nil(t, err)
}

func TestGivenABlockToken_AndClientTokenIsEqualBlockToken_AndTokenIsLimited_ShouldReturnErr(t *testing.T) {
	rateLimitMock := &MockRateLimit{IsLimited: true}
	blockIp := ""
	clientIp := ""
	blockIpLimit := int64(0)
	blockIpTime := 0
	blockToken := "block_token"
	clientToken := "block_token"
	blockTokenLimit := int64(10)
	blockTokenTime := 60

	err := Handle(rateLimitMock, blockIp, blockIpLimit, blockIpTime, clientIp, blockToken, blockTokenLimit, blockTokenTime, clientToken)
	assert.EqualError(t, err, ErrFoo.Error())
}

func TestGivenABlockIp_AndClientIpIsEqualBlockIp_ButIpIsLimited_ShouldReturnErr(t *testing.T) {
	rateLimitMock := &MockRateLimit{IsLimited: true}
	blockIp := "block_ip"
	clientIp := "block_ip"
	blockIpLimit := int64(10)
	blockIpTime := 60
	blockToken := ""
	clientToken := ""
	blockTokenLimit := int64(0)
	blockTokenTime := 0

	err := Handle(rateLimitMock, blockIp, blockIpLimit, blockIpTime, clientIp, blockToken, blockTokenLimit, blockTokenTime, clientToken)
	assert.EqualError(t, err, ErrFoo.Error())
}

func TestGivenBlockIpLimitLessThanBlockTokenLimit_AndTokenLimitIsReached_ShouldReturnErr(t *testing.T) {
	rateLimitMock := &MockRateLimit{IsLimited: true}
	blockIp := "block_ip"
	clientIp := "block_ip"
	blockIpLimit := int64(5)
	blockIpTime := 60
	blockToken := "block_token"
	clientToken := "block_token"
	blockTokenLimit := int64(10)
	blockTokenTime := 60

	err := Handle(rateLimitMock, blockIp, blockIpLimit, blockIpTime, clientIp, blockToken, blockTokenLimit, blockTokenTime, clientToken)
	assert.EqualError(t, err, ErrFoo.Error())
}

func TestGivenBlockIpLimitMoreThanBlockTokenLimit_AndTokenLimitIsReached_ShouldReturnErr(t *testing.T) {
	rateLimitMock := &MockRateLimit{IsLimited: true}
	blockIp := "block_ip"
	clientIp := "block_ip"
	blockIpLimit := int64(10)
	blockIpTime := 60
	blockToken := "block_token"
	clientToken := "block_token"
	blockTokenLimit := int64(5)
	blockTokenTime := 60

	err := Handle(rateLimitMock, blockIp, blockIpLimit, blockIpTime, clientIp, blockToken, blockTokenLimit, blockTokenTime, clientToken)
	assert.EqualError(t, err, ErrFoo.Error())
}

func TestGivenBlockIp_AndEmptyBlockToken_AndBlockIpLimitIsReached_ShouldReturnErr(t *testing.T) {
	rateLimitMock := &MockRateLimit{IsLimited: true}
	blockIp := "block_ip"
	clientIp := "block_ip"
	blockIpLimit := int64(5)
	blockIpTime := 60
	blockToken := ""
	clientToken := ""
	blockTokenLimit := int64(0)
	blockTokenTime := 0

	err := Handle(rateLimitMock, blockIp, blockIpLimit, blockIpTime, clientIp, blockToken, blockTokenLimit, blockTokenTime, clientToken)
	assert.EqualError(t, err, ErrFoo.Error())
}
