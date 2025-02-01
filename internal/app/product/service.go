package product

import (
	"context"
	"fmt"
	"indico/internal/repositories"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

type IProductService interface {
	AddProduct(ctx context.Context, req CreateProductRequest) (err error)
	GetListProduct(ctx context.Context) (list []Product, err error)
	GetProduct(ctx context.Context, id string) (product Product, err error)
	UpdateProduct(ctx context.Context, id string, req UpdateProductRequest) (err error)
	DeleteProduct(ctx context.Context, id string) (err error)
}

type ProductService struct {
	repo repositories.Querier
}

// DeleteProduct implements IProductService.
func (p *ProductService) DeleteProduct(ctx context.Context, id string) (err error) {

	err = p.repo.DeleteProduct(ctx, pgtype.UUID{Bytes: uuid.MustParse(id), Valid: true})
	if err != nil {
		log.Error().Err(err).Send()
		return
	}
	return
}

// UpdateProduct implements IProductService.
func (p *ProductService) UpdateProduct(ctx context.Context, id string, req UpdateProductRequest) (err error) {

	currentData, err := p.repo.SelectOneProduct(ctx, pgtype.UUID{Bytes: uuid.MustParse(id), Valid: true})
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	if req.Name == "" {
		req.Name = currentData.Name
	}

	if req.Sku == "" {
		req.Sku = currentData.Sku
	}

	if req.Quantity == 0 {
		req.Quantity = int(currentData.Quantity)
	}

	var locationId pgtype.UUID
	if req.LocationId != nil {

		idLoc, errP := uuid.Parse(*req.LocationId)
		if errP != nil {
			log.Error().Err(errP).Send()
			err = errP
			return
		}

		locationId = pgtype.UUID{Bytes: idLoc, Valid: true}

	} else {

		if currentData.WarehouseID.Valid {

			locationId = pgtype.UUID{Bytes: currentData.WarehouseID.Bytes, Valid: true}
		} else {
			locationId = pgtype.UUID{Bytes: [16]byte{}, Valid: false}
		}

	}

	err = p.repo.UpdateProduct(ctx, repositories.UpdateProductParams{
		ProductID:  pgtype.UUID{Bytes: uuid.MustParse(id), Valid: true},
		Sku:        req.Sku,
		Name:       req.Name,
		Quantity:   int32(req.Quantity),
		LocationID: locationId,
	})
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	return

}

// GetProduct implements IProductService.
func (p *ProductService) GetProduct(ctx context.Context, id string) (product Product, err error) {

	data, err := p.repo.SelectOneProduct(ctx, pgtype.UUID{Bytes: uuid.MustParse(id), Valid: true})
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	product = Product{
		ProductID:     data.ProductID.String(),
		Sku:           data.Sku,
		Name:          data.Name,
		Quantity:      int(data.Quantity),
		WarehouseID:   data.WarehouseID.String(),
		WarehouseName: data.WarehouseName.String,
		Capacity:      int(data.Capacity.Int32),
	}
	return

}

// GetListProduct implements IProductService.
func (p *ProductService) GetListProduct(ctx context.Context) (list []Product, err error) {

	data, err := p.repo.SelectProducts(ctx)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	for _, v := range data {
		product := Product{
			ProductID:     v.ProductID.String(),
			Sku:           v.Sku,
			Name:          v.Name,
			Quantity:      int(v.Quantity),
			WarehouseID:   v.WarehouseID.String(),
			WarehouseName: v.WarehouseName.String,
			Capacity:      int(v.Capacity.Int32),
		}
		list = append(list, product)
	}

	return
}

// AddProduct implements IProductService.
func (p *ProductService) AddProduct(ctx context.Context, req CreateProductRequest) (err error) {

	id, err := uuid.NewV7()
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	var location pgtype.UUID

	if req.LocationId != nil {
		location = pgtype.UUID{Bytes: uuid.MustParse(*req.LocationId), Valid: true}
	}

	sku := fmt.Sprintf("%s-%s-%04d", strings.ToUpper(req.Name[:3]), time.Now().Format("20060102"), req.Quantity)

	err = p.repo.InsertProduct(ctx, repositories.InsertProductParams{
		ProductID:  pgtype.UUID{Bytes: id, Valid: true},
		Sku:        sku,
		Name:       req.Name,
		Quantity:   int32(req.Quantity),
		LocationID: location,
	},
	)

	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	return
}

func NewProductService(repo repositories.Querier) IProductService {
	return &ProductService{
		repo: repo,
	}
}
