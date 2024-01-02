package main

import (
	"context"
	"log"
	"net"

	pb "github.com/MichalPitr/exchange/protos"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedOrderServiceServer
}

func (s *server) SendOrder(ctx context.Context, in *pb.OrderRequest) (*pb.OrderResponse, error) {
	log.Printf("Received: %v", in)
	// Process the order here
	return &pb.OrderResponse{Status: "Success", Details: "Order processed"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
