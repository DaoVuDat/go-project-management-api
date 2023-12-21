package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
	"os/signal"
	"project-management/common"
	"project-management/domain"
	apiv1 "project-management/features/v1"

	"syscall"
	"time"
)

func main() {
	//============================ Loading config
	globalEnvConfig, err := loadConfig(".")
	if err != nil {
		fmt.Println("cannot load config:", err)
	}

	//============================ Load DB Service (Pg Pool)
	dbPool, err := connectDb(context.Background(), *globalEnvConfig)
	if err != nil {
		fmt.Println("cannot connect postgres:", err)
	}
	defer dbPool.Close()

	//============================ Set up Logger
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format(time.RFC1123Z))
	}
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	//============================ Setup AppCtx
	appCtx := common.AppContext{
		Logger:   logger,
		Pool:     dbPool,
		GbConfig: globalEnvConfig,
	}

	//============================ Create Mux Router
	r := echo.New()
	r.Pre(middleware.RemoveTrailingSlash())
	//r.Validator = validatorReq
	r.Logger.SetLevel(log.INFO)

	// Setup Default Error Handling
	r.HTTPErrorHandler = func(err error, c echo.Context) {

		err = c.JSON(http.StatusInternalServerError, domain.ErrInternalResponse(err))
		if err != nil {
			c.Logger().Error(err)
		}
	}

	// Setup Global Middleware
	r.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogHost:      true,
		LogMethod:    true,
		LogURIPath:   true,
		LogError:     true,
		LogURI:       true,
		LogStatus:    true,
		LogUserAgent: true,
		LogProtocol:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("Request",
				zap.String("Host", v.Host),
				zap.String("URI", v.URI),
				zap.String("URIPath", v.URIPath),
				zap.String("Method", v.Method),
				zap.Int("Status", v.Status),
				zap.String("UserAgent", v.UserAgent),
				zap.String("Protocol", v.Protocol),
				zap.Error(v.Error),
			)
			return nil
		},
	}))

	// Setup handlers
	rApiGroup := r.Group("/api")
	//=== Version 1
	apiv1.SetupRestVersion1Api(appCtx, rApiGroup)

	//============================ Create Server
	server := http.Server{
		Addr:         globalEnvConfig.ServerAddress,
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	//============================ Start Server
	go func() {
		if err := r.StartServer(&server); err != nil && errors.Is(err, http.ErrServerClosed) {
			r.Logger.Info("main: shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		r.Logger.Fatal("main: Server was shutdown gracefully")
	}

}

func loadConfig(path string) (*domain.GlobalEnvConfig, error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	var config *domain.GlobalEnvConfig

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	return config, nil
}

func connectDb(ctx context.Context, appConfig domain.GlobalEnvConfig) (*pgxpool.Pool, error) {
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
