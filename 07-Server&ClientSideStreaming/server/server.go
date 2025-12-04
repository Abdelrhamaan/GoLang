package main

import (
	"bufio"
	calcpb "grpcstreams/proto/gen"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	calcpb.UnimplementedCaculatorServiceServer
}

func (s *server) FabonacciStreams(req *calcpb.AddFabonachiStreamRequest, stream calcpb.CaculatorService_FabonacciStreamsServer) error {
	n := req.N
	a, b := 0, 1

	for i := 0; i < int(n); i++ {
		err := stream.Send(
			&calcpb.AddFabonachiStreamResponse{
				Number: int32(a),
			},
		)
		if err != nil {
			return err
		}
		a, b = b, a+b
		time.Sleep(time.Second)
	}
	return nil

}

func (s *server) ClientSendNumber(stream calcpb.CaculatorService_ClientSendNumberServer) error {
	var sum int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Client finished sending, send back the final response
			return stream.SendAndClose(&calcpb.SendResponseNum{
				Sum: sum,
			})
		}
		if err != nil {
			// Handle other errors
			return err
		}

		// Accumulate the sum
		sum += req.Number
		log.Printf("Received number: %d, Current sum: %d", req.Number, sum)
	}
}

func (s *server) Chat(stream calcpb.CaculatorService_ChatServer) error {
	// Goroutine to receive messages from client
	go func() {
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				log.Println("Client finished sending messages")
				return
			}
			if err != nil {
				log.Printf("Error receiving: %v", err)
				return
			}
			log.Printf("Client says: %s", req.GetMessage())
		}
	}()

	// Main thread sends messages from server stdin
	reader := bufio.NewReader(os.Stdin)
	for {
		log.Print("Server, enter message (or 'quit' to exit): ")
		str, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		input := strings.TrimSpace(str)

		if input == "quit" {
			log.Println("Server ending chat session")
			return nil
		}

		err = stream.Send(&calcpb.Message{
			Message: input,
		})
		if err != nil {
			log.Printf("Error sending: %v", err)
			return err
		}
		log.Printf("Server sent: %s", input)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln("Error happens when creating server", err)
	}

	grpcServer := grpc.NewServer()
	calcpb.RegisterCaculatorServiceServer(grpcServer, &server{})

	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
