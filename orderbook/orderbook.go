package orderbook

type Order struct {
	UserID     int32
	Type       string
	OrderType  string
	Amount     int32
	Price      int64
	Time       int64
	ResultChan chan OrderResult
}

type OrderResult struct {
	Success bool
	Message string
}
