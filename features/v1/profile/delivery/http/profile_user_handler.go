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
	g.PATCH("/profile", handler.UpdateUserProfileHandler)
	g.PATCH("/profile-image", handler.UpdateUserProfileUrlImageHandler)
}

func (handler *ProfileUserHandler) GetUserProfileHandler(c echo.Context) error {
	// get PayLoad
	payload := c.Get(auth.AuthorizedPayloadKey).(*auth.JwtCustomPayload)
	ctx := c.Request().Context()

	// get user profile
	userProfileResponse, err := handler.profileUserUC.GetUserProfile(ctx, payload.UserId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, userProfileResponse)
}

func (handler *ProfileUserHandler) UpdateUserProfileHandler(c echo.Context) error {
	payload := c.Get(auth.AuthorizedPayloadKey).(*auth.JwtCustomPayload)
	ctx := c.Request().Context()

	var data domain.UserProfileUpdateRequest
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequestResponse(err))
	}

	if err := data.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidRequestResponse(err))
	}

	var updateUserProfile domain.UserProfileUpdate
	updateUserProfile.MapUserProfileUpdateRequestToUserProfileUpdate(payload.UserId, data)

	userProfileResponse, err := handler.profileUserUC.UpdateUserProfile(ctx, updateUserProfile)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, userProfileResponse)
}

func (handler *ProfileUserHandler) UpdateUserProfileUrlImageHandler(c echo.Context) error {
	payload := c.Get(auth.AuthorizedPayloadKey).(*auth.JwtCustomPayload)
	ctx := c.Request().Context()

	var data domain.UserProfileImageUrlRequest
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrInvalidLoginResponse(err))
	}

	var updateUserProfileImageUrl domain.UserProfileImageUrlUpdate
	updateUserProfileImageUrl.MapUserProfileImageUrlUpdateRequestToUserProfileImageUrlUpdate(payload.UserId, data)

	userProfileResponse, err := handler.profileUserUC.UpdateUserProfileImageUrl(ctx, updateUserProfileImageUrl)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, userProfileResponse)
}
