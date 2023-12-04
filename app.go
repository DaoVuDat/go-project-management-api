package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/graceful"
	"github.com/spf13/viper"
	"log"
	"net/http"
	httpuseracc "project-management/features/account/delivery/http"
	"project-management/features/account/repository/postgres"
	"project-management/features/account/usecase"
	glbmiddleware "project-management/features/middleware"
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

	//queries := db.New(dbPool)
	//projects, err := queries.ListProjects(context.Background())
	//if err != nil {
	//	fmt.Println(err)
	//}
	//log.Println(projects)

	//============================ Create Mux Router
	r := chi.NewRouter()
	// Setup Global Middleware
	r.Use(middleware.Logger)

	// Setup handlers
	r.Route("/api/", func(r chi.Router) {
		// Version 1
		r.Route("/v1/", func(r chi.Router) {
			r.Use(glbmiddleware.ApiVersionCtxMiddleware("v1"))
			// Repo
			accountRepo := postgres.NewPostgresAccountUserRepository(dbPool)

			// Use case
			accountUseCase := usecase.NewAccountUserUseCase(accountRepo)

			// Setup Handlers
			httpuseracc.SetupAccountUserHandler(r, accountUseCase)
		})
	})

	//============================ Create Server
	server := graceful.WithDefaults(&http.Server{
		Addr:         globalEnvConfig.ServerAddress,
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	})

	//============================ Start Server
	log.Printf("main: Listening on %s\n", globalEnvConfig.ServerAddress)
	if err := graceful.Graceful(server.ListenAndServe, server.Shutdown); err != nil {
		log.Fatalln("main: Failed to gracefully shutdown")
	}
	log.Println("main: Server was shutdown gracefully")
}
