// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: product.sql

package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM i_products WHERE product_id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, productID pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteProduct, productID)
	return err
}

const insertProduct = `-- name: InsertProduct :exec
INSERT INTO i_products(product_id, sku, name, quantity, location_id) VALUES($1, $2, $3, $4, $5)
`

type InsertProductParams struct {
	ProductID  pgtype.UUID `db:"product_id" json:"product_id"`
	Sku        string      `db:"sku" json:"sku"`
	Name       string      `db:"name" json:"name"`
	Quantity   int32       `db:"quantity" json:"quantity"`
	LocationID pgtype.UUID `db:"location_id" json:"location_id"`
}

func (q *Queries) InsertProduct(ctx context.Context, arg InsertProductParams) error {
	_, err := q.db.Exec(ctx, insertProduct,
		arg.ProductID,
		arg.Sku,
		arg.Name,
		arg.Quantity,
		arg.LocationID,
	)
	return err
}

const selectOneProduct = `-- name: SelectOneProduct :one
SELECT a.product_id, a.sku, a.name, a.quantity, b.warehouse_id, b.name as warehouse_name, b.capacity
FROM i_products a
LEFT JOIN i_warehouses b ON b.warehouse_id = a.location_id WHERE a.product_id = $1
`

type SelectOneProductRow struct {
	ProductID     pgtype.UUID `db:"product_id" json:"product_id"`
	Sku           string      `db:"sku" json:"sku"`
	Name          string      `db:"name" json:"name"`
	Quantity      int32       `db:"quantity" json:"quantity"`
	WarehouseID   pgtype.UUID `db:"warehouse_id" json:"warehouse_id"`
	WarehouseName pgtype.Text `db:"warehouse_name" json:"warehouse_name"`
	Capacity      pgtype.Int4 `db:"capacity" json:"capacity"`
}

func (q *Queries) SelectOneProduct(ctx context.Context, productID pgtype.UUID) (SelectOneProductRow, error) {
	row := q.db.QueryRow(ctx, selectOneProduct, productID)
	var i SelectOneProductRow
	err := row.Scan(
		&i.ProductID,
		&i.Sku,
		&i.Name,
		&i.Quantity,
		&i.WarehouseID,
		&i.WarehouseName,
		&i.Capacity,
	)
	return i, err
}

const selectProducts = `-- name: SelectProducts :many
SELECT a.product_id, a.sku, a.name, a.quantity, b.warehouse_id, b.name as warehouse_name, b.capacity
FROM i_products a
LEFT JOIN i_warehouses b ON b.warehouse_id = a.location_id
`

type SelectProductsRow struct {
	ProductID     pgtype.UUID `db:"product_id" json:"product_id"`
	Sku           string      `db:"sku" json:"sku"`
	Name          string      `db:"name" json:"name"`
	Quantity      int32       `db:"quantity" json:"quantity"`
	WarehouseID   pgtype.UUID `db:"warehouse_id" json:"warehouse_id"`
	WarehouseName pgtype.Text `db:"warehouse_name" json:"warehouse_name"`
	Capacity      pgtype.Int4 `db:"capacity" json:"capacity"`
}

func (q *Queries) SelectProducts(ctx context.Context) ([]SelectProductsRow, error) {
	rows, err := q.db.Query(ctx, selectProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SelectProductsRow
	for rows.Next() {
		var i SelectProductsRow
		if err := rows.Scan(
			&i.ProductID,
			&i.Sku,
			&i.Name,
			&i.Quantity,
			&i.WarehouseID,
			&i.WarehouseName,
			&i.Capacity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :exec
UPDATE i_products SET sku = $2, name = $3, quantity = $4, location_id = $5, updated_at=now() WHERE product_id = $1
`

type UpdateProductParams struct {
	ProductID  pgtype.UUID `db:"product_id" json:"product_id"`
	Sku        string      `db:"sku" json:"sku"`
	Name       string      `db:"name" json:"name"`
	Quantity   int32       `db:"quantity" json:"quantity"`
	LocationID pgtype.UUID `db:"location_id" json:"location_id"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.db.Exec(ctx, updateProduct,
		arg.ProductID,
		arg.Sku,
		arg.Name,
		arg.Quantity,
		arg.LocationID,
	)
	return err
}

const updateProductQuantity = `-- name: UpdateProductQuantity :exec
UPDATE i_products SET quantity = $2, updated_at=now() WHERE product_id = $1
`

type UpdateProductQuantityParams struct {
	ProductID pgtype.UUID `db:"product_id" json:"product_id"`
	Quantity  int32       `db:"quantity" json:"quantity"`
}

func (q *Queries) UpdateProductQuantity(ctx context.Context, arg UpdateProductQuantityParams) error {
	_, err := q.db.Exec(ctx, updateProductQuantity, arg.ProductID, arg.Quantity)
	return err
}
