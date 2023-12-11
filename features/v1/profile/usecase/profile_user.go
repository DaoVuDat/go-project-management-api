package usecaseprofileuser

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
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
	userProfile, err := uc.userProfileRepository.GetUserProfile(ctx, id)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return domain.UserProfileResponse{}, domain.ErrInvalidUserAccountId
		}
		return domain.UserProfileResponse{}, err
	}

	userProfileResponse := domain.UserProfileResponse{
		Id:        int(userProfile.ID),
		FirstName: userProfile.FirstName,
		LastName:  userProfile.LastName,
		ImageUrl:  userProfile.ImageUrl.String,
	}

	return userProfileResponse, nil
}

func (uc *profileUserUC) UpdateUserProfile(ctx context.Context, userProfileUpdate domain.UserProfileUpdate) (domain.UserProfileResponse, error) {
	userProfile, err := uc.userProfileRepository.UpdateUserProfile(ctx, userProfileUpdate)
	if err != nil {
		return domain.UserProfileResponse{}, err
	}

	userProfileResponse := domain.UserProfileResponse{
		Id:        int(userProfile.ID),
		FirstName: userProfile.FirstName,
		LastName:  userProfile.LastName,
		ImageUrl:  userProfile.ImageUrl.String,
	}

	return userProfileResponse, nil
}

func (uc *profileUserUC) UpdateUserProfileImageUrl(ctx context.Context, userProfileImageUrlUpdate domain.UserProfileImageUrlUpdate) (domain.UserProfileResponse, error) {
	userProfile, err := uc.userProfileRepository.UpdateUserProfileImageUrl(ctx, userProfileImageUrlUpdate)
	if err != nil {
		return domain.UserProfileResponse{}, err
	}

	userProfileResponse := domain.UserProfileResponse{
		Id:        int(userProfile.ID),
		FirstName: userProfile.FirstName,
		LastName:  userProfile.LastName,
		ImageUrl:  userProfile.ImageUrl.String,
	}
	return userProfileResponse, nil
}
