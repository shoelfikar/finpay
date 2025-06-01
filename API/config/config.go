package config

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"

	"github.com/spf13/viper"
)

type Config struct {
	BaseURL  string
	Port     string
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
		return nil, fmt.Errorf("failed read config file: %v", err)
	}

	cfg := &Config{
		BaseURL: viper.GetString("BASE_URL_PATH"),
		Port:    viper.GetString("PORT"),

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

	log.Info("success load configuration...")

	return cfg, nil
}
