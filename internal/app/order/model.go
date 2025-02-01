package order

type CreateOrderRequest struct {
	UserId      string `json:"user_id"`
	ProductId   string `json:"product_id" binding:"required"`
	OrderStatus int    `json:"order_status"  binding:"required"`
	Quantity    int    `json:"quantity"  binding:"required"`
}

type Order struct {
	OrderId     string `json:"order_id"`
	UserId      string `json:"user_id"`
	ProductId   string `json:"product_id"`
	OrderStatus int    `json:"order_status"`
	OrderType   string `json:"order_type"`
	Quantity    int    `json:"quantity"`
}
