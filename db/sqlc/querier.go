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
	CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error)
	CreateUserProfile(ctx context.Context, arg CreateUserProfileParams) (UserProfile, error)
	GetProject(ctx context.Context, id int64) (Project, error)
	GetProjectByUser(ctx context.Context, userProfile pgtype.Int8) (Project, error)
	GetUserNameAccount(ctx context.Context, username string) (UserAccount, error)
	GetUserProfileById(ctx context.Context, id int64) (UserProfile, error)
	ListProjects(ctx context.Context) ([]Project, error)
	UpdateImageUrlUserProfile(ctx context.Context, arg UpdateImageUrlUserProfileParams) (UserProfile, error)
	UpdateProjectName(ctx context.Context, arg UpdateProjectNameParams) (Project, error)
	UpdateProjectPaid(ctx context.Context, arg UpdateProjectPaidParams) (Project, error)
	UpdateProjectTimeWorking(ctx context.Context, arg UpdateProjectTimeWorkingParams) (Project, error)
	UpdateUserAccount(ctx context.Context, arg UpdateUserAccountParams) (UserAccount, error)
	UpdateUserProfile(ctx context.Context, arg UpdateUserProfileParams) (UserProfile, error)
}

var _ Querier = (*Queries)(nil)
