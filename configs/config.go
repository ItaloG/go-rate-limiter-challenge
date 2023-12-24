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
}

var (
	ErrBlockIpRequired        = errors.New("você deve passar um IP para realizar o bloqueio")
	ErrBlockIpTimeRequired    = errors.New("você deve passar um tempo limite de bloqueio para o IP")
	ErrBlockTokenRequired     = errors.New("você deve passar um TOKEN para realizar o bloqueio")
	ErrBlockTokenTimeRequired = errors.New("você deve passar um tempo limite de bloqueio para o TOKEN")
	ErrRateLimitEnvsRequired  = errors.New("você deve passar as variaveis de RATE LIMIT")
)

func LoadConfig(path string) (*conf, error) {

	blockIp := os.Getenv("BLOCK_IP")
	blockIpTime := os.Getenv("BLOCK_IP_TIME")
	blockToken := os.Getenv("BLOCK_TOKEN")
	blockTokenTime := os.Getenv("BLOCK_TOKEN_TIME")

	envConfig, err := validateEnvs(blockIp, blockIpTime, blockToken, blockTokenTime)
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

	viperConfig, err = validateEnvs(viperConfig.BlockIp, viperConfig.BlockIpTime, viperConfig.BlockToken, viperConfig.BlockTokenTime)

	return viperConfig, err
}

func validateEnvs(blockIp string, blockIpTime string, blockToken string, blockTokenTime string) (*conf, error) {
	if blockIp != "" && blockIpTime == "" {
		return nil, ErrBlockIpTimeRequired
	}

	if blockIp == "" && blockIpTime != "" {
		return nil, ErrBlockIpRequired
	}

	if blockIp != "" && blockIpTime != "" {
		return &conf{BlockIp: blockIp, BlockIpTime: blockIpTime}, nil
	}

	if blockToken != "" && blockTokenTime == "" {
		return nil, ErrBlockTokenTimeRequired
	}

	if blockToken == "" && blockTokenTime != "" {
		return nil, ErrBlockTokenRequired
	}

	if blockToken != "" && blockTokenTime != "" {
		return &conf{BlockToken: blockToken, BlockTokenTime: blockTokenTime}, nil
	}

	return nil, ErrRateLimitEnvsRequired
}
