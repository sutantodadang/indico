// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const insertUser = `-- name: InsertUser :exec
INSERT INTO i_users(user_id, full_name, email, password, role_id) VALUES($1, $2, $3, $4, $5)
`

type InsertUserParams struct {
	UserID   pgtype.UUID `db:"user_id" json:"user_id"`
	FullName string      `db:"full_name" json:"full_name"`
	Email    string      `db:"email" json:"email"`
	Password string      `db:"password" json:"password"`
	RoleID   pgtype.UUID `db:"role_id" json:"role_id"`
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) error {
	_, err := q.db.Exec(ctx, insertUser,
		arg.UserID,
		arg.FullName,
		arg.Email,
		arg.Password,
		arg.RoleID,
	)
	return err
}

const selectOneUserByEmail = `-- name: SelectOneUserByEmail :one
SELECT user_id, full_name, email, password, role_id FROM i_users WHERE email = $1
`

type SelectOneUserByEmailRow struct {
	UserID   pgtype.UUID `db:"user_id" json:"user_id"`
	FullName string      `db:"full_name" json:"full_name"`
	Email    string      `db:"email" json:"email"`
	Password string      `db:"password" json:"password"`
	RoleID   pgtype.UUID `db:"role_id" json:"role_id"`
}

func (q *Queries) SelectOneUserByEmail(ctx context.Context, email string) (SelectOneUserByEmailRow, error) {
	row := q.db.QueryRow(ctx, selectOneUserByEmail, email)
	var i SelectOneUserByEmailRow
	err := row.Scan(
		&i.UserID,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.RoleID,
	)
	return i, err
}

const selectOneUserById = `-- name: SelectOneUserById :one
SELECT a.user_id, a.full_name, a.email, a.password, b.name, b.unique_name
FROM i_users a
JOIN i_users_roles b ON b.user_role_id = a.role_id
WHERE user_id = $1
`

type SelectOneUserByIdRow struct {
	UserID     pgtype.UUID `db:"user_id" json:"user_id"`
	FullName   string      `db:"full_name" json:"full_name"`
	Email      string      `db:"email" json:"email"`
	Password   string      `db:"password" json:"password"`
	Name       string      `db:"name" json:"name"`
	UniqueName UserRole    `db:"unique_name" json:"unique_name"`
}

func (q *Queries) SelectOneUserById(ctx context.Context, userID pgtype.UUID) (SelectOneUserByIdRow, error) {
	row := q.db.QueryRow(ctx, selectOneUserById, userID)
	var i SelectOneUserByIdRow
	err := row.Scan(
		&i.UserID,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.Name,
		&i.UniqueName,
	)
	return i, err
}

const selectUserByRole = `-- name: SelectUserByRole :many
SELECT a.user_id, a.full_name, a.email, a.password, b.user_role_id, b.name, b.unique_name
FROM i_users a
JOIN i_users_roles b ON b.user_role_id = a.role_id
WHERE b.unique_name = $1
`

type SelectUserByRoleRow struct {
	UserID     pgtype.UUID `db:"user_id" json:"user_id"`
	FullName   string      `db:"full_name" json:"full_name"`
	Email      string      `db:"email" json:"email"`
	Password   string      `db:"password" json:"password"`
	UserRoleID pgtype.UUID `db:"user_role_id" json:"user_role_id"`
	Name       string      `db:"name" json:"name"`
	UniqueName UserRole    `db:"unique_name" json:"unique_name"`
}

func (q *Queries) SelectUserByRole(ctx context.Context, uniqueName UserRole) ([]SelectUserByRoleRow, error) {
	rows, err := q.db.Query(ctx, selectUserByRole, uniqueName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SelectUserByRoleRow
	for rows.Next() {
		var i SelectUserByRoleRow
		if err := rows.Scan(
			&i.UserID,
			&i.FullName,
			&i.Email,
			&i.Password,
			&i.UserRoleID,
			&i.Name,
			&i.UniqueName,
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
