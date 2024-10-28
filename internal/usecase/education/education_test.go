package education

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mockeducation "github.com/ramasuryananda/dummy-cv/gomock/repository/mockEducation"
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

var educationData = entity.Education{
	ID:          1,
	ProfileCode: userProfile.ProfileCode,
	School:      "test school",
	Degree:      "test degree",
	StartDate:   time.Date(2024, 12, 12, 0, 0, 0, 0, time.Now().Location()),
	EndDate:     sql.NullTime{},
	City:        "test city",
	Description: "test description",
}

func TestUseCase_GetUserEducation(t *testing.T) {
	type fields struct {
		educationRepository *mockeducation.MockRepositoryProvider
		profileRepository   *mockprofile.MockRepositoryProvider
	}
	type args struct {
		ctx context.Context
		req request.GetEducationRequest
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
				req: request.GetEducationRequest{
					ProfileCode: userProfile.ProfileCode,
				},
			},
			mock: func(mocks fields) {
				mocks.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(userProfile, nil)

				mocks.educationRepository.EXPECT().GetEducationByProfileCode(context.Background(), userProfile.ProfileCode).Return([]entity.Education{educationData}, nil)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseSuccess.Code,
				Message: constant.ResponseSuccess.Description,
				Data: []response.EducationDataResponse{
					{
						ID:          educationData.ID,
						School:      educationData.School,
						Degree:      educationData.Degree,
						StartDate:   general.YMDDate(educationData.StartDate),
						EndDate:     general.YMDDate(educationData.EndDate.Time),
						City:        educationData.City,
						Description: educationData.Description,
					},
				},
			},
			wantHttpStatus: constant.ResponseSuccess.Status,
		},
		{
			name: "failed get data education",
			args: args{
				ctx: context.Background(),
				req: request.GetEducationRequest{
					ProfileCode: userProfile.ProfileCode,
				},
			},
			mock: func(mocks fields) {
				mocks.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(userProfile, nil)

				mocks.educationRepository.EXPECT().GetEducationByProfileCode(context.Background(), userProfile.ProfileCode).Return([]entity.Education{}, assert.AnError)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseInternalServerError.Code,
				Message: constant.ResponseInternalServerError.Description,
			},
			wantHttpStatus: constant.ResponseInternalServerError.Status,
		},
		{
			name: "failed get data profile",
			args: args{
				ctx: context.Background(),
				req: request.GetEducationRequest{
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
			name: "failed get data profile data not found",
			args: args{
				ctx: context.Background(),
				req: request.GetEducationRequest{
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
				profileRepository:   mockprofile.NewMockRepositoryProvider(ctrl),
				educationRepository: mockeducation.NewMockRepositoryProvider(ctrl),
			}

			tt.mock(fields)

			uc := &UseCase{
				profileRepository:   fields.profileRepository,
				educationRepository: fields.educationRepository,
			}
			gotResp, gotHttpStatus := uc.GetUserEducation(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UseCase.GetUserEducation() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if gotHttpStatus != tt.wantHttpStatus {
				t.Errorf("UseCase.GetUserEducation() gotHttpStatus = %v, want %v", gotHttpStatus, tt.wantHttpStatus)
			}
		})
	}
}

func TestUseCase_CreateEducationData(t *testing.T) {
	type fields struct {
		educationRepository *mockeducation.MockRepositoryProvider
		profileRepository   *mockprofile.MockRepositoryProvider
	}

	type args struct {
		ctx context.Context
		req request.CreateEducationRequest
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
				req: request.CreateEducationRequest{
					ProfileCode: userProfile.ProfileCode,
					School:      "test school",
					Degree:      "test degree",
					StartDate:   "12-12-2024",
					EndDate:     "13-12-2024",
					City:        "test city",
					Description: "test description",
				},
			},
			mock: func(mocks fields) {
				mocks.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(userProfile, nil)

				mocks.educationRepository.EXPECT().CreateEducationData(context.Background(), entity.Education{
					ProfileCode: userProfile.ProfileCode,
					School:      "test school",
					Degree:      "test degree",
					StartDate:   time.Date(2024, 12, 12, 0, 0, 0, 0, time.Now().Location()),
					EndDate: sql.NullTime{
						Valid: true,
						Time:  time.Date(2024, 12, 13, 0, 0, 0, 0, time.Now().Location()),
					},
					City:        "test city",
					Description: "test description",
				}).Return(uint64(1), nil)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseSuccess.Code,
				Message: constant.ResponseSuccess.Description,
				Data: response.CreateEducationResponse{
					Id:          uint64(1),
					ProfileCode: educationData.ProfileCode,
				},
			},
			wantHttpStatus: constant.ResponseSuccess.Status,
		},
		{
			name: "failed create data",
			args: args{
				ctx: context.Background(),
				req: request.CreateEducationRequest{
					ProfileCode: userProfile.ProfileCode,
					School:      "test school",
					Degree:      "test degree",
					StartDate:   "12-12-2024",
					EndDate:     "13-12-2024",
					City:        "test city",
					Description: "test description",
				},
			},
			mock: func(mocks fields) {
				mocks.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(userProfile, nil)

				mocks.educationRepository.EXPECT().CreateEducationData(context.Background(), entity.Education{
					ProfileCode: userProfile.ProfileCode,
					School:      "test school",
					Degree:      "test degree",
					StartDate:   time.Date(2024, 12, 12, 0, 0, 0, 0, time.Now().Location()),
					EndDate: sql.NullTime{
						Valid: true,
						Time:  time.Date(2024, 12, 13, 0, 0, 0, 0, time.Now().Location()),
					},
					City:        "test city",
					Description: "test description",
				}).Return(uint64(1), assert.AnError)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseInternalServerError.Code,
				Message: constant.ResponseInternalServerError.Description,
			},
			wantHttpStatus: constant.ResponseInternalServerError.Status,
		},
		{
			name: "failed parse end date",
			args: args{
				ctx: context.Background(),
				req: request.CreateEducationRequest{
					ProfileCode: userProfile.ProfileCode,
					School:      "test school",
					Degree:      "test degree",
					StartDate:   "12-12-2024",
					EndDate:     "13-12-20245",
					City:        "test city",
					Description: "test description",
				},
			},
			mock: func(mocks fields) {
				mocks.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(userProfile, nil)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseBadRequest.Code,
				Message: constant.ResponseBadRequest.Description,
			},
			wantHttpStatus: constant.ResponseBadRequest.Status,
		},
		{
			name: "failed parse start date",
			args: args{
				ctx: context.Background(),
				req: request.CreateEducationRequest{
					ProfileCode: userProfile.ProfileCode,
					School:      "test school",
					Degree:      "test degree",
					StartDate:   "12-12-20245",
					EndDate:     "13-12-2024",
					City:        "test city",
					Description: "test description",
				},
			},
			mock: func(mocks fields) {
				mocks.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(userProfile, nil)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseBadRequest.Code,
				Message: constant.ResponseBadRequest.Description,
			},
			wantHttpStatus: constant.ResponseBadRequest.Status,
		},
		{
			name: "failed get data",
			args: args{
				ctx: context.Background(),
				req: request.CreateEducationRequest{
					ProfileCode: userProfile.ProfileCode,
					School:      "test school",
					Degree:      "test degree",
					StartDate:   "12-12-2024",
					EndDate:     "13-12-2024",
					City:        "test city",
					Description: "test description",
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
			name: "failed get data not found",
			args: args{
				ctx: context.Background(),
				req: request.CreateEducationRequest{
					ProfileCode: userProfile.ProfileCode,
					School:      "test school",
					Degree:      "test degree",
					StartDate:   "12-12-2024",
					EndDate:     "13-12-2024",
					City:        "test city",
					Description: "test description",
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
				profileRepository:   mockprofile.NewMockRepositoryProvider(ctrl),
				educationRepository: mockeducation.NewMockRepositoryProvider(ctrl),
			}

			tt.mock(fields)

			uc := &UseCase{
				profileRepository:   fields.profileRepository,
				educationRepository: fields.educationRepository,
			}
			gotResp, gotHttpStatus := uc.CreateEducationData(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UseCase.CreateEducationData() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if gotHttpStatus != tt.wantHttpStatus {
				t.Errorf("UseCase.CreateEducationData() gotHttpStatus = %v, want %v", gotHttpStatus, tt.wantHttpStatus)
			}
		})
	}
}

func TestUseCase_DeleteEducationData(t *testing.T) {
	type fields struct {
		educationRepository *mockeducation.MockRepositoryProvider
		profileRepository   *mockprofile.MockRepositoryProvider
	}
	type args struct {
		ctx context.Context
		req request.DeleteEducationRequest
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
				req: request.DeleteEducationRequest{
					ProfileCode: userProfile.ProfileCode,
					ID:          1,
				},
			},
			mock: func(mocks fields) {
				mocks.educationRepository.EXPECT().GetFirstEducationByProfileCodeandID(context.Background(), userProfile.ProfileCode, uint64(1)).Return(educationData, nil)

				mocks.educationRepository.EXPECT().DeleteEducationData(context.Background(), educationData.ProfileCode, educationData.ID).Return(nil)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseSuccess.Code,
				Message: constant.ResponseSuccess.Description,
				Data: response.DeleteEducationResponse{
					ProfileCode: userProfile.ProfileCode,
				},
			},
			wantHttpStatus: constant.ResponseSuccess.Status,
		},
		{
			name: "failed delete",
			args: args{
				ctx: context.Background(),
				req: request.DeleteEducationRequest{
					ProfileCode: userProfile.ProfileCode,
					ID:          1,
				},
			},
			mock: func(mocks fields) {
				mocks.educationRepository.EXPECT().GetFirstEducationByProfileCodeandID(context.Background(), userProfile.ProfileCode, uint64(1)).Return(educationData, nil)

				mocks.educationRepository.EXPECT().DeleteEducationData(context.Background(), educationData.ProfileCode, educationData.ID).Return(assert.AnError)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseInternalServerError.Code,
				Message: constant.ResponseInternalServerError.Description,
			},
			wantHttpStatus: constant.ResponseInternalServerError.Status,
		},
		{
			name: "failed delete not found",
			args: args{
				ctx: context.Background(),
				req: request.DeleteEducationRequest{
					ProfileCode: userProfile.ProfileCode,
					ID:          1,
				},
			},
			mock: func(mocks fields) {
				mocks.educationRepository.EXPECT().GetFirstEducationByProfileCodeandID(context.Background(), userProfile.ProfileCode, uint64(1)).Return(entity.Education{}, constant.ErrorDatabaseNotFound)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseErrorNotFound.Code,
				Message: constant.ResponseErrorNotFound.Description,
			},
			wantHttpStatus: constant.ResponseErrorNotFound.Status,
		},
		{
			name: "failed delete get data",
			args: args{
				ctx: context.Background(),
				req: request.DeleteEducationRequest{
					ProfileCode: userProfile.ProfileCode,
					ID:          1,
				},
			},
			mock: func(mocks fields) {
				mocks.educationRepository.EXPECT().GetFirstEducationByProfileCodeandID(context.Background(), userProfile.ProfileCode, uint64(1)).Return(entity.Education{}, assert.AnError)
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
				profileRepository:   mockprofile.NewMockRepositoryProvider(ctrl),
				educationRepository: mockeducation.NewMockRepositoryProvider(ctrl),
			}

			tt.mock(fields)

			uc := &UseCase{
				profileRepository:   fields.profileRepository,
				educationRepository: fields.educationRepository,
			}
			gotResp, gotHttpStatus := uc.DeleteEducationData(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UseCase.DeleteEducationData() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if gotHttpStatus != tt.wantHttpStatus {
				t.Errorf("UseCase.DeleteEducationData() gotHttpStatus = %v, want %v", gotHttpStatus, tt.wantHttpStatus)
			}
		})
	}
}
