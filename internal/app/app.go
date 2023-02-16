package app

import (
	"context"
	"fmt"
	"github/architecture/config"
	"github/architecture/genproto/file_processing"
	grpc_server "github/architecture/internal/delivery/grpc/server"
	grpc_service "github/architecture/internal/delivery/grpc/service"
	"github/architecture/internal/entity"
	"github/architecture/internal/usecase"
	"github/architecture/internal/usecase/repo"
	"github/architecture/pkg/logger"
	"github/architecture/pkg/postgres"
	"net"

	"log"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func Run(cfg *config.Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ContextTimeout)
	defer cancel()

	logger, err := logger.New(cfg.LogLevel, cfg.Environment, cfg.App+".log")
	if err != nil {
		return err
	}

	db, err := postgres.New(ctx, cfg)
	if err != nil {
		return err
	}

	customerRepo := repo.NewCustomerRepo(db)
	fileRepo := repo.NewFileRepo(db)

	customerUseCase := usecase.NewCustomer(cfg.ContextTimeout, customerRepo)
	fileUseCase := usecase.NewFile(cfg.ContextTimeout, fileRepo)

	ff := &entity.File{
		UserId:   "2adebc26-5538-43c6-81d4-1c2c283f86d0",
		FileName: "asdasd.name",
	}
	err = fileUseCase.Create(ctx, ff)
	if err != nil {
		fmt.Println("error create")
		return err
	}

	ff, err = fileUseCase.Get(ctx, ff.Guid)
	if err != nil {
		fmt.Println("error get")
		return err
	}
	fmt.Printf("=========== %+v\n", ff)

	ll, err := fileUseCase.List(ctx, map[string]string{"user_id": ff.UserId})
	if err != nil {
		fmt.Println("error list")
		return err
	}
	fmt.Printf("=========== %+v\n", ll[0])

	s, err := grpc_server.NewGRPCServer(cfg, logger)
	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", cfg.RpcPort)
	if err != nil {
		return err
	}

	rpc := grpc_service.NewRPC(
		logger,
		customerUseCase,
	)

	file_processing.RegisterFileProcessingServiceServer(s, rpc)

	logger.Info("service is running...", zap.String("port", cfg.RpcPort))
	if err = s.Serve(l); err != nil {
		return err
	}

	c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	db.Close()
	logger.Sync()
	s.GracefulStop()

	log.Println("service stop")
	return nil
}
