package handler

import (
	"app/config"
	"app/storage"
)

type Handler struct {
	storage storage.StorageI
	cfg     *config.Config

}

func NewHandler(cfg *config.Config, storage storage.StorageI) *Handler {
	return &Handler{
		cfg:     cfg,
		storage: storage,
	}
}
