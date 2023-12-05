package httpuseracc

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"project-management/domain"
)

type AccountUserHandler struct {
	AccountUserUC domain.AccountUseCase
}

func SetupAccountUserHandler(group *echo.Group, accountUserHandler domain.AccountUseCase) {
	handler := &AccountUserHandler{
		AccountUserUC: accountUserHandler,
	}

	group.POST("/account", handler.CreateUserAccountHandler)
	//router.Patch("/account", handler.UpdateUserAccountHandler)
	//router.Patch("/account", handler.UpdateUserAccountHandler)
}

/*
	Handlers
*/

func (handler *AccountUserHandler) CreateUserAccountHandler(c echo.Context) error {
	log.Info("CreateUserAccountHandler")
	ctx := c.Request().Context()
	data := &domain.AccountCreateAndLoginRequest{}

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequest(err))
	}

	accountResponse, err := handler.AccountUserUC.CreateUserAccount(ctx, data.Username, data.Password)
	if err != nil {
		if errors.Is(err, domain.ErrUsernameExists) {
			return c.JSON(http.StatusConflict, domain.ErrResourceConflict(err))
		}
		return err
	}
	fmt.Println(c.Request().Context().Value("api.version"))
	return c.JSON(http.StatusOK, accountResponse)
}

//func (handler *AccountUserHandler) UpdateUserAccountHandler(w http.ResponseWriter, r *http.Request) {
//	ctx := r.Context()
//	data := &domain.AccountUpdateRequest{}
//
//	if err := render.Bind(r, data); err != nil {
//		err := render.Render(w, r, domain.ErrInvalidRequest(err))
//		if err != nil {
//			log.Fatalf("Render Response failed: %s", err.Error())
//			return
//		}
//		return
//	}
//
//}
