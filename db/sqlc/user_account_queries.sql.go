// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: user_account_queries.sql

package db

import (
	"context"
	"time"
)

const addUserAccount = `-- name: AddUserAccount :one
INSERT INTO user_account (username, password)
VALUES ($1, $2)
RETURNING user_id, username, password, type, status, created_at, updated_at
`

type AddUserAccountParams struct {
	Username string
	Password string
}

func (q *Queries) AddUserAccount(ctx context.Context, arg AddUserAccountParams) (UserAccount, error) {
	row := q.db.QueryRow(ctx, addUserAccount, arg.Username, arg.Password)
	var i UserAccount
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Password,
		&i.Type,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserNameAccount = `-- name: GetUserNameAccount :one
SELECT user_id, username, password, type, status, created_at, updated_at
FROM user_account
WHERE username = $1
LIMIT 1
`

func (q *Queries) GetUserNameAccount(ctx context.Context, username string) (UserAccount, error) {
	row := q.db.QueryRow(ctx, getUserNameAccount, username)
	var i UserAccount
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Password,
		&i.Type,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserAccount = `-- name: UpdateUserAccount :one
UPDATE user_account
SET status     = COALESCE($3, status),
    type       = COALESCE($4, type),
    updated_at = $2
WHERE user_id = $1
RETURNING user_id, username, password, type, status, created_at, updated_at
`

type UpdateUserAccountParams struct {
	UserID    int32
	UpdatedAt time.Time
	Status    NullAccountStatus
	Type      NullAccountType
}

func (q *Queries) UpdateUserAccount(ctx context.Context, arg UpdateUserAccountParams) (UserAccount, error) {
	row := q.db.QueryRow(ctx, updateUserAccount,
		arg.UserID,
		arg.UpdatedAt,
		arg.Status,
		arg.Type,
	)
	var i UserAccount
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Password,
		&i.Type,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
