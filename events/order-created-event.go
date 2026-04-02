package events

type OrderCreatedEvent struct {
	OrderID   uint    `json:"order_id"`
	UserID    uint    `json:"user_id"`
	Amount    float64 `json:"amount"`
	CreatedAt int64   `json:"created_at"`
}
