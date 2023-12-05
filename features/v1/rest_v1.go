package apiv1

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	glbmiddleware "project-management/features/middleware"
	httpuseracc "project-management/features/v1/account/delivery/http"
	"project-management/features/v1/account/repository/postgres"
	"project-management/features/v1/account/usecase"
)

func SetupRestVersion1Api(groupRoute *echo.Group, dbPool *pgxpool.Pool) {
	rV1Group := groupRoute.Group("/v1/")
	rV1Group.Use(glbmiddleware.ApiVersionCtxMiddleware("v1"))
	// Repo
	accountRepo := postgres.NewPostgresAccountUserRepository(dbPool)

	// Use case
	accountUseCase := usecase.NewAccountUserUseCase(accountRepo)

	// Setup Handlers
	httpuseracc.SetupAccountUserHandler(rV1Group, accountUseCase)
}
