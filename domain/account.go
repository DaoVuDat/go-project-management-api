package domain

import (
	"context"
	db "project-management/db/sqlc"
)

type AccountCreateAndLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccountUpdateRequest struct {
	Type     string `json:"accountType"`
	Status   string `json:"accountStatus"`
	Password string `json:"password"`
}

type AccountResponse struct {
	Id       int              `json:"id"`
	Username string           `json:"username"`
	Type     db.AccountType   `json:"type"`
	Status   db.AccountStatus `json:"status"`
}

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
