package configs

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type conf struct {
	BlockIp        string `mapstructure:"BLOCK_IP"`
	BlockIpTime    string `mapstructure:"BLOCK_IP_TIME"`
	BlockToken     string `mapstructure:"BLOCK_TOKEN"`
	BlockTokenTime string `mapstructure:"BLOCK_TOKEN_TIME"`
	RedisHost      string `mapstructure:"REDIS_HOST"`
}

var (
	ErrBlockIpRequired        = errors.New("você deve passar um IP para realizar o bloqueio")
	ErrBlockIpTimeRequired    = errors.New("você deve passar um tempo limite de bloqueio para o IP")
	ErrBlockTokenRequired     = errors.New("você deve passar um TOKEN para realizar o bloqueio")
	ErrBlockTokenTimeRequired = errors.New("você deve passar um tempo limite de bloqueio para o TOKEN")
	ErrRateLimitEnvsRequired  = errors.New("você deve passar as variaveis de RATE LIMIT")
	ErrRedisHostRequired      = errors.New("você deve passar o host do REDIS")
)

func LoadConfig(path string) (*conf, error) {

	blockIp := os.Getenv("BLOCK_IP")
	blockIpTime := os.Getenv("BLOCK_IP_TIME")
	blockToken := os.Getenv("BLOCK_TOKEN")
	blockTokenTime := os.Getenv("BLOCK_TOKEN_TIME")
	redisHost := os.Getenv("REDIS_HOST")

	envConfig, err := validateEnvs(blockIp, blockIpTime, blockToken, blockTokenTime, redisHost)
	if envConfig != nil {
		return envConfig, err
	}

	var viperConfig *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&viperConfig)
	if err != nil {
		panic(err)
	}

	viperConfig, err = validateEnvs(viperConfig.BlockIp, viperConfig.BlockIpTime, viperConfig.BlockToken, viperConfig.BlockTokenTime, viperConfig.RedisHost)

	if err != nil {
		return nil, err
	}

	os.Setenv("REDIS_HOST", viperConfig.RedisHost)

	return viperConfig, err
}

func validateEnvs(blockIp string, blockIpTime string, blockToken string, blockTokenTime string, redisHost string) (*conf, error) {
	if redisHost == "" {
		return nil, ErrRedisHostRequired
	}

	if blockIp != "" && blockIpTime == "" {
		return nil, ErrBlockIpTimeRequired
	}

	if blockIp == "" && blockIpTime != "" {
		return nil, ErrBlockIpRequired
	}

	if blockToken != "" && blockTokenTime == "" {
		return nil, ErrBlockTokenTimeRequired
	}

	if blockToken == "" && blockTokenTime != "" {
		return nil, ErrBlockTokenRequired
	}

	if blockIp != "" && blockIpTime != "" && blockToken == "" && blockTokenTime == "" {
		return &conf{BlockIp: blockIp, BlockIpTime: blockIpTime, RedisHost: redisHost}, nil
	}

	if blockToken != "" && blockTokenTime != "" && blockIp == "" && blockIpTime == "" {
		return &conf{BlockToken: blockToken, BlockTokenTime: blockTokenTime, RedisHost: redisHost}, nil
	}

	if blockToken == "" && blockTokenTime == "" && blockIp == "" && blockIpTime == "" {
		return nil, ErrRateLimitEnvsRequired
	}

	return &conf{BlockToken: blockToken, BlockTokenTime: blockTokenTime, BlockIp: blockIp, BlockIpTime: blockIpTime, RedisHost: redisHost}, nil
}
