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

const getProject = `-- name: GetProject :one
SELECT id, user_profile, name, description, price, paid, status, start_time, end_time, created_at, updated_at
FROM project
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetProject(ctx context.Context, id int32) (Project, error) {
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

const getProjectByUser = `-- name: GetProjectByUser :one
SELECT id, user_profile, name, description, price, paid, status, start_time, end_time, created_at, updated_at
FROM project
WHERE user_profile = $1
ORDER BY name
`

func (q *Queries) GetProjectByUser(ctx context.Context, userProfile pgtype.Int4) (Project, error) {
	row := q.db.QueryRow(ctx, getProjectByUser, userProfile)
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

const updateProjectPaid = `-- name: UpdateProjectPaid :one
UPDATE project
SET paid = $2
WHERE id = $1
RETURNING id, user_profile, name, description, price, paid, status, start_time, end_time, created_at, updated_at
`

type UpdateProjectPaidParams struct {
	ID   int32
	Paid int32
}

func (q *Queries) UpdateProjectPaid(ctx context.Context, arg UpdateProjectPaidParams) (Project, error) {
	row := q.db.QueryRow(ctx, updateProjectPaid, arg.ID, arg.Paid)
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

const updateProjectStatus = `-- name: UpdateProjectStatus :one
UPDATE project
SET status = $2
WHERE id = $1
RETURNING id, user_profile, name, description, price, paid, status, start_time, end_time, created_at, updated_at
`

type UpdateProjectStatusParams struct {
	ID     int32
	Status ProjectStatus
}

func (q *Queries) UpdateProjectStatus(ctx context.Context, arg UpdateProjectStatusParams) (Project, error) {
	row := q.db.QueryRow(ctx, updateProjectStatus, arg.ID, arg.Status)
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
SET start_time = $2,
    end_time   = $3
WHERE id = $1
RETURNING id, user_profile, name, description, price, paid, status, start_time, end_time, created_at, updated_at
`

type UpdateProjectTimeWorkingParams struct {
	ID        int32
	StartTime time.Time
	EndTime   time.Time
}

func (q *Queries) UpdateProjectTimeWorking(ctx context.Context, arg UpdateProjectTimeWorkingParams) (Project, error) {
	row := q.db.QueryRow(ctx, updateProjectTimeWorking, arg.ID, arg.StartTime, arg.EndTime)
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
