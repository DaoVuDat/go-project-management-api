package domain

import "time"

type GlobalEnvConfig struct {
	DBSource         string        `mapstructure:"DB_SOURCE"`
	ServerAddress    string        `mapstructure:"SERVER_ADDRESS"`
	MaxConnLifetime  time.Duration `mapstructure:"MAX_CONN_LIFETIME"`
	MaxConnIdleTime  time.Duration `mapstructure:"MAX_CONN_IDLE_TIME"`
	MaxConn          int32         `mapstructure:"MAX_CONN"`
	MinConn          int32         `mapstructure:"MIN_CONN"`
	TokenPrivateKey  string        `mapstructure:"TOKEN_PRIVATE_KEY"`
	TokenExpiredTime time.Duration `mapstructure:"TOKEN_EXPIRE_TIME"`
}
