package apiv1

import (
	"github.com/labstack/echo/v4"
	"project-management/common"
	glbmiddleware "project-management/features/middleware"
	httpuseracc "project-management/features/v1/account/delivery/http"
	"project-management/features/v1/account/repository/postgres"
	"project-management/features/v1/account/usecase"
)

func SetupRestVersion1Api(appCtx common.AppContext, groupRoute *echo.Group) {
	rV1Group := groupRoute.Group("/v1")
	rV1Group.Use(glbmiddleware.ApiVersionCtxMiddleware("v1"))
	// Repo
	accountRepo := postgres.NewPostgresAccountUserRepository(appCtx)

	// Use case
	accountUseCase := usecase.NewAccountUserUseCase(appCtx, accountRepo)

	// Setup Handlers
	httpuseracc.SetupAccountUserHandler(rV1Group, accountUseCase)
}
