package redis

import (
	"github.com/go-redis/redis"

	"app/config"
	"app/storage"
)

type CacheStore struct {
	db  *redis.Client
	car *carRepo
}

func NewRedisCacheStorage(cfg config.Config) (storage.CacheStorageI, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &CacheStore{
		db:  client,
		car: NewCarRepo(client),
	}, nil
}

func (c *CacheStore) CloseDB() {
	c.db.Close()
}

func (c *CacheStore) CarCache() storage.CarCachaRepoI {
	if c.car == nil {
		c.car = NewCarRepo(c.db)
	}

	return c.car
}
