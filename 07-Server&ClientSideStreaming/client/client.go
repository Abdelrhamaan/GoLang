package main

import (
	"context"
	calcpb "grpcstreams/proto/gen"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/encoding/gzip"
)

func main() {
	// conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// compressing data with gzip
	// this will compressing all requests
	// also we can make compressing for specific requests
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client := calcpb.NewCaculatorServiceClient(conn)

	ctx := context.Background()
	// server side streaming
	req := &calcpb.AddFabonachiStreamRequest{
		N: 10,
	}

	// stream, err := client.FabonacciStreams(ctx, req)
	stream, err := client.FabonacciStreams(ctx, req, grpc.UseCompressor(gzip.Name))

	if err != nil {
		log.Fatal(err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("Stream ended - received all Fibonacci numbers")
			break
		}
		if err != nil {
			log.Fatalf("Error receiving from stream: %v", err)
		}
		log.Println("Fibonacci number is:", resp.GetNumber())
	}

	// client side streaming
	stream1, err := client.ClientSendNumber(ctx)
	if err != nil {
		log.Fatalln("error happens while start straming:", err)
	}

	for num := range 9 {
		err := stream1.Send(&calcpb.SendRequestNum{Number: int32(num)})
		if err != nil {
			log.Fatalln("error happens while streaming: ", err)
		}
		time.Sleep(time.Second)

	}
	res, err := stream1.CloseAndRecv()
	if err != nil {
		log.Fatalln("error rec reponse: ", res)
	}
	log.Println("response: ", res.Sum)

	// bidirectional streaming
	chatStream, err := client.Chat(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	doneChan := make(chan struct{})

	// go routine for sending messages
	go func() {
		msgs := []string{"Hello !", "How are you?", "WHats your name?", "How old are you?"}
		for _, m := range msgs {
			err := chatStream.Send(&calcpb.Message{Message: m})
			if err != nil {
				log.Fatalln("error happens", err)
			}
			log.Printf("Sent: %s", m)
			time.Sleep(time.Second)
		}
		log.Println("Finished sending all messages, waiting for server responses...")
		// Don't close - keep listening
	}()

	// go routine for receiving msgs
	go func() {
		// Give server time to respond (adjust as needed)
		timeout := time.After(15 * time.Second)
		msgCount := 0

		for {
			select {
			case <-timeout:
				log.Println("Timeout waiting for server messages")
				doneChan <- struct{}{}
				return
			default:
				rec, err := chatStream.Recv()
				if err == io.EOF {
					log.Println("Server closed the stream")
					doneChan <- struct{}{}
					return
				}
				if err != nil {
					log.Printf("Error receiving: %v", err)
					doneChan <- struct{}{}
					return
				}
				log.Printf("Received from server: %s", rec.GetMessage())
				msgCount++
			}
		}
	}()

	// Wait for receiver to finish
	<-doneChan
	log.Println("Chat completed successfully")
}
