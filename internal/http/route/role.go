package route

import (
	"indico/internal/app/role"
	"indico/internal/constants"
	"indico/internal/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoleRoute(app *gin.Engine, handler role.IRoleHandler, middleware *middlewares.Middleware) {
	roleRoute := app.Group("/api/v1")

	roleRoute.POST("/roles", middleware.Auth([]string{constants.ADMIN}), handler.RegisterRole)
	roleRoute.GET("/roles", handler.ListRole)

}
