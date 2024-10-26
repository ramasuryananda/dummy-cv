package profile

import (
	"github.com/labstack/echo/v4"
	"github.com/ramasuryananda/dummy-cv/internal/constant"
	"github.com/ramasuryananda/dummy-cv/internal/dto/request"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/logger"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/translator"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/writer"
)

func (h *Handler) HanmdleGetProfile(c echo.Context) error {
	ctx := c.Request().Context()

	var input request.GetUserProfileRequest
	if err := c.Bind(&input); err != nil {
		logger.Error(ctx, nil, err, "c.Bind() error - HandleUpdateProfile")
		response := writer.APIResponse(constant.ResponseBadRequest.Code, constant.ResponseBadRequest.Description, nil)
		return c.JSON(constant.ResponseBadRequest.Status, response)
	}

	if err := c.Validate(input); err != nil {
		trans := translator.TranslateError(err, request.GetUserProfileRequest{})
		logger.Warning(ctx, nil, err, "c.Validate() error - HandleUpdateProfile")
		response := writer.APIValidationResponse(nil, trans)
		return c.JSON(constant.ResponseValidationError.Status, response)
	}

	response, status := h.profileUsecase.GetUserProfile(ctx, input.ProfileCode)

	return c.JSON(status, response)
}

func (handler *Handler) HandleCreateProfile(c echo.Context) error {
	ctx := c.Request().Context()

	var input request.CreateProfileRequest
	if err := c.Bind(&input); err != nil {
		logger.Error(ctx, nil, err, "c.Bind() error - HandleCreateProfile")
		response := writer.APIResponse(constant.ResponseBadRequest.Code, constant.ResponseBadRequest.Description, nil)
		return c.JSON(constant.ResponseBadRequest.Status, response)
	}

	if err := c.Validate(input); err != nil {
		trans := translator.TranslateError(err, request.CreateProfileRequest{})
		logger.Warning(ctx, nil, err, "c.Validate() error - HandleCreateProfile")
		response := writer.APIValidationResponse(nil, trans)
		return c.JSON(constant.ResponseValidationError.Status, response)
	}

	resp, status := handler.profileUsecase.CreateUserProfile(ctx, input)
	return c.JSON(status, resp)
}

func (handler *Handler) HandleUpdateProfile(c echo.Context) error {
	ctx := c.Request().Context()

	var input request.UpdateProfileRequest
	if err := c.Bind(&input); err != nil {
		logger.Error(ctx, nil, err, "c.Bind() error - HandleUpdateProfile")
		response := writer.APIResponse(constant.ResponseBadRequest.Code, constant.ResponseBadRequest.Description, nil)
		return c.JSON(constant.ResponseBadRequest.Status, response)
	}

	if err := c.Validate(input); err != nil {
		trans := translator.TranslateError(err, request.UpdateProfileRequest{})
		logger.Warning(ctx, nil, err, "c.Validate() error - HandleUpdateProfile")
		response := writer.APIValidationResponse(nil, trans)
		return c.JSON(constant.ResponseValidationError.Status, response)
	}

	resp, status := handler.profileUsecase.UpdateUserProfile(ctx, input)
	return c.JSON(status, resp)
}
