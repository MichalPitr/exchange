package engine

import (
	"log"

	"github.com/MichalPitr/exchange/orderbook"
)

func ProcessOrders(queue <-chan orderbook.Order) {
	for order := range queue {
		// Process the order
		log.Printf("Processing order: %v", order)
		order.ResultChan <- orderbook.OrderResult{Message: "Processed", Success: true}
	}
}
