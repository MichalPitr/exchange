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

func (s *engine.Server) SendOrder(ctx context.Context, in *pb.OrderRequest) (*pb.OrderResponse, error) {
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

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	e := engine.New(32)

	go engine.ProcessOrders(e)

	pb.RegisterOrderServiceServer(s, e)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
