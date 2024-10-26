package working_experience

import (
	"github.com/labstack/echo/v4"
	"github.com/ramasuryananda/dummy-cv/internal/constant"
	"github.com/ramasuryananda/dummy-cv/internal/dto/request"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/logger"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/translator"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/writer"
)

func (handler *Handler) HandleUpsertWorkingExperience(c echo.Context) error {
	ctx := c.Request().Context()

	var input request.UpsertWorkingExperienceRequest
	if err := c.Bind(&input); err != nil {
		logger.Error(ctx, nil, err, "c.Bind() error - HandleUpsertWorkingExperience")
		response := writer.APIResponse(constant.ResponseBadRequest.Code, constant.ResponseBadRequest.Description, nil)
		return c.JSON(constant.ResponseBadRequest.Status, response)
	}

	if err := c.Validate(input); err != nil {
		trans := translator.TranslateError(err, request.UpsertWorkingExperienceRequest{})
		logger.Warning(ctx, nil, err, "c.Validate() error - HandleUpsertWorkingExperience")
		response := writer.APIValidationResponse(nil, trans)
		return c.JSON(constant.ResponseValidationError.Status, response)
	}

	resp, status := handler.workingExperienceUsecase.UpsertUserWorkingExperience(ctx, input)
	return c.JSON(status, resp)
}

func (handler *Handler) HandleGetWorkingExperience(c echo.Context) error {
	ctx := c.Request().Context()

	var input request.GetUserWorkingExperienceRequest
	if err := c.Bind(&input); err != nil {
		logger.Error(ctx, nil, err, "c.Bind() error - HandleGetWorkingExperience")
		response := writer.APIResponse(constant.ResponseBadRequest.Code, constant.ResponseBadRequest.Description, nil)
		return c.JSON(constant.ResponseBadRequest.Status, response)
	}

	if err := c.Validate(input); err != nil {
		trans := translator.TranslateError(err, request.GetUserWorkingExperienceRequest{})
		logger.Warning(ctx, nil, err, "c.Validate() error - HandleGetWorkingExperience")
		response := writer.APIValidationResponse(nil, trans)
		return c.JSON(constant.ResponseValidationError.Status, response)
	}

	resp, status := handler.workingExperienceUsecase.GetUserWorkingExperience(ctx, input)
	return c.JSON(status, resp)
}
