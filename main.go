package main

import (
	"log"
	"net"

	pb "github.com/MichalPitr/exchange/protos"
	"google.golang.org/grpc"

	"github.com/MichalPitr/exchange/engine"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	e, err := engine.New(1000)
	if err != nil {
		log.Fatalf("Failed to start engine: %v", err)
	}

	go engine.ProcessOrders(e)

	pb.RegisterOrderServiceServer(s, e)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
