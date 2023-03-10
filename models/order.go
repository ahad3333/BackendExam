package models

type OrderPrimeryKey struct {
	Id 		  string `json:"id"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	TypeU     string `json:"typeU"`
}

type CreateOrder struct {
	CarId    string `json:"car_id"`
	ClientId string `json:"client_id"`
	DayCount int    `json:"day_count"`
	BranchId  string `json:"branch_id"`
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
	BranchId  string `json:"branch_id"`
	ClientFullName string   `json:"client_full_name"`
	TotalPrice     float64  `json:"total_price"`
	PaidPrice      float64  `json:"paid_price"`
	DayCount       float64      `json:"day_count"`
	GiveKm         float64      `json:"give_km"`
	ReceiveKm      float64      `json:"recieve_km"`
	Status         string   `json:"status"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
}

type UpdateOrder struct {
	Id         string  `json:"id"`
	CarId      string  `json:"car_id"`
	ClientId   string  `json:"client_id"`
	BranchId  string `json:"branch_id"`
	TotalPrice float64 `json:"total_price"`
	PaidPrice  float64 `json:"paid_price"`
	DayCount   int     `json:"day_count"`
	GiveKm     int     `json:"give_km"`
	ReceiveKm  int     `json:"recieve_km"`
	Status     string  `json:"status"`
}

type UpdateOrderSwag struct {
	Car_id  		string      `json:"car_id"`
	Client_id   	string      `json:"client_id"`
	BranchId  string `json:"branch_id"`
	Paid_price      float64      `json:"paid_price"`
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

type GetListOrderInvestorResponse struct {
	Count  int64    `json:"count"`
	Orders []*OrderInvestor `json:"orders"`
}

type OrderInvestor struct {
	Id             string   `json:"id"`
	CarId          string   `json:"car_id"`
	Car            CarOrder `json:"car"`
	ClientId       string   `json:"client_id"`
	ClientFullName string   `json:"client_full_name"`
	BranchId  string `json:"branch_id"`
	TotalPrice     float64  `json:"total_price"`
	DayCount       float64  `json:"day_count"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
}