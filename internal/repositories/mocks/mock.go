package mocks

import (
	"context"
	"indico/internal/repositories"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

// SelectSumQuantityProductsByLocation implements repositories.Querier.
func (m *MockRepo) SelectSumQuantityProductsByLocation(ctx context.Context, locationID pgtype.UUID) (int64, error) {
	args := m.Called(ctx, locationID)
	return args.Get(0).(int64), args.Error(1)
}

// SelectRoles implements repositories.Querier.
func (m *MockRepo) SelectRoles(ctx context.Context) ([]repositories.SelectRolesRow, error) {
	panic("unimplemented")
}

// InsertOrder implements repositories.Querier.
func (m *MockRepo) InsertOrder(ctx context.Context, arg repositories.InsertOrderParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
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
	args := m.Called(ctx, orderID)
	return args.Get(0).(repositories.SelectOneOrderRow), args.Error(1)
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
	args := m.Called(ctx)
	return args.Get(0).([]repositories.SelectOrdersRow), args.Error(1)
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
	args := m.Called(ctx, arg)
	return args.Error(0)
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
