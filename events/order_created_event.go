package events

import "github.com/thangtranitwork/kafka-demo-definitions/enum"

type OrderItemCreated struct {
	ItemID   uint `json:"item_id"`
	Quantity int  `json:"quantity"`
}

type OrderCreatedEvent struct {
	OrderID   uint               `json:"order_id"`
	UserID    uint               `json:"user_id"`
	Status    enum.OrderStatus   `json:"status"`
	CreatedAt int64              `json:"created_at"`
	Items     []OrderItemCreated `json:"items"`
}
