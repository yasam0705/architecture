package server

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	limitConcurrentRequests = map[string]int{
		"/customer_service.CustomerService/Get": 2, // save to redis,
	}
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		var (
			userId string
			method string = info.FullMethod
		)
		// REFACTOR
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			userId = md.Get("user_id")[0] // get from context, not metadata
			ctx = metadata.NewOutgoingContext(ctx, md)
		}

		client := NewClient(userId)

		maxCountRequests, ok := limitConcurrentRequests[method]
		if !ok {
			return handler(ctx, req)
		}

		clientCountRequest := client.ConcurrentRequests[method]
		fmt.Println(clientCountRequest, maxCountRequests)
		if clientCountRequest > maxCountRequests-1 {
			return nil, fmt.Errorf("error max concurrent requests, max: %d", maxCountRequests)
		}

		count := client.Increment(method)
		fmt.Println(count, client.ConcurrentRequests)

		resp, err = handler(ctx, req)

		count = client.Decrement(method)
		fmt.Println(count, client.ConcurrentRequests)

		if count == 0 {
			defer RemoveClient(client)
		}
		return resp, err
	}
}