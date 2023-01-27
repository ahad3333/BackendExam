package models

type BranchPrimeryKey struct {
	Id string `json:"id"`
}

type CreateBranch struct {
	Name string `json:"name"`
}

type Branch struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateBranch struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UpdateBranchSwag struct {
	Name string `json:"name"`
}

type GetListBranchRequest struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type GetListBranchResponse struct {
	Count     int64     `json:"count"`
	Branchs []*Branch 	`json:"branch"`
}
