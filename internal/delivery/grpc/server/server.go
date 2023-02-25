package server

import (
	"github/architecture/config"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func NewGRPCServer(cfg *config.Config, log *zap.Logger) (*grpc.Server, error) {
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_zap.UnaryServerInterceptor(log),
			grpc_recovery.UnaryServerInterceptor(),
			grpc_ctxtags.UnaryServerInterceptor(),
			UnaryServerInterceptor(),
			otelgrpc.UnaryServerInterceptor(),
			UnaryServerInterceptorErrorHandling(),
		),
	)

	return s, nil
}
