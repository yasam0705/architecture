package server

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Error(err error) error {
	var st *status.Status
	if err == nil {
		return nil
	}

	switch {
	default:
		st = status.New(codes.Internal, err.Error())
	}
	return st.Err()
}
