package storage

import (
	"app/models"
)

type CacheStorageI interface {
	CloseDB()
	CarCache() CarCachaRepoI
}

type CarCachaRepoI interface {
	Create(*models.GetListCarResponse) error
	GetList() (*models.GetListCarResponse, error)
	Delete() error
	Exists() (bool, error)
}
