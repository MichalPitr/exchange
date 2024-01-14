package engine

import (
	"container/heap"
	"testing"

	"github.com/MichalPitr/exchange/orderbook"
)

func TestMatchOneBuyTwoSells(t *testing.T) {
	engine := New(32)

	orders := []orderbook.Order{
		{
			UserID:     1,
			Type:       "SELL",
			OrderType:  "LIMIT",
			Amount:     10,
			Price:      100,
			Time:       1641016800, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
		{
			UserID:     2,
			Type:       "SELL",
			OrderType:  "LIMIT",
			Amount:     10,
			Price:      200,
			Time:       1641103200, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
		{
			UserID:     3,
			Type:       "SELL",
			OrderType:  "LIMIT",
			Amount:     10,
			Price:      150,
			Time:       1641189600, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
	}

	// Add items to both heaps
	for _, order := range orders {
		heap.Push(engine.sellBook, orderbook.Item{Order: order})
	}

	buy := orderbook.Order{
		UserID:     1,
		Type:       "BUY",
		OrderType:  "LIMIT",
		Amount:     20,
		Price:      200,
		Time:       1641016800, // Example Unix timestamp
		ResultChan: nil,        // Or initialize as appropriate
	}

	if !match(buy, engine) {
		t.Error("Expected a match")
	}
}

func TestMatchOneBuyPartialSells(t *testing.T) {
	engine := New(32)

	orders := []orderbook.Order{
		{
			UserID:     1,
			Type:       "SELL",
			OrderType:  "LIMIT",
			Amount:     10,
			Price:      100,
			Time:       1641016800, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
		{
			UserID:     2,
			Type:       "SELL",
			OrderType:  "LIMIT",
			Amount:     10,
			Price:      200,
			Time:       1641103200, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
		{
			UserID:     3,
			Type:       "SELL",
			OrderType:  "LIMIT",
			Amount:     10,
			Price:      150,
			Time:       1641189600, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
	}

	// Add items to both heaps
	for _, order := range orders {
		heap.Push(engine.sellBook, orderbook.Item{Order: order})
	}

	buy := orderbook.Order{
		UserID:     1,
		Type:       "BUY",
		OrderType:  "LIMIT",
		Amount:     15,
		Price:      200,
		Time:       1641016800, // Example Unix timestamp
		ResultChan: nil,        // Or initialize as appropriate
	}

	if !match(buy, engine) {
		t.Error("Expected a match")
	}
	if order, ok := engine.sellBook.Peek(); ok {
		if order.Amount != 5 {
			t.Errorf("Expected top sell order to have quantity 5, but has %d", order.Amount)
		}
	}
}

func TestMatchBuyNoMatch(t *testing.T) {
	engine := New(32)

	orders := []orderbook.Order{
		{
			UserID:     1,
			Type:       "SELL",
			OrderType:  "LIMIT",
			Amount:     10,
			Price:      100,
			Time:       1641016800, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
	}

	// Add items to both heaps
	for _, order := range orders {
		heap.Push(engine.sellBook, orderbook.Item{Order: order})
	}

	buy := orderbook.Order{
		UserID:     1,
		Type:       "BUY",
		OrderType:  "LIMIT",
		Amount:     10,
		Price:      50,         // Price set below sell price.
		Time:       1641016800, // Example Unix timestamp
		ResultChan: nil,        // Or initialize as appropriate
	}

	if match(buy, engine) {
		t.Error("Expected no match")
	}
	if order, ok := engine.sellBook.Peek(); ok {
		if order.Amount != 10 {
			t.Errorf("Sellbook shouldn't change when match fails.")
		}
	}
}

func TestMatchOneSellTwoBuys(t *testing.T) {
	engine := New(32)

	orders := []orderbook.Order{
		{
			UserID:     1,
			Type:       "BUY",
			OrderType:  "LIMIT",
			Amount:     10,
			Price:      100,
			Time:       1641016800, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
		{
			UserID:     2,
			Type:       "BUY",
			OrderType:  "LIMIT",
			Amount:     10,
			Price:      200,
			Time:       1641103200, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
		{
			UserID:     3,
			Type:       "BUY",
			OrderType:  "LIMIT",
			Amount:     10,
			Price:      150,
			Time:       1641189600, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
	}

	// Add items to both heaps
	for _, order := range orders {
		heap.Push(engine.buyBook, orderbook.Item{Order: order})
	}

	sell := orderbook.Order{
		UserID:     1,
		Type:       "SELL",
		OrderType:  "LIMIT",
		Amount:     20,
		Price:      150,
		Time:       1641016800, // Example Unix timestamp
		ResultChan: nil,        // Or initialize as appropriate
	}

	if !match(sell, engine) {
		t.Error("Expected a match")
	}
}
