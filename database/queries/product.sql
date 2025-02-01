-- name: InsertProduct :exec
INSERT INTO i_products(product_id, sku, name, quantity, location_id) VALUES($1, $2, $3, $4, $5);

-- name: SelectProducts :many
SELECT a.product_id, a.sku, a.name, a.quantity, b.warehouse_id, b.name as warehouse_name, b.capacity
FROM i_products a
LEFT JOIN i_warehouses b ON b.warehouse_id = a.location_id;

-- name: SelectOneProduct :one
SELECT a.product_id, a.sku, a.name, a.quantity, b.warehouse_id, b.name as warehouse_name, b.capacity
FROM i_products a
LEFT JOIN i_warehouses b ON b.warehouse_id = a.location_id WHERE a.product_id = $1;

-- name: UpdateProduct :exec
UPDATE i_products SET sku = $2, name = $3, quantity = $4, location_id = $5, updated_at=now() WHERE product_id = $1;

-- name: UpdateProductQuantity :exec
UPDATE i_products SET quantity = $2, updated_at=now() WHERE product_id = $1;

-- name: DeleteProduct :exec
DELETE FROM i_products WHERE product_id = $1;