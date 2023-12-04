package usecase

import (
	db "project-management/db/sqlc"
	"project-management/domain"
)

type accountUserUseCase struct {
	accountUserRepo domain.AccountRepository
}

func NewAccountUserUseCase(accountUserRepo domain.AccountRepository) domain.AccountUseCase {
	return &accountUserUseCase{
		accountUserRepo: accountUserRepo,
	}
}

/*
IMPLEMENT USE CASE INTERFACE
*/

func (accountUserUC *accountUserUseCase) CreateUserAccount(
	username string,
	password string,
) (domain.AccountResponse, error) {
	panic(1)
}

func (accountUserUC *accountUserUseCase) LoginAccount(
	username string,
	password string,
) (domain.AccountResponse, error) {
	panic(1)
}

func (accountUserUC *accountUserUseCase) UpdateUserAccount(
	typeAccount *db.AccountType,
	statusAccount *db.AccountStatus,
) (domain.AccountResponse, error) {
	panic(1)
}
