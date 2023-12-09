package httpuseracc

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"project-management/common"
	"project-management/domain"
	"strconv"
)

type accountUserHandler struct {
	appCtx        common.AppContext
	accountUserUC domain.AccountUseCase
}

func SetupAccountUserHandler(group *echo.Group, appContext common.AppContext, accountUserUC domain.AccountUseCase) {
	handler := &accountUserHandler{
		appCtx:        appContext,
		accountUserUC: accountUserUC,
	}

	g := group.Group("/account")
	g.PATCH("/:id", handler.UpdateUserAccountHandler)
	g.POST("/login", handler.LoginUserAccountHandler)
	g.POST("/", handler.CreateUserAccountHandler)
}

/*
	Handlers
*/

func (handler *accountUserHandler) CreateUserAccountHandler(c echo.Context) error {
	ctx := c.Request().Context()
	var data domain.AccountCreateAndLoginRequest

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequestResponse(err))
	}

	accountResponse, err := handler.accountUserUC.CreateUserAccount(ctx, data.Username, data.Password)
	if err != nil {
		if errors.Is(err, domain.ErrUsernameExists) {
			return c.JSON(http.StatusConflict, domain.ErrResourceConflictResponse(err))
		}
		return err
	}
	return c.JSON(http.StatusOK, accountResponse)
}

func (handler *accountUserHandler) UpdateUserAccountHandler(c echo.Context) error {

	ctx := c.Request().Context()

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequestResponse(domain.ErrInvalidUserAccountId))
	}

	var data domain.AccountUpdateRequest

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequestResponse(err))
	}
	if err := c.Validate(data); err != nil {
		return err
	}

	// Prepare data to process
	accountToUpdate := domain.AccountUpdate{}
	accountToUpdate.MapAccountUpdateRequestToAccountUpdate(userId, data)

	// Process data
	updateUserAccount, err := handler.accountUserUC.UpdateUserAccount(ctx, accountToUpdate)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, updateUserAccount)

}

func (handler *accountUserHandler) LoginUserAccountHandler(c echo.Context) error {
	handler.appCtx.Logger.Info("LoginUserAccountHandler", zap.String("Begin", "Here"))
	ctx := c.Request().Context()
	var data domain.AccountCreateAndLoginRequest

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequestResponse(err))
	}

	accountResponse, err := handler.accountUserUC.LoginAccount(ctx, data.Username, data.Password)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidLogin) {
			return c.JSON(http.StatusUnauthorized, domain.ErrInvalidLoginResponse(err))
		}
	}

	return c.JSON(http.StatusOK, accountResponse)
}
