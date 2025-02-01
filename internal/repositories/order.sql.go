// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: order.sql

package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const insertOrder = `-- name: InsertOrder :exec
INSERT INTO i_orders (order_id, user_id, product_id, quantity, order_type, order_status) VALUES($1, $2, $3, $4, $5, $6)
`

type InsertOrderParams struct {
	OrderID     pgtype.UUID `db:"order_id" json:"order_id"`
	UserID      pgtype.UUID `db:"user_id" json:"user_id"`
	ProductID   pgtype.UUID `db:"product_id" json:"product_id"`
	Quantity    int32       `db:"quantity" json:"quantity"`
	OrderType   TypeOrder   `db:"order_type" json:"order_type"`
	OrderStatus int32       `db:"order_status" json:"order_status"`
}

func (q *Queries) InsertOrder(ctx context.Context, arg InsertOrderParams) error {
	_, err := q.db.Exec(ctx, insertOrder,
		arg.OrderID,
		arg.UserID,
		arg.ProductID,
		arg.Quantity,
		arg.OrderType,
		arg.OrderStatus,
	)
	return err
}

const selectOneOrder = `-- name: SelectOneOrder :one
SELECT order_id, user_id, product_id, quantity, order_type, order_status FROM i_orders WHERE order_id = $1
`

type SelectOneOrderRow struct {
	OrderID     pgtype.UUID `db:"order_id" json:"order_id"`
	UserID      pgtype.UUID `db:"user_id" json:"user_id"`
	ProductID   pgtype.UUID `db:"product_id" json:"product_id"`
	Quantity    int32       `db:"quantity" json:"quantity"`
	OrderType   TypeOrder   `db:"order_type" json:"order_type"`
	OrderStatus int32       `db:"order_status" json:"order_status"`
}

func (q *Queries) SelectOneOrder(ctx context.Context, orderID pgtype.UUID) (SelectOneOrderRow, error) {
	row := q.db.QueryRow(ctx, selectOneOrder, orderID)
	var i SelectOneOrderRow
	err := row.Scan(
		&i.OrderID,
		&i.UserID,
		&i.ProductID,
		&i.Quantity,
		&i.OrderType,
		&i.OrderStatus,
	)
	return i, err
}

const selectOrders = `-- name: SelectOrders :many
SELECT order_id, user_id, product_id, quantity, order_type, order_status FROM i_orders
`

type SelectOrdersRow struct {
	OrderID     pgtype.UUID `db:"order_id" json:"order_id"`
	UserID      pgtype.UUID `db:"user_id" json:"user_id"`
	ProductID   pgtype.UUID `db:"product_id" json:"product_id"`
	Quantity    int32       `db:"quantity" json:"quantity"`
	OrderType   TypeOrder   `db:"order_type" json:"order_type"`
	OrderStatus int32       `db:"order_status" json:"order_status"`
}

func (q *Queries) SelectOrders(ctx context.Context) ([]SelectOrdersRow, error) {
	rows, err := q.db.Query(ctx, selectOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SelectOrdersRow
	for rows.Next() {
		var i SelectOrdersRow
		if err := rows.Scan(
			&i.OrderID,
			&i.UserID,
			&i.ProductID,
			&i.Quantity,
			&i.OrderType,
			&i.OrderStatus,
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
