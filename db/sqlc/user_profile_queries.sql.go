// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: user_profile_queries.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const getUserProfileById = `-- name: GetUserProfileById :one
SELECT id, first_name, last_name, created_at, updated_at, image_url
FROM user_profile
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetUserProfileById(ctx context.Context, id int64) (UserProfile, error) {
	row := q.db.QueryRow(ctx, getUserProfileById, id)
	var i UserProfile
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ImageUrl,
	)
	return i, err
}

const updateUserProfile = `-- name: UpdateUserProfile :one
UPDATE user_profile
SET first_name = $2,
    last_name  = $3,
    image_url  = $4,
    updated_at = $5
WHERE id = $1
RETURNING id, first_name, last_name, created_at, updated_at, image_url
`

type UpdateUserProfileParams struct {
	ID        int64       `json:"id"`
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	ImageUrl  pgtype.Text `json:"imageUrl"`
	UpdatedAt time.Time   `json:"updatedAt"`
}

func (q *Queries) UpdateUserProfile(ctx context.Context, arg UpdateUserProfileParams) (UserProfile, error) {
	row := q.db.QueryRow(ctx, updateUserProfile,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.ImageUrl,
		arg.UpdatedAt,
	)
	var i UserProfile
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ImageUrl,
	)
	return i, err
}
