package models

type Debtor struct {
	Price    float64 `json:"price"`
	FullName string  `json:"full_name"`
}

type GetListDebtorResponse struct {
	Count   int32     `json:"count"`
	Debtors []*Debtor `json:"debtors"`
}

type InvestorShare struct {
	Id       string  `json:"id"`
	FullName string  `json:"full_name"`
	Price    float64 `json:"price"`
}

type GetListInvestorShareResponse struct {
	Count     int32            `json:"count"`
	Investors []*InvestorShare `json:"investors"`
}

type BranchShare struct {
	Id      string  `json:"id"`
	Name 	string  `json:"name"`
	Price   float64 `json:"price"`
}

type GetListBranchShareResponse struct {
	Count     int32           `json:"count"`
	Branchs []*BranchShare    `json:"branchs"`
}