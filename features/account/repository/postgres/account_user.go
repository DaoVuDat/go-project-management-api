package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	db "project-management/db/sqlc"
	"project-management/domain"
	"time"
)

type accountUserRepository struct {
	ConnPool *pgxpool.Pool
}

func NewPostgresAccountUserRepository(connPool *pgxpool.Pool) domain.AccountRepository {
	return &accountUserRepository{
		ConnPool: connPool,
	}
}

/*
IMPLEMENT REPOSITORY INTERFACE
*/

func (accountUserRepo *accountUserRepository) GetUserAccount(
	ctx context.Context,
	username string,
) (*db.UserAccount, error) {
	query := db.New(accountUserRepo.ConnPool)
	account, err := query.GetUserNameAccount(ctx, username)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (accountUserRepo *accountUserRepository) InsertUserAccount(
	ctx context.Context,
	username,
	password string,
) (*db.UserAccount, error) {
	query := db.New(accountUserRepo.ConnPool)
	account, err := query.AddUserAccount(ctx, db.AddUserAccountParams{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (accountUserRepo *accountUserRepository) UpdateUserAccount(
	ctx context.Context,
	userId int,
	typeAccount *db.AccountType,
	statusAccount *db.AccountStatus,
) (*db.UserAccount, error) {
	accountStatus := db.NullAccountStatus{}
	accountType := db.NullAccountType{}

	if typeAccount != nil {
		accountType.AccountType = *typeAccount
		accountType.Valid = true
	}

	if statusAccount != nil {
		accountStatus.AccountStatus = *statusAccount
		accountType.Valid = true
	}

	query := db.New(accountUserRepo.ConnPool)
	account, err := query.UpdateUserAccount(ctx, db.UpdateUserAccountParams{
		UserID:    int32(userId),
		UpdatedAt: time.Now(),
		Status:    accountStatus,
		Type:      accountType,
	})
	if err != nil {
		return nil, err
	}
	return &account, nil
}
