package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"app/api"
	"app/config"
	"app/storage/postgres"
	"app/storage/redis"
)

func main() {

	cfg := config.Load()

	storage, err := postgres.NewPostgres(context.Background(), cfg)
	fmt.Println(cfg)
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	defer storage.CloseDB()

	cache, err := redis.NewRedisCacheStorage(cfg)
	if err != nil {
		log.Fatal("failed to redis connect database:", err)
	}
	defer cache.CloseDB()

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	api.NewApi(&cfg, r, storage, cache)

	err = r.Run(cfg.HTTPPort)
	if err != nil {
		panic(err)
	}
}
