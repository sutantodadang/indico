package role

import (
	"indico/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type IRoleHandler interface {
	RegisterRole(c *gin.Context)
	ListRole(c *gin.Context)
}

type RoleHandler struct {
	service IRoleService
}

// ListRole implements IRoleHandler.
func (r *RoleHandler) ListRole(c *gin.Context) {

	role, err := r.service.ListRole(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "list role successfully", "data": role})
}

// RegisterRole implements IRoleHandler.
func (r *RoleHandler) RegisterRole(c *gin.Context) {

	req := new(RegisterRoleRequest)

	if err := c.ShouldBindJSON(&req); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(400, gin.H{"error": utils.NewValidationError(errs)})
			return
		}

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := r.service.AddRole(c, *req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "register role successfully"})
}

func NewRoleHandler(service IRoleService) IRoleHandler {
	return &RoleHandler{
		service: service,
	}
}
