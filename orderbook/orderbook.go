package orderbook

import (
	"container/heap"
	"fmt"
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
	order Order

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
		if b.orders[i].order.Price == b.orders[j].order.Price {
			return b.orders[i].order.Time < b.orders[j].order.Time
		}
		return b.orders[i].order.Price < b.orders[j].order.Price
	}

	if b.orders[i].order.Price == b.orders[j].order.Price {
		return b.orders[i].order.Time < b.orders[j].order.Time
	}
	return b.orders[i].order.Price > b.orders[j].order.Price
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

func main() {
	sellBook := New(true)
	buyBook := New(false)

	// Example orders
	orders := []Order{
		{
			UserID:     1,
			Type:       "Limit",
			OrderType:  "Sell",
			Amount:     10,
			Price:      100,
			Time:       1641016800, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
		{
			UserID:     2,
			Type:       "Limit",
			OrderType:  "Sell",
			Amount:     5,
			Price:      200,
			Time:       1641103200, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
		{
			UserID:     3,
			Type:       "Limit",
			OrderType:  "Sell",
			Amount:     15,
			Price:      50,
			Time:       1641189600, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
		{
			UserID:     1,
			Type:       "Limit",
			OrderType:  "Buy",
			Amount:     10,
			Price:      100,
			Time:       1641016800, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
		{
			UserID:     2,
			Type:       "Limit",
			OrderType:  "Buy",
			Amount:     5,
			Price:      200,
			Time:       1641103200, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
		{
			UserID:     3,
			Type:       "Limit",
			OrderType:  "Buy",
			Amount:     15,
			Price:      50,
			Time:       1641189600, // Example Unix timestamp
			ResultChan: nil,        // Or initialize as appropriate
		},
	}

	// Add items to both heaps
	for _, order := range orders {
		if order.OrderType == "Sell" {
			heap.Push(sellBook, Item{order: order})
		} else {
			heap.Push(buyBook, Item{order: order})
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
