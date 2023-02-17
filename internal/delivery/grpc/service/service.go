package service

import (
	"context"
	gp "github/architecture/genproto/file_processing"
	"github/architecture/internal/entity"
	"github/architecture/internal/usecase"
	time_pkg "github/architecture/pkg/time"

	"go.uber.org/zap"
)

type service struct {
	log         *zap.Logger
	fileUseCase usecase.FileService
	gp.UnimplementedFileProcessingServiceServer
}

func NewRPC(log *zap.Logger, fileUseCase usecase.FileService) *service {
	return &service{
		log:         log,
		fileUseCase: fileUseCase,
	}
}

func (s *service) Create(ctx context.Context, req *gp.CreateRequest) (*gp.CreateResponse, error) {
	file := &entity.File{
		FileName: req.GetFile(),
		UserId:   req.GetUserId(),
	}

	err := s.fileUseCase.Create(ctx, file)
	if err != nil {
		s.log.Error("error on fileUseCase.Create", zap.Error(err))
		return nil, err
	}
	return &gp.CreateResponse{
		Guid: file.Guid,
	}, nil
}

func (s *service) List(ctx context.Context, req *gp.ListRequest) (*gp.ListResponse, error) {
	list, err := s.fileUseCase.List(ctx, req.GetLimit(), req.GetOffset(), req.GetFilter())
	if err != nil {
		s.log.Error("error on fileUseCase.List", zap.Error(err))
		return nil, err
	}

	files := make([]*gp.File, 0, len(list))
	for _, v := range list {
		files = append(files, &gp.File{
			Guid:      v.Guid,
			UserId:    v.UserId,
			FileName:  v.FileName,
			CreatedAt: time_pkg.DateToStringRFC3339(v.CreatedAt),
			UpdatedAt: time_pkg.DateToStringRFC3339(v.UpdatedAt),
		})
	}

	return &gp.ListResponse{
		List: files,
	}, nil
}

func (s *service) Get(ctx context.Context, req *gp.GetRequest) (*gp.GetResponse, error) {
	file, err := s.fileUseCase.Get(ctx, req.GetGuid())
	if err != nil {
		s.log.Error("error on fileUseCase.Get", zap.Error(err))
		return nil, err
	}

	return &gp.GetResponse{
		File: &gp.File{
			Guid:      file.Guid,
			UserId:    file.UserId,
			FileName:  file.FileName,
			CreatedAt: time_pkg.DateToStringRFC3339(file.CreatedAt),
			UpdatedAt: time_pkg.DateToStringRFC3339(file.UpdatedAt),
		},
	}, nil
}

func (s *service) Update(ctx context.Context, req *gp.UpdateRequest) (*gp.UpdateResponse, error) {
	file := &entity.File{
		Guid:     req.GetFile().GetGuid(),
		FileName: req.GetFile().GetFileName(),
	}

	err := s.fileUseCase.Update(ctx, file)
	if err != nil {
		s.log.Error("error on fileUseCase.Update", zap.Error(err))
		return nil, err
	}
	return &gp.UpdateResponse{
		UpdatedAt: time_pkg.DateToStringRFC3339(file.UpdatedAt),
	}, nil
}
