-- name: InsertRole :exec
INSERT INTO i_users_roles (user_role_id, unique_name, name) VALUES($1, $2, $3);

-- name: SelectOneUserByRoleId :one
SELECT user_role_id, unique_name, name, status FROM i_users_roles WHERE user_role_id = $1;

-- name: SelectRoles :many
SELECT user_role_id, unique_name, name, status FROM i_users_roles;