package main

import (
	"context"
	"log"
	calculatorpb "simplegrpcclient/gen"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// lets read cert from file and pass it to read it
	const CERT = "cert.pem"

	cert, err := credentials.NewClientTLSFromFile(CERT, "")

	if err != nil {
		log.Println("Client cannot get certicate: ", err)
	}

	// it need tls connection with authorized credentials so we told him here that we will not use credentials
	// conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(cert))

	if err != nil {
		log.Fatalln("Cannot connect ot server: ", err)
	}

	defer conn.Close()

	// 1️⃣ Calculate client – used for the Add RPC
	calcClient := calculatorpb.NewCalculateClient(conn)
	doAdd(calcClient)

	// 2️⃣ Greeter client – used for the Greet RPC
	greeterClient := calculatorpb.NewGreeterClient(conn)
	doGreet(greeterClient)

	// just to check connection
	state := conn.GetState()

	log.Println("Server State is: ", state)

}

func doAdd(c calculatorpb.CalculateClient) {
	log.Println("Starting to do Add...")
	req := &calculatorpb.AddRequest{
		A: 5,
		B: 10,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Add(ctx, req)
	if err != nil {
		log.Fatalln("Cannot add: ", err)
	}

	log.Println("adding two function: ", res.Sum)
}

func doGreet(c calculatorpb.GreeterClient) {
	log.Println("Starting to do Greet...")
	req := &calculatorpb.GreeterRequest{
		Name: "Ali",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.Greet(ctx, req)
	if err != nil {
		log.Fatalln("Cannot greet: ", err)
	}
	log.Println("Greeter message is : ", res.Message)
}
