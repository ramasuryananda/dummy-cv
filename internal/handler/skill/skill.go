package skill

import (
	"github.com/labstack/echo/v4"
	"github.com/ramasuryananda/dummy-cv/internal/constant"
	"github.com/ramasuryananda/dummy-cv/internal/dto/request"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/logger"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/translator"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/writer"
)

func (h *Handler) HandleGetUserSkill(c echo.Context) error {
	ctx := c.Request().Context()

	var input request.GetSkillRequest
	if err := c.Bind(&input); err != nil {
		logger.Error(ctx, nil, err, "c.Bind() error - HandleGetUserSkill")
		response := writer.APIResponse(constant.ResponseBadRequest.Code, constant.ResponseBadRequest.Description, nil)
		return c.JSON(constant.ResponseBadRequest.Status, response)
	}

	if err := c.Validate(input); err != nil {
		trans := translator.TranslateError(err, request.GetSkillRequest{})
		logger.Warning(ctx, nil, err, "c.Validate() error - HandleGetUserSkill")
		response := writer.APIValidationResponse(nil, trans)
		return c.JSON(constant.ResponseValidationError.Status, response)
	}

	response, status := h.skillUsecase.GetUserSkill(ctx, input)

	return c.JSON(status, response)
}

func (h *Handler) HandleCreateSkill(c echo.Context) error {
	ctx := c.Request().Context()

	var input request.CreateSkillRequest
	if err := c.Bind(&input); err != nil {
		logger.Error(ctx, nil, err, "c.Bind() error - HandleCreateSkill")
		response := writer.APIResponse(constant.ResponseBadRequest.Code, constant.ResponseBadRequest.Description, nil)
		return c.JSON(constant.ResponseBadRequest.Status, response)
	}

	if err := c.Validate(input); err != nil {
		trans := translator.TranslateError(err, request.CreateSkillRequest{})
		logger.Warning(ctx, nil, err, "c.Validate() error - HandleCreateSkill")
		response := writer.APIValidationResponse(nil, trans)
		return c.JSON(constant.ResponseValidationError.Status, response)
	}

	response, status := h.skillUsecase.CreateSkillData(ctx, input)

	return c.JSON(status, response)
}

func (h *Handler) HandleDeleteSkill(c echo.Context) error {
	ctx := c.Request().Context()

	var input request.DeleteSkillRequest
	if err := c.Bind(&input); err != nil {
		logger.Error(ctx, nil, err, "c.Bind() error - HandleDeleteSkill")
		response := writer.APIResponse(constant.ResponseBadRequest.Code, constant.ResponseBadRequest.Description, nil)
		return c.JSON(constant.ResponseBadRequest.Status, response)
	}

	if err := c.Validate(input); err != nil {
		trans := translator.TranslateError(err, request.DeleteSkillRequest{})
		logger.Warning(ctx, nil, err, "c.Validate() error - HandleDeleteSkill")
		response := writer.APIValidationResponse(nil, trans)
		return c.JSON(constant.ResponseValidationError.Status, response)
	}

	response, status := h.skillUsecase.DeleteSkillData(ctx, input)

	return c.JSON(status, response)
}
