package profile_photo

import (
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ramasuryananda/dummy-cv/internal/constant"
	"github.com/ramasuryananda/dummy-cv/internal/dto/entity"
	"github.com/ramasuryananda/dummy-cv/internal/dto/request"
	"github.com/ramasuryananda/dummy-cv/internal/dto/response"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/writer"
)

func (uc *UseCase) UpsertUserPhotoProfile(ctx context.Context, req request.UpsertPhotoProfileRequest) (resp writer.Response, httpStatus int) {

	userProfileData, err := uc.profileRepository.GetUserByProfileCode(ctx, req.ProfileCode)
	if err != nil {
		if errors.Is(err, constant.ErrorDatabaseNotFound) {
			resp = writer.APIErrorResponse(constant.ResponseErrorNotFound.Code, constant.ResponseErrorNotFound.Description, err)
			httpStatus = constant.ResponseErrorNotFound.Status
			return
		}

		resp = writer.APIErrorResponse(constant.ResponseInternalServerError.Code, constant.ResponseInternalServerError.Description, err)
		httpStatus = constant.ResponseInternalServerError.Status
		return
	}
	base64Split := strings.Split(req.Base64Image, ",")

	// Decode the Base64 data
	decodedData, err := base64.StdEncoding.DecodeString(base64Split[1])
	if err != nil {
		resp = writer.APIErrorResponse(constant.ResponseInvalidBase64.Code, constant.ResponseInvalidBase64.Description, err)
		httpStatus = constant.ResponseInvalidBase64.Status
		return
	}

	filename := fmt.Sprintf("./app/upload/photo/%d.png", userProfileData.ProfileCode)

	err = os.MkdirAll("./app/upload/photo/", os.ModePerm) // Create directory if it doesn't exist
	if err != nil {
		resp = writer.APIErrorResponse(constant.ResponseInternalServerError.Code, constant.ResponseInternalServerError.Description, err)
		httpStatus = constant.ResponseInternalServerError.Status
		return
	}

	// Write the decoded data to a file
	err = os.WriteFile(filename, decodedData, 0644)
	if err != nil {
		resp = writer.APIErrorResponse(constant.ResponseInternalServerError.Code, constant.ResponseInternalServerError.Description, err)
		httpStatus = constant.ResponseInternalServerError.Status
		return
	}

	userProfilePhoto, err := uc.profilePhotoRepository.GetUserProfilePhotoByProfileCode(ctx, req.ProfileCode)
	if err != nil {
		if !errors.Is(err, constant.ErrorDatabaseNotFound) {
			resp = writer.APIErrorResponse(constant.ResponseInternalServerError.Code, constant.ResponseInternalServerError.Description, err)
			httpStatus = constant.ResponseInternalServerError.Status
			return
		}

	}

	profilePhotoData := entity.ProfilePhoto{
		ID:          userProfilePhoto.ID,
		ProfileCode: userProfileData.ProfileCode,
		PhotoURL:    filename,
		CreatedAt:   userProfilePhoto.CreatedAt,
		UpdatedAt:   sql.NullTime{Valid: true, Time: time.Now()},
	}

	err = uc.profilePhotoRepository.SaveUserProfilePhoto(ctx, profilePhotoData)
	if err != nil {
		resp = writer.APIErrorResponse(constant.ResponseInternalServerError.Code, constant.ResponseInternalServerError.Description, err)
		httpStatus = constant.ResponseInternalServerError.Status
		return
	}

	respData := response.UpdatePhotoProfileResponse{
		ProfileCode: userProfileData.ProfileCode,
		PhotoURL:    strings.TrimLeft(filename, "."),
	}

	resp = writer.APIResponse(constant.ResponseSuccess.Code, constant.ResponseSuccess.Description, respData)
	httpStatus = constant.ResponseSuccess.Status
	return
}

func (uc *UseCase) DownloadPhotoProfile(ctx context.Context, req request.DownloadPhotoProfileRequest) (resp writer.Response, httpStatus int) {
	userProfilePhoto, err := uc.profilePhotoRepository.GetUserProfilePhotoByProfileCode(ctx, req.ProfileCode)
	if err != nil {
		if errors.Is(err, constant.ErrorDatabaseNotFound) {
			resp = writer.APIErrorResponse(constant.ResponseErrorNotFound.Code, constant.ResponseErrorNotFound.Description, err)
			httpStatus = constant.ResponseErrorNotFound.Status
			return
		}
		resp = writer.APIErrorResponse(constant.ResponseInternalServerError.Code, constant.ResponseInternalServerError.Description, err)
		httpStatus = constant.ResponseInternalServerError.Status
		return
	}

	data, err := os.ReadFile(userProfilePhoto.PhotoURL)
	if err != nil {
		resp = writer.APIErrorResponse(constant.ResponseInternalServerError.Code, constant.ResponseInternalServerError.Description, err)
		httpStatus = constant.ResponseInternalServerError.Status
		return
	}

	base64Data := base64.StdEncoding.EncodeToString(data)

	respData := response.DowndloadPhotoProfile{
		Base64String: "data:image/png;base64," + base64Data,
		PhotoURL:     userProfilePhoto.PhotoURL,
	}

	resp = writer.APIResponse(constant.ResponseSuccess.Code, constant.ResponseSuccess.Description, respData)
	httpStatus = constant.ResponseSuccess.Status
	return
}

func (uc *UseCase) DeletePhotoProfile(ctx context.Context, req request.DeletePhotoProfileRequest) (resp writer.Response, httpStatus int) {
	userProfilePhoto, err := uc.profilePhotoRepository.GetUserProfilePhotoByProfileCode(ctx, req.ProfileCode)
	if err != nil {
		if errors.Is(err, constant.ErrorDatabaseNotFound) {
			resp = writer.APIErrorResponse(constant.ResponseErrorNotFound.Code, constant.ResponseErrorNotFound.Description, err)
			httpStatus = constant.ResponseErrorNotFound.Status
			return
		}
		resp = writer.APIErrorResponse(constant.ResponseInternalServerError.Code, constant.ResponseInternalServerError.Description, err)
		httpStatus = constant.ResponseInternalServerError.Status
		return
	}

	err = uc.profilePhotoRepository.DeleteUserProfilePhoto(ctx, userProfilePhoto)
	if err != nil {
		resp = writer.APIErrorResponse(constant.ResponseInternalServerError.Code, constant.ResponseInternalServerError.Description, err)
		httpStatus = constant.ResponseInternalServerError.Status
		return
	}

	err = os.Remove(userProfilePhoto.PhotoURL)
	if err != nil {
		resp = writer.APIErrorResponse(constant.ResponseInternalServerError.Code, constant.ResponseInternalServerError.Description, err)
		httpStatus = constant.ResponseInternalServerError.Status
		return
	}

	respData := response.DeletePhotoProfileResponse{
		ProfileCode: userProfilePhoto.ProfileCode,
	}

	resp = writer.APIResponse(constant.ResponseSuccess.Code, constant.ResponseSuccess.Description, respData)
	httpStatus = constant.ResponseSuccess.Status
	return
}
