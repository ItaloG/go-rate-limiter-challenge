package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyRedisHost_ShouldReturnAnRedisHostIsRequiredErr(t *testing.T) {
	_, err := ValidateEnvs("blockIp", "blockIpLimit", "blockIpTime", "blockToken", "blockTokenLimit", "blockTokenTime", "")
	assert.EqualError(t, err, "você deve passar o host do REDIS")
}

func TestGivenAnBlockIp_AndGivenAnEmptyBlockIpTime_ShouldReturnAnBlockIpTimeRequiredErr(t *testing.T) {
	_, err := ValidateEnvs("blockIp", "blockIpLimit", "", "", "", "", "redisHost")
	assert.EqualError(t, err, "você deve passar um tempo limite de bloqueio para o IP")
}

func TestGivenAnBlockIp_AndGivenAnEmptyBlockIpLimit_ShouldReturnAnBlockIpLimitRequiredErr(t *testing.T) {
	_, err := ValidateEnvs("blockIp", "", "blockIpTime", "", "", "", "redisHost")
	assert.EqualError(t, err, "você deve passar um LIMITE DE TENTATIVAS para realizar o bloqueio por IP")
}

func TestGivenAnEmptyBlockIp_AndGivenAnBlockIpTime_ShouldReturnAnBlockIpRequiredErr(t *testing.T) {
	_, err := ValidateEnvs("", "blockIpLimit", "blockIpTime", "", "", "", "redisHost")
	assert.EqualError(t, err, "você deve passar um IP para realizar o bloqueio")
}

func TestGivenAnBlockToken_AndGivenAnEmptyBlockTokenTime_ShouldReturnAnBlockTokenTimeRequiredErr(t *testing.T) {
	_, err := ValidateEnvs("", "", "", "blockToken", "blockTokenLimit", "", "redisHost")
	assert.EqualError(t, err, "você deve passar um tempo limite de bloqueio para o TOKEN")
} //ok

func TestGivenAnBlockToken_AndGivenAnEmptyBlockTokenLimit_ShouldReturnAnBlockTokenLimitRequiredErr(t *testing.T) {
	_, err := ValidateEnvs("", "", "", "blockToken", "", "blockTokenTime", "redisHost")
	assert.EqualError(t, err, "você deve passar um LIMITE DE TENTATIVAS para realizar o bloqueio por TOKEN")
}

func TestGivenAnEmptyBlockToken_AndGivenAnBlockTokenTime_ShouldReturnAnBlockTokenRequiredErr(t *testing.T) {
	_, err := ValidateEnvs("", "", "", "", "blockTokenLimit", "blockTokenTime", "redisHost")
	assert.EqualError(t, err, "você deve passar um TOKEN para realizar o bloqueio")
}

func TestAnEmptyBlockTokenEnv_AndEmptyBlockIpEnv_ShouldReturnRateLimitEnvsRequiredErr(t *testing.T) {
	_, err := ValidateEnvs("", "", "", "", "", "", "redisHost")
	assert.EqualError(t, err, "você deve passar as variaveis de RATE LIMIT")
}

func TestGivenAnBlockIpEnvs_AndEmptyBlockTokenEnvs_ShouldReturnAConfigWithIpEnvs(t *testing.T) {
	config, err := ValidateEnvs("blockIp", "blockIpLimit", "blockIpTime", "", "", "", "redisHost")
	assert.Nil(t, err)
	assert.Equal(t, "redisHost", config.RedisHost)
	assert.Equal(t, "blockIp", config.BlockIp)
	assert.Equal(t, "blockIpLimit", config.BlockIpLimit)
	assert.Equal(t, "blockIpTime", config.BlockIpTime)
	assert.Equal(t, "", config.BlockToken)
	assert.Equal(t, "", config.BlockTokenLimit)
	assert.Equal(t, "", config.BlockTokenTime)
}

func TestGivenAnEmptyBlockIpEnvs_AndBlockTokenEnvs_ShouldReturnAConfigWithTokenEnvs(t *testing.T) {
	config, err := ValidateEnvs("", "", "", "blockToken", "blockTokenLimit", "blockTokenTime", "redisHost")
	assert.Nil(t, err)
	assert.Equal(t, "redisHost", config.RedisHost)
	assert.Equal(t, "", config.BlockIp)
	assert.Equal(t, "", config.BlockIpLimit)
	assert.Equal(t, "", config.BlockIpTime)
	assert.Equal(t, "blockToken", config.BlockToken)
	assert.Equal(t, "blockTokenLimit", config.BlockTokenLimit)
	assert.Equal(t, "blockTokenTime", config.BlockTokenTime)
}

func TestGivenAnBlockIpEnvs_AndBlockTokenEnvs_ShouldReturnAConfigWithBothEnvs(t *testing.T) {
	config, err := ValidateEnvs("blockIp", "blockIpLimit", "blockIpTime", "blockToken", "blockTokenLimit", "blockTokenTime", "redisHost")
	assert.Nil(t, err)
	assert.Equal(t, "redisHost", config.RedisHost)
	assert.Equal(t, "blockIp", config.BlockIp)
	assert.Equal(t, "blockIpLimit", config.BlockIpLimit)
	assert.Equal(t, "blockIpTime", config.BlockIpTime)
	assert.Equal(t, "blockToken", config.BlockToken)
	assert.Equal(t, "blockTokenLimit", config.BlockTokenLimit)
	assert.Equal(t, "blockTokenTime", config.BlockTokenTime)
}
