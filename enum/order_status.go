package enum

type OrderStatus string

const (
	OrderStatusCreated    OrderStatus = "created"
	OrderStatusApproved   OrderStatus = "approved"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusOnDelivery OrderStatus = "on_delivery"
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusCompleted  OrderStatus = "completed"
)
