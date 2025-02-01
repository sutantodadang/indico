-- name: InsertWarehouse :exec
INSERT INTO i_warehouses(warehouse_id, name, capacity) VALUES($1, $2, $3);

-- name: SelectWarehouses :many
SELECT warehouse_id, name, capacity FROM i_warehouses;
