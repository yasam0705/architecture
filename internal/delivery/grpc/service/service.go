package service

import (
	"context"
	gp "github/architecture/genproto/customer_service"
	"github/architecture/internal/usecase"
	"github/architecture/pkg/logger"
)

type service struct {
	log             logger.Logger
	customerUseCase usecase.CustomerService
	gp.UnsafeCustomerServiceServer
}

func NewRPC(log logger.Logger, customerUseCase usecase.CustomerService) *service {
	return &service{
		log:             log,
		customerUseCase: customerUseCase,
	}
}

func (s *service) Get(ctx context.Context, req *gp.GetRequest) (*gp.GetResponse, error) {
	return &gp.GetResponse{}, nil
}
