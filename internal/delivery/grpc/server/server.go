package server

import (
	"github/architecture/config"

	"google.golang.org/grpc"
)

func NewGRPCServer(cfg *config.Config) (*grpc.Server, error) {
	s := grpc.NewServer(
	// grpc.CallCustomCodec(grpc.),
	)

	return s, nil
}
