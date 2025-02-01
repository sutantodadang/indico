package main

import (
	"context"
	"indico/database"
	"indico/internal/app/order"
	"indico/internal/app/product"
	"indico/internal/app/role"
	"indico/internal/app/user"
	"indico/internal/app/warehouse"
	"indico/internal/http/middlewares"
	"indico/internal/http/route"
	"indico/internal/repositories"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = zerolog.New(os.Stdout).With().Caller().Timestamp().Logger()
}

func main() {

	app := gin.Default()

	db := database.ConnectPG()

	defer db.Close()

	app.Use(middlewares.Trace())
	app.Use(middlewares.RequestLoggerMiddleware(), middlewares.ResponseLoggerMiddleware())

	setupContainer(app, db)

	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: app,
	}

	go func() {
		log.Info().Msgf("Starting server... on port %s", os.Getenv("PORT"))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Server failed")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exiting")

}

func setupContainer(app *gin.Engine, db *pgxpool.Pool) {

	repo := repositories.New(db)

	middleware := middlewares.NewMiddleware(repo)

	userService := user.NewUserService(repo)
	productService := product.NewProductService(repo)
	warehouseService := warehouse.NewWarehouseService(repo)
	orderService := order.NewOrderService(repo)
	roleService := role.NewRoleService(repo)

	orderHandler := order.NewOrderHandler(orderService)
	productHandler := product.NewProductHandler(productService)
	warehouseHandler := warehouse.NewWarehouseHandler(warehouseService)
	userHandler := user.NewUserHandler(userService)
	roleHandler := role.NewRoleHandler(roleService)

	route.RegisterRoleRoute(app, roleHandler, middleware)
	route.RegisterUserRoute(app, userHandler, middleware)
	route.RegisterProductRoute(app, productHandler, middleware)
	route.RegisterWarehouseRoute(app, warehouseHandler, middleware)
	route.RegisterOrderRoute(app, orderHandler, middleware)

}
