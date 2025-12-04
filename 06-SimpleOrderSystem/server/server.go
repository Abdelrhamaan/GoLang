package main

import (
	"context"
	"fmt"
	"log"
	"net"
	orderpb "simpleordersystem/server/gen"

	"google.golang.org/grpc"
)

type server struct {
	orderpb.UnimplementedOrderServiceServer
	// Make simple map to make crud operations on it
	orders map[int32]*orderpb.Order
}

func (s *server) ListOrders(ctx context.Context, req *orderpb.ListOrdersRequest) (*orderpb.ListOrdersResponse, error) {
	var orderList []*orderpb.Order
	for _, o := range s.orders {
		orderList = append(orderList, o)
	}
	return &orderpb.ListOrdersResponse{
		Orders: orderList,
	}, nil
}

func (s *server) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {
	// 1- create new id
	newId := int32(len(s.orders) + 1)
	// 2- create orderct
	newOrder := &orderpb.Order{
		OrderId:  int32(newId),
		Customer: req.GetCustomer(),
		Items:    req.GetOrderItem(),
		Status:   "Created",
	}

	// 3-add it to our map
	s.orders[newId] = newOrder

	log.Println("Order has been created", newOrder)

	return &orderpb.CreateOrderResponse{
		Order: newOrder,
	}, nil

}

func (s *server) GetOrder(ctx context.Context, req *orderpb.GetOrderRequest) (*orderpb.GetOrderResponse, error) {
	order, ok := s.orders[req.OrderId]
	if !ok {
		return nil, fmt.Errorf("order not found")
	}

	return &orderpb.GetOrderResponse{
		Order: order,
	}, nil

}

func main() {

	const PORT = ":50053"

	// Listen on that port
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("Cannot listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// regiseter grpc server
	srv := &server{
		orders: make(map[int32]*orderpb.Order),
	}
	orderpb.RegisterOrderServiceServer(grpcServer, srv)

	log.Printf("Server started on %s", PORT)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
