package employment

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ramasuryananda/dummy-cv/internal/constant"
	"github.com/ramasuryananda/dummy-cv/internal/dto/entity"
	"github.com/ramasuryananda/dummy-cv/internal/dto/general"
	"github.com/ramasuryananda/dummy-cv/internal/dto/request"
	"github.com/ramasuryananda/dummy-cv/internal/dto/response"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/writer"
)

func (uc *UseCase) GetUserEmployment(ctx context.Context, req request.GetEmploymentRequest) (resp writer.Response, httpStatus int) {

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

	employmentData, err := uc.employmentRepository.GetEmploymentByProfileCode(ctx, userProfileData.ProfileCode)
	if err != nil {
		resp = writer.APIErrorResponse(constant.ResponseInternalServerError.Code, constant.ResponseInternalServerError.Description, err)
		httpStatus = constant.ResponseInternalServerError.Status
		return
	}

	respData := make([]response.EmploymentDataResponse, 0, len(employmentData))
	for _, data := range employmentData {
		respData = append(respData, response.EmploymentDataResponse{
			ID:          data.ID,
			JobTitle:    data.JobTitle,
			Employer:    data.Employer,
			StartDate:   general.YMDDate(data.StartDate),
			EndDate:     general.YMDDate(data.EndDate.Time),
			City:        data.City,
			Description: data.Description,
		})
	}

	resp = writer.APIResponse(constant.ResponseSuccess.Code, constant.ResponseSuccess.Description, respData)
	httpStatus = constant.ResponseSuccess.Status
	return
}

func (uc *UseCase) CreateEmploymentData(ctx context.Context, req request.CreateEmploymentRequest) (resp writer.Response, httpStatus int) {

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

	startDate, err := time.ParseInLocation(constant.DateFormatDDMMYYY, req.StartDate, time.Now().Location())
	if err != nil {
		resp = writer.APIErrorResponse(constant.ResponseBadRequest.Code, constant.ResponseBadRequest.Description, err)
		httpStatus = constant.ResponseBadRequest.Status
		return
	}

	var endDate sql.NullTime

	if req.EndDate != "" {
		date, err := time.ParseInLocation(constant.DateFormatDDMMYYY, req.EndDate, time.Now().Location())
		if err != nil {
			resp = writer.APIErrorResponse(constant.ResponseBadRequest.Code, constant.ResponseBadRequest.Description, err)
			httpStatus = constant.ResponseBadRequest.Status
			return
		}

		endDate = sql.NullTime{
			Valid: true,
			Time:  date,
		}
	}

	employmentData := entity.Employment{
		ProfileCode: userProfileData.ProfileCode,
		JobTitle:    req.JobTitle,
		Employer:    req.Employer,
		StartDate:   startDate,
		EndDate:     endDate,
		City:        req.City,
		Description: req.Description,
	}

	lastID, err := uc.employmentRepository.CreateEmploymentData(ctx, employmentData)
	if err != nil {
		resp = writer.APIErrorResponse(constant.ResponseInternalServerError.Code, constant.ResponseInternalServerError.Description, err)
		httpStatus = constant.ResponseInternalServerError.Status
		return
	}

	respData := response.CreateEmploymentResponse{
		Id:          lastID,
		ProfileCode: userProfileData.ProfileCode,
	}

	resp = writer.APIResponse(constant.ResponseSuccess.Code, constant.ResponseSuccess.Description, respData)
	httpStatus = constant.ResponseSuccess.Status
	return
}

func (uc *UseCase) DeleteEmploymentData(ctx context.Context, req request.DeleteEmploymentRequest) (resp writer.Response, httpStatus int) {

	employmentData, err := uc.employmentRepository.GetFirstEmploymentByProfileCodeandID(ctx, req.ProfileCode, req.ID)
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

	err = uc.employmentRepository.DeleteEmploymentData(ctx, employmentData.ProfileCode, employmentData.ID)
	if err != nil {
		resp = writer.APIErrorResponse(constant.ResponseInternalServerError.Code, constant.ResponseInternalServerError.Description, err)
		httpStatus = constant.ResponseInternalServerError.Status
		return
	}

	respData := response.DeleteEmploymentResponse{
		ProfileCode: req.ProfileCode,
	}

	resp = writer.APIResponse(constant.ResponseSuccess.Code, constant.ResponseSuccess.Description, respData)
	httpStatus = constant.ResponseSuccess.Status
	return
}
