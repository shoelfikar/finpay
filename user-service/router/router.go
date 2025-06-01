package router

import (
	"fmt"
	"net"
	"sync"

	"github.com/shoelfikar/finpay/user-service/helper"
	"github.com/shoelfikar/finpay/user-service/middleware"
	"google.golang.org/grpc"
)

type Routes struct {
	GrpcServer *grpc.Server
	Listener   net.Listener

}

func (r *Routes) RunGRPC(port string, wg *sync.WaitGroup) {
	defer wg.Done()

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		msg := fmt.Sprintf("Failed to create listener: %v", err)
		helper.LoggingError(msg)
	}

	msg := fmt.Sprintf("gRPC server is starting on port %s", port)
	helper.LoggingInfo(msg)

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(middleware.UnaryLoggingInterceptor, middleware.AuthInterceptor),
	)

	r.GrpcServer = grpcServer
	r.Listener = listener


	err = r.GrpcServer.Serve(r.Listener)
	if err != nil {
		helper.LoggingError(err.Error())
	}
}
