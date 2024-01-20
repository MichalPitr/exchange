package orderbook

import (
	"container/heap"
	"log"
)

type Order struct {
	Id         uint64
	UserID     int32
	Type       string // BUY or SELL
	OrderType  string // MARKET or LIMIT
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
	old := *b
	n := len(old.orders)
	item := old.orders[n-1]
	old.orders[n-1] = nil // avoid memory leak
	item.index = -1       // for safety
	(*b).orders = old.orders[0 : n-1]
	return item
}

func New(asc bool) *Book {
	b := &Book{asc: asc}
	heap.Init(b)
	return b
}

func (b Book) Peek() (*Order, bool) {
	if b.Len() == 0 {
		log.Println("Peeking empty orderbook.")
		return &Order{}, false
	}
	return &b.orders[0].Order, true
}
