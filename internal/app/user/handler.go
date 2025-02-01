package user

import (
	"errors"
	"indico/internal/constants"
	"indico/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type IUserHandler interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	InfoUser(c *gin.Context)
	UserAdmin(c *gin.Context)
}

type UserHandler struct {
	service IUserService
}

// UserAdmin implements IUserHandler.
func (u *UserHandler) UserAdmin(c *gin.Context) {

	userAdmin, err := u.service.GetUserAdmin(c, constants.ADMIN)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "get user admin successfully", "data": userAdmin})
}

// InfoUser implements IUserHandler.
func (u *UserHandler) InfoUser(c *gin.Context) {

	val, ok := c.Get(constants.USER_ID)
	if !ok {
		c.JSON(400, gin.H{"error": errors.New("user id not found")})

		return
	}

	info, err := u.service.GetUserInfo(c, val.(string))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "get user info successfully", "data": info})

}

// LoginUser implements IUserHandler.
func (u *UserHandler) LoginUser(c *gin.Context) {
	req := new(LoginUserRequest)

	if err := c.ShouldBindJSON(&req); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(400, gin.H{"error": utils.NewValidationError(errs)})
			return
		}

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := u.service.LoginUser(c, *req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "login successfully", "token": token})

}

// RegisterUser implements IUserHandler.
func (u *UserHandler) RegisterUser(c *gin.Context) {

	req := new(RegisterUserRequest)

	if err := c.ShouldBindJSON(&req); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(400, gin.H{"error": utils.NewValidationError(errs)})
			return
		}

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := u.service.RegisterUser(c, *req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "register successfully"})
}

func NewUserHandler(service IUserService) IUserHandler {
	return &UserHandler{
		service: service,
	}
}
