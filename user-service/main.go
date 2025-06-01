package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/shoelfikar/finpay/user-service/config"
	"github.com/shoelfikar/finpay/user-service/helper"
	"github.com/shoelfikar/finpay/user-service/router"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		helper.LoggingError(err.Error())
		return
	}

	db, err := config.InitPostresSQL(cfg.Postgres)
	if err != nil {
		helper.LoggingError(err.Error())
		return
	}
	defer db.Close()

	router := initDepedencies()

	var wg sync.WaitGroup
	wg.Add(1)
	go router.RunGRPC(cfg.GrpcPort, &wg)

	wg.Wait()

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM)

	<-stopSignal
	helper.LoggingInfo("Received shutdown signal, shutting down servers...")

	router.GrpcServer.GracefulStop()
	helper.LoggingInfo("gRPC server stopped...")

	if err := router.Listener.Close(); err != nil {
		msg := fmt.Sprintf("Error closing listener: %v", err)
		helper.LoggingError(msg)
	}

	helper.LoggingInfo("HTTP server stopped")

	if err := db.Close(); err != nil {
		msg := fmt.Sprintf("Error closing database connection: %v", err)
		helper.LoggingError(msg)
	}
}

func initDepedencies() *router.Routes {
	// add handlers here
	return &router.Routes{}
}