package postgresprofileuser

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"project-management/common"
	db "project-management/db/sqlc"
	"project-management/domain"
	"time"
)

type profileUserRepo struct {
	appContext common.AppContext
	connPool   *pgxpool.Pool
}

func NewPostgresUserProfileRepo(appCtx common.AppContext) domain.UserProfileRepository {
	return &profileUserRepo{
		appContext: appCtx,
		connPool:   appCtx.Pool,
	}
}

func (repo *profileUserRepo) GetUserProfile(ctx context.Context, id int) (*db.UserProfile, error) {

	query := db.New(repo.connPool)
	userProfile, err := query.GetUserProfileById(ctx, int64(id))
	if err != nil {
		return nil, err
	}
	return &userProfile, nil
}

func (repo *profileUserRepo) UpdateUserProfile(ctx context.Context, userProfileUpdate domain.UserProfileUpdate) (*db.UserProfile, error) {
	query := db.New(repo.connPool)
	userProfile, err := query.UpdateUserProfile(ctx, db.UpdateUserProfileParams{
		ID:        int64(userProfileUpdate.UserId),
		UpdatedAt: time.Now(),
		FirstName: userProfileUpdate.FirstName,
		LastName:  userProfileUpdate.LastName,
	})

	if err != nil {
		return nil, err
	}
	return &userProfile, nil
}

func (repo *profileUserRepo) UpdateUserProfileImageUrl(ctx context.Context, userProfileImageUrlUpdate domain.UserProfileImageUrlUpdate) (*db.UserProfile, error) {
	query := db.New(repo.connPool)
	userProfile, err := query.UpdateImageUrlUserProfile(ctx, db.UpdateImageUrlUserProfileParams{
		ID:       int64(userProfileImageUrlUpdate.UserId),
		ImageUrl: userProfileImageUrlUpdate.ImageUrl,
	})
	if err != nil {
		return nil, err
	}

	return &userProfile, nil
}
