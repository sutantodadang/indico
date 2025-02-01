package order

import (
	"indico/internal/constants"
	"indico/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type IOrderHandler interface {
	CreateOrderReceive(c *gin.Context)
	CreateOrderShip(c *gin.Context)
	ListOrder(c *gin.Context)
	GetOrder(c *gin.Context)
}

type OrderHandler struct {
	service IOrderService
}

// CreateOrderShip implements IOrderHandler.
func (o *OrderHandler) CreateOrderShip(c *gin.Context) {
	req := new(CreateOrderRequest)
	if err := c.ShouldBindJSON(&req); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(400, gin.H{"error": utils.NewValidationError(errs)})
			return
		}

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	val, ok := c.Get(constants.USER_ID)
	if !ok {
		c.JSON(400, gin.H{"error": "user id not found"})
		return
	}

	req.UserId = val.(string)

	err := o.service.ShipOrder(c, *req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "order ship created successfully"})
}

// GetOrder implements IOrderHandler.
func (o *OrderHandler) GetOrder(c *gin.Context) {

	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	data, err := o.service.GetOrder(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "order detail successfully", "data": data})
}

// ListOrder implements IOrderHandler.
func (o *OrderHandler) ListOrder(c *gin.Context) {

	data, err := o.service.ListOrder(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "list order successfully", "data": data})

}

// CreateOrder implements IOrderHandler.
func (o *OrderHandler) CreateOrderReceive(c *gin.Context) {
	req := new(CreateOrderRequest)
	if err := c.ShouldBindJSON(&req); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(400, gin.H{"error": utils.NewValidationError(errs)})
			return
		}

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	val, ok := c.Get(constants.USER_ID)
	if !ok {
		c.JSON(400, gin.H{"error": "user id not found"})
		return
	}

	req.UserId = val.(string)

	err := o.service.ReceiveOrder(c, *req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "order receive created successfully"})

}

func NewOrderHandler(service IOrderService) IOrderHandler {
	return &OrderHandler{
		service: service,
	}
}

// func (h *TodoHandler) CreateTodo(c *gin.Context) {

// 	var req CreateTodoRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {

// 		if errs, ok := err.(validator.ValidationErrors); ok {
// 			c.JSON(400, gin.H{"error": utils.NewValidationError(errs)})
// 			return
// 		}

// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	todo, err := h.service.CreateTodo(c, req)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(201, gin.H{"message": "Task created successfully", "task": todo})
// }

// func (h *TodoHandler) ListTodo(c *gin.Context) {

// 	var req ListTodoRequestParams
// 	if err := c.ShouldBindQuery(&req); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	todos, countData, currentPage, currentLimit, err := h.service.GetListTodos(c, req)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	var response struct {
// 		Tasks      []Todo `json:"tasks"`
// 		Pagination struct {
// 			CurrentPage int   `json:"current_page"`
// 			TotalPage   int   `json:"total_page"`
// 			TotalTasks  int64 `json:"total_tasks"`
// 		} `json:"pagination"`
// 	}

// 	response.Tasks = todos
// 	response.Pagination.CurrentPage = currentPage
// 	response.Pagination.TotalPage = int(countData) / currentLimit
// 	response.Pagination.TotalTasks = countData

// 	c.JSON(200, response)
// }

// func (h *TodoHandler) GetTodoById(c *gin.Context) {

// 	id := c.Param("id")

// 	if err := utils.ValidateId(id); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return

// 	}

// 	todo, err := h.service.GetTodo(c, id)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(200, todo)

// }

// func (h *TodoHandler) UpdateTodo(c *gin.Context) {

// 	id := c.Param("id")

// 	if err := utils.ValidateId(id); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return

// 	}

// 	var req UpdateTodoRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		if errs, ok := err.(validator.ValidationErrors); ok {
// 			c.JSON(400, gin.H{"error": utils.NewValidationError(errs)})
// 			return
// 		}

// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	todo, err := h.service.UpdateTodo(c, req, id)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(200, gin.H{"message": "Task updated successfully", "task": todo})
// }

// func (h *TodoHandler) DeleteTodo(c *gin.Context) {

// 	id := c.Param("id")

// 	if err := utils.ValidateId(id); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return

// 	}

// 	err := h.service.DeleteTodo(c, id)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(200, gin.H{"message": "Task deleted successfully"})
// }

// func (h *TodoHandler) GetToken(c *gin.Context) {

// 	token, err := h.service.GetToken(c)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(200, gin.H{"token": token})
// }
