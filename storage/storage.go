package storage

import (
	"context"
	"exam/api/models"
)

type IStorage interface {
	CloseDB()
	CustomersStorage() CustomersStorage
}

type CustomersStorage interface {
	Create(ctx context.Context, customers models.Customers) (string, error)
	GetAll(ctx context.Context, req models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error)
	Update(ctx context.Context, id string,customers models.Customers) error
	GetById(ctx context.Context, id string) (models.GetCustomers, error)
	Delete(ctx context.Context, id string) error
}