package service

import (
	"exam/pkg/logger"
	"exam/storage"
)

type IServiceManager interface {
	Customers() CustomersService
}

type Service struct {
	CustomersService CustomersService

	logger logger.ILogger
}

func New(storage storage.IStorage, log logger.ILogger) Service {
	return Service{
		CustomersService: NewCustomersService(storage, log),

		logger: log,
	}
}

func (s Service) Customers() CustomersService {
	return s.CustomersService
}
