package redis

import (
	"app/models"
	"encoding/json"

	"github.com/go-redis/redis"
)

type carRepo struct {
	db *redis.Client
}

func NewCarRepo(db *redis.Client) *carRepo {
	return &carRepo{
		db: db,
	}
}

func (r *carRepo) Create(req *models.GetListCarResponse) error {

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = r.db.Set("car_list", body, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *carRepo) GetList() (*models.GetListCarResponse, error) {

	resp := &models.GetListCarResponse{}

	body, err := r.db.Get("car_list").Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *carRepo) Delete() error {

	err := r.db.Del("car_list").Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *carRepo) Exists() (bool, error) {

	exist, err := r.db.Exists("car_list").Result()
	if err != nil {
		return false, err
	}

	if exist <= 0 {
		return false, nil
	}

	return true, nil
}
