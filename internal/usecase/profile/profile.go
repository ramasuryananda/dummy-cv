package profile

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ramasuryananda/dummy-cv/internal/constant"
	"github.com/ramasuryananda/dummy-cv/internal/dto/entity"
	"github.com/ramasuryananda/dummy-cv/internal/dto/general"
	"github.com/ramasuryananda/dummy-cv/internal/dto/request"
	"github.com/ramasuryananda/dummy-cv/internal/dto/response"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/writer"
)

func (uc *UseCase) GetUserProfile(ctx context.Context, code int) (resp writer.Response, httpStatus int) {

	userProfileData, err := uc.profileRepository.GetUserByProfileCode(ctx, code)
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

	responseData := response.GetProfileResponse{
		WantedJobTitle: userProfileData.WantedJobTitle,
		FirstName:      userProfileData.FirstName,
		LastName:       userProfileData.LastName,
		Email:          userProfileData.Email,
		Phone:          userProfileData.Phone,
		Country:        userProfileData.Country,
		City:           userProfileData.City,
		Address:        userProfileData.Address,
		PostalCode:     userProfileData.PostalCode,
		DrivingLicense: userProfileData.DrivingLicense,
		Nationality:    userProfileData.Nationality,
		PlaceOfBirth:   userProfileData.PlaceOfBirth,
		DateOfBirth:    general.YMDDate(userProfileData.DateOfBirth),
	}

	resp = writer.APIResponse(constant.ResponseSuccess.Code, constant.ResponseSuccess.Description, responseData)
	httpStatus = constant.ResponseSuccess.Status
	return
}

func (uc *UseCase) CreateUserProfile(ctx context.Context, req request.CreateProfileRequest) (resp writer.Response, httpStatus int) {

	daterOfBirth, err := time.ParseInLocation(constant.DateFormatDDMMYYY, req.DateOfBirth, time.Now().Location())
	if err != nil {
		resp = writer.APIErrorResponse(constant.ResponseValidationError.Code, constant.ResponseValidationError.Description, err)
		httpStatus = constant.ResponseValidationError.Status
		return
	}

	userProfile := entity.Profile{
		WantedJobTitle: req.WantedJobTitle,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		Phone:          req.Phone,
		Country:        req.Country,
		City:           req.City,
		Address:        req.Address,
		PostalCode:     req.PostalCode,
		DrivingLicense: req.DrivingLicense,
		Nationality:    req.Nationality,
		PlaceOfBirth:   req.PlaceOfBirth,
		DateOfBirth:    daterOfBirth,
	}

	profileCode, err := uc.profileRepository.InsertProfile(ctx, userProfile)
	if err != nil {
		fmt.Println(err)
		resp = writer.APIErrorResponse(constant.ResponseInternalServerError.Code, constant.ResponseInternalServerError.Description, err)
		httpStatus = constant.ResponseInternalServerError.Status
		return
	}

	respData := response.CreateProfileResponse{
		ProfileCode: profileCode,
	}

	resp = writer.APIResponse(constant.ResponseSuccess.Code, constant.ResponseSuccess.Description, respData)
	httpStatus = constant.ResponseSuccess.Status
	return
}

func (uc *UseCase) UpdateUserProfile(ctx context.Context, req request.UpdateProfileRequest) (resp writer.Response, httpStatus int) {

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

	daterOfBirth, err := time.ParseInLocation(constant.DateFormatDDMMYYY, req.DateOfBirth, time.Now().Location())
	if err != nil {
		resp = writer.APIErrorResponse(constant.ResponseValidationError.Code, constant.ResponseValidationError.Description, err)
		httpStatus = constant.ResponseValidationError.Status
		return
	}

	updatedUserProfile := entity.Profile{
		ProfileCode:    userProfileData.ProfileCode,
		WantedJobTitle: req.WantedJobTitle,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		Phone:          req.Phone,
		Country:        req.Country,
		City:           req.City,
		Address:        req.Address,
		PostalCode:     req.PostalCode,
		DrivingLicense: req.DrivingLicense,
		Nationality:    req.Nationality,
		PlaceOfBirth:   req.PlaceOfBirth,
		DateOfBirth:    daterOfBirth,
	}

	profileCode, err := uc.profileRepository.UpdateProfile(ctx, updatedUserProfile)
	if err != nil {
		fmt.Println(err)
		resp = writer.APIErrorResponse(constant.ResponseInternalServerError.Code, constant.ResponseInternalServerError.Description, err)
		httpStatus = constant.ResponseInternalServerError.Status
		return
	}

	respData := response.CreateProfileResponse{
		ProfileCode: profileCode,
	}

	resp = writer.APIResponse(constant.ResponseSuccess.Code, constant.ResponseSuccess.Description, respData)
	httpStatus = constant.ResponseSuccess.Status
	return
}
