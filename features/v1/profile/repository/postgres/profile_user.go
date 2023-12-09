package postgresprofileuser

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"project-management/common"
	db "project-management/db/sqlc"
	"project-management/domain"
)

type profileUserRepo struct {
	appContext common.AppContext
	connPool   *pgxpool.Pool
}

func NewPostgresUserProfileRepo(appCtx common.AppContext, connPool *pgxpool.Pool) domain.UserProfileRepository {
	return &profileUserRepo{
		appContext: appCtx,
		connPool:   connPool,
	}
}

func (repo *profileUserRepo) GetUserProfile(ctx context.Context, id int) (*db.UserProfile, error) {
	panic(1)
}

func (repo *profileUserRepo) UpdateUserProfile(ctx context.Context, userProfileUpdate domain.UserProfileUpdate) (*db.UserProfile, error) {
	panic(1)
}

func (repo *profileUserRepo) UpdateUserProfileImageUrl(ctx context.Context, userProfileImageUrlUpdate domain.UserProfileImageUrlUpdate) (*db.UserProfile, error) {
	panic(1)
}
