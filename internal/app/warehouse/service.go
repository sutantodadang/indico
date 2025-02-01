package warehouse

import (
	"context"
	"indico/internal/repositories"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

type IWarehouseService interface {
	AddWarehouse(ctx context.Context, req RegisterWarehouseRequest) (err error)
	ListWarehouse(ctx context.Context) (warehouse []Warehouse, err error)
}

type WarehouseService struct {
	repo repositories.Querier
}

// ListWarehouse implements IWarehouseService.
func (w *WarehouseService) ListWarehouse(ctx context.Context) (warehouse []Warehouse, err error) {

	warehouseData, err := w.repo.SelectWarehouses(ctx)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	for _, v := range warehouseData {

		warehouse = append(warehouse, Warehouse{
			WarehouseId: v.WarehouseID.String(),
			Name:        v.Name,
			Capacity:    int(v.Capacity),
		})
	}

	return
}

// AddWarehouse implements IWarehouseService.
func (w *WarehouseService) AddWarehouse(ctx context.Context, req RegisterWarehouseRequest) (err error) {

	id, err := uuid.NewV7()
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	err = w.repo.InsertWarehouse(ctx, repositories.InsertWarehouseParams{
		WarehouseID: pgtype.UUID{Bytes: id, Valid: true},
		Name:        req.Name,
		Capacity:    int32(req.Capacity),
	})
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	return

}

func NewWarehouseService(repo repositories.Querier) IWarehouseService {
	return &WarehouseService{
		repo: repo,
	}
}
