package domain

import (
	"context"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jackc/pgx/v5/pgtype"

	//"project-management/common"
	db "project-management/db/sqlc"
)

// Request and Response model

type AccountCreateAndLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccountUpdateRequest struct {
	Type     db.AccountType   `json:"accountType,omitempty"`
	Status   db.AccountStatus `json:"accountStatus,omitempty"`
	Password string           `json:"password,omitempty"`
}

type AccountResponse struct {
	Id       int              `json:"id"`
	Username string           `json:"username"`
	Type     db.AccountType   `json:"type"`
	Status   db.AccountStatus `json:"status"`
}

type AccountResponseWithToken struct {
	AccountResponse
	Token string `json:"accessToken"`
}

// UC Layer and Repo Layer

type AccountUseCase interface {
	CreateUserAccount(ctx context.Context, username string, password string) (AccountResponseWithToken, error)
	LoginAccount(ctx context.Context, username string, password string) (AccountResponseWithToken, error)
	UpdateUserAccount(ctx context.Context, updateUserAccount AccountUpdate) (AccountResponseWithToken, error)
}

type AccountRepository interface {
	GetUserAccount(ctx context.Context, username string) (*db.UserAccount, error)
	InsertUserAccount(ctx context.Context, queries *db.Queries, username string, password string) (*db.UserAccount, error)
	UpdateUserAccount(ctx context.Context, updateUserAccount AccountUpdate) (*db.UserAccount, error)
}

// Utils

type AccountUpdate struct {
	UserId   int
	Type     db.NullAccountType
	Status   db.NullAccountStatus
	Password pgtype.Text
}

func (a *AccountUpdate) MapAccountUpdateRequestToAccountUpdate(userId int, data AccountUpdateRequest) {
	a.UserId = userId
	a.Type = db.NullAccountType{
		AccountType: data.Type,
		Valid:       len(data.Type) > 0,
	}
	a.Status = db.NullAccountStatus{
		AccountStatus: data.Status,
		Valid:         len(data.Status) > 0,
	}
	a.Password = pgtype.Text{
		String: data.Password,
		Valid:  len(data.Password) > 0,
	}
}

// Setup validators

func (req AccountCreateAndLoginRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Username,
			validation.Required.Error("userName is required"),
			validation.Length(6, 20).Error("the userName must be in range 6 to 20 characters"),
		),
		validation.Field(&req.Password,
			validation.Required.Error("password is required"),
			validation.Min(6).Error("the password must be at least 8 characters"),
		),
	)
}

func (req AccountUpdateRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Type, validation.When(
			req.Type != "",
			validation.In(
				db.AccountTypeAdmin,
				db.AccountTypeClient,
			).Error(fmt.Sprintf("must be %s or %s", db.AccountTypeAdmin, db.AccountTypeClient)),
		)),
		validation.Field(&req.Status,
			validation.When(
				req.Status != "",
				validation.In(
					db.AccountStatusActivated,
					db.AccountStatusPending,
				).Error(fmt.Sprintf("must be %s or %s", db.AccountStatusActivated, db.AccountStatusPending)),
			),
		),
		validation.Field(&req.Password,
			validation.When(
				req.Password != "",
				validation.Length(6, 100).Error("must be at least 8 characters"),
			),
		),
	)
}
