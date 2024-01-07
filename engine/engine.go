package engine

import (
	"context"
	"log"
	"time"

	"github.com/MichalPitr/exchange/orderbook"

	pb "github.com/MichalPitr/exchange/protos"
)

type Engine struct {
	pb.UnimplementedOrderServiceServer
	orderQueue chan orderbook.Order
	sellBook   *orderbook.Book
	buyBook    *orderbook.Book
}

func (s *Engine) SendOrder(ctx context.Context, in *pb.OrderRequest) (*pb.OrderResponse, error) {
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

func New(queueSize int) *Engine {
	return &Engine{
		orderQueue: make(chan orderbook.Order, queueSize),
		sellBook:   orderbook.New(true),
		buyBook:    orderbook.New(false),
	}
}

func ProcessOrders(e *Engine) {
	for order := range e.orderQueue {
		// Process the order
		log.Printf("Current buybook: %v\n", e.buyBook)
		log.Printf("Current sellbook: %v\n", e.sellBook)

		log.Printf("Processing order: %v\n", order)
		if !match(order, e) {
			if order.OrderType == "BUY" {
				e.buyBook.Push(orderbook.Item{Order: order})
			} else {
				e.sellBook.Push(orderbook.Item{Order: order})
			}
		}
		order.ResultChan <- orderbook.OrderResult{Message: "Processed", Success: true}
	}
}

func match(order orderbook.Order, e *Engine) bool {
	// Check if order can be served by existing orders in the orderbook. Might have to combine multiple existing orders together.
	log.Printf("Matching order %v\n", order)
	if order.Type == "BUY" {
		if top, ok := e.sellBook.Peek(); ok {
			log.Printf("Top of sellbook: %v", top)
			if top.Price <= order.Price {
				log.Printf("Found matching order for %v: %v\nSettlement price: %d", order, top, top.Price)
				return true
			}
		}
	} else if order.Type == "SELL" {
		if top, ok := e.buyBook.Peek(); ok {
			log.Printf("Top of sellbook: %v", top)
			if top.Price >= order.Price {
				log.Printf("Found matching order for %v: %v\nSettlement price: %d", order, top, top.Price)
				return true
			}
		}
	}
	return false
}
