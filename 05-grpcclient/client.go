package main

import (
	"context"
	"log"
	calculatorpb "simplegrpcclient/gen"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// it need tls connection with authorized credentials so we told him here that we will not use credentials
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("Cannot connect ot server: ", err)
	}

	defer conn.Close()

	// create client from main_grpc.pb.go file
	client := calculatorpb.NewCalculateClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second) // that is a dead line

	defer cancel()

	// create request to pass it in add rpc service
	req := calculatorpb.AddRequest{
		A: 5,
		B: 10,
	}

	res, err := client.Add(ctx, &req)

	if err != nil {
		log.Fatalln("Cannot add: ", err)
	}

	log.Println("adding two function: ", res.Sum)

}
