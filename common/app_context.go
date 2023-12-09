package common

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"project-management/domain"
)

type AppContext struct {
	Logger   *zap.Logger
	Pool     *pgxpool.Pool
	GbConfig *domain.GlobalEnvConfig
}
