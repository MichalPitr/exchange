package engine

import (
	"container/heap"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MichalPitr/exchange/orderbook"

	pb "github.com/MichalPitr/exchange/protos"
)

type Engine struct {
	pb.UnimplementedOrderServiceServer
	orderQueue  chan orderbook.Order
	sellBook    *orderbook.Book
	buyBook     *orderbook.Book
	nextOrderId uint64
}

type Match struct {
	buyId  uint64
	sellId uint64
	amount int32
	price  int64 // Technically unnecessary as it can be reconstructed from the orders and picking whichever was older but convenient.
}

func New(queueSize int) *Engine {
	return &Engine{
		orderQueue:  make(chan orderbook.Order, queueSize),
		sellBook:    orderbook.New(true),
		buyBook:     orderbook.New(false),
		nextOrderId: 0,
	}
}

func (s *Engine) SendOrder(ctx context.Context, in *pb.OrderRequest) (*pb.OrderResponse, error) {
	log.Printf("Received: %v", in)
	resultChan := make(chan orderbook.OrderResult, 1)
	// Process the order here
	order := orderbook.Order{
		Id:         s.nextOrderId,
		UserID:     in.UserId,
		Type:       in.Type,
		OrderType:  in.OrderType,
		Amount:     in.Amount,
		Price:      in.Price,
		Time:       time.Now().UnixNano(),
		ResultChan: resultChan,
	}
	s.nextOrderId++

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
		processOrder(e, order)
		order.ResultChan <- orderbook.OrderResult{Message: "Processed", Success: true}
	}
}

func processOrder(e *Engine, order orderbook.Order) {
	log.Printf("Processing order: %v\n", order)
	if remainder, matches := match(e, order); remainder == 0 {
		log.Printf("Fully matched order with: %v", matches)
	} else {
		if len(matches) > 0 {
			fmt.Printf("Partially matched ordered with: %v", matches)
		}

		if order.OrderType == "MARKET" {
			// Unfilled part of market order does not enter orderbook.
			return
		} else if order.OrderType == "LIMIT" {
			order.Amount = remainder
			if order.Type == "BUY" {
				e.buyBook.Push(orderbook.Item{Order: order})
			} else {
				e.sellBook.Push(orderbook.Item{Order: order})
			}
		}
	}
}

func match(e *Engine, order orderbook.Order) (int32, []Match) {
	// Check if order can be served by existing orders in the orderbook. Might have to combine multiple existing orders together.
	log.Printf("Matching order %v\n", order)
	remainingAmount := order.Amount
	matches := make([]Match, 0)
	if order.Type == "BUY" {
		for e.sellBook.Len() > 0 && remainingAmount > 0 {
			if top, ok := e.sellBook.Peek(); ok {
				if top.Price > order.Price {
					return remainingAmount, matches
				}
				if top.Amount > remainingAmount {
					matches = append(matches, Match{order.Id, top.Id, remainingAmount, top.Price})
					// Top SELL is larger than remaining BUY, so update existing SELL.
					top.Amount -= remainingAmount
					return 0, matches
				} else {
					matches = append(matches, Match{order.Id, top.Id, top.Amount, top.Price})
					remainingAmount -= top.Amount
					heap.Pop(e.sellBook)
				}
			}
		}
	} else if order.Type == "SELL" {
		for e.buyBook.Len() > 0 && remainingAmount > 0 {
			if top, ok := e.buyBook.Peek(); ok {
				if top.Price < order.Price {
					return remainingAmount, matches
				}
				if top.Amount > remainingAmount {
					matches = append(matches, Match{top.Id, order.Id, remainingAmount, top.Price})
					top.Amount -= remainingAmount
					return 0, matches
				} else {
					matches = append(matches, Match{top.Id, order.Id, top.Amount, top.Price})
					remainingAmount -= top.Amount
					heap.Pop(e.buyBook)
				}
			}
		}
	}
	return remainingAmount, matches
}
