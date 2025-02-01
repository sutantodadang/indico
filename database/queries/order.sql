-- name: InsertOrder :exec
INSERT INTO i_orders (order_id, user_id, product_id, quantity, order_type, order_status) VALUES($1, $2, $3, $4, $5, $6);

-- name: SelectOrders :many
SELECT order_id, user_id, product_id, quantity, order_type, order_status FROM i_orders;

-- name: SelectOneOrder :one
SELECT order_id, user_id, product_id, quantity, order_type, order_status FROM i_orders WHERE order_id = $1;