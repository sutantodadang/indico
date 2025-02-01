package route

import (
	"indico/internal/app/user"
	"indico/internal/constants"
	"indico/internal/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoute(app *gin.Engine, handler user.IUserHandler, middleware *middlewares.Middleware) {
	userRoute := app.Group("/api/v1")

	userRoute.POST("/register", handler.RegisterUser)
	userRoute.POST("/login", handler.LoginUser)
	userRoute.GET("/users", middleware.Auth([]string{constants.ADMIN}), handler.UserAdmin)
	userRoute.GET("/users/me", middleware.Auth([]string{}), handler.InfoUser)

}
