package main

import (
	"context"
	"exam/api"
	"exam/config"
	"exam/pkg/logger"
	"exam/service"
	"exam/storage/postgres"
	"fmt"
)

func main() {
	cfg := config.Load()

	store, err := postgres.New(context.Background(), cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}

	defer store.CloseDB()

	log := logger.New(cfg.ServiceName)

	service := service.New(store, log)

	c := api.New(service, log)

	fmt.Println("programm is running on localhost:8008...")
	c.Run(":8080")
}
