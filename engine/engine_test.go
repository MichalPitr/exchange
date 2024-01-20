package engine

import (
	"container/heap"
	"testing"

	"github.com/MichalPitr/exchange/orderbook"
)

func TestMatchOneBuyTwoSells(t *testing.T) {
	engine, err := New(32)
	if err != nil {
		t.Fatal(err)
	}
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
	remainder, matches := match(engine, buy)
	if remainder != 0 {
		t.Errorf("Expected 0 remainder, but got %d", remainder)
	}
	if len(matches) != 2 {
		t.Errorf("Expected 2 matches, but got %d", len(matches))
	}
}

func TestMatchOneBuyPartialSells(t *testing.T) {
	engine, err := New(32)
	if err != nil {
		t.Fatal(err)
	}
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

	remainder, matches := match(engine, buy)
	if remainder != 0 {
		t.Errorf("Expected 0 remainder, but got %d", remainder)
	}
	if len(matches) != 2 {
		t.Errorf("Expected 2 matches, but got %d", len(matches))
	}

	if order, ok := engine.sellBook.Peek(); ok {
		if order.Amount != 5 {
			t.Errorf("Expected top sell order to have quantity 5, but has %d", order.Amount)
		}
	}
}

func TestMatchBuyNoMatch(t *testing.T) {
	engine, err := New(32)
	if err != nil {
		t.Fatal(err)
	}
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
	remainder, matches := match(engine, buy)
	if remainder != 10 {
		t.Errorf("Expected 10 remainder, but got %d", remainder)
	}
	if len(matches) != 0 {
		t.Errorf("Expected 0 matches, but got %d", len(matches))
	}

	if order, ok := engine.sellBook.Peek(); ok {
		if order.Amount != 10 {
			t.Errorf("Sellbook shouldn't change when match fails.")
		}
	}
}

func TestMatchOneSellTwoBuys(t *testing.T) {
	engine, err := New(32)
	if err != nil {
		t.Fatal(err)
	}
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

	if remainder, _ := match(engine, sell); remainder != 0 {
		t.Error("Expected a match")
	}
}

func TestProcessLimitOrder(t *testing.T) {
	engine, err := New(32)
	if err != nil {
		t.Fatal(err)
	}
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
		Amount:     20, // Amount higher than demand in orderbook
		Price:      100,
		Time:       1641016800, // Example Unix timestamp
		ResultChan: nil,        // Or initialize as appropriate
	}

	processOrder(engine, sell)

	if order, ok := engine.sellBook.Peek(); ok {
		if order.Amount != 10 {
			t.Errorf("Expected 10 units of order to be added to orderbook, instead got %d", order.Amount)
		}
	} else {
		t.Error("Expected order to be added to sell orderbook.")
	}
}

func TestProcessMarketOrder(t *testing.T) {
	engine, err := New(32)
	if err != nil {
		t.Fatal(err)
	}
	orders := []orderbook.Order{
		{
			UserID:     1,
			Type:       "BUY",
			OrderType:  "LIMIT", // Buy is limit so that it is added to orderbook.
			Amount:     10,
			Price:      100,
			Time:       1641016800, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
	}

	for _, order := range orders {
		heap.Push(engine.buyBook, orderbook.Item{Order: order})
	}

	if engine.sellBook.Len() != 0 {
		t.Error("Expected empty sellbook.")
	}

	sell := orderbook.Order{
		UserID:     1,
		Type:       "SELL",
		OrderType:  "MARKET",
		Amount:     20, // Amount higher than demand in orderbook
		Price:      100,
		Time:       1641016800, // Example Unix timestamp
		ResultChan: nil,        // Or initialize as appropriate
	}

	processOrder(engine, sell)

	if order, ok := engine.sellBook.Peek(); ok {
		t.Errorf("Expected market order to not be added to sell orderbook. %v", order)
	}
}
