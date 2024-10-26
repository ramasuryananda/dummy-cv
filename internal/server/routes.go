package server

import (
	app "github.com/ramasuryananda/dummy-cv/internal/app"

	"github.com/labstack/echo/v4"
)

type Router struct {
	Echo       *echo.Echo
	Handler    *app.Handlers
	Middleware *app.Middleware
}

func NewRouter(e *echo.Echo, handler *app.Handlers, middleware *app.Middleware) *Router {
	router := &Router{
		Echo:       e,
		Handler:    handler,
		Middleware: middleware,
	}
	e.Use(middleware.LogRequest.LogRequest())
	e.Use(middleware.PanicHandler.HandlePanic())
	e.Use(middleware.HttpWrapper.HttpWrapper())

	RegisterProfileRoutes(router)

	return router
}

func RegisterProfileRoutes(router *Router) {
	profileGroup := router.Echo.Group("api/profile")
	profileGroup.GET("/:code", router.Handler.Profile.HanmdleGetProfile)
	profileGroup.POST("", router.Handler.Profile.HandleCreateProfile)
	profileGroup.PUT("/:code", router.Handler.Profile.HandleUpdateProfile)

}
