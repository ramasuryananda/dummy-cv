package working_experience

import (
	"context"
	"errors"
	"fmt"

	"github.com/ramasuryananda/dummy-cv/internal/constant"
	"github.com/ramasuryananda/dummy-cv/internal/dto/request"
	"github.com/ramasuryananda/dummy-cv/internal/dto/response"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/writer"
)

func (uc *UseCase) UpsertUserWorkingExperience(ctx context.Context, req request.UpsertWorkingExperienceRequest) (resp writer.Response, httpStatus int) {

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

	workingExperienceData, err := uc.workingExperienceRepository.GetWorkingExperienceByProfileCode(ctx, userProfileData.ProfileCode)
	if err != nil {
		if !errors.Is(err, constant.ErrorDatabaseNotFound) {
			resp = writer.APIErrorResponse(constant.ResponseInternalServerError.Code, constant.ResponseInternalServerError.Description, err)
			httpStatus = constant.ResponseInternalServerError.Status
			return
		}
	}

	workingExperienceData.ProfileCode = userProfileData.ProfileCode
	workingExperienceData.WorkingExperience = req.WorkingExperience

	err = uc.workingExperienceRepository.SaveWorkingExperience(ctx, workingExperienceData)
	if err != nil {
		resp = writer.APIErrorResponse(constant.ResponseInternalServerError.Code, constant.ResponseInternalServerError.Description, err)
		httpStatus = constant.ResponseInternalServerError.Status
		return
	}

	respData := response.UpsertWorkingExperienceResponse{
		WorkingExperience: fmt.Sprintf("update %s", req.WorkingExperience),
	}

	resp = writer.APIResponse(constant.ResponseSuccess.Code, constant.ResponseSuccess.Description, respData)
	httpStatus = constant.ResponseSuccess.Status
	return
}

func (uc *UseCase) GetUserWorkingExperience(ctx context.Context, req request.GetUserWorkingExperienceRequest) (resp writer.Response, httpStatus int) {

	workingExperienceData, err := uc.workingExperienceRepository.GetWorkingExperienceByProfileCode(ctx, req.ProfileCode)
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

	respData := response.GetUserWorkingExperienceResponse{
		WorkingExperience: workingExperienceData.WorkingExperience,
	}

	resp = writer.APIResponse(constant.ResponseSuccess.Code, constant.ResponseSuccess.Description, respData)
	httpStatus = constant.ResponseSuccess.Status
	return
}
