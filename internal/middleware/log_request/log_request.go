package log_request

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

func (middleware *Middleware) LogRequest() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Printf("%v - %v %v\n", time.Now().Format("2006-01-02 15:04:05"), c.Request().Method, c.Path())
			return next(c)
		}
	}
}
