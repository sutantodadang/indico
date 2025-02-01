package product_test

import (
	"context"
	"errors"
	"indico/internal/app/product"
	"indico/internal/repositories"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

// SelectSumQuantityProductsByLocation implements repositories.Querier.
func (m *MockRepo) SelectSumQuantityProductsByLocation(ctx context.Context, locationID pgtype.UUID) (int64, error) {
	panic("unimplemented")
}

// SelectRoles implements repositories.Querier.
func (m *MockRepo) SelectRoles(ctx context.Context) ([]repositories.SelectRolesRow, error) {
	panic("unimplemented")
}

// InsertOrder implements repositories.Querier.
func (m *MockRepo) InsertOrder(ctx context.Context, arg repositories.InsertOrderParams) error {
	panic("unimplemented")
}

// InsertRole implements repositories.Querier.
func (m *MockRepo) InsertRole(ctx context.Context, arg repositories.InsertRoleParams) error {
	panic("unimplemented")
}

// InsertUser implements repositories.Querier.
func (m *MockRepo) InsertUser(ctx context.Context, arg repositories.InsertUserParams) error {
	panic("unimplemented")
}

// InsertWarehouse implements repositories.Querier.
func (m *MockRepo) InsertWarehouse(ctx context.Context, arg repositories.InsertWarehouseParams) error {
	panic("unimplemented")
}

// SelectOneOrder implements repositories.Querier.
func (m *MockRepo) SelectOneOrder(ctx context.Context, orderID pgtype.UUID) (repositories.SelectOneOrderRow, error) {
	panic("unimplemented")
}

// SelectOneUserByEmail implements repositories.Querier.
func (m *MockRepo) SelectOneUserByEmail(ctx context.Context, email string) (repositories.SelectOneUserByEmailRow, error) {
	panic("unimplemented")
}

// SelectOneUserById implements repositories.Querier.
func (m *MockRepo) SelectOneUserById(ctx context.Context, userID pgtype.UUID) (repositories.SelectOneUserByIdRow, error) {
	panic("unimplemented")
}

// SelectOneUserByRoleId implements repositories.Querier.
func (m *MockRepo) SelectOneUserByRoleId(ctx context.Context, userRoleID pgtype.UUID) (repositories.SelectOneUserByRoleIdRow, error) {
	panic("unimplemented")
}

// SelectOrders implements repositories.Querier.
func (m *MockRepo) SelectOrders(ctx context.Context) ([]repositories.SelectOrdersRow, error) {
	panic("unimplemented")
}

// SelectUserByRole implements repositories.Querier.
func (m *MockRepo) SelectUserByRole(ctx context.Context, uniqueName repositories.UserRole) ([]repositories.SelectUserByRoleRow, error) {
	panic("unimplemented")
}

// SelectWarehouses implements repositories.Querier.
func (m *MockRepo) SelectWarehouses(ctx context.Context) ([]repositories.SelectWarehousesRow, error) {
	panic("unimplemented")
}

// UpdateProductQuantity implements repositories.Querier.
func (m *MockRepo) UpdateProductQuantity(ctx context.Context, arg repositories.UpdateProductQuantityParams) error {
	panic("unimplemented")
}

func (m *MockRepo) InsertProduct(ctx context.Context, params repositories.InsertProductParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *MockRepo) SelectProducts(ctx context.Context) ([]repositories.SelectProductsRow, error) {
	args := m.Called(ctx)
	return args.Get(0).([]repositories.SelectProductsRow), args.Error(1)
}

func (m *MockRepo) SelectOneProduct(ctx context.Context, id pgtype.UUID) (repositories.SelectOneProductRow, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(repositories.SelectOneProductRow), args.Error(1)
}

func (m *MockRepo) UpdateProduct(ctx context.Context, params repositories.UpdateProductParams) error {
	args := m.Called(ctx, params)
	return args.Error(0)
}

func (m *MockRepo) DeleteProduct(ctx context.Context, id pgtype.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestAddProduct(t *testing.T) {
	mockRepo := new(MockRepo)
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
	mockRepo := new(MockRepo)
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
	mockRepo := new(MockRepo)
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
	mockRepo := new(MockRepo)
	service := product.NewProductService(mockRepo)

	id := uuid.New()
	mockRepo.On("DeleteProduct", mock.Anything, pgtype.UUID{Bytes: id, Valid: true}).Return(nil)

	err := service.DeleteProduct(context.Background(), id.String())
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
