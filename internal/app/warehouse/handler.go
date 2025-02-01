package warehouse

import (
	"indico/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type IWarehouseHandler interface {
	RegisterWarehouse(c *gin.Context)
	ListWarehouse(c *gin.Context)
}

type WarehouseHandler struct {
	service IWarehouseService
}

// ListWarehouse implements IWarehouseHandler.
func (w *WarehouseHandler) ListWarehouse(c *gin.Context) {

	data, err := w.service.ListWarehouse(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "warehouse list successfully", "data": data})
}

// RegisterWarehouse implements IWarehouseHandler.
func (w *WarehouseHandler) RegisterWarehouse(c *gin.Context) {

	req := new(RegisterWarehouseRequest)
	if err := c.ShouldBindJSON(&req); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(400, gin.H{"error": utils.NewValidationError(errs)})
			return
		}

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := w.service.AddWarehouse(c, *req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "warehouse created successfully"})

}

func NewWarehouseHandler(service IWarehouseService) IWarehouseHandler {
	return &WarehouseHandler{
		service: service,
	}
}
