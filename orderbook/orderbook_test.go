package orderbook

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestBuyOrderbook(t *testing.T) {
	buyBook := New(false)

	// Example orders
	orders := []Order{
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
			Amount:     5,
			Price:      200,
			Time:       1641103200, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
		{
			UserID:     3,
			Type:       "BUY",
			OrderType:  "LIMIT",
			Amount:     15,
			Price:      50,
			Time:       1641189600, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
	}

	expectPrice := []int64{200, 100, 50}

	// Add items to both heaps
	for _, order := range orders {
		heap.Push(buyBook, Item{Order: order})
	}

	// Pop items from both heaps
	fmt.Println("\nbuy book:")
	for _, p := range expectPrice {
		order, _ := buyBook.Peek()
		if order.Price != p {
			t.Errorf("Expected peek price: %d but got %d\n", p, order.Price)
		}
		item := heap.Pop(buyBook).(*Item)
		if item.Order.Price != p {
			t.Errorf("Expected price: %d but got %d\n", p, item.Order.Price)
		}
	}
}

func TestSellOrderbook(t *testing.T) {
	sellbook := New(true)

	// Example orders
	orders := []Order{
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
			Amount:     5,
			Price:      200,
			Time:       1641103200, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
		{
			UserID:     3,
			Type:       "SELL",
			OrderType:  "LIMIT",
			Amount:     15,
			Price:      50,
			Time:       1641189600, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
	}

	expectPrice := []int64{50, 100, 200}

	// Add items to both heaps
	for _, order := range orders {
		heap.Push(sellbook, Item{Order: order})
	}

	// Pop items from both heaps
	fmt.Println("\nsell book:")
	for _, p := range expectPrice {
		order, _ := sellbook.Peek()
		if order.Price != p {
			t.Errorf("Expected peek price: %d but got %d\n", p, order.Price)
		}
		item := heap.Pop(sellbook).(*Item)
		if item.Order.Price != p {
			t.Errorf("Expected top price: %d but got %d\n", p, item.Order.Price)
		}
	}
}
