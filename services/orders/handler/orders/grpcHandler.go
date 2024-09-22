package handler

import (
	"context"

	"github.com/nanasintown/gomcrsv/services/common/genproto/orders"
	"github.com/nanasintown/gomcrsv/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}
 
func NewGrpcOrdersService(grpc *grpc.Server, orderService types.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{
		orderService: orderService,
	} 
	orders.RegisterOrderServiceServer(grpc, gRPCHandler)

}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID: 32,
		CustomerID: 1,
		ProductID: 12,
		Quantity: 23,
	}
	err := h.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}
	return res, nil
	// return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}


func (h *OrdersGrpcHandler) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) (*orders.GetOrdersResponse, error) {
	o := h.orderService.GetOrders(ctx)
	res := &orders.GetOrdersResponse{
		Orders: o,
	}

	return res, nil
}