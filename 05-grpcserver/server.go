package main

import (
	// Context package provides request-scoped values, cancellation signals, and deadlines.
	"context"
	// Log package for simple logging.
	"log"
	// Net package for network operations, used to listen on a TCP port.
	"net"
	// Import the generated protobuf code; alias to calculatorpb for clarity.
	calculatorpb "simplegrpcserver/gen"

	// gRPC library for creating server and handling RPCs.
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// server implements the Calculate service defined in the protobuf.
// It embeds the unimplemented server to ensure forward compatibility.
type server struct {
	calculatorpb.UnimplementedCalculateServer
	calculatorpb.UnimplementedGreeterServer
}

// Add implements the Add RPC method.
// It receives a context and an AddRequest, and returns an AddResponse.
func (s *server) Add(ctx context.Context, req *calculatorpb.AddRequest) (*calculatorpb.AddResponse, error) {
	return &calculatorpb.AddResponse{
		Sum: req.A + req.B,
	}, nil
}

func (s *server) Greet(ctx context.Context, greet_req *calculatorpb.GreeterRequest) (*calculatorpb.GreeterResponse, error) {
	return &calculatorpb.GreeterResponse{
		Message: "Hello My Friend: " + greet_req.Name,
	}, nil
}

func main() {

	// lets passing cred and cert
	const CERT = "cert.pem"
	const KEY = "cred.pem"

	// Define the port the server will listen on.
	const PORT = ":50051"
	// Create a TCP listener on the specified port.
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatal("Cannot listen: ", err)
	}

	creds, err := credentials.NewServerTLSFromFile(CERT, KEY)

	if err != nil {
		log.Println("Cannot read creds: ", err)
	}

	// Create a new gRPC server instance.
	grpcServer := grpc.NewServer(grpc.Creds(creds))

	// Register the Calculate service implementation with the gRPC server.
	calculatorpb.RegisterCalculateServer(grpcServer, &server{})
	calculatorpb.RegisterGreeterServer(grpcServer, &server{})

	// Log that the server has started.
	log.Println("Server is started on port: ", PORT)
	// Start serving incoming connections.
	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatal("Cannot serve on port 50051: ", err)
	}

}
