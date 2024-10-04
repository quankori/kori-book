package main

import (
	"log"
	"net"

	pb "github.com/quankori/go-grpc/server/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.StreamServiceServer
}

func (s *server) BidirectionalStream(stream pb.StreamService_BidirectionalStreamServer) error {
	for {
		request, err := stream.Recv()
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			break
		}
		log.Printf("Received message: %s", request.Message)

		response := &pb.StreamResponse{Message: "Server received: " + request.Message}
		if err := stream.Send(response); err != nil {
			log.Printf("Error sending message: %v", err)
			break
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStreamServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
