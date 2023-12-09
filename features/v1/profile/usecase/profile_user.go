package usecaseprofileuser

import (
	"context"
	"project-management/common"
	"project-management/domain"
)

type profileUserUC struct {
	appContext            common.AppContext
	userProfileRepository domain.UserProfileRepository
}

func NewUserProfileUseCase(appCtx common.AppContext, userProfileRepo domain.UserProfileRepository) domain.UserProfileUseCase {
	return &profileUserUC{
		appContext:            appCtx,
		userProfileRepository: userProfileRepo,
	}
}

func (uc *profileUserUC) GetUserProfile(ctx context.Context, id int) (domain.UserProfileResponse, error) {
	panic(1)
}

func (uc *profileUserUC) UpdateUserProfile(ctx context.Context, userProfileUpdate domain.UserProfileUpdate) (domain.UserProfileResponse, error) {
	panic(1)
}

func (uc *profileUserUC) UpdateUserProfileImageUrl(ctx context.Context, userProfileImageUrlUpdate domain.UserProfileImageUrlUpdate) (domain.UserProfileResponse, error) {
	panic(1)
}
