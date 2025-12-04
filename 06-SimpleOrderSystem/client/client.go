package main

import (
	"context"
	"log"
	orderpb "simpleordersystem/server/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Something wrong happens when create client:", err)
	}
	defer conn.Close()

	state := conn.GetState()

	client := orderpb.NewOrderServiceClient(conn)
	res, err := client.ListOrders(context.Background(), &orderpb.ListOrdersRequest{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Orders:", res.Orders)
	// creating items
	res0, err0 := client.CreateOrder(context.Background(), &orderpb.CreateOrderRequest{
		Customer: &orderpb.Customer{Id: 2, Name: "Omar"},
		// outer * to create slice of pointers
		OrderItem: []*orderpb.OrderItem{
			// & here to get the memmory address(pointer) to save it
			&orderpb.OrderItem{
				ProductId: 1, Quantity: 2, Price: 20,
			},
			&orderpb.OrderItem{
				ProductId: 2, Quantity: 5, Price: 100,
			},
		},
	})
	if err0 != nil {
		log.Fatalln(err0)
	}
	log.Println("Created Order:", res0.Order)

	res1, err1 := client.CreateOrder(context.Background(), &orderpb.CreateOrderRequest{
		Customer: &orderpb.Customer{Id: 1, Name: "Ali"},
		// outer * to create slice of pointers
		OrderItem: []*orderpb.OrderItem{
			// & here to get the memmory address(pointer) to save it
			&orderpb.OrderItem{
				ProductId: 1, Quantity: 2, Price: 20,
			},
			&orderpb.OrderItem{
				ProductId: 2, Quantity: 5, Price: 100,
			},
		},
	})
	if err1 != nil {
		log.Fatalln(err1)
	}
	log.Println("Created Order:", res1.Order)
	// list items
	res3, _ := client.ListOrders(context.Background(), &orderpb.ListOrdersRequest{})
	log.Println("Orders:", res3.Orders)

	// get item by id
	res4, err4 := client.GetOrder(context.Background(), &orderpb.GetOrderRequest{OrderId: 1})
	if err4 != nil {
		log.Fatalln(err4)
	}

	// delete item
	res5, err5 := client.DeleteOrder(context.Background(), &orderpb.DeleteOrderRequest{
		OrderId: 2,
	})
	if err5 != nil {
		log.Fatalln(err5)
	}
	log.Println("Deleted res: ", res5.Message, res5.Order)

	log.Println("Get Order:", res4.Order)

	log.Println("connection state is:", state)

}
