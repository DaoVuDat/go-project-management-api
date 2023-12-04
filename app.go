package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	//============================ Create Server
	server := http.Server{
		Addr:         globalEnvConfig.ServerAddress,
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	//============================ Shutdown Gracefully Handling
	// Create a shutdownError channel. We will use this to receive any errors returned
	// by the graceful Shutdown() function.
	shutdownError := make(chan error)
	go func() {
		// Create a quit channel which carries os.Signal values.
		quit := make(chan os.Signal, 1)

		// Use signal.Notify() to listen for incoming SIGINT and SIGTERM signals and
		// relay them to the quit channel. Any other signals will not be caught by
		// signal.Notify() and will retain their default behavior.
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// Read the signal from the quit channel. This code will block until a signal is
		// received.
		s := <-quit

		// Log a message to say that the signal has been caught. Notice that we also
		// call the String() method on the signal to get the signal name and include it
		// in the log entry properties.
		log.Printf("Signal: %s", s.String())

		// Create a context with a 5-second timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Call Shutdown() on our server, passing in the context we just made.
		// Shutdown() will return nil if the graceful shutdown was successful, or an
		// error (which may happen because of a problem closing the listeners, or
		// because the shutdown didn't complete before the 5-second context deadline is
		// hit). We relay this return value to the shutdownError channel.
		shutdownError <- server.Shutdown(ctx)
	}()

	//============================ Start Server
	log.Printf("Listening on %s\n", globalEnvConfig.ServerAddress)
	err = server.ListenAndServe()

	// Calling Shutdown() on our server will cause ListenAndServe() to immediately
	// return a http.ErrServerClosed error. So if we see this error, it is actually a
	// good thing and an indication that the graceful shutdown has started. So we check
	// specifically for this, only returning the error if it is NOT http.ErrServerClosed.
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}

	// Otherwise, we wait to receive the return value from Shutdown() on the
	// shutdownError channel. If return value is an error, we know that there was a
	// problem with the graceful shutdown and we return the error.
	err = <-shutdownError
	if err != nil {
		log.Fatal(err)
	}

	// At this point we know that the graceful shutdown completed successfully and we
	// log a "stopped server" message.
	log.Println("Stopped Server Gracefully")
}
