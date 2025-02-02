package order_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"indico/internal/app/order"
	"indico/internal/repositories"
	"indico/internal/repositories/mocks"
)

func TestGetOrder(t *testing.T) {
	mockRepo := new(mocks.MockRepo)
	service := order.NewOrderService(mockRepo)

	id := uuid.New()
	userId := uuid.New()
	productId := uuid.New()

	expectedOrder := order.Order{
		OrderId:     id.String(),
		UserId:      userId.String(),
		ProductId:   productId.String(),
		OrderStatus: 1,
		OrderType:   "RECEIVE",
		Quantity:    10,
	}

	returnData := repositories.SelectOneOrderRow{
		OrderID:     pgtype.UUID{Bytes: id, Valid: true},
		UserID:      pgtype.UUID{Bytes: userId, Valid: true},
		ProductID:   pgtype.UUID{Bytes: productId, Valid: true},
		OrderStatus: 1,
		OrderType:   "RECEIVE",
		Quantity:    10,
	}

	mockRepo.On("SelectOneOrder", mock.Anything, pgtype.UUID{Bytes: id, Valid: true}).Return(returnData, nil)

	order, err := service.GetOrder(context.Background(), id.String())
	assert.NoError(t, err)
	assert.Equal(t, expectedOrder, order)
	mockRepo.AssertExpectations(t)
}

func TestListOrder(t *testing.T) {
	mockRepo := new(mocks.MockRepo)
	service := order.NewOrderService(mockRepo)

	id := uuid.New()
	userId := uuid.New()
	productId := uuid.New()

	returnData := []repositories.SelectOrdersRow{
		{
			OrderID:     pgtype.UUID{Bytes: id, Valid: true},
			UserID:      pgtype.UUID{Bytes: userId, Valid: true},
			ProductID:   pgtype.UUID{Bytes: productId, Valid: true},
			OrderStatus: 1,
			OrderType:   "RECEIVE",
			Quantity:    10,
		},
	}

	mockRepo.On("SelectOrders", mock.Anything).Return(returnData, nil)

	orders, err := service.ListOrder(context.Background())
	assert.NoError(t, err)
	assert.Len(t, orders, 1)
	assert.Equal(t, id.String(), orders[0].OrderId)
	mockRepo.AssertExpectations(t)
}

func TestReceiveOrder_Success(t *testing.T) {
	mockRepo := new(mocks.MockRepo)
	service := order.NewOrderService(mockRepo)

	productId := uuid.New()
	userId := uuid.New()

	req := order.CreateOrderRequest{
		UserId:      userId.String(),
		ProductId:   productId.String(),
		OrderStatus: 1,
		Quantity:    10,
	}

	productData := repositories.SelectOneProductRow{
		ProductID:   pgtype.UUID{Bytes: productId, Valid: true},
		WarehouseID: pgtype.UUID{Bytes: uuid.New(), Valid: true},
		Quantity:    20,
		Capacity:    pgtype.Int4{Int32: 100, Valid: true},
	}

	mockRepo.On("SelectOneProduct", mock.Anything, pgtype.UUID{Bytes: productId, Valid: true}).Return(productData, nil)
	mockRepo.On("SelectSumQuantityProductsByLocation", mock.Anything, productData.WarehouseID).Return(int64(50), nil)
	mockRepo.On("InsertOrder", mock.Anything, mock.AnythingOfType("repositories.InsertOrderParams")).Return(nil)
	mockRepo.On("UpdateProductQuantity", mock.Anything, mock.AnythingOfType("repositories.UpdateProductQuantityParams")).Return(nil)

	err := service.ReceiveOrder(context.Background(), req)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestReceiveOrder_ExceededCapacity(t *testing.T) {
	mockRepo := new(mocks.MockRepo)
	service := order.NewOrderService(mockRepo)

	productId := uuid.New()
	userId := uuid.New()

	req := order.CreateOrderRequest{
		UserId:      userId.String(),
		ProductId:   productId.String(),
		OrderStatus: 1,
		Quantity:    10,
	}

	productData := repositories.SelectOneProductRow{
		ProductID:   pgtype.UUID{Bytes: productId, Valid: true},
		WarehouseID: pgtype.UUID{Bytes: uuid.New(), Valid: true},
		Quantity:    20,
		Capacity:    pgtype.Int4{Int32: 50, Valid: true},
	}

	mockRepo.On("SelectOneProduct", mock.Anything, pgtype.UUID{Bytes: productId, Valid: true}).Return(productData, nil)
	mockRepo.On("SelectSumQuantityProductsByLocation", mock.Anything, productData.WarehouseID).Return(int64(45), nil)

	err := service.ReceiveOrder(context.Background(), req)
	assert.Error(t, err)
	assert.Equal(t, "exceeded capacity", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestShipOrder_Success(t *testing.T) {
	mockRepo := new(mocks.MockRepo)
	service := order.NewOrderService(mockRepo)

	productId := uuid.New()
	userId := uuid.New()

	req := order.CreateOrderRequest{
		UserId:      userId.String(),
		ProductId:   productId.String(),
		OrderStatus: 1,
		Quantity:    10,
	}

	productData := repositories.SelectOneProductRow{
		ProductID:   pgtype.UUID{Bytes: productId, Valid: true},
		WarehouseID: pgtype.UUID{Bytes: uuid.New(), Valid: true},
		Quantity:    20,
		Capacity:    pgtype.Int4{Int32: 100, Valid: true},
	}

	mockRepo.On("SelectOneProduct", mock.Anything, pgtype.UUID{Bytes: productId, Valid: true}).Return(productData, nil)
	mockRepo.On("SelectSumQuantityProductsByLocation", mock.Anything, productData.WarehouseID).Return(int64(50), nil)
	mockRepo.On("InsertOrder", mock.Anything, mock.AnythingOfType("repositories.InsertOrderParams")).Return(nil)
	mockRepo.On("UpdateProductQuantity", mock.Anything, mock.AnythingOfType("repositories.UpdateProductQuantityParams")).Return(nil)

	err := service.ShipOrder(context.Background(), req)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestShipOrder_ExceededCapacity(t *testing.T) {
	mockRepo := new(mocks.MockRepo)
	service := order.NewOrderService(mockRepo)

	productId := uuid.New()
	userId := uuid.New()

	req := order.CreateOrderRequest{
		UserId:      userId.String(),
		ProductId:   productId.String(),
		OrderStatus: 1,
		Quantity:    30,
	}

	productData := repositories.SelectOneProductRow{
		ProductID:   pgtype.UUID{Bytes: productId, Valid: true},
		WarehouseID: pgtype.UUID{Bytes: uuid.New(), Valid: true},
		Quantity:    20,
		Capacity:    pgtype.Int4{Int32: 50, Valid: true},
	}

	mockRepo.On("SelectOneProduct", mock.Anything, pgtype.UUID{Bytes: productId, Valid: true}).Return(productData, nil)
	mockRepo.On("SelectSumQuantityProductsByLocation", mock.Anything, productData.WarehouseID).Return(int64(100), nil)

	err := service.ShipOrder(context.Background(), req)
	assert.Error(t, err)
	assert.Equal(t, "exceeded capacity", err.Error())
	mockRepo.AssertExpectations(t)
}
