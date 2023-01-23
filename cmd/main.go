package main

import (
	"add/api"
	"context"
	"add/config"
	"add/storage/postgres"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	cfg := config.Load()

	storage, err := postgres.NewPostgres(context.Background(), cfg)
		if err != nil {
			log.Fatal("failed to connect database:", err)
		}
	defer storage.CloseDB()

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	api.NewApi(r, storage)

	err = r.Run(cfg.HTTPPort)
		if err != nil {
			panic(err)
		}
}
