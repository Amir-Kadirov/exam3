package service

import (
	"context"
	"exam/api/models"
	"exam/pkg/logger"
	"exam/storage"
	"fmt"
)

type CustomersService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewCustomersService(storage storage.IStorage, logger logger.ILogger) CustomersService {
	return CustomersService{
		storage: storage,
		logger:  logger,
	}
}

func (s CustomersService) Create(ctx context.Context, Customers models.Customers) (string, error) {
	id, err := s.storage.CustomersStorage().Create(ctx, Customers)
	if err != nil {
		fmt.Println("error while creating Customers, err: ", err)
		return "", err
	}

	return id, nil
}

func (s CustomersService) GetAllCustomers(ctx context.Context, req models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error) {
	Customers := models.GetAllCustomersResponse{}

	Customers, err := s.storage.CustomersStorage().GetAll(ctx, req)
	if err != nil {
		fmt.Println("error while get all Customerss err: ", err)
		return Customers, err
	}

	return Customers, nil
}

func (s CustomersService) UpdateCustomers(ctx context.Context,id string ,Customers models.Customers) error {
	err := s.storage.CustomersStorage().Update(ctx, id,Customers)
	if err != nil {
		return err
	}

	return nil
}

func (s CustomersService) GetByIdCustomers(ctx context.Context, id string) (models.GetCustomers, error) {
	customers := models.GetCustomers{}
	customers, err := s.storage.CustomersStorage().GetById(ctx, id)
	if err != nil {
		fmt.Println("error while get by id customers err: ", err)
		return customers, err
	}

	return customers, err
}

func (s CustomersService) DeleteCustomers(ctx context.Context, id string) error {
	err := s.storage.CustomersStorage().Delete(ctx, id)
	if err != nil {
		fmt.Println("error while deleting Customers", err)
		return err
	}

	return nil
}