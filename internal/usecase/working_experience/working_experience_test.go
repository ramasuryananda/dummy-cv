package working_experience

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mockprofile "github.com/ramasuryananda/dummy-cv/gomock/repository/mockProfile"
	mockworkingexperience "github.com/ramasuryananda/dummy-cv/gomock/repository/mockWorkingExperience"
	"github.com/ramasuryananda/dummy-cv/internal/constant"
	"github.com/ramasuryananda/dummy-cv/internal/dto/entity"
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

var workingExperience = entity.WorkingExperience{
	ID:                1,
	ProfileCode:       userProfile.ProfileCode,
	WorkingExperience: "test working experience",
}

func TestUseCase_UpsertUserWorkingExperience(t *testing.T) {
	type fields struct {
		workingExperienceRepository *mockworkingexperience.MockRepositoryProvider
		profileRepository           *mockprofile.MockRepositoryProvider
	}
	type args struct {
		ctx context.Context
		req request.UpsertWorkingExperienceRequest
	}
	tests := []struct {
		name           string
		mock           func(mocks fields)
		args           args
		wantResp       writer.Response
		wantHttpStatus int
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: request.UpsertWorkingExperienceRequest{
					ProfileCode:       userProfile.ProfileCode,
					WorkingExperience: "test update working experience",
				},
			},
			mock: func(mocks fields) {
				mocks.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(userProfile, nil)

				mocks.workingExperienceRepository.EXPECT().GetWorkingExperienceByProfileCode(context.Background(), userProfile.ProfileCode).Return(workingExperience, nil)

				mocks.workingExperienceRepository.EXPECT().SaveWorkingExperience(context.Background(), entity.WorkingExperience{
					ID:                workingExperience.ID,
					ProfileCode:       userProfile.ProfileCode,
					WorkingExperience: "test update working experience",
					CreatedAt:         workingExperience.CreatedAt,
					UpdatedAt:         workingExperience.UpdatedAt,
				}).Return(nil)

			},
			wantResp: writer.Response{
				Code:    constant.ResponseSuccess.Code,
				Message: constant.ResponseSuccess.Description,
				Data: response.UpsertWorkingExperienceResponse{
					WorkingExperience: "update test update working experience",
				},
			},
			wantHttpStatus: constant.ResponseSuccess.Status,
		},
		{
			name: "failed saving data",
			args: args{
				ctx: context.Background(),
				req: request.UpsertWorkingExperienceRequest{
					ProfileCode:       userProfile.ProfileCode,
					WorkingExperience: "test update working experience",
				},
			},
			mock: func(mocks fields) {
				mocks.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(userProfile, nil)

				mocks.workingExperienceRepository.EXPECT().GetWorkingExperienceByProfileCode(context.Background(), userProfile.ProfileCode).Return(workingExperience, nil)

				mocks.workingExperienceRepository.EXPECT().SaveWorkingExperience(context.Background(), entity.WorkingExperience{
					ID:                workingExperience.ID,
					ProfileCode:       userProfile.ProfileCode,
					WorkingExperience: "test update working experience",
					CreatedAt:         workingExperience.CreatedAt,
					UpdatedAt:         workingExperience.UpdatedAt,
				}).Return(assert.AnError)

			},
			wantResp: writer.Response{
				Code:    constant.ResponseInternalServerError.Code,
				Message: constant.ResponseInternalServerError.Description,
			},
			wantHttpStatus: constant.ResponseInternalServerError.Status,
		},
		{
			name: "failed getting data return other than not found",
			args: args{
				ctx: context.Background(),
				req: request.UpsertWorkingExperienceRequest{
					ProfileCode:       userProfile.ProfileCode,
					WorkingExperience: "test update working experience",
				},
			},
			mock: func(mocks fields) {
				mocks.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(userProfile, nil)

				mocks.workingExperienceRepository.EXPECT().GetWorkingExperienceByProfileCode(context.Background(), userProfile.ProfileCode).Return(workingExperience, assert.AnError)

			},
			wantResp: writer.Response{
				Code:    constant.ResponseInternalServerError.Code,
				Message: constant.ResponseInternalServerError.Description,
			},
			wantHttpStatus: constant.ResponseInternalServerError.Status,
		},
		{
			name: "failed getting data profile return other than not found",
			args: args{
				ctx: context.Background(),
				req: request.UpsertWorkingExperienceRequest{
					ProfileCode:       userProfile.ProfileCode,
					WorkingExperience: "test update working experience",
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
			name: "failed getting data profile return  not found",
			args: args{
				ctx: context.Background(),
				req: request.UpsertWorkingExperienceRequest{
					ProfileCode:       userProfile.ProfileCode,
					WorkingExperience: "test update working experience",
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
				profileRepository:           mockprofile.NewMockRepositoryProvider(ctrl),
				workingExperienceRepository: mockworkingexperience.NewMockRepositoryProvider(ctrl),
			}

			tt.mock(fields)

			uc := &UseCase{
				profileRepository:           fields.profileRepository,
				workingExperienceRepository: fields.workingExperienceRepository,
			}
			gotResp, gotHttpStatus := uc.UpsertUserWorkingExperience(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UseCase.UpsertUserWorkingExperience() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if gotHttpStatus != tt.wantHttpStatus {
				t.Errorf("UseCase.UpsertUserWorkingExperience() gotHttpStatus = %v, want %v", gotHttpStatus, tt.wantHttpStatus)
			}
		})
	}
}

func TestUseCase_GetUserWorkingExperience(t *testing.T) {
	type fields struct {
		workingExperienceRepository *mockworkingexperience.MockRepositoryProvider
		profileRepository           *mockprofile.MockRepositoryProvider
	}
	type args struct {
		ctx context.Context
		req request.GetUserWorkingExperienceRequest
	}
	tests := []struct {
		name           string
		args           args
		mock           func(mocks fields)
		wantResp       writer.Response
		wantHttpStatus int
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: request.GetUserWorkingExperienceRequest{
					ProfileCode: userProfile.ProfileCode,
				},
			},
			mock: func(mocks fields) {
				mocks.workingExperienceRepository.EXPECT().GetWorkingExperienceByProfileCode(context.Background(), userProfile.ProfileCode).Return(workingExperience, nil)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseSuccess.Code,
				Message: constant.ResponseSuccess.Description,
				Data: response.GetUserWorkingExperienceResponse{
					WorkingExperience: workingExperience.WorkingExperience,
				},
			},
			wantHttpStatus: constant.ResponseSuccess.Status,
		},
		{
			name: "failed not found",
			args: args{
				ctx: context.Background(),
				req: request.GetUserWorkingExperienceRequest{
					ProfileCode: userProfile.ProfileCode,
				},
			},
			mock: func(mocks fields) {
				mocks.workingExperienceRepository.EXPECT().GetWorkingExperienceByProfileCode(context.Background(), userProfile.ProfileCode).Return(entity.WorkingExperience{}, constant.ErrorDatabaseNotFound)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseErrorNotFound.Code,
				Message: constant.ResponseErrorNotFound.Description,
			},
			wantHttpStatus: constant.ResponseErrorNotFound.Status,
		},
		{
			name: "failed get data",
			args: args{
				ctx: context.Background(),
				req: request.GetUserWorkingExperienceRequest{
					ProfileCode: userProfile.ProfileCode,
				},
			},
			mock: func(mocks fields) {
				mocks.workingExperienceRepository.EXPECT().GetWorkingExperienceByProfileCode(context.Background(), userProfile.ProfileCode).Return(entity.WorkingExperience{}, assert.AnError)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseInternalServerError.Code,
				Message: constant.ResponseInternalServerError.Description,
			},
			wantHttpStatus: constant.ResponseInternalServerError.Status,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := fields{
				profileRepository:           mockprofile.NewMockRepositoryProvider(ctrl),
				workingExperienceRepository: mockworkingexperience.NewMockRepositoryProvider(ctrl),
			}

			tt.mock(fields)

			uc := &UseCase{
				profileRepository:           fields.profileRepository,
				workingExperienceRepository: fields.workingExperienceRepository,
			}
			gotResp, gotHttpStatus := uc.GetUserWorkingExperience(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UseCase.GetUserWorkingExperience() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if gotHttpStatus != tt.wantHttpStatus {
				t.Errorf("UseCase.GetUserWorkingExperience() gotHttpStatus = %v, want %v", gotHttpStatus, tt.wantHttpStatus)
			}
		})
	}
}
