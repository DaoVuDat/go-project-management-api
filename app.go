package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log"
	"net/http"
	db "project-management/db/sqlc"
	"time"
)

type GlobalEnvConfig struct {
	DBSource        string        `mapstructure:"DB_SOURCE"`
	ServerAddress   string        `mapstructure:"SERVER_ADDRESS"`
	MaxConnLifetime time.Duration `mapstructure:"MAX_CONN_LIFETIME"`
	MaxConnIdleTime time.Duration `mapstructure:"MAX_CONN_IDLE_TIME"`
	MaxConn         int32         `mapstructure:"MAX_CONN"`
	MinConn         int32         `mapstructure:"MIN_CONN"`
}

func loadConfig(path string) (*GlobalEnvConfig, error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	var config *GlobalEnvConfig

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	return config, nil
}

func connectDb(ctx context.Context, appConfig GlobalEnvConfig) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(appConfig.DBSource)
	// Apply config
	config.MaxConnLifetime = appConfig.MaxConnLifetime
	config.MaxConns = appConfig.MaxConn
	config.MinConns = appConfig.MinConn
	config.MaxConnIdleTime = appConfig.MaxConnIdleTime

	if err != nil {
		return nil, err
	}

	dbPool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return dbPool, err
}

func main() {
	//============================ Loading config
	globalEnvConfig, err := loadConfig(".")
	if err != nil {
		fmt.Println("cannot load config:", err)
	}

	//fmt.Println(globalEnvConfig)

	//============================ Load DB Service (Pg Pool)
	dbPool, err := connectDb(context.Background(), *globalEnvConfig)
	if err != nil {
		fmt.Println("cannot connect postgres:", err)
	}
	defer dbPool.Close()

	// Repo

	// Use case
	queries := db.New(dbPool)
	projects, err := queries.ListProjects(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	log.Println(projects)

	//============================ Create Mux Router
	router := echo.New()

	router.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Hello string `json:"hello"`
		}{
			Hello: "fine",
		})
	})

	router.Logger.Fatal(router.Start(globalEnvConfig.ServerAddress))
}
