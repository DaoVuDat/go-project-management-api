package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	db "project-management/db/sqlc"
	"project-management/domain"
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
	username string,
) (db.UserAccount, error) {
	panic(1)
}

func (accountUserRepo *accountUserRepository) InsertUserAccount(
	username,
	password string,
) (db.UserAccount, error) {
	panic(1)
}

func (accountUserRepo *accountUserRepository) UpdateUserAccount(
	typeAccount *db.AccountType,
	statusAccount *db.AccountStatus,
) (db.UserAccount, error) {
	panic(1)
}
