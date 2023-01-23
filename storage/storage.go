package storage

import (
	"context"
	"add/models"
)

type StorageI interface {
	CloseDB()
	Car() 	   CarRepoI
	Order()    OrderRepoI
	Client()   ClientRepoI
	Investor() InvestorRepoI
	Calc() CalcRepoI
}

type  ClientRepoI interface {
	Insert(context.Context, *models.CreateClient) (string, error)
	GetByID(context.Context, *models.ClientPrimeryKey) (*models.Client, error)
	GetList(ctx context.Context, req *models.GetListClientRequest) (*models.GetListClientResponse, error)
	Update(ctx context.Context, client *models.UpdateClient) error
	Delete(ctx context.Context, req *models.ClientPrimeryKey) error 
}

type InvestorRepoI interface {
	Insert(context.Context, *models.CreateInvestor) (string, error)
	GetByID(context.Context, *models.InvestorPrimeryKey) (*models.Investor, error)
	GetList(ctx context.Context, req *models.GetListInvestorRequest) (*models.GetListInvestorResponse, error)
	Update(ctx context.Context, category *models.UpdateInvestor) error
	Delete(ctx context.Context, req *models.InvestorPrimeryKey) error
}

type  CarRepoI interface {
	Insert(context.Context, *models.CreateCar) (string, error)
	GetByID(context.Context, *models.CarPrimeryKey) (*models.Car, error)
	GetList(ctx context.Context, req *models.GetListCarRequest) (*models.GetListCarResponse, error)
	Update(ctx context.Context, car *models.UpdateCar) error
	Delete(ctx context.Context, req *models.CarPrimeryKey) error 
}

type  OrderRepoI interface {
	Insert(context.Context, *models.CreateOrder) (string, error)
	GetByID(context.Context, *models.OrderPrimeryKey) (*models.Order, error)
	GetList(ctx context.Context, req *models.GetListOrderRequest) (*models.GetListOrderResponse, error)
	Update(ctx context.Context, car *models.Order) error
	Return(ctx context.Context, car *models.Order,)error
	Delete(ctx context.Context, req *models.OrderPrimeryKey) error 
}

type CalcRepoI interface{
	GetListInves(ctx context.Context, req *models.GetListInvesRequest) (*models.GetListInves, error)
	GetListBebt(ctx context.Context, req *models.GetListBebtRequest) (*models.GetListBebt, error)
}