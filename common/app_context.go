package common

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type AppContext struct {
	Logger *zap.Logger
	Pool   *pgxpool.Pool
}
