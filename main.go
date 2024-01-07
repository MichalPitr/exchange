package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/MichalPitr/exchange/protos"
	"google.golang.org/grpc"

	"github.com/MichalPitr/exchange/engine"
	"github.com/MichalPitr/exchange/orderbook"
)

type server struct {
	pb.UnimplementedOrderServiceServer
	orderQueue chan orderbook.Order
}

func (s *server) SendOrder(ctx context.Context, in *pb.OrderRequest) (*pb.OrderResponse, error) {
	log.Printf("Received: %v", in)
	resultChan := make(chan orderbook.OrderResult, 1)
	// Process the order here
	order := orderbook.Order{
		UserID:     in.UserId,
		Type:       in.Type,
		OrderType:  in.OrderType,
		Amount:     in.Amount,
		Price:      in.Price,
		Time:       time.Now().UnixNano(),
		ResultChan: resultChan,
	}

	// Enqueue the order.
	select {
	case s.orderQueue <- order:
		// Order enqueued successfully
	case <-ctx.Done():
		// Handle context cancellation
		return nil, ctx.Err()
	}

	var result orderbook.OrderResult
	select {
	case result = <-resultChan:
		// Received the processing result
	case <-ctx.Done():
		// Handle context cancellation
		return nil, ctx.Err()
	}

	if result.Success {
		return &pb.OrderResponse{Status: "Success", Details: result.Message}, nil
	} else {
		return &pb.OrderResponse{Status: "Failed", Details: result.Message}, nil
	}
}

func newServer(queueSize int) *server {
	return &server{
		orderQueue: make(chan orderbook.Order, queueSize),
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	srv := newServer(32)

	go engine.ProcessOrders(srv.orderQueue)

	pb.RegisterOrderServiceServer(s, srv)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
