package handler

import (
	"add/storage"
)

type Handler struct {
	storage storage.StorageI
}

func NewHandler(storage storage.StorageI) *Handler {
	return &Handler{
		storage: storage,
	}
}