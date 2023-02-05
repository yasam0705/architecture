package app

import (
	"context"
	"github/architecture/config"
	"github/architecture/genproto/customer_service"
	grpc_server "github/architecture/internal/delivery/grpc/server"
	grpc_service "github/architecture/internal/delivery/grpc/service"
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

	customerUseCase := usecase.NewCustomer(cfg.ContextTimeout, customerRepo)
	_ = customerUseCase

	s, err := grpc_server.NewGRPCServer(cfg)
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

	customer_service.RegisterCustomerServiceServer(s, rpc)

	logger.Info("service is running...", zap.String("port", cfg.RpcPort))
	if err = s.Serve(l); err != nil {
		return err
	}

	c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	db.Close()
	logger.Close()

	log.Println("service stop")
	return nil
}
