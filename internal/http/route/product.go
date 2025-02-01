package route

import (
	"indico/internal/app/product"
	"indico/internal/constants"
	"indico/internal/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoute(app *gin.Engine, handler product.IProductHandler, middleware *middlewares.Middleware) {
	productRoute := app.Group("/api/v1")

	productRoute.POST("/products", middleware.Auth([]string{constants.ADMIN}), handler.CreateProduct)
	productRoute.GET("/products", middleware.Auth([]string{}), handler.ListProduct)
	productRoute.GET("/products/:id", middleware.Auth([]string{}), handler.GetProduct)
	productRoute.PUT("/products/:id", middleware.Auth([]string{constants.ADMIN}), handler.UpdateProduct)
	productRoute.DELETE("/products/:id", middleware.Auth([]string{constants.ADMIN}), handler.DeleteProduct)

}
