package domain

import (
	"context"
	"net/http"
	db "project-management/db/sqlc"
)

type AccountCreateAndLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccountUpdateRequest struct {
	Type   string `json:"accountType"`
	Status string `json:"accountStatus"`
}

type AccountUpdatePasswordRequest struct {
	Password string `json:"password"`
}

type AccountResponse struct {
	Id       int              `json:"id"`
	Username string           `json:"username"`
	Type     db.AccountType   `json:"type"`
	Status   db.AccountStatus `json:"status"`
}

// IMPLEMENTATION

func (a *AccountCreateAndLoginRequest) Bind(r *http.Request) error {
	// Post Process after decode
	return nil
}

func (a *AccountUpdateRequest) Bind(r *http.Request) error {
	// Post Process after decode
	return nil
}

func (a *AccountUpdatePasswordRequest) Bind(r *http.Request) error {
	// Post Process after decode
	return nil
}

func (a *AccountResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

type AccountUseCase interface {
	CreateUserAccount(ctx context.Context, username string, password string) (AccountResponse, error)
	LoginAccount(ctx context.Context, username string, password string) (AccountResponse, error)
	UpdateUserAccount(ctx context.Context, userId int, typeAccount *db.AccountType, statusAccount *db.AccountStatus) (AccountResponse, error)
}

type AccountRepository interface {
	GetUserAccount(ctx context.Context, username string) (*db.UserAccount, error)
	InsertUserAccount(ctx context.Context, username string, password string) (*db.UserAccount, error)
	UpdateUserAccount(ctx context.Context, userId int, typeAccount *db.AccountType, statusAccount *db.AccountStatus) (*db.UserAccount, error)
}
