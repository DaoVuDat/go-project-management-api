package domain

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgx/v5/pgtype"
	db "project-management/db/sqlc"
)

// Request and Response model

type UserProfileUpdateRequest struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

type UserProfileImageUrlRequest struct {
	ImageURL string `json:"imageUrl"`
}

type UserProfileResponse struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	ImageUrl  string `json:"imageUrl"`
}

// UC Layer and Repo Layer

type UserProfileUseCase interface {
	GetUserProfile(ctx context.Context, id int) (UserProfileResponse, error)
	UpdateUserProfile(ctx context.Context, userProfileUpdate UserProfileUpdate) (UserProfileResponse, error)
	UpdateUserProfileImageUrl(ctx context.Context, userProfileImageUrlUpdate UserProfileImageUrlUpdate) (UserProfileResponse, error)
}

type UserProfileRepository interface {
	CreateUserProfile(ctx context.Context, queries *db.Queries, userProfileCreate UserProfileCreate) (*db.UserProfile, error)
	GetUserProfile(ctx context.Context, id int) (*db.UserProfile, error)
	UpdateUserProfile(ctx context.Context, userProfileUpdate UserProfileUpdate) (*db.UserProfile, error)
	UpdateUserProfileImageUrl(ctx context.Context, userProfileImageUrlUpdate UserProfileImageUrlUpdate) (*db.UserProfile, error)
}

// Utils

type UserProfileCreate struct {
	UserId    int
	FirstName string
	LastName  string
}

type UserProfileUpdate struct {
	UserId    int
	FirstName pgtype.Text
	LastName  pgtype.Text
}

type UserProfileImageUrlUpdate struct {
	UserId   int
	ImageUrl pgtype.Text
}

func (user *UserProfileUpdate) MapUserProfileUpdateRequestToUserProfileUpdate(id int, data UserProfileUpdateRequest) {
	user.UserId = id
	if data.FirstName != nil {
		user.FirstName = pgtype.Text{
			String: *data.FirstName,
			Valid:  true,
		}
	} else {
		user.FirstName = pgtype.Text{
			Valid: false,
		}
	}

	if data.LastName != nil {
		user.LastName = pgtype.Text{
			String: *data.LastName,
			Valid:  true,
		}
	} else {
		user.LastName = pgtype.Text{
			Valid: false,
		}
	}

}

func (user *UserProfileImageUrlUpdate) MapUserProfileImageUrlUpdateRequestToUserProfileImageUrlUpdate(id int, data UserProfileImageUrlRequest) {
	user.UserId = id
	user.ImageUrl = pgtype.Text{
		String: data.ImageURL,
		Valid:  true,
	}
}

// Validations

func (req UserProfileUpdateRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.FirstName, validation.When(
			req.FirstName != nil,
			validation.Length(1, 100).Error("must be at least 1")),
		),
		validation.Field(&req.LastName, validation.When(
			req.LastName != nil,
			validation.Length(1, 100).Error("must be at least 1")),
		),
	)
}
