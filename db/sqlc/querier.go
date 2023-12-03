// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	AddUserAccount(ctx context.Context, arg AddUserAccountParams) (UserAccount, error)
	GetProject(ctx context.Context, id int32) (Project, error)
	GetProjectByUser(ctx context.Context, userProfile pgtype.Int4) (Project, error)
	GetUserNameAccount(ctx context.Context, username string) (UserAccount, error)
	GetUserProfileById(ctx context.Context, id int32) (UserProfile, error)
	ListProjects(ctx context.Context) ([]Project, error)
	UpdateProjectPaid(ctx context.Context, arg UpdateProjectPaidParams) (Project, error)
	UpdateProjectStatus(ctx context.Context, arg UpdateProjectStatusParams) (Project, error)
	UpdateProjectTimeWorking(ctx context.Context, arg UpdateProjectTimeWorkingParams) (Project, error)
	UpdateUserAccount(ctx context.Context, arg UpdateUserAccountParams) (UserAccount, error)
	UpdateUserProfile(ctx context.Context, arg UpdateUserProfileParams) (UserProfile, error)
}

var _ Querier = (*Queries)(nil)
