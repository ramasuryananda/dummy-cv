package employment

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mockemployment "github.com/ramasuryananda/dummy-cv/gomock/repository/mockEmployment"
	mockprofile "github.com/ramasuryananda/dummy-cv/gomock/repository/mockProfile"
	"github.com/ramasuryananda/dummy-cv/internal/constant"
	"github.com/ramasuryananda/dummy-cv/internal/dto/entity"
	"github.com/ramasuryananda/dummy-cv/internal/dto/general"
	"github.com/ramasuryananda/dummy-cv/internal/dto/request"
	"github.com/ramasuryananda/dummy-cv/internal/dto/response"
	"github.com/ramasuryananda/dummy-cv/internal/pkg/writer"
	"github.com/stretchr/testify/assert"
)

var userProfile = entity.Profile{
	ProfileCode:    1,
	WantedJobTitle: "test",
	FirstName:      "test",
	LastName:       "test",
	Email:          "test@gmail.com",
	Phone:          "12345677",
	Country:        "test country",
	City:           "test test city",
	Address:        "test address",
	PostalCode:     1234,
	DrivingLicense: "12345456",
	Nationality:    "test nationality",
	PlaceOfBirth:   "test place",
	DateOfBirth:    time.Date(2024, 12, 12, 0, 0, 0, 0, time.Now().Location()),
}

var employmentData = entity.Employment{
	ID:          1,
	ProfileCode: userProfile.ProfileCode,
	JobTitle:    "test job title",
	Employer:    "test employer",
	StartDate:   time.Date(2024, 12, 12, 0, 0, 0, 0, time.Now().Location()),
	EndDate:     sql.NullTime{},
	City:        "test city",
	Description: "test description",
}

func TestUseCase_GetUserEmployment(t *testing.T) {
	type fields struct {
		employmentRepository *mockemployment.MockRepositoryProvider
		profileRepository    *mockprofile.MockRepositoryProvider
	}
	type args struct {
		ctx context.Context
		req request.GetEmploymentRequest
	}
	tests := []struct {
		name           string
		args           args
		mock           func(mockfield fields)
		wantResp       writer.Response
		wantHttpStatus int
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: request.GetEmploymentRequest{
					ProfileCode: userProfile.ProfileCode,
				},
			},
			mock: func(mocks fields) {
				mocks.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(userProfile, nil)

				mocks.employmentRepository.EXPECT().GetEmploymentByProfileCode(context.Background(), userProfile.ProfileCode).Return([]entity.Employment{employmentData}, nil)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseSuccess.Code,
				Message: constant.ResponseSuccess.Description,
				Data: []response.EmploymentDataResponse{
					{
						ID:          employmentData.ID,
						JobTitle:    employmentData.JobTitle,
						Employer:    employmentData.Employer,
						StartDate:   general.YMDDate(employmentData.StartDate),
						EndDate:     general.YMDDate(employmentData.EndDate.Time),
						City:        employmentData.City,
						Description: employmentData.Description,
					},
				},
			},
			wantHttpStatus: constant.ResponseSuccess.Status,
		},
		{
			name: "failed get employment",
			args: args{
				ctx: context.Background(),
				req: request.GetEmploymentRequest{
					ProfileCode: userProfile.ProfileCode,
				},
			},
			mock: func(mocks fields) {
				mocks.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(userProfile, nil)

				mocks.employmentRepository.EXPECT().GetEmploymentByProfileCode(context.Background(), userProfile.ProfileCode).Return([]entity.Employment{}, assert.AnError)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseInternalServerError.Code,
				Message: constant.ResponseInternalServerError.Description,
			},
			wantHttpStatus: constant.ResponseInternalServerError.Status,
		},
		{
			name: "failed get user profile ",
			args: args{
				ctx: context.Background(),
				req: request.GetEmploymentRequest{
					ProfileCode: userProfile.ProfileCode,
				},
			},
			mock: func(mocks fields) {
				mocks.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(entity.Profile{}, assert.AnError)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseInternalServerError.Code,
				Message: constant.ResponseInternalServerError.Description,
			},
			wantHttpStatus: constant.ResponseInternalServerError.Status,
		},
		{
			name: "failed get user profile not found",
			args: args{
				ctx: context.Background(),
				req: request.GetEmploymentRequest{
					ProfileCode: userProfile.ProfileCode,
				},
			},
			mock: func(mocks fields) {
				mocks.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(entity.Profile{}, constant.ErrorDatabaseNotFound)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseErrorNotFound.Code,
				Message: constant.ResponseErrorNotFound.Description,
			},
			wantHttpStatus: constant.ResponseErrorNotFound.Status,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := fields{
				profileRepository:    mockprofile.NewMockRepositoryProvider(ctrl),
				employmentRepository: mockemployment.NewMockRepositoryProvider(ctrl),
			}

			tt.mock(fields)

			uc := &UseCase{
				profileRepository:    fields.profileRepository,
				employmentRepository: fields.employmentRepository,
			}

			gotResp, gotHttpStatus := uc.GetUserEmployment(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UseCase.GetUserEmployment() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if gotHttpStatus != tt.wantHttpStatus {
				t.Errorf("UseCase.GetUserEmployment() gotHttpStatus = %v, want %v", gotHttpStatus, tt.wantHttpStatus)
			}
		})
	}
}
