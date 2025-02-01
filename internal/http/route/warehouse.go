package route

import (
	"indico/internal/app/warehouse"
	"indico/internal/constants"
	"indico/internal/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterWarehouseRoute(app *gin.Engine, handler warehouse.IWarehouseHandler, middleware *middlewares.Middleware) {
	warehouseRoute := app.Group("/api/v1")

	warehouseRoute.POST("/locations", middleware.Auth([]string{constants.ADMIN}), handler.RegisterWarehouse)
	warehouseRoute.GET("/locations", middleware.Auth([]string{}), handler.ListWarehouse)

}
