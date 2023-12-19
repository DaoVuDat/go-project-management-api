package httpproject

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project-management/auth"
	"project-management/common"
	db "project-management/db/sqlc"
	"project-management/domain"
)

type projectHandler struct {
	appContext common.AppContext
	projectUC  domain.ProjectUseCase
}

func SetupProjectHandler(
	group *echo.Group,
	ctx common.AppContext,
	projectUC domain.ProjectUseCase,
) {
	handler := &projectHandler{
		appContext: ctx,
		projectUC:  projectUC,
	}

	g := group.Group("/projects")
	g.Use(auth.AuthorizationRestrictedMiddleware(ctx.GbConfig.TokenPrivateKey))
	g.POST("/", handler.CreateANewProject)
}

func (handler *projectHandler) CreateANewProject(c echo.Context) error {
	// Get payloads
	payload := c.Get(auth.AuthorizedPayloadKey).(*auth.JwtCustomPayload)

	if payload.Role != string(db.AccountTypeAdmin) {
		return c.JSON(http.StatusUnauthorized, domain.ErrUnauthorizedResponse(domain.ErrInvalidAuthorization))
	}

	ctx := c.Request().Context()

	var dataRequest domain.ProjectCreateRequest
	if err := c.Bind(&dataRequest); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequestResponse(err))
	}

	if err := dataRequest.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequestResponse(err))
	}

	var createProject domain.ProjectCreate
	createProject.MapProjectCreateRequestToProjectCreate(dataRequest)

	projectResponse, err := handler.projectUC.CreateANewProject(ctx, createProject)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, projectResponse)
}
