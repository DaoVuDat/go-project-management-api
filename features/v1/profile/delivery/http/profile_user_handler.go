package httpprofileuser

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project-management/auth"
	"project-management/common"
	"project-management/domain"
)

type ProfileUserHandler struct {
	appCtx        common.AppContext
	profileUserUC domain.UserProfileUseCase
}

func SetupProfileUserHandler(group *echo.Group,
	ctx common.AppContext,
	userProfileUC domain.UserProfileUseCase,
) {
	handler := &ProfileUserHandler{
		appCtx:        ctx,
		profileUserUC: userProfileUC,
	}
	g := group.Group("/user")
	g.Use(auth.AuthorizationRestrictedMiddleware(ctx.GbConfig.TokenPrivateKey))
	g.GET("/profile", handler.GetUserProfileHandler)
}

func (handler *ProfileUserHandler) GetUserProfileHandler(c echo.Context) error {
	// get PayLoad
	payload := c.Get(auth.AuthorizedPayloadKey)
	//ctx := c.Request().Context()

	return c.JSON(http.StatusOK, payload)
}
