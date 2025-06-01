package main

import (
	"github.com/gofiber/fiber/v2/log"

	"github.com/shoelfikar/API/config"
	"github.com/shoelfikar/API/router"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("SERVER ERROR] - %v", err)
		return
	}

	db, err := config.InitPostresSQL(cfg.Postgres)
	if err != nil {
		return
	}
	defer db.Close()

	router := initDependencies()
	router.Setup(cfg.BaseURL)
	router.Run(cfg.Port)
}

func initDependencies() *router.Router {
	return &router.Router{}
}