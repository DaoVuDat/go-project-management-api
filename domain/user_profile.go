package domain

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgx/v5/pgtype"
	db "project-management/db/sqlc"
)

// Request and Response model

type UserProfileUpdateRequest struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
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
	GetUserProfile(ctx context.Context, id int) (*db.UserProfile, error)
	UpdateUserProfile(ctx context.Context, userProfileUpdate UserProfileUpdate) (*db.UserProfile, error)
	UpdateUserProfileImageUrl(ctx context.Context, userProfileImageUrlUpdate UserProfileImageUrlUpdate) (*db.UserProfile, error)
}

// Utils

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
	user.FirstName = pgtype.Text{
		String: data.FirstName,
		Valid:  len(data.FirstName) > 0,
	}
	user.LastName = pgtype.Text{
		String: data.LastName,
		Valid:  len(data.LastName) > 0,
	}
}

func (user *UserProfileImageUrlUpdate) MapUserProfileImageUrlUpdateRequestToUserProfileImageUrlUpdate(id int, data UserProfileImageUrlRequest) {
	user.UserId = id
	user.ImageUrl = pgtype.Text{
		String: data.ImageURL,
		Valid:  len(data.ImageURL) > 0,
	}
}

// Validations

func (req UserProfileUpdateRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.FirstName, validation.When(
			req.FirstName != "",
			validation.Length(1, 100).Error("must be at least 1")),
		),
		validation.Field(&req.LastName, validation.When(
			req.LastName != "",
			validation.Length(1, 100).Error("must be at least 1")),
		),
	)
}
