package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyRedisHost_ShouldReturnAnRedisHostIsRequiredErr(t *testing.T) {
	_, err := ValidateEnvs("blockIp", "blockIpLimit", "blockIpTime", "blockToken", "blockTokenLimit", "blockTokenTime", "")
	assert.EqualError(t, err, "você deve passar o host do REDIS")
} // ok

func TestGivenAnBlockIp_AndGivenAnEmptyBlockIpTime_ShouldReturnAnBlockIpTimeRequiredErr(t *testing.T) {
	_, err := ValidateEnvs("blockIp", "blockIpLimit", "", "", "", "", "redisHost")
	assert.EqualError(t, err, "você deve passar um tempo limite de bloqueio para o IP")
} // ok

func TestGivenAnBlockIp_AndGivenAnEmptyBlockIpLimit_ShouldReturnAnBlockIpLimitRequiredErr(t *testing.T) {
	_, err := ValidateEnvs("blockIp", "", "blockIpTime", "", "", "", "redisHost")
	assert.EqualError(t, err, "você deve passar um LIMITE DE TENTATIVAS para realizar o bloqueio por IP")
} // ok

func TestGivenAnEmptyBlockIp_AndGivenAnBlockIpLimit_ShouldReturnAnBlockIpRequiredErr(t *testing.T) {
	_, err := ValidateEnvs("", "blockIpLimit", "blockIpTime", "", "", "", "redisHost")
	assert.EqualError(t, err, "você deve passar um IP para realizar o bloqueio")
} // ok

func TestGivenAnEmptyBlockIp_AndGivenAnBlockIpTime_ShouldReturnAnBlockIpRequiredErr(t *testing.T) {
	_, err := ValidateEnvs("", "blockIpLimit", "blockIpTime", "", "", "", "redisHost")
	assert.EqualError(t, err, "você deve passar um IP para realizar o bloqueio")
} // ok

func TestGivenAnBlockToken_AndGivenAnEmptyBlockTokenTime_ShouldReturnAnBlockTokenTimeRequiredErr(t *testing.T) {
	_, err := ValidateEnvs("", "", "", "blockToken", "blockTokenLimit", "", "redisHost")
	assert.EqualError(t, err, "você deve passar um tempo limite de bloqueio para o TOKEN")
} //ok

func TestGivenAnEmptyBlockToken_AndGivenAnBlockTokenLimit_ShouldReturnAnBlockTokenRequiredErr(t *testing.T) {
	_, err := ValidateEnvs("", "", "", "blockToken", "blockTokenLimit", "", "redisHost")
	assert.EqualError(t, err, "você deve passar um tempo limite de bloqueio para o TOKEN")
} // ok

func TestGivenAnBlockToken_AndGivenAnEmptyBlockTokenLimit_ShouldReturnAnBlockTokenLimitRequiredErr(t *testing.T) {
	_, err := ValidateEnvs("", "", "", "blockToken", "", "blockTokenTime", "redisHost")
	assert.EqualError(t, err, "você deve passar um LIMITE DE TENTATIVAS para realizar o bloqueio por TOKEN")
} // ok

func TestGivenAnEmptyBlockToken_AndGivenAnBlockTokenTime_ShouldReturnAnBlockTokenRequiredErr(t *testing.T) {
	_, err := ValidateEnvs("", "", "", "", "blockTokenLimit", "blockTokenTime", "redisHost")
	assert.EqualError(t, err, "você deve passar um TOKEN para realizar o bloqueio")
}
