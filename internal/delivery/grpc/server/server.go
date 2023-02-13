package server

import (
	"github/architecture/config"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func NewGRPCServer(cfg *config.Config) (*grpc.Server, error) {
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(),
			UnaryServerInterceptor(),
			grpc_ctxtags.UnaryServerInterceptor(),
			otelgrpc.UnaryServerInterceptor(),
		),
	)

	return s, nil
}
