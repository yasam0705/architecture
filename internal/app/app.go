package app

import (
	"github/architecture/config"
	"github/architecture/pkg/logger"

	"go.uber.org/zap"
)

func Run(cfg *config.Config) error {

	logger, err := logger.New(cfg.LogLevel, cfg.Environment, cfg.App+".log")
	if err != nil {
		return err
	}

	logger.Info("service is running...", zap.String("port", cfg.RpcPort))

	return nil
}
