package domain

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
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

// UC Layer and Repo Layer

type AccountUseCase interface {
	CreateUserAccount(ctx context.Context, username string, password string) (AccountResponse, error)
	LoginAccount(ctx context.Context, username string, password string) (AccountResponse, error)
	UpdateUserAccount(ctx context.Context, updateUserAccount AccountUpdate) (AccountResponse, error)
}

type AccountRepository interface {
	GetUserAccount(ctx context.Context, username string) (*db.UserAccount, error)
	InsertUserAccount(ctx context.Context, username string, password string) (*db.UserAccount, error)
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

// Setup custom validators

func (cv *CustomValidator) SetUpAccountUserValidator() {
	cv.Validator.RegisterStructValidation(func(sl validator.StructLevel) {
		fmt.Println("Validate")
		accountUpdateRequest := sl.Current().Interface().(AccountUpdateRequest)

		if len(accountUpdateRequest.Status) > 0 {
			status := accountUpdateRequest.Status
			if status != db.AccountStatusPending && status != db.AccountStatusActivated {
				sl.ReportError(accountUpdateRequest.Status, "accountStatus", "Status", "", "invalid status account")
			}
		}

		if len(accountUpdateRequest.Type) > 0 {
			typeAccount := accountUpdateRequest.Type
			if typeAccount != db.AccountTypeClient && typeAccount != db.AccountTypeAdmin {
				sl.ReportError(accountUpdateRequest.Type, "accountType", "Type", "", "invalid type account")
			}
		}

	}, AccountUpdateRequest{})
}
