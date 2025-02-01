package warehouse

type RegisterWarehouseRequest struct {
	Name     string `json:"name" binding:"required"`
	Capacity int    `json:"capacity" binding:"required"`
}

type Warehouse struct {
	WarehouseId string `json:"warehouse_id"`
	Name        string `json:"name"`
	Capacity    int    `json:"capacity"`
}
