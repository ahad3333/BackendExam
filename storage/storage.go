package storage

import (
	"app/models"
	"context"
)

type StorageI interface {
	CloseDB()
	Investor() InvestorRepoI
	Car() CarRepoI
	Client() ClientRepoI
	Order() OrderRepoI
	Report() ReportRepoI
	Branch() BranchRepoI
	User() UserRepoI
}

type InvestorRepoI interface {
	Insert(context.Context, *models.CreateInvestor) (string, error)
	GetByID(context.Context, *models.InvestorPrimeryKey) (*models.Investor, error)
	GetList(context.Context, *models.GetListInvestorRequest) (*models.GetListInvestorResponse, error)
	Update(context.Context, *models.UpdateInvestor) error
	Delete(context.Context, *models.InvestorPrimeryKey) error
}

type CarRepoI interface {
	Insert(context.Context, *models.CreateCar) (string, error)
	GetByID(context.Context, *models.CarPrimeryKey) (*models.Car, error)
	GetList(context.Context, *models.GetListCarRequest) (*models.GetListCarResponse, error)
	Update(context.Context, *models.UpdateCar) (int64, error)
	Delete(context.Context, *models.CarPrimeryKey) error
}

type ClientRepoI interface {
	Insert(context.Context, *models.CreateClient) (string, error)
	GetByID(context.Context, *models.ClientPrimeryKey) (*models.Client, error)
	GetList(context.Context, *models.GetListClientRequest) (*models.GetListClientResponse, error)
	Update(context.Context, *models.UpdateClient) (int64, error)
	Delete(context.Context, *models.ClientPrimeryKey) error
}

type OrderRepoI interface {
	Insert(context.Context, *models.CreateOrder) (string, error)
	GetByID(context.Context, *models.OrderPrimeryKey) (*models.Order, error)
	GetList(context.Context, *models.GetListOrderRequest) (*models.GetListOrderResponse, error)
	Update(context.Context, *models.UpdateOrder) error
	UpdatePatch(context.Context, *models.UpdatePatch) error
	Delete(context.Context, *models.OrderPrimeryKey) error
}

type ReportRepoI interface {
	GetListDebtors(context.Context) (*models.GetListDebtorResponse, error)
	GetListInvestorShare(context.Context) (*models.GetListInvestorShareResponse, error)
	GetListBranchShare(context.Context) (*models.GetListBranchShareResponse, error)
}

type BranchRepoI interface {
	Insert(context.Context, *models.CreateBranch) (string, error)
	GetByID(context.Context, *models.BranchPrimeryKey) (*models.Branch, error)
	GetList(context.Context, *models.GetListBranchRequest) (*models.GetListBranchResponse, error)
	Update(context.Context, *models.UpdateBranch) error
	Delete(context.Context, *models.BranchPrimeryKey) error
}

type UserRepoI interface {
	Create(ctx context.Context, req *models.CreateUser) (string, error)
	GetByPKey(ctx context.Context, req *models.UserPrimarKey) (*models.User, error)
	GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error)
	// Update(ctx context.Context, req *models.UpdateUser) (int64, error)
	// Delete(ctx context.Context, req *models.UserPrimarKey) error
}