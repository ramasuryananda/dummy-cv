package profile_photo

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ramasuryananda/dummy-cv/internal/constant"
	"github.com/ramasuryananda/dummy-cv/internal/dto/request"
	"github.com/ramasuryananda/dummy-cv/internal/dto/response"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/logger"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/translator"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/writer"
)

func (handler *Handler) HandleUpsertPhotoProfile(c echo.Context) error {
	ctx := c.Request().Context()

	var input request.UpsertPhotoProfileRequest
	if err := c.Bind(&input); err != nil {
		logger.Error(ctx, nil, err, "c.Bind() error - HandleUpsertPhotoProfile")
		response := writer.APIResponse(constant.ResponseBadRequest.Code, constant.ResponseBadRequest.Description, nil)
		return c.JSON(constant.ResponseBadRequest.Status, response)
	}

	if err := c.Validate(input); err != nil {
		trans := translator.TranslateError(err, request.UpsertPhotoProfileRequest{})
		logger.Warning(ctx, nil, err, "c.Validate() error - HandleUpsertPhotoProfile")
		response := writer.APIValidationResponse(nil, trans)
		return c.JSON(constant.ResponseValidationError.Status, response)
	}

	resp, status := handler.photoProfileUsecase.UpsertUserPhotoProfile(ctx, input)
	return c.JSON(status, resp)
}

func (handler *Handler) HandleDownloadPhotoData(c echo.Context) error {
	ctx := c.Request().Context()

	var input request.DownloadPhotoProfileRequest
	if err := c.Bind(&input); err != nil {
		logger.Error(ctx, nil, err, "c.Bind() error - HandleDownloadPhotoData")
		response := writer.APIResponse(constant.ResponseBadRequest.Code, constant.ResponseBadRequest.Description, nil)
		return c.JSON(constant.ResponseBadRequest.Status, response)
	}

	if err := c.Validate(input); err != nil {
		trans := translator.TranslateError(err, request.DownloadPhotoProfileRequest{})
		logger.Warning(ctx, nil, err, "c.Validate() error - HandleDownloadPhotoData")
		response := writer.APIValidationResponse(nil, trans)
		return c.JSON(constant.ResponseValidationError.Status, response)
	}

	resp, status := handler.photoProfileUsecase.DownloadPhotoProfile(ctx, input)
	if status != http.StatusOK {
		return c.JSON(status, resp)
	}

	data := resp.Data.(response.DowndloadPhotoProfile)

	c.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.base64", "image"))
	c.Response().Header().Set("Content-Type", "text/plain")

	// Write the Base64 data to the response
	return c.String(http.StatusOK, data.Base64String)

}

func (handler *Handler) HandleDeleteProfilePhoto(c echo.Context) error {
	ctx := c.Request().Context()

	var input request.DeletePhotoProfileRequest
	if err := c.Bind(&input); err != nil {
		logger.Error(ctx, nil, err, "c.Bind() error - HandleDeleteProfilePhoto")
		response := writer.APIResponse(constant.ResponseBadRequest.Code, constant.ResponseBadRequest.Description, nil)
		return c.JSON(constant.ResponseBadRequest.Status, response)
	}

	if err := c.Validate(input); err != nil {
		trans := translator.TranslateError(err, request.DeletePhotoProfileRequest{})
		logger.Warning(ctx, nil, err, "c.Validate() error - HandleDeleteProfilePhoto")
		response := writer.APIValidationResponse(nil, trans)
		return c.JSON(constant.ResponseValidationError.Status, response)
	}

	resp, status := handler.photoProfileUsecase.DeletePhotoProfile(ctx, input)
	return c.JSON(status, resp)

}
