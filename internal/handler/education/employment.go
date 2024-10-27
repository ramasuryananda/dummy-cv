package education

import (
	"github.com/labstack/echo/v4"
	"github.com/ramasuryananda/dummy-cv/internal/constant"
	"github.com/ramasuryananda/dummy-cv/internal/dto/request"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/logger"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/translator"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/writer"
)

func (h *Handler) HandleGetUserEducation(c echo.Context) error {
	ctx := c.Request().Context()

	var input request.GetEducationRequest
	if err := c.Bind(&input); err != nil {
		logger.Error(ctx, nil, err, "c.Bind() error - HandleGetUserEducation")
		response := writer.APIResponse(constant.ResponseBadRequest.Code, constant.ResponseBadRequest.Description, nil)
		return c.JSON(constant.ResponseBadRequest.Status, response)
	}

	if err := c.Validate(input); err != nil {
		trans := translator.TranslateError(err, request.GetEducationRequest{})
		logger.Warning(ctx, nil, err, "c.Validate() error - HandleGetUserEducation")
		response := writer.APIValidationResponse(nil, trans)
		return c.JSON(constant.ResponseValidationError.Status, response)
	}

	response, status := h.educationUsecase.GetUserEducation(ctx, input)

	return c.JSON(status, response)
}

func (h *Handler) HandleCreateEducation(c echo.Context) error {
	ctx := c.Request().Context()

	var input request.CreateEducationRequest
	if err := c.Bind(&input); err != nil {
		logger.Error(ctx, nil, err, "c.Bind() error - HandleCreateEducation")
		response := writer.APIResponse(constant.ResponseBadRequest.Code, constant.ResponseBadRequest.Description, nil)
		return c.JSON(constant.ResponseBadRequest.Status, response)
	}

	if err := c.Validate(input); err != nil {
		trans := translator.TranslateError(err, request.CreateEducationRequest{})
		logger.Warning(ctx, nil, err, "c.Validate() error - HandleCreateEducation")
		response := writer.APIValidationResponse(nil, trans)
		return c.JSON(constant.ResponseValidationError.Status, response)
	}

	response, status := h.educationUsecase.CreateEducationData(ctx, input)

	return c.JSON(status, response)
}

func (h *Handler) HandleDeleteEducation(c echo.Context) error {
	ctx := c.Request().Context()

	var input request.DeleteEducationRequest
	if err := c.Bind(&input); err != nil {
		logger.Error(ctx, nil, err, "c.Bind() error - HandleDeleteEducation")
		response := writer.APIResponse(constant.ResponseBadRequest.Code, constant.ResponseBadRequest.Description, nil)
		return c.JSON(constant.ResponseBadRequest.Status, response)
	}

	if err := c.Validate(input); err != nil {
		trans := translator.TranslateError(err, request.DeleteEducationRequest{})
		logger.Warning(ctx, nil, err, "c.Validate() error - HandleDeleteEducation")
		response := writer.APIValidationResponse(nil, trans)
		return c.JSON(constant.ResponseValidationError.Status, response)
	}

	response, status := h.educationUsecase.DeleteEducationData(ctx, input)

	return c.JSON(status, response)
}
