package employment

import (
	"github.com/labstack/echo/v4"
	"github.com/ramasuryananda/dummy-cv/internal/constant"
	"github.com/ramasuryananda/dummy-cv/internal/dto/request"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/logger"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/translator"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/writer"
)

func (h *Handler) HandleGetUserEmployment(c echo.Context) error {
	ctx := c.Request().Context()

	var input request.GetEmploymentRequest
	if err := c.Bind(&input); err != nil {
		logger.Error(ctx, nil, err, "c.Bind() error - HandleGetUserEmployment")
		response := writer.APIResponse(constant.ResponseBadRequest.Code, constant.ResponseBadRequest.Description, nil)
		return c.JSON(constant.ResponseBadRequest.Status, response)
	}

	if err := c.Validate(input); err != nil {
		trans := translator.TranslateError(err, request.GetEmploymentRequest{})
		logger.Warning(ctx, nil, err, "c.Validate() error - HandleGetUserEmployment")
		response := writer.APIValidationResponse(nil, trans)
		return c.JSON(constant.ResponseValidationError.Status, response)
	}

	response, status := h.employmentUsecase.GetUserEmployment(ctx, input)

	return c.JSON(status, response)
}

func (h *Handler) HandleCreateEmployment(c echo.Context) error {
	ctx := c.Request().Context()

	var input request.CreateEmploymentRequest
	if err := c.Bind(&input); err != nil {
		logger.Error(ctx, nil, err, "c.Bind() error - HandleCreateEmployment")
		response := writer.APIResponse(constant.ResponseBadRequest.Code, constant.ResponseBadRequest.Description, nil)
		return c.JSON(constant.ResponseBadRequest.Status, response)
	}

	if err := c.Validate(input); err != nil {
		trans := translator.TranslateError(err, request.CreateEmploymentRequest{})
		logger.Warning(ctx, nil, err, "c.Validate() error - HandleCreateEmployment")
		response := writer.APIValidationResponse(nil, trans)
		return c.JSON(constant.ResponseValidationError.Status, response)
	}

	response, status := h.employmentUsecase.CreateEmploymentData(ctx, input)

	return c.JSON(status, response)
}

func (h *Handler) HandleDeleteEmployment(c echo.Context) error {
	ctx := c.Request().Context()

	var input request.DeleteEmploymentRequest
	if err := c.Bind(&input); err != nil {
		logger.Error(ctx, nil, err, "c.Bind() error - HandleDeleteEmployment")
		response := writer.APIResponse(constant.ResponseBadRequest.Code, constant.ResponseBadRequest.Description, nil)
		return c.JSON(constant.ResponseBadRequest.Status, response)
	}

	if err := c.Validate(input); err != nil {
		trans := translator.TranslateError(err, request.DeleteEmploymentRequest{})
		logger.Warning(ctx, nil, err, "c.Validate() error - HandleDeleteEmployment")
		response := writer.APIValidationResponse(nil, trans)
		return c.JSON(constant.ResponseValidationError.Status, response)
	}

	response, status := h.employmentUsecase.DeleteEmploymentData(ctx, input)

	return c.JSON(status, response)
}
