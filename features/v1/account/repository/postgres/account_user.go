package postgresuseracc

import (
	"context"
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
	queries *db.Queries,
	username,
	password string,
) (*db.UserAccount, error) {
	account, err := queries.AddUserAccount(ctx, db.AddUserAccountParams{
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
	updateUserAccount domain.AccountUpdate,
) (*db.UserAccount, error) {

	query := db.New(accountUserRepo.connPool)
	account, err := query.UpdateUserAccount(ctx, db.UpdateUserAccountParams{
		UserID:    int64(updateUserAccount.UserId),
		UpdatedAt: time.Now(),
		Status:    updateUserAccount.Status,
		Type:      updateUserAccount.Type,
		Password:  updateUserAccount.Password,
	})
	if err != nil {
		return nil, err
	}
	return &account, nil
}
