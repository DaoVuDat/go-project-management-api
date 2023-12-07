package usecase

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"project-management/common"
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
) (domain.AccountResponseWithToken, error) {
	accountUserUC.appContext.Logger.Debug("CreateUserAccount UC")

	// Check username is existed ?
	account, err := accountUserUC.accountUserRepo.GetUserAccount(ctx, username)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			accountUserUC.appContext.Logger.Debug("CreateUserAccount UC", zap.String("error", "general error"))
			return domain.AccountResponseWithToken{}, err
		}
	}

	if account != nil {
		accountUserUC.appContext.Logger.Debug("CreateUserAccount UC", zap.String("error", "user existed"))

		return domain.AccountResponseWithToken{}, domain.ErrUsernameExists
	}

	// If it does not exist, create the new one
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return domain.AccountResponseWithToken{}, err
	}

	newUserAccount, err := accountUserUC.accountUserRepo.InsertUserAccount(ctx, username, hashedPassword)
	if err != nil {
		return domain.AccountResponseWithToken{}, err
	}
	newUserAccountResponseWithToken := domain.AccountResponseWithToken{
		AccountResponse: domain.AccountResponse{
			Id:       int(newUserAccount.UserID),
			Username: newUserAccount.Username,
			Type:     newUserAccount.Type,
			Status:   newUserAccount.Status,
		},
		Token: "",
	}
	return newUserAccountResponseWithToken, nil
}

func (accountUserUC *accountUserUseCase) LoginAccount(
	ctx context.Context,
	username string,
	password string,
) (domain.AccountResponseWithToken, error) {
	account, err := accountUserUC.accountUserRepo.GetUserAccount(ctx, username)
	if err != nil {
		accountUserUC.appContext.Logger.Debug("LoginAccount UC", zap.String("error", "dasdas"))
		if !errors.Is(err, pgx.ErrNoRows) {
			accountUserUC.appContext.Logger.Debug("LoginAccount UC", zap.String("error", "general error"))
			return domain.AccountResponseWithToken{}, domain.ErrInvalidLogin
		}
	}

	// If user exists, check password
	err = util.ComparePassword(account.Password, password)
	if err != nil {
		return domain.AccountResponseWithToken{}, domain.ErrInvalidLogin
	}

	// We should return token here too
	accountResponse := domain.AccountResponseWithToken{
		AccountResponse: domain.AccountResponse{
			Id:       int(account.UserID),
			Username: account.Username,
			Type:     account.Type,
			Status:   account.Status,
		},
		Token: "",
	}

	return accountResponse, nil
}

func (accountUserUC *accountUserUseCase) UpdateUserAccount(
	ctx context.Context,
	updateUserAccount domain.AccountUpdate,
) (domain.AccountResponse, error) {

	// Hash update password if update password does exist
	if updateUserAccount.Password.Valid {
		hp, err := util.HashPassword(updateUserAccount.Password.String)
		if err != nil {
			return domain.AccountResponse{}, err
		}
		updateUserAccount.Password.String = hp
	}

	updatedAccount, err := accountUserUC.accountUserRepo.UpdateUserAccount(ctx, updateUserAccount)
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
