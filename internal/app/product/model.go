package product

type CreateProductRequest struct {
	Name       string  `json:"name" binding:"required,min=3"`
	Quantity   int     `json:"quantity" binding:"required,min=1"`
	LocationId *string `json:"location_id"`
}

type Product struct {
	ProductID     string `json:"product_id"`
	Sku           string `json:"sku"`
	Name          string `json:"name"`
	Quantity      int    `json:"quantity"`
	WarehouseID   string `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name"`
	Capacity      int    `json:"capacity"`
}

type UpdateProductRequest struct {
	Name       string  `json:"name"`
	Sku        string  `json:"sku"`
	Quantity   int     `json:"quantity"`
	LocationId *string `json:"location_id"`
}
