package handler

import (
	"app/config"
	"app/storage"
)

type Handler struct {
	cfg     *config.Config
	storage storage.StorageI
	cache   storage.CacheStorageI
}

func NewHandler(cfg *config.Config, storage storage.StorageI, cache storage.CacheStorageI) *Handler {
	return &Handler{
		cfg:     cfg,
		storage: storage,
		cache:   cache,
	}
}
