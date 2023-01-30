package app

import (
	"context"
	"github/architecture/config"
	"github/architecture/internal/usecase"
	"github/architecture/internal/usecase/repo"
	"github/architecture/pkg/logger"
	"github/architecture/pkg/postgres"
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

	logger.Info("service is running...", zap.String("port", cfg.RpcPort))

	c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	db.Close()
	logger.Close()

	log.Println("service stop")
	return nil
}
