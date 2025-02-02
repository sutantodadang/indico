package product_test

import (
	"context"
	"errors"
	"indico/internal/app/product"
	"indico/internal/repositories"
	"indico/internal/repositories/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddProduct(t *testing.T) {
	mockRepo := new(mocks.MockRepo)
	service := product.NewProductService(mockRepo)

	req := product.CreateProductRequest{
		Name:     "Test Product",
		Quantity: 10,
	}

	mockRepo.On("InsertProduct", mock.Anything, mock.Anything).Return(nil)

	err := service.AddProduct(context.Background(), req)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetProduct(t *testing.T) {
	mockRepo := new(mocks.MockRepo)
	service := product.NewProductService(mockRepo)

	id := uuid.New()
	expectedProduct := product.Product{
		ProductID: id.String(),
		Name:      "Test Product",
		Sku:       "SKU123",
		Quantity:  20,
	}

	returnData := repositories.SelectOneProductRow{
		ProductID: pgtype.UUID{Bytes: id, Valid: true},
		Name:      "Test Product",
		Sku:       "SKU123",
		Quantity:  20,
	}

	mockRepo.On("SelectOneProduct", mock.Anything, pgtype.UUID{Bytes: id, Valid: true}).Return(returnData, nil)

	product, err := service.GetProduct(context.Background(), id.String())
	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, product)
	mockRepo.AssertExpectations(t)
}

func TestUpdateProduct_Error(t *testing.T) {
	mockRepo := new(mocks.MockRepo)
	service := product.NewProductService(mockRepo)

	id := uuid.New()

	returnData := repositories.SelectOneProductRow{
		ProductID: pgtype.UUID{Bytes: id, Valid: true},
		Name:      "Test Product",
		Sku:       "SKU123",
		Quantity:  20,
	}

	mockRepo.On("SelectOneProduct", mock.Anything, pgtype.UUID{Bytes: id, Valid: true}).Return(returnData, nil)
	mockRepo.On("UpdateProduct", mock.Anything, mock.Anything).Return(errors.New("update failed"))

	err := service.UpdateProduct(context.Background(), id.String(), product.UpdateProductRequest{Name: "New Product"})
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteProduct(t *testing.T) {
	mockRepo := new(mocks.MockRepo)
	service := product.NewProductService(mockRepo)

	id := uuid.New()
	mockRepo.On("DeleteProduct", mock.Anything, pgtype.UUID{Bytes: id, Valid: true}).Return(nil)

	err := service.DeleteProduct(context.Background(), id.String())
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
