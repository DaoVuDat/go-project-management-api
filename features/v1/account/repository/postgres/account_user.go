package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"project-management/common"
	db "project-management/db/sqlc"
	"project-management/domain"
	"time"
)

type accountUserRepository struct {
	appCtx   common.AppContext
	connPool *pgxpool.Pool
}

func NewPostgresAccountUserRepository(appCtx common.AppContext) domain.AccountRepository {
	return &accountUserRepository{
		appCtx:   appCtx,
		connPool: appCtx.Pool,
	}
}

/*
IMPLEMENT REPOSITORY INTERFACE
*/

func (accountUserRepo *accountUserRepository) GetUserAccount(
	ctx context.Context,
	username string,
) (*db.UserAccount, error) {
	query := db.New(accountUserRepo.connPool)
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
	query := db.New(accountUserRepo.connPool)
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
	password *string,
) (*db.UserAccount, error) {
	accountStatus := db.NullAccountStatus{}
	accountType := db.NullAccountType{}
	passwordType := pgtype.Text{}

	if typeAccount != nil {
		accountType.AccountType = *typeAccount
		accountType.Valid = true
	}

	if statusAccount != nil {
		accountStatus.AccountStatus = *statusAccount
		accountType.Valid = true
	}

	if password != nil {
		passwordType.String = *password
		passwordType.Valid = true
	}

	query := db.New(accountUserRepo.connPool)
	account, err := query.UpdateUserAccount(ctx, db.UpdateUserAccountParams{
		UserID:    int64(userId),
		UpdatedAt: time.Now(),
		Status:    accountStatus,
		Type:      accountType,
		Password:  passwordType,
	})
	if err != nil {
		return nil, err
	}
	return &account, nil
}
