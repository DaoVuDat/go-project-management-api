// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package db

import (
	"context"
)

type Querier interface {
	GetProject(ctx context.Context, id int32) (Project, error)
	ListProjects(ctx context.Context) ([]Project, error)
}

var _ Querier = (*Queries)(nil)
