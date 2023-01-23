package models

type ClientPrimeryKey struct {
	Id string `json:"id"`
}

type CreateClient struct {
	Id          	string      `json:"id"`
	First_name  	string      `json:"first_name"`
	Last_name   	string      `json:"last_name"`
	Address     	string      `json:"address"`
	Phone_number    string      `json:"phone_number"`
	CreatedAt   	string      `json:"created_at"`
	UpdatedAt   	string      `json:"updated_at"`
}
type Client struct {
	Id          	string      `json:"id"`
	First_name  	string      `json:"first_name"`
	Last_name   	string      `json:"last_name"`
	Address     	string      `json:"address"`
	Phone_number    string      `json:"phone_number"`
	CreatedAt   	string      `json:"created_at"`
	UpdatedAt   	string      `json:"updated_at"`
}

type UpdateClient struct {
	Id          	string      `json:"id"`
	First_name  	string      `json:"first_name"`
	Last_name   	string      `json:"last_name"`
	Address     	string      `json:"address"`
	Phone_number    string      `json:"phone_number"`
	CreatedAt   	string      `json:"created_at"`
	UpdatedAt   	string      `json:"updated_at"`
}

type UpdateClientSwag struct {
	First_name  	string      `json:"first_name"`
	Last_name   	string      `json:"last_name"`
	Address     	string      `json:"address"`
	Phone_number    string      `json:"phone_number"`
}

type GetListClientRequest struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
	Search string `json:"search"`
}

type GetListClientResponse struct {
	Count int64   `json:"count"`
	Client []*CreateClient `json:"client"`
}

