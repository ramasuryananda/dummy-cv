package skill

import (
	"context"
	"errors"

	"github.com/ramasuryananda/dummy-cv/internal/constant"
	"github.com/ramasuryananda/dummy-cv/internal/dto/entity"
	"github.com/ramasuryananda/dummy-cv/internal/dto/request"
	"github.com/ramasuryananda/dummy-cv/internal/dto/response"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/writer"
)

func (uc *UseCase) GetUserSkill(ctx context.Context, req request.GetSkillRequest) (resp writer.Response, httpStatus int) {

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

	skillData, err := uc.skillRepository.GetSkillByProfileCode(ctx, userProfileData.ProfileCode)
	if err != nil {
		resp = writer.APIErrorResponse(constant.ResponseErrorNotFound.Code, constant.ResponseErrorNotFound.Description, err)
		httpStatus = constant.ResponseErrorNotFound.Status
		return
	}

	respData := make([]response.SkillDataResponse, 0, len(skillData))
	for _, data := range skillData {
		respData = append(respData, response.SkillDataResponse{
			ID:    data.ID,
			Skill: data.Skill,
			Level: data.Level,
		})
	}

	resp = writer.APIResponse(constant.ResponseSuccess.Code, constant.ResponseSuccess.Description, respData)
	httpStatus = constant.ResponseSuccess.Status
	return
}

func (uc *UseCase) CreateSkillData(ctx context.Context, req request.CreateSkillRequest) (resp writer.Response, httpStatus int) {

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

	skill := entity.Skill{
		ProfileCode: userProfileData.ProfileCode,
		Skill:       req.Skill,
		Level:       req.Level,
	}

	lastID, err := uc.skillRepository.CreateSkillData(ctx, skill)
	if err != nil {
		resp = writer.APIErrorResponse(constant.ResponseInternalServerError.Code, constant.ResponseInternalServerError.Description, err)
		httpStatus = constant.ResponseInternalServerError.Status
		return
	}

	respData := response.CreateSkillResponse{
		Id:          lastID,
		ProfileCode: userProfileData.ProfileCode,
	}

	resp = writer.APIResponse(constant.ResponseSuccess.Code, constant.ResponseSuccess.Description, respData)
	httpStatus = constant.ResponseSuccess.Status
	return
}

func (uc *UseCase) DeleteSkillData(ctx context.Context, req request.DeleteSkillRequest) (resp writer.Response, httpStatus int) {

	skillData, err := uc.skillRepository.GetFirstSkillByProfileCodeandID(ctx, req.ProfileCode, req.ID)
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

	err = uc.skillRepository.DeleteSkillData(ctx, skillData.ProfileCode, skillData.ID)
	if err != nil {
		resp = writer.APIErrorResponse(constant.ResponseInternalServerError.Code, constant.ResponseInternalServerError.Description, err)
		httpStatus = constant.ResponseInternalServerError.Status
		return
	}

	respData := response.DeleteSkillResponse{
		ProfileCode: req.ProfileCode,
	}

	resp = writer.APIResponse(constant.ResponseSuccess.Code, constant.ResponseSuccess.Description, respData)
	httpStatus = constant.ResponseSuccess.Status
	return
}
