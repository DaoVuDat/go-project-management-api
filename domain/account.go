package domain

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"

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
	Password string           `json:"password"`
}

type AccountResponse struct {
	Id       int              `json:"id"`
	Username string           `json:"username"`
	Type     db.AccountType   `json:"type"`
	Status   db.AccountStatus `json:"status"`
}

// Setup custom validators

func (cv *CustomValidator) SetUpAccountUserValidator() {
	cv.Validator.RegisterStructValidation(func(sl validator.StructLevel) {
		accountUpdateRequest := sl.Current().Interface().(AccountUpdateRequest)
		fmt.Println(len(accountUpdateRequest.Status))

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

// UC Layer and Repo Layer

type AccountUseCase interface {
	CreateUserAccount(ctx context.Context, username string, password string) (AccountResponse, error)
	LoginAccount(ctx context.Context, username string, password string) (AccountResponse, error)
	UpdateUserAccount(ctx context.Context, userId int, typeAccount *db.AccountType, statusAccount *db.AccountStatus, password *string) (AccountResponse, error)
}

type AccountRepository interface {
	GetUserAccount(ctx context.Context, username string) (*db.UserAccount, error)
	InsertUserAccount(ctx context.Context, username string, password string) (*db.UserAccount, error)
	UpdateUserAccount(ctx context.Context, userId int, typeAccount *db.AccountType, statusAccount *db.AccountStatus, password *string) (*db.UserAccount, error)
}
