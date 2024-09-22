package main

import (
	"log"
	"net"

	handler "github.com/nanasintown/gomcrsv/services/orders/handler/orders"
	"github.com/nanasintown/gomcrsv/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(grpcServer, orderService)
	log.Println("starting gRPC server on: ", s.addr)
	return grpcServer.Serve(lis)
}