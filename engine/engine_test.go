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

	if remainder, _ := match(buy, engine); remainder != 0 {
		t.Errorf("Expected 0 remainder, but got %d", remainder)
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

	if remainder, _ := match(buy, engine); remainder != 0 {
		t.Errorf("Expected 0 remainder, but got %d", remainder)
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
	if remainder, _ := match(buy, engine); remainder != 10 {
		t.Errorf("Expected 10 remainder, but got %d", remainder)
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

	if remainder, _ := match(sell, engine); remainder != 0 {
		t.Error("Expected a match")
	}
}

func TestMatchInsufficientDemand(t *testing.T) {
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
	}

	// Add items to both heaps
	for _, order := range orders {
		heap.Push(engine.buyBook, orderbook.Item{Order: order})
	}

	if engine.sellBook.Len() != 0 {
		t.Error("Expected empty sellbook.")
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

	if remainder, _ := match(sell, engine); remainder != 10 {
		t.Errorf("Expected 10 remainder, but got %d", remainder)
	}
	if order, ok := engine.sellBook.Peek(); ok {
		if order.Amount != 10 {
			t.Errorf("Expected remainder of order to be 10, but got %d", order.Amount)
		}
	} else {
		t.Error("Expected remainder of order to be added to sellbook.")
	}

}
