package usecase

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"project-management/common"
	db "project-management/db/sqlc"
	"project-management/domain"
	"project-management/util"
)

type accountUserUseCase struct {
	appContext      common.AppContext
	accountUserRepo domain.AccountRepository
}

func NewAccountUserUseCase(appCtx common.AppContext, accountUserRepo domain.AccountRepository) domain.AccountUseCase {
	return &accountUserUseCase{
		appContext:      appCtx,
		accountUserRepo: accountUserRepo,
	}
}

/*
IMPLEMENT USE CASE INTERFACE
*/

func (accountUserUC *accountUserUseCase) CreateUserAccount(
	ctx context.Context,
	username string,
	password string,
) (domain.AccountResponse, error) {
	// Check username is existed ?
	account, err := accountUserUC.accountUserRepo.GetUserAccount(ctx, username)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return domain.AccountResponse{}, err
		}
	}

	if account != nil {
		return domain.AccountResponse{}, domain.ErrUsernameExists
	}

	// If it does not exist, create the new one
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return domain.AccountResponse{}, err
	}

	newUserAccount, err := accountUserUC.accountUserRepo.InsertUserAccount(ctx, username, hashedPassword)
	if err != nil {
		return domain.AccountResponse{}, err
	}
	newUserAccountResponse := domain.AccountResponse{
		Id:       int(newUserAccount.UserID),
		Username: newUserAccount.Username,
		Type:     newUserAccount.Type,
		Status:   newUserAccount.Status,
	}
	return newUserAccountResponse, nil
}

func (accountUserUC *accountUserUseCase) LoginAccount(
	ctx context.Context,
	username string,
	password string,
) (domain.AccountResponse, error) {
	panic(1)
}

func (accountUserUC *accountUserUseCase) UpdateUserAccount(
	ctx context.Context,
	userId int,
	typeAccount *db.AccountType,
	statusAccount *db.AccountStatus,
	password *string,
) (domain.AccountResponse, error) {
	updatedAccount, err := accountUserUC.accountUserRepo.UpdateUserAccount(ctx, userId, typeAccount, statusAccount, password)
	if err != nil {
		return domain.AccountResponse{}, err
	}

	updatedAccountResponse := domain.AccountResponse{
		Id:       int(updatedAccount.UserID),
		Username: updatedAccount.Username,
		Type:     updatedAccount.Type,
		Status:   updatedAccount.Status,
	}

	return updatedAccountResponse, nil
}
