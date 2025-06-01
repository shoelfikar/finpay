package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/shoelfikar/finpay/user-service/helper"
)

type PostresSQLConfig struct {
	Host          string
	Port          string
	DBName        string
	Username      string
	Password      string
	DbMaxOpenConn int
	DbMaxIdleConn int
	DbMaxLifeTime time.Duration
}

func InitPostresSQL(cfg PostresSQLConfig) (*sql.DB, error) {
	conn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
	cfg.Username,
	cfg.Password,
	cfg.DBName,
	cfg.Host,
	cfg.Port)

	if cfg.Password == "" {
		conn = fmt.Sprintf("user=%s dbname=%s host=%s port=%s sslmode=disable",
			cfg.Username,
			cfg.DBName,
			cfg.Host,
			cfg.Port)
	}

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(fmt.Sprintf("[SERVER ERROR] Failed to connect postgres database %v", err))
	}

	db.SetMaxOpenConns(cfg.DbMaxOpenConn)
	db.SetMaxIdleConns(cfg.DbMaxIdleConn)
	db.SetConnMaxLifetime(time.Minute * cfg.DbMaxLifeTime)

	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("[SERVER ERROR] Failed to ping postgres database %v", err))
	}

	msg := fmt.Sprintf("Connected to PostgreSQL at %s:%s", cfg.Host, cfg.Port)
	helper.LoggingInfo(msg)
	return db, nil
}
