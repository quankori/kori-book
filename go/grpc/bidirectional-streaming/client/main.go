package main

import (
	"context"
	"log"
	"time"

	"github.com/quankori/go-grpc/server/services"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := services.NewStreamServiceClient(conn)

	stream, err := client.BidirectionalStream(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				log.Fatalf("Error receiving message: %v", err)
			}
			log.Printf("Received message: %s", res.GetMessage())
		}
	}()

	for i := 0; i < 10; i++ {
		if err := stream.Send(&services.StreamRequest{Message: "Hello"}); err != nil {
			log.Fatalf("Failed to send a note: %v", err)
		}
		time.Sleep(time.Second)
	}

	stream.CloseSend()
}
