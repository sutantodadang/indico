package order

import (
	"context"
	"fmt"
	"indico/internal/repositories"
	"sync"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

type IOrderService interface {
	ReceiveOrder(ctx context.Context, req CreateOrderRequest) (err error)
	ShipOrder(ctx context.Context, req CreateOrderRequest) (err error)
	ListOrder(ctx context.Context) (order []Order, err error)
	GetOrder(ctx context.Context, id string) (order Order, err error)
}

type OrderService struct {
	repo repositories.Querier
	wg   sync.WaitGroup
	mut  sync.Mutex
}

// GetOrder implements IOrderService.
func (o *OrderService) GetOrder(ctx context.Context, id string) (order Order, err error) {

	dataOrder, err := o.repo.SelectOneOrder(ctx, pgtype.UUID{Bytes: uuid.MustParse(id), Valid: true})
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	order = Order{
		OrderId:     dataOrder.OrderID.String(),
		UserId:      dataOrder.UserID.String(),
		ProductId:   dataOrder.ProductID.String(),
		OrderStatus: int(dataOrder.OrderStatus),
		OrderType:   string(dataOrder.OrderType),
		Quantity:    int(dataOrder.Quantity),
	}
	return
}

// ListOrder implements IOrderService.
func (o *OrderService) ListOrder(ctx context.Context) (order []Order, err error) {

	dataOrder, err := o.repo.SelectOrders(ctx)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	for _, data := range dataOrder {
		order = append(order, Order{
			OrderId:     data.OrderID.String(),
			UserId:      data.UserID.String(),
			ProductId:   data.ProductID.String(),
			OrderStatus: int(data.OrderStatus),
			OrderType:   string(data.OrderType),
			Quantity:    int(data.Quantity),
		})
	}
	return

}

// ShipOrder implements IOrderService.
func (o *OrderService) ShipOrder(ctx context.Context, req CreateOrderRequest) (err error) {

	errChan := make(chan error, 1)

	o.wg.Add(1)

	go func(ctx context.Context, req CreateOrderRequest) {
		defer o.wg.Done()
		defer close(errChan)

		o.mut.Lock()
		defer o.mut.Unlock()

		defer func() {
			if r := recover(); r != nil {
				log.Error().Interface("", r).Send()
				errChan <- fmt.Errorf("panic: %v", r)
				return
			}
		}()

		id, err := uuid.NewV7()
		if err != nil {
			log.Error().Err(err).Send()
			errChan <- err
			return
		}

		dataProduct, err := o.repo.SelectOneProduct(ctx, pgtype.UUID{Bytes: uuid.MustParse(req.ProductId), Valid: true})
		if err != nil {
			log.Error().Err(err).Send()
			errChan <- err
			return
		}

		err = o.repo.InsertOrder(ctx, repositories.InsertOrderParams{
			OrderID:     pgtype.UUID{Bytes: id, Valid: true},
			UserID:      pgtype.UUID{Bytes: uuid.MustParse(req.UserId), Valid: true},
			OrderType:   repositories.TypeOrderSHIP,
			OrderStatus: int32(req.OrderStatus),
			ProductID:   dataProduct.ProductID,
			Quantity:    int32(req.Quantity),
		})
		if err != nil {
			log.Error().Err(err).Send()
			errChan <- err
			return
		}

		err = o.repo.UpdateProductQuantity(ctx, repositories.UpdateProductQuantityParams{
			ProductID: dataProduct.ProductID,
			Quantity:  int32(int(dataProduct.Quantity) - req.Quantity),
		})
		if err != nil {
			log.Error().Err(err).Send()
			errChan <- err
			return
		}

		errChan <- nil
	}(ctx, req)

	// Wait for either context cancellation or goroutine completion
	select {
	case err = <-errChan:
		return err
	case <-ctx.Done():
		o.wg.Wait() // Wait for goroutine to finish
		return ctx.Err()
	}

}

// ReceiveOrder implements IOrderService.
func (o *OrderService) ReceiveOrder(ctx context.Context, req CreateOrderRequest) (err error) {

	errChan := make(chan error, 1)

	o.wg.Add(1)

	go func(ctx context.Context, req CreateOrderRequest) {
		defer o.wg.Done()
		defer close(errChan)

		o.mut.Lock()
		defer o.mut.Unlock()

		defer func() {
			if r := recover(); r != nil {
				log.Error().Interface("", r).Send()
				errChan <- fmt.Errorf("panic: %v", r)
				return
			}
		}()

		id, err := uuid.NewV7()
		if err != nil {
			log.Error().Err(err).Send()
			errChan <- err
			return
		}

		dataProduct, err := o.repo.SelectOneProduct(ctx, pgtype.UUID{Bytes: uuid.MustParse(req.ProductId), Valid: true})
		if err != nil {
			log.Error().Err(err).Send()
			errChan <- err
			return
		}

		err = o.repo.InsertOrder(ctx, repositories.InsertOrderParams{
			OrderID:     pgtype.UUID{Bytes: id, Valid: true},
			UserID:      pgtype.UUID{Bytes: uuid.MustParse(req.UserId), Valid: true},
			OrderType:   repositories.TypeOrderRECEIVE,
			OrderStatus: int32(req.OrderStatus),
			ProductID:   dataProduct.ProductID,
			Quantity:    int32(req.Quantity),
		})
		if err != nil {
			log.Error().Err(err).Send()
			errChan <- err
			return
		}

		err = o.repo.UpdateProductQuantity(ctx, repositories.UpdateProductQuantityParams{
			ProductID: dataProduct.ProductID,
			Quantity:  int32(int(dataProduct.Quantity) + req.Quantity),
		})
		if err != nil {
			log.Error().Err(err).Send()
			errChan <- err
			return
		}

		errChan <- nil
	}(ctx, req)

	select {
	case err = <-errChan:
		return err
	case <-ctx.Done():
		o.wg.Wait()
		return ctx.Err()
	}

}

func NewOrderService(repo repositories.Querier) IOrderService {
	return &OrderService{
		repo: repo,
	}
}
