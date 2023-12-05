package glbmiddleware

import (
	"context"
	"github.com/labstack/echo/v4"
)

func ApiVersionCtxMiddleware(version string) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request().WithContext(context.WithValue(c.Request().Context(), "api.version", version))
			c.SetRequest(r)
			return next(c)
		}
	}
}
