package service

import (
	"context"
	gp "github/architecture/genproto/file_processing"
	"github/architecture/internal/usecase"

	"go.uber.org/zap"
)

type service struct {
	log             *zap.Logger
	customerUseCase usecase.CustomerService
	gp.UnimplementedFileProcessingServiceServer
}

func NewRPC(log *zap.Logger, customerUseCase usecase.CustomerService) *service {
	return &service{
		log:             log,
		customerUseCase: customerUseCase,
	}
}

func (s *service) Create(ctx context.Context, req *gp.CreateRequest) (*gp.CreateResponse, error) {
	return &gp.CreateResponse{}, nil
}

func (s *service) List(ctx context.Context, req *gp.ListRequest) (*gp.ListResponse, error) {
	return &gp.ListResponse{}, nil
}

func (s *service) Get(ctx context.Context, req *gp.GetRequest) (*gp.GetResponse, error) {
	return &gp.GetResponse{}, nil
}

func (s *service) Update(ctx context.Context, req *gp.UpdateRequest) (*gp.UpdateResponse, error) {
	return &gp.UpdateResponse{}, nil
}
