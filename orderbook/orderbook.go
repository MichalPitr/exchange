package orderbook

import (
	"container/heap"
	"fmt"
	"log"
)

type Order struct {
	UserID     int32
	Type       string
	OrderType  string
	Amount     int32
	Price      int64
	Time       int64
	ResultChan chan OrderResult
}

type Item struct {
	Order Order

	// Necessary for heap.interface methods
	index int
}

type OrderResult struct {
	Success bool
	Message string
}

type Book struct {
	orders []*Item
	asc    bool // Denotes if orders are ordered in asc or desc order. asc is for Sells, desc is for Buys
}

func (b Book) Len() int { return len(b.orders) }

func (b Book) Less(i, j int) bool {
	// Min-heap if asc is true, else max-heap
	if b.asc {
		if b.orders[i].Order.Price == b.orders[j].Order.Price {
			return b.orders[i].Order.Time < b.orders[j].Order.Time
		}
		return b.orders[i].Order.Price < b.orders[j].Order.Price
	}

	if b.orders[i].Order.Price == b.orders[j].Order.Price {
		return b.orders[i].Order.Time < b.orders[j].Order.Time
	}
	return b.orders[i].Order.Price > b.orders[j].Order.Price
}

func (b Book) Swap(i, j int) { b.orders[i], b.orders[j] = b.orders[j], b.orders[i] }

func (b *Book) Push(x any) {
	n := len(b.orders)
	item := x.(Item)
	item.index = n
	b.orders = append(b.orders, &item)
}

func (b *Book) Pop() any {
	old := b.orders
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	b.orders = old[0 : n-1]
	return item
}

func New(asc bool) *Book {
	b := &Book{asc: asc}
	heap.Init(b)
	return b
}

func (b Book) Peek() (Order, bool) {
	if b.Len() == 0 {
		log.Println("Peeking empty orderbook.")
		return Order{}, false
	}
	return b.orders[0].Order, true
}

func main() {
	sellBook := New(true)
	buyBook := New(false)

	// Example orders
	orders := []Order{
		{
			UserID:     1,
			Type:       "Limit",
			OrderType:  "SELL",
			Amount:     10,
			Price:      100,
			Time:       1641016800, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
		{
			UserID:     2,
			Type:       "Limit",
			OrderType:  "SELL",
			Amount:     5,
			Price:      200,
			Time:       1641103200, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
		{
			UserID:     3,
			Type:       "Limit",
			OrderType:  "SELL",
			Amount:     15,
			Price:      50,
			Time:       1641189600, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
		{
			UserID:     1,
			Type:       "Limit",
			OrderType:  "BUY",
			Amount:     10,
			Price:      100,
			Time:       1641016800, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
		{
			UserID:     2,
			Type:       "Limit",
			OrderType:  "BUY",
			Amount:     5,
			Price:      200,
			Time:       1641103200, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
		{
			UserID:     3,
			Type:       "Limit",
			OrderType:  "BUY",
			Amount:     15,
			Price:      50,
			Time:       1641189600, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
	}

	// Add items to both heaps
	for _, order := range orders {
		if order.OrderType == "SELL" {
			heap.Push(sellBook, Item{Order: order})
		} else {
			heap.Push(buyBook, Item{Order: order})
		}
	}

	// Pop items from both heaps
	fmt.Println("Sell book:")
	for sellBook.Len() > 0 {
		item := heap.Pop(sellBook).(*Item)
		fmt.Printf("%v ", item)
	}
	fmt.Println("\nbuy book:")
	for buyBook.Len() > 0 {
		item := heap.Pop(buyBook).(*Item)
		fmt.Printf("%v ", item)
	}
}
