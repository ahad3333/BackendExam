package models

type CarPrimeryKey struct {
	Id string `json:"id"`
}

type CreateCar struct {
	State_number string  `json:"state_number"`
	Model 		 string  `json:"model"`
	Price 		 float64 `json:"price"`
	Daily_limit  float64 `json:"daily_limit"`
	Over_limit   float64 `json:"over_limit"`
	Investor_id  string  `json:"investor_id"`
	Km   		 float64 `json:"km"`
	UpdatedAt    string  `json:"updated_at"`
}
type Car struct {
	Id        			string 	`json:"id"`
	State_number 		string 	`json:"state_number"`
	Model 				string 	`json:"model"`
	Status 				string 	`json:"status"`
	Price 				float64 `json:"price"`
	Daily_limit 		float64 `json:"daily_limit"`
	Over_limit 			float64 `json:"over_limit"`
	Investor_percentage float64 `json:"investor_percentage"`
	Investor_id 		string  `json:"investor_id"`
	Km 					float64 `json:"km"`
	CreatedAt   		string  `json:"created_at"`
	UpdatedAt   		string  `json:"updated_at"`
}

type UpdateCar struct {
	Id 			 string  `json:"id"`
	State_number string  `json:"state_number"`
	Model 		 string  `json:"model"`
	Price 		 float64 `json:"price"`
	Daily_limit  float64 `json:"daily_limit"`
	Over_limit   float64 `json:"over_limit"`
	Investor_id  string  `json:"investor_id"`
	Km 			 float64 `json:"km"`
	UpdatedAt    string  `json:"updated_at"`
}

type UpdateCarSwag struct {
	State_number string  `json:"state_number"`
	Model 		 string  `json:"model"`
	Price 		 float64 `json:"price"`
	Daily_limit  float64 `json:"daily_limit"`
	Over_limit   float64 `json:"over_limit"`
	Investor_id  string  `json:"investor_id"`
	Km 			 float64 `json:"km"`
}

type GetListCarRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
}

type GetListCarResponse struct {
	Count int64 `json:"count"`
	Cars []*Car `json:"Cars"`
}
