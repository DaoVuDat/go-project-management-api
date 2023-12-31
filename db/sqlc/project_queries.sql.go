// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: project_queries.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createProject = `-- name: CreateProject :one
INSERT INTO project(user_profile, price, paid)
VALUES ($1, $2, $3)
RETURNING id, user_profile, name, description, price, paid, status, start_time, end_time, created_at, updated_at
`

type CreateProjectParams struct {
	UserProfile pgtype.Int8 `db:"user_profile"`
	Price       int32       `db:"price"`
	Paid        int32       `db:"paid"`
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error) {
	row := q.db.QueryRow(ctx, createProject, arg.UserProfile, arg.Price, arg.Paid)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.UserProfile,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Paid,
		&i.Status,
		&i.StartTime,
		&i.EndTime,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAProjectByUserId = `-- name: GetAProjectByUserId :one
SELECT id, user_profile, name, description, price, paid, status, start_time, end_time, created_at, updated_at
FROM project
WHERE id = $1
  AND user_profile = $2
LIMIT 1
`

type GetAProjectByUserIdParams struct {
	ID          int64       `db:"id"`
	UserProfile pgtype.Int8 `db:"user_profile"`
}

func (q *Queries) GetAProjectByUserId(ctx context.Context, arg GetAProjectByUserIdParams) (Project, error) {
	row := q.db.QueryRow(ctx, getAProjectByUserId, arg.ID, arg.UserProfile)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.UserProfile,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Paid,
		&i.Status,
		&i.StartTime,
		&i.EndTime,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProject = `-- name: GetProject :one
SELECT id, user_profile, name, description, price, paid, status, start_time, end_time, created_at, updated_at
FROM project
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetProject(ctx context.Context, id int64) (Project, error) {
	row := q.db.QueryRow(ctx, getProject, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.UserProfile,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Paid,
		&i.Status,
		&i.StartTime,
		&i.EndTime,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listProjects = `-- name: ListProjects :many
SELECT id, user_profile, name, description, price, paid, status, start_time, end_time, created_at, updated_at
FROM project
ORDER BY created_at
`

func (q *Queries) ListProjects(ctx context.Context) ([]Project, error) {
	rows, err := q.db.Query(ctx, listProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Project{}
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.UserProfile,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.Paid,
			&i.Status,
			&i.StartTime,
			&i.EndTime,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const listProjectsByUserId = `-- name: ListProjectsByUserId :many
SELECT id, user_profile, name, description, price, paid, status, start_time, end_time, created_at, updated_at
FROM project
WHERE user_profile = $1
ORDER BY created_at
`

func (q *Queries) ListProjectsByUserId(ctx context.Context, userProfile pgtype.Int8) ([]Project, error) {
	rows, err := q.db.Query(ctx, listProjectsByUserId, userProfile)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Project{}
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.UserProfile,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.Paid,
			&i.Status,
			&i.StartTime,
			&i.EndTime,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateProjectName = `-- name: UpdateProjectName :one
UPDATE project
SET name        = COALESCE($4, name),
    description = COALESCE($5, description),
    updated_at  = $3
WHERE id = $1 AND user_profile = $2
RETURNING id, user_profile, name, description, price, paid, status, start_time, end_time, created_at, updated_at
`

type UpdateProjectNameParams struct {
	ID          int64       `db:"id"`
	UserProfile pgtype.Int8 `db:"user_profile"`
	UpdatedAt   time.Time   `db:"updated_at"`
	Name        pgtype.Text `db:"name"`
	Description pgtype.Text `db:"description"`
}

func (q *Queries) UpdateProjectName(ctx context.Context, arg UpdateProjectNameParams) (Project, error) {
	row := q.db.QueryRow(ctx, updateProjectName,
		arg.ID,
		arg.UserProfile,
		arg.UpdatedAt,
		arg.Name,
		arg.Description,
	)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.UserProfile,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Paid,
		&i.Status,
		&i.StartTime,
		&i.EndTime,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateProjectPaid = `-- name: UpdateProjectPaid :one
UPDATE project
SET paid   = COALESCE($4, paid),
    status = COALESCE($5, status),
    updated_at = $3
WHERE id = $1 AND user_profile = $2
RETURNING id, user_profile, name, description, price, paid, status, start_time, end_time, created_at, updated_at
`

type UpdateProjectPaidParams struct {
	ID          int64             `db:"id"`
	UserProfile pgtype.Int8       `db:"user_profile"`
	UpdatedAt   time.Time         `db:"updated_at"`
	Paid        pgtype.Int4       `db:"paid"`
	Status      NullProjectStatus `db:"status"`
}

func (q *Queries) UpdateProjectPaid(ctx context.Context, arg UpdateProjectPaidParams) (Project, error) {
	row := q.db.QueryRow(ctx, updateProjectPaid,
		arg.ID,
		arg.UserProfile,
		arg.UpdatedAt,
		arg.Paid,
		arg.Status,
	)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.UserProfile,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Paid,
		&i.Status,
		&i.StartTime,
		&i.EndTime,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateProjectTimeWorking = `-- name: UpdateProjectTimeWorking :one
UPDATE project
SET start_time = COALESCE($4, start_time),
    end_time   = COALESCE($5, end_time),
    updated_at = $3
WHERE id = $1 AND user_profile = $2
RETURNING id, user_profile, name, description, price, paid, status, start_time, end_time, created_at, updated_at
`

type UpdateProjectTimeWorkingParams struct {
	ID          int64              `db:"id"`
	UserProfile pgtype.Int8        `db:"user_profile"`
	UpdatedAt   time.Time          `db:"updated_at"`
	StartTime   pgtype.Timestamptz `db:"start_time"`
	EndTime     pgtype.Timestamptz `db:"end_time"`
}

func (q *Queries) UpdateProjectTimeWorking(ctx context.Context, arg UpdateProjectTimeWorkingParams) (Project, error) {
	row := q.db.QueryRow(ctx, updateProjectTimeWorking,
		arg.ID,
		arg.UserProfile,
		arg.UpdatedAt,
		arg.StartTime,
		arg.EndTime,
	)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.UserProfile,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Paid,
		&i.Status,
		&i.StartTime,
		&i.EndTime,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
