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
	RegisterProfilePhotoRoutes(router)
	RegisterWorkingExperienceRoutes(router)
	RegisterEmploymentRoutes(router)
	RegisterEducationRoute(router)
	RegisterSkillRoute(router)

	return router
}

func RegisterProfileRoutes(router *Router) {
	profileGroup := router.Echo.Group("api/profile")
	profileGroup.GET("/:code", router.Handler.Profile.HanmdleGetProfile)
	profileGroup.POST("", router.Handler.Profile.HandleCreateProfile)
	profileGroup.PUT("/:code", router.Handler.Profile.HandleUpdateProfile)

}

func RegisterProfilePhotoRoutes(router *Router) {
	profilePhotoGroup := router.Echo.Group("api/photo")
	profilePhotoGroup.PUT("/:code", router.Handler.ProfilePhoto.HandleUpsertPhotoProfile)
	profilePhotoGroup.GET("/:code", router.Handler.ProfilePhoto.HandleDownloadPhotoData)
	profilePhotoGroup.DELETE("/:code", router.Handler.ProfilePhoto.HandleDeleteProfilePhoto)

}

func RegisterWorkingExperienceRoutes(router *Router) {
	workingExperienceGroup := router.Echo.Group("api/working-experience")
	workingExperienceGroup.PUT("/:code", router.Handler.WorkingExperience.HandleUpsertWorkingExperience)
	workingExperienceGroup.GET("/:code", router.Handler.WorkingExperience.HandleGetWorkingExperience)
}

func RegisterEmploymentRoutes(router *Router) {
	employmentGroup := router.Echo.Group("api/employment")
	employmentGroup.GET("/:code", router.Handler.Employment.HandleGetUserEmployment)
	employmentGroup.POST("/:code", router.Handler.Employment.HandleCreateEmployment)
	employmentGroup.DELETE("/:code", router.Handler.Employment.HandleDeleteEmployment)
}

func RegisterEducationRoute(router *Router) {
	educationGroup := router.Echo.Group("api/education")
	educationGroup.GET("/:code", router.Handler.Education.HandleGetUserEducation)
	educationGroup.POST("/:code", router.Handler.Education.HandleCreateEducation)
	educationGroup.DELETE("/:code", router.Handler.Education.HandleDeleteEducation)
}

func RegisterSkillRoute(router *Router) {
	skillGroup := router.Echo.Group("api/skill")
	skillGroup.GET("/:code", router.Handler.Skill.HandleGetUserSkill)
	skillGroup.POST("/:code", router.Handler.Skill.HandleCreateSkill)
	skillGroup.DELETE("/:code", router.Handler.Skill.HandleDeleteSkill)
}
