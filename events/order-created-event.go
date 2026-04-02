package events

type OrderCreatedEvent struct {
	OrderID   int     `json:"order_id"`
	UserID    int     `json:"user_id"`
	Amount    float64 `json:"amount"`
	CreatedAt int64   `json:"created_at"`
}
