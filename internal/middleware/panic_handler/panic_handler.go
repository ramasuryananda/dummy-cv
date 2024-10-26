package panichandler

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/ramasuryananda/dummy-cv/internal/constant"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/logger"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/writer"
)

type Field struct {
	Title string
	Value string
}

// HandlePanic handle panic and send error message to Discord.
func (middleware *Middleware) HandlePanic() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			defer func() {
				// Recover the panic if exists
				r := recover()
				if r == nil {
					return
				}
				// Get the body
				body, err := io.ReadAll(c.Request().Body)
				if err != nil {
					c.String(http.StatusInternalServerError, "Error reading request body")
					return
				}

				// Convert the body to a string
				bodyString := string(body)
				if bodyString == "" {
					bodyString = "-"
				}

				// Get the query params
				queryParams := c.QueryString()
				if queryParams == "" {
					queryParams = "-"
				}

				// Get the stack trace information
				stackTrace := make([]uintptr, 1)
				n := runtime.Callers(3, stackTrace[:])
				stackFrames := runtime.CallersFrames(stackTrace[:n])
				frame, _ := stackFrames.Next()

				// Create fields for the mattermost message
				fields := []Field{
					{
						Title: "IP",
						Value: c.RealIP(),
					},
					{
						Title: "Method",
						Value: c.Request().Method,
					},
					{
						Title: "Path",
						Value: c.Request().URL.String(),
					},
					{
						Title: "Function",
						Value: frame.Function,
					},
					{
						Title: "Line",
						Value: fmt.Sprintf("%v", frame.Line),
					},
					{
						Title: "File",
						Value: frame.File,
					},
					{
						Title: "Query Params",
						Value: queryParams,
					},
					{
						Title: "Body Payload",
						Value: bodyString,
					},
				}

				errMsg := fmt.Sprintf("%v", r)
				message := "**[PANIC]** " + errMsg
				logger.Error(context.Background(), fields, errors.New(message), "panic error")

				c.JSON(constant.ResponseInternalServerError.Status, writer.APIResponse(constant.ResponseInternalServerError.Code, constant.ResponseInternalServerError.Description, nil))
			}()

			return next(c)
		}
	}
}
