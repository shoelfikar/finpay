package config

import (
	"fmt"

	"github.com/shoelfikar/finpay/user-service/helper"
	"github.com/spf13/viper"
)


type Config struct {
	BaseURL  string
	Port     string
	GrpcPort string
	
	Postgres PostresSQLConfig
}

func Load() (*Config, error) {
	viper.New()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		helper.LoggingError(err.Error())
		return nil, fmt.Errorf("failed read config file: %v", err)
	}

	cfg := &Config{
		BaseURL: viper.GetString("BASE_URL_PATH"),
		Port:    viper.GetString("PORT"),
		GrpcPort: viper.GetString("GRPC_PORT"),

		Postgres: PostresSQLConfig{
			Host:          viper.GetString("DB_HOST"),
			Port:          viper.GetString("DB_PORT"),
			Username:      viper.GetString("DB_USERNAME"),
			Password:      viper.GetString("DB_PASSWORD"),
			DBName:        viper.GetString("DB_NAME"),
			DbMaxOpenConn: viper.GetInt("DB_MAX_OPEN_CONN"),
			DbMaxIdleConn: viper.GetInt("DB_MAX_IDLE_CONN"),
			DbMaxLifeTime: viper.GetDuration("DB_MAX_LIFE_TIME"),
		},
	}

	helper.LoggingInfo("success load configuration...")

	return cfg, nil
}
