package events

type OrderItemApproved struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
type OrderApprovedEvent struct {
	OrderID uint `json:"order_id"`
	Items   []OrderItemApproved
}
