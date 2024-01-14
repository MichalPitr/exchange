package engine

import (
	"container/heap"
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

func New(queueSize int) *Engine {
	return &Engine{
		orderQueue: make(chan orderbook.Order, queueSize),
		sellBook:   orderbook.New(true),
		buyBook:    orderbook.New(false),
	}
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

func ProcessOrders(e *Engine) {
	for order := range e.orderQueue {
		// Process the order
		log.Printf("Current buybook: %v\n", e.buyBook)
		log.Printf("Current sellbook: %v\n", e.sellBook)

		log.Printf("Processing order: %v\n", order)
		if !match(order, e) {
			if order.Type == "BUY" {
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
	remainingAmount := order.Amount
	if order.Type == "BUY" {
		for e.sellBook.Len() > 0 && remainingAmount > 0 {
			if top, ok := e.sellBook.Peek(); ok {
				log.Printf("Top of sellbook: %v", *top)
				if top.Price > order.Price {
					return false
				}
				if top.Amount > remainingAmount {
					log.Printf("Found matching order for %v: %v\nSettlement price: %d\nAmount: %d", order, top, top.Price, remainingAmount)
					// Top SELL is larger than remaining BUY, so update existing SELL.
					top.Amount -= remainingAmount
					return true
				} else {
					log.Printf("Found partial matching order for %v: %v\nSettlement price: %d\nAmount: %d", order, top, top.Price, top.Amount)
					remainingAmount -= top.Amount
					heap.Pop(e.sellBook)
				}
			}
		}
	} else if order.Type == "SELL" {
		for e.buyBook.Len() > 0 && remainingAmount > 0 {
			if top, ok := e.buyBook.Peek(); ok {
				log.Printf("Top of buyBook: %v", *top)
				if top.Price < order.Price {
					return false
				}
				if top.Amount > remainingAmount {
					log.Printf("Found matching order for %v: %v\nSettlement price: %d\nAmount: %d", order, top, top.Price, remainingAmount)
					top.Amount -= remainingAmount
					return true
				} else {
					log.Printf("Found partial matching order for %v: %v\nSettlement price: %d\nAmount: %d", order, top, top.Price, top.Amount)
					remainingAmount -= top.Amount
					heap.Pop(e.buyBook)
				}
			}
		}
	}
	return remainingAmount == 0
}
