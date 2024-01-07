package engine

import (
	"log"

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

func ProcessOrders(e Engine) {
	for order := range e.orderQueue {
		// Process the order
		log.Printf("Processing order: %v", order)
		order.ResultChan <- orderbook.OrderResult{Message: "Processed", Success: true}
	}
}

func match() {}
