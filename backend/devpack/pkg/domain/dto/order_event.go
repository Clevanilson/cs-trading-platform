package pkgdto

type OrderEvent struct {
	OrderID   string  `json:"order_id"`
	AccountID string  `json:"account_id"`
	MarketID  string  `json:"market_id"`
	Side      string  `json:"side"`
	Amount    float64 `json:"amount"`
	Price     float64 `json:"price"`
	Status    string  `json:"status"`
}
