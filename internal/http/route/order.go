package route

import (
	"indico/internal/app/order"
	"indico/internal/constants"
	"indico/internal/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoute(app *gin.Engine, handler order.IOrderHandler, middleware *middlewares.Middleware) {
	orderRoute := app.Group("/api/v1")

	orderRoute.POST("/orders/receive", middleware.Auth([]string{constants.STAFF}), handler.CreateOrderReceive)
	orderRoute.POST("/orders/ship", middleware.Auth([]string{constants.STAFF}), handler.CreateOrderShip)
	orderRoute.GET("/orders", middleware.Auth([]string{}), handler.ListOrder)
	orderRoute.GET("/orders/:id", middleware.Auth([]string{}), handler.GetOrder)

}
