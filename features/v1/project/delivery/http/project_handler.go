package httpproject

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"project-management/auth"
	"project-management/common"
	db "project-management/db/sqlc"
	"project-management/domain"
	"strconv"
	"strings"
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
	g.GET("/:id", handler.GetAProject)              // admin and client
	g.PATCH("/:id/name", handler.UpdateProjectName) // for admin only
	g.PATCH("/:id/time", handler.UpdateProjectTime) // for admin only
	g.PATCH("/:id/paid", handler.UpdateProjectPaid) // for admin only
	g.POST("", handler.CreateANewProject)           // for admin only
	g.GET("", handler.ListAllProject)               // admin and client
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

func (handler *projectHandler) UpdateProjectName(c echo.Context) error {
	// get payload
	payload := c.Get(auth.AuthorizedPayloadKey).(*auth.JwtCustomPayload)

	if payload.Role != string(db.AccountTypeAdmin) {
		return c.JSON(http.StatusUnauthorized, domain.ErrUnauthorizedResponse(domain.ErrInvalidAuthorization))
	}

	ctx := c.Request().Context()
	var dataRequest domain.ProjectUpdateNameRequest
	if err := c.Bind(&dataRequest); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequestResponse(err))
	}

	if err := dataRequest.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequestResponse(err))
	}

	pathValue := c.Param("id")
	projectId, err := strconv.Atoi(pathValue)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequestResponse(domain.ErrInvalidProjectId))
	}

	var updateProjectName domain.ProjectUpdateName
	updateProjectName.MapProjectUpdateRequestToProjectUpdate(projectId, dataRequest)

	projectResponse, err := handler.projectUC.UpdateAProjectName(ctx, updateProjectName)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, projectResponse)
}

func (handler *projectHandler) UpdateProjectTime(c echo.Context) error {
	// get payload
	payload := c.Get(auth.AuthorizedPayloadKey).(*auth.JwtCustomPayload)

	if payload.Role != string(db.AccountTypeAdmin) {
		return c.JSON(http.StatusUnauthorized, domain.ErrUnauthorizedResponse(domain.ErrInvalidAuthorization))
	}

	ctx := c.Request().Context()
	var dataRequest domain.ProjectUpdateTimeWorkingRequest
	if err := c.Bind(&dataRequest); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequestResponse(err))
	}

	if err := dataRequest.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequestResponse(err))
	}

	pathValue := c.Param("id")
	projectId, err := strconv.Atoi(pathValue)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequestResponse(domain.ErrInvalidProjectId))
	}

	var updateProjectTime domain.ProjectUpdateTimeWorking
	err = updateProjectTime.MapProjectUpdateTimeWorkingRequestToProjectUpdateTimeWorking(projectId, dataRequest)
	if err != nil {
		return err
	}
	handler.appContext.Logger.Debug("Handler Update Name", zap.Any("updateProjectTime", updateProjectTime))

	projectResponse, err := handler.projectUC.UpdateAProjectTimeWorking(ctx, updateProjectTime)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, projectResponse)
}

func (handler *projectHandler) UpdateProjectPaid(c echo.Context) error {
	// get payload
	payload := c.Get(auth.AuthorizedPayloadKey).(*auth.JwtCustomPayload)

	if payload.Role != string(db.AccountTypeAdmin) {
		return c.JSON(http.StatusUnauthorized, domain.ErrUnauthorizedResponse(domain.ErrInvalidAuthorization))
	}

	ctx := c.Request().Context()
	var dataRequest domain.ProjectUpdatePaidRequest
	if err := c.Bind(&dataRequest); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequestResponse(err))
	}

	if err := dataRequest.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequestResponse(err))
	}

	pathValue := c.Param("id")
	projectId, err := strconv.Atoi(pathValue)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequestResponse(domain.ErrInvalidProjectId))
	}

	var updateProjectPaid domain.ProjectUpdatePaid
	updateProjectPaid.MapProjectUpdatePaidRequestToProjectUpdatePaid(projectId, dataRequest)

	projectResponse, err := handler.projectUC.UpdateAProjectPaid(ctx, updateProjectPaid)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, projectResponse)
}

func (handler *projectHandler) GetAProject(c echo.Context) error {
	// get payload
	payload := c.Get(auth.AuthorizedPayloadKey).(*auth.JwtCustomPayload)

	ctx := c.Request().Context()

	// get id
	projectId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequestResponse(domain.ErrInvalidProjectId))
	}

	if strings.Compare(payload.Role, string(db.AccountTypeAdmin)) == 0 {
		// use Admin service
		projectResponse, err := handler.projectUC.ListAProject(ctx, projectId)
		if err != nil {
			if errors.Is(err, domain.ErrNoRowsPG) {
				return c.JSON(http.StatusOK, []interface{}{})
			}
			return err
		}
		return c.JSON(http.StatusOK, projectResponse)
	} else if strings.Compare(payload.Role, string(db.AccountTypeClient)) == 0 {
		// use Client service
		projectResponse, err := handler.projectUC.ListAProjectByUserId(ctx, payload.UserId, projectId)
		if err != nil {
			if errors.Is(err, domain.ErrNoRowsPG) {
				return c.JSON(http.StatusOK, []interface{}{})
			}
			return c.JSON(http.StatusNotFound, domain.ErrNotFoundResponse)
		}
		return c.JSON(http.StatusOK, projectResponse)
	} else {
		return c.JSON(http.StatusUnauthorized, domain.ErrUnauthorizedResponse(domain.ErrInvalidAuthorization))
	}
}

func (handler *projectHandler) ListAllProject(c echo.Context) error {
	// get payload
	payload := c.Get(auth.AuthorizedPayloadKey).(*auth.JwtCustomPayload)

	ctx := c.Request().Context()

	if strings.Compare(payload.Role, string(db.AccountTypeAdmin)) == 0 {
		// use Admin service

		// Get query userId if exists
		userIdStr := c.QueryParam("userId")
		if strings.Compare(userIdStr, "") == 0 {
			// List all
			projectsResponse, err := handler.projectUC.ListAllProjects(ctx)
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, projectsResponse)
		} else {
			// List all by userId
			userId, err := strconv.Atoi(userIdStr)
			if err != nil {
				return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequestResponse(domain.ErrInvalidUserAccountId))
			}

			projectsResponse, err := handler.projectUC.ListAllProjectsByUserId(ctx, userId)
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, projectsResponse)
		}

	} else if strings.Compare(payload.Role, string(db.AccountTypeClient)) == 0 {
		// use Client service
		projectsResponse, err := handler.projectUC.ListAllProjectsByUserId(ctx, payload.UserId)
		if err != nil {
			return c.JSON(http.StatusNotFound, domain.ErrNotFoundResponse)
		}
		return c.JSON(http.StatusOK, projectsResponse)
	} else {
		return c.JSON(http.StatusUnauthorized, domain.ErrUnauthorizedResponse(domain.ErrInvalidAuthorization))
	}
}
