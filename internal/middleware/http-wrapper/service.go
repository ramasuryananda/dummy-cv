package httpwrapper

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/logger"
)

// HttpWrapper wraps handlers for HTTP service.
func (middleware *Middleware) HttpWrapper() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := logger.InitLogCtx(context.Background())
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
