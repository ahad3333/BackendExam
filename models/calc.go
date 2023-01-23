package models

type InvestorBenefit struct {
	Name  		string      `json:"name"`
	Price       float64     `json:"price"`
	CreatedAt   string      `json:"created_at"`
	UpdatedAt  	string      `json:"updated_at"`
}

type GetListInves struct {
	Count int64   `json:"count"`
	Benefits []*InvestorBenefit `json:"order"`
}

type GetListInvesRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
}

type Bebtors struct {
	Name  			string       `json:"name"`
	Bebt     		float64      `json:"bebt"`
	CreatedAt   	string       `json:"created_at"`
	UpdatedAt  		string       `json:"updated_at"`
}

type GetListBebt struct {
	Count int64   `json:"count"`
	Bebtors []*Bebtors `json:"order"`
}

type GetListBebtRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
}