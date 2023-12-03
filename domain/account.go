package domain

import db "project-management/db/sqlc"

type AccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccountResponse struct {
	Id       int              `json:"id"`
	Username string           `json:"username"`
	Type     db.AccountType   `json:"type"`
	Status   db.AccountStatus `json:"status"`
}

type AccountUseCase interface {
	CreateUserAccount(username string, password string) (AccountResponse, error)
	LoginAccount(username string, password string) (AccountResponse, error)
	UpdateUserAccount(typeAccount *db.AccountType, statusAccount *db.AccountStatus) (AccountResponse, error)
}

type AccountRepository interface {
	GetUserAccount(username string) (db.UserAccount, error)
	InsertUserAccount(username string, password string) (db.UserAccount, error)
	UpdateUserAccount(typeAccount *db.AccountType, statusAccount *db.AccountStatus) (db.UserAccount, error)
}
