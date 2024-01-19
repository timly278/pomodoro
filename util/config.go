package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	HttpServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GrpcServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	AppEmail             string        `mapstructure:"APP_EMAIL"`
	AppPassword          string        `mapstructure:"APP_PASSWORD"`
	AppSmtpHost          string        `mapstructure:"APP_SMTP_HOST"`
	AppSmtpPort          int           `mapstructure:"APP_SMTP_PORT"`
	RedisClientAddress   string        `mapstructure:"REDIS_CLIENT_ADDRESS"`
	RedisDbPassword      string        `mapstructure:"REDIS_DB_PASSWORD"`
	RedisDb              int           `mapstructure:"REDIS_DB"`
	LogFilesPath         string        `mapstructure:"LOG_FILES_PATH"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") // you can use json/xml here if you want so as long it has correct format

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil

}
