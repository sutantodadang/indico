-- name: InsertUser :exec
INSERT INTO i_users(user_id, full_name, email, password, role_id) VALUES($1, $2, $3, $4, $5);

-- name: SelectOneUserByEmail :one
SELECT user_id, full_name, email, password, role_id FROM i_users WHERE email = $1;

-- name: SelectOneUserById :one
SELECT a.user_id, a.full_name, a.email, a.password, b.name, b.unique_name
FROM i_users a
JOIN i_users_roles b ON b.user_role_id = a.role_id
WHERE user_id = $1;

-- name: SelectUserByRole :many
SELECT a.user_id, a.full_name, a.email, a.password, b.user_role_id, b.name, b.unique_name
FROM i_users a
JOIN i_users_roles b ON b.user_role_id = a.role_id
WHERE b.unique_name = $1;