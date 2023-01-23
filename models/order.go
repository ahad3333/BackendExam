package models

type OrderPrimeryKey struct {
	Id string `json:"id"`
}

type Payment struct {
	Id string `json:"id"`
}

type CreateOrder struct {
	Id          	string      `json:"id"`
	Car_id  		string      `json:"car_id"`
	Client_id   	string      `json:"client_id"`
	Day_count       float64      `json:"day_count"`
	UpdatedAt  		string      `json:"updated_at"`
}
type CreateOrderSwag struct {
	Car_id  		string      `json:"car_id"`
	Client_id   	string      `json:"client_id"`
	Day_count       float64      `json:"day_count"`
}

type Order struct {
	Id          	string      `json:"id"`
	Car_id  		string      `json:"car_id"`
	Client_id   	string      `json:"client_id"`
	Total_price     float64      `json:"total_price"`
	Paid_price      float64      `json:"paid_price"`
	Day_count       float64      `json:"day_count"`
	Give_km       	float64      `json:"give_km"`
	Receive_km      float64      `json:"receive_km"`
	Status       	string      `json:"status"`
	CreatedAt   	string      `json:"created_at"`
	UpdatedAt  		string      `json:"updated_at"`
}

type UpdateOrder struct {
	Id          	string      `json:"id"`
	Car_id  		string      `json:"car_id"`
	Client_id   	string      `json:"client_id"`
	Total_price     float64      `json:"total_price"`
	Paid_price      float64      `json:"paid_price"`
	Day_count       float64      `json:"day_count"`
	Give_km       	float64      `json:"give_km"`
	Receive_km      float64      `json:"receive_km"`
	UpdatedAt  		string      `json:"updated_at"`
}

type UpdateOrderSwag struct {
	Car_id  		string      `json:"car_id"`
	Client_id   	string      `json:"client_id"`
	Paid_price      float64      `json:"paid_price"`
}

type GetListOrderRequest struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
	Search string `json:"search"`
}

type GetListOrderResponse struct {
	Count int64   `json:"count"`
	Order []*Order `json:"order"`
}

type ReturnCar struct {
	Car_id  		string      `json:"car_id"`
	Client_id   	string      `json:"client_id"`
	Receive_km      float64      `json:"receive_km"`
}





