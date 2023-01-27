package models

type OrderPrimeryKey struct {
	Id string `json:"id"`
}

type CreateOrder struct {
	CarId    string `json:"car_id"`
	ClientId string `json:"client_id"`
	DayCount int    `json:"day_count"`
}

type CarOrder struct {
	Id          string `json:"id"`
	StateNumber string `json:"state_number"`
	Model       string `json:"model"`
}

type Order struct {
	Id             string   `json:"id"`
	CarId          string   `json:"car_id"`
	Car            CarOrder `json:"car"`
	ClientId       string   `json:"client_id"`
	ClientFullName string   `json:"client_full_name"`
	TotalPrice     float64  `json:"total_price"`
	PaidPrice      float64  `json:"paid_price"`
	DayCount       int      `json:"day_count"`
	GiveKm         int      `json:"give_km"`
	ReceiveKm      int      `json:"recieve_km"`
	Status         string   `json:"status"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
}

type UpdateOrder struct {
	Id         string  `json:"id"`
	CarId      string  `json:"car_id"`
	ClientId   string  `json:"client_id"`
	TotalPrice float64 `json:"total_price"`
	PaidPrice  float64 `json:"paid_price"`
	DayCount   int     `json:"day_count"`
	GiveKm     int     `json:"give_km"`
	ReceiveKm  int     `json:"recieve_km"`
	Status     string  `json:"status"`
}

type UpdateOrderSwag struct {
	CarId      string  `json:"car_id"`
	ClientId   string  `json:"client_id"`
	TotalPrice float64 `json:"total_price"`
	PaidPrice  float64 `json:"paid_price"`
	DayCount   int     `json:"day_count"`
	GiveKm     int     `json:"give_km"`
	ReceiveKm  int     `json:"recieve_km"`
	Status     string  `json:"status"`
}

type UpdatePatch struct {
	Id   string                 `json:"id"`
	Data map[string]interface{} `json:"data"`
}

type GetListOrderRequest struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type GetListOrderResponse struct {
	Count  int64    `json:"count"`
	Orders []*Order `json:"orders"`
}
