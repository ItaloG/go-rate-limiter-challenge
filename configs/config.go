package configs

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type conf struct {
	BlockIp         string `mapstructure:"BLOCK_IP"`
	BlockIpLimit    string `mapstructure:"BLOCK_IP_LIMIT"`
	BlockIpTime     string `mapstructure:"BLOCK_IP_TIME"`
	BlockToken      string `mapstructure:"BLOCK_TOKEN"`
	BlockTokenLimit string `mapstructure:"BLOCK_TOKEN_LIMIT"`
	BlockTokenTime  string `mapstructure:"BLOCK_TOKEN_TIME"`
	RedisHost       string `mapstructure:"REDIS_HOST"`
}

var (
	ErrBlockIpRequired         = errors.New("você deve passar um IP para realizar o bloqueio")
	ErrBlockIpLimitRequired    = errors.New("você deve passar um LIMITE DE TENTATIVAS para realizar o bloqueio por IP")
	ErrBlockIpTimeRequired     = errors.New("você deve passar um tempo limite de bloqueio para o IP")
	ErrBlockTokenRequired      = errors.New("você deve passar um TOKEN para realizar o bloqueio")
	ErrBlockTokenLimitRequired = errors.New("você deve passar um LIMITE DE TENTATIVAS para realizar o bloqueio por TOKEN")
	ErrBlockTokenTimeRequired  = errors.New("você deve passar um tempo limite de bloqueio para o TOKEN")
	ErrRateLimitEnvsRequired   = errors.New("você deve passar as variaveis de RATE LIMIT")
	ErrRedisHostRequired       = errors.New("você deve passar o host do REDIS")
)

func LoadConfig(path string) (*conf, error) {

	_, err := os.ReadFile(".env")
	if err != nil {
		blockIp := os.Getenv("BLOCK_IP")
		blockIpLimit := os.Getenv("BLOCK_IP_LIMIT")
		blockIpTime := os.Getenv("BLOCK_IP_TIME")
		blockToken := os.Getenv("BLOCK_TOKEN")
		blockTokenLimit := os.Getenv("BLOCK_TOKEN_LIMIT")
		blockTokenTime := os.Getenv("BLOCK_TOKEN_TIME")
		redisHost := os.Getenv("REDIS_HOST")

		envConfig, err := ValidateEnvs(blockIp, blockIpLimit, blockIpTime, blockToken, blockTokenLimit, blockTokenTime, redisHost)

		if err != nil {
			return nil, err
		}
		if envConfig != nil {
			return envConfig, nil
		}
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

	viperConfig, err = ValidateEnvs(viperConfig.BlockIp, viperConfig.BlockIpLimit, viperConfig.BlockIpTime, viperConfig.BlockToken, viperConfig.BlockTokenLimit, viperConfig.BlockTokenTime, viperConfig.RedisHost)

	if err != nil {
		return nil, err
	}

	os.Setenv("REDIS_HOST", viperConfig.RedisHost)
	os.Setenv("BLOCK_IP", viperConfig.BlockIp)
	os.Setenv("BLOCK_IP_LIMIT", viperConfig.BlockIpLimit)
	os.Setenv("BLOCK_IP_TIME", viperConfig.BlockIpTime)
	os.Setenv("BLOCK_TOKEN", viperConfig.BlockToken)
	os.Setenv("BLOCK_TOKEN_LIMIT", viperConfig.BlockTokenLimit)
	os.Setenv("BLOCK_TOKEN_TIME", viperConfig.BlockTokenTime)

	return viperConfig, err
}

func ValidateEnvs(blockIp string, blockIpLimit string, blockIpTime string, blockToken string, blockTokenLimit string, blockTokenTime string, redisHost string) (*conf, error) {
	if redisHost == "" {
		return nil, ErrRedisHostRequired
	}

	if blockIp != "" && blockIpTime == "" {
		return nil, ErrBlockIpTimeRequired
	}

	if blockIp != "" && blockIpLimit == "" {
		return nil, ErrBlockIpLimitRequired
	}

	if blockIp == "" && blockIpLimit != "" {
		return nil, ErrBlockIpRequired
	}

	if blockToken != "" && blockTokenTime == "" {
		return nil, ErrBlockTokenTimeRequired
	}

	if blockToken == "" && blockTokenLimit != "" {
		return nil, ErrBlockTokenRequired
	}

	if blockToken != "" && blockTokenLimit == "" {
		return nil, ErrBlockTokenLimitRequired
	}

	if blockToken == "" && blockTokenTime == "" && blockIp == "" && blockIpTime == "" {
		return nil, ErrRateLimitEnvsRequired
	}

	return &conf{BlockToken: blockToken, BlockTokenLimit: blockTokenLimit, BlockTokenTime: blockTokenTime, BlockIp: blockIp, BlockIpLimit: blockIpLimit, BlockIpTime: blockIpTime, RedisHost: redisHost}, nil
}
