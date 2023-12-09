package auth

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"project-management/domain"
	"strings"
)

const (
	authorizationHeader     = "authorization"
	authorizationTypeBearer = "bearer"
	AuthorizedPayloadKey    = "authorization_payload"
)

func AuthorizationRestrictedMiddleware(signingKey string) func(next echo.HandlerFunc) echo.HandlerFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorizationHeader := c.Request().Header.Get(authorizationHeader)
			if len(authorizationHeader) == 0 {
				err := errors.New("authorization header is not provided")
				return c.JSON(http.StatusUnauthorized, domain.ErrUnauthorizedResponse(err))
			}

			fields := strings.Fields(authorizationHeader)
			if len(fields) < 2 {
				err := errors.New("invalid authorization header format")
				return c.JSON(http.StatusUnauthorized, domain.ErrUnauthorizedResponse(err))
			}

			authorizationType := strings.ToLower(fields[0])
			if authorizationType != authorizationTypeBearer {
				err := fmt.Errorf("unsupported authorization type %s", authorizationType)
				return c.JSON(http.StatusUnauthorized, domain.ErrUnauthorizedResponse(err))
			}

			accessToken := fields[1]
			payload, err := VerifyToken(accessToken, signingKey)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, domain.ErrUnauthorizedResponse(err))
			}

			c.Set(AuthorizedPayloadKey, payload)
			return next(c)
		}
	}
}
