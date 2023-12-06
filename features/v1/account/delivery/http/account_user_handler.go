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

	group.PATCH("/account/:id", handler.UpdateUserAccountHandler)
	group.POST("/account", handler.CreateUserAccountHandler)
	//router.Patch("/account", handler.UpdateUserAccountHandler)
}

/*
	Handlers
*/

func (handler *accountUserHandler) CreateUserAccountHandler(c echo.Context) error {
	handler.appCtx.Logger.Debug("CreateUserAccountHandler")
	//handler.appCtx.Logger.Info("CreateUserAccountHandler", zap.Any("api.version", c.Request().Context().Value("api.version")))
	ctx := c.Request().Context()
	var data domain.AccountCreateAndLoginRequest

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequest(err))
	}

	accountResponse, err := handler.accountUserUC.CreateUserAccount(ctx, data.Username, data.Password)
	if err != nil {
		if errors.Is(err, domain.ErrUsernameExists) {
			return c.JSON(http.StatusConflict, domain.ErrResourceConflict(err))
		}
		return err
	}
	return c.JSON(http.StatusOK, accountResponse)
}

func (handler *accountUserHandler) UpdateUserAccountHandler(c echo.Context) error {

	ctx := c.Request().Context()

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequest(domain.ErrInvalidUserAccountId))
	}

	handler.appCtx.Logger.Debug("UpdateUserAccountHandler", zap.Int("userId", userId))
	var data domain.AccountUpdateRequest

	if err := c.Bind(&data); err != nil {
		handler.appCtx.Logger.Debug("UpdateUserAccountHandler", zap.String("error", "binding error"))
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequest(err))
	}
	if err := c.Validate(data); err != nil {
		handler.appCtx.Logger.Debug("UpdateUserAccountHandler", zap.Any("data", data))
		handler.appCtx.Logger.Debug("UpdateUserAccountHandler", zap.Any("err", err))

		return err
	}
	handler.appCtx.Logger.Debug("UpdateUserAccountHandler", zap.Any("data", data))
	handler.appCtx.Logger.Debug("UpdateUserAccountHandler", zap.Any("type", data.Type))
	handler.appCtx.Logger.Debug("UpdateUserAccountHandler", zap.Any("status", data.Status))
	handler.appCtx.Logger.Debug("UpdateUserAccountHandler", zap.Any("password", data.Password))

	// Prepare data to process
	accountToUpdate := domain.AccountUpdate{}
	accountToUpdate.MapAccountUpdateRequestToAccountUpdate(userId, data)

	// Process data
	updateUserAccount, err := handler.accountUserUC.UpdateUserAccount(ctx, accountToUpdate)
	if err != nil {
		handler.appCtx.Logger.Debug("UpdateUserAccountHandler", zap.String("error", "update user account error"))
		return err
	}

	handler.appCtx.Logger.Debug("UpdateUserAccountHandler", zap.String("success", "updated account"))
	return c.JSON(http.StatusOK, updateUserAccount)

}
