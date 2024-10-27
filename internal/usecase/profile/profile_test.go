package profile

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
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

func TestUseCase_GetUserProfile(t *testing.T) {
	type fields struct {
		profileRepository *mockprofile.MockRepositoryProvider
	}

	type args struct {
		ctx  context.Context
		code uint64
	}
	tests := []struct {
		name           string
		args           args
		wantResp       writer.Response
		wantHttpStatus int
		mock           func(mock fields)
	}{
		{
			name: "success",
			args: args{
				ctx:  context.Background(),
				code: userProfile.ProfileCode,
			},
			mock: func(mock fields) {
				mock.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(userProfile, nil)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseSuccess.Code,
				Message: constant.ResponseSuccess.Description,
				Data: response.GetProfileResponse{
					WantedJobTitle: userProfile.WantedJobTitle,
					FirstName:      userProfile.FirstName,
					LastName:       userProfile.LastName,
					Email:          userProfile.Email,
					Phone:          userProfile.Phone,
					Country:        userProfile.Country,
					City:           userProfile.City,
					Address:        userProfile.Address,
					PostalCode:     userProfile.PostalCode,
					DrivingLicense: userProfile.DrivingLicense,
					Nationality:    userProfile.Nationality,
					PlaceOfBirth:   userProfile.PlaceOfBirth,
					DateOfBirth:    general.YMDDate(userProfile.DateOfBirth),
				},
			},
			wantHttpStatus: constant.ResponseSuccess.Status,
		},
		{
			name: "failed data not found",
			args: args{
				ctx:  context.Background(),
				code: userProfile.ProfileCode,
			},
			mock: func(mock fields) {
				mock.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(entity.Profile{}, constant.ErrorDatabaseNotFound)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseErrorNotFound.Code,
				Message: constant.ResponseErrorNotFound.Description,
			},
			wantHttpStatus: constant.ResponseErrorNotFound.Status,
		},
		{
			name: "failed database",
			args: args{
				ctx:  context.Background(),
				code: userProfile.ProfileCode,
			},
			mock: func(mock fields) {
				mock.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(entity.Profile{}, assert.AnError)
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
				profileRepository: mockprofile.NewMockRepositoryProvider(ctrl),
			}

			tt.mock(fields)

			uc := &UseCase{
				profileRepository: fields.profileRepository,
			}
			gotResp, gotHttpStatus := uc.GetUserProfile(tt.args.ctx, tt.args.code)
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UseCase.GetUserProfile() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if gotHttpStatus != tt.wantHttpStatus {
				t.Errorf("UseCase.GetUserProfile() gotHttpStatus = %v, want %v", gotHttpStatus, tt.wantHttpStatus)
			}
		})
	}
}

func TestUseCase_CreateUserProfile(t *testing.T) {
	type fields struct {
		profileRepository *mockprofile.MockRepositoryProvider
	}
	type args struct {
		ctx context.Context
		req request.CreateProfileRequest
	}
	tests := []struct {
		name           string
		args           args
		wantResp       writer.Response
		wantHttpStatus int
		mock           func(mock fields)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: request.CreateProfileRequest{
					WantedJobTitle: userProfile.WantedJobTitle,
					FirstName:      userProfile.FirstName,
					LastName:       userProfile.LastName,
					Email:          userProfile.Email,
					Phone:          userProfile.Phone,
					Country:        userProfile.Country,
					City:           userProfile.City,
					Address:        userProfile.Address,
					PostalCode:     userProfile.PostalCode,
					DrivingLicense: userProfile.DrivingLicense,
					Nationality:    userProfile.Nationality,
					PlaceOfBirth:   userProfile.PlaceOfBirth,
					DateOfBirth:    "12-12-2024",
				},
			},
			mock: func(mock fields) {
				mock.profileRepository.EXPECT().InsertProfile(context.Background(), entity.Profile{
					WantedJobTitle: userProfile.WantedJobTitle,
					FirstName:      userProfile.FirstName,
					LastName:       userProfile.LastName,
					Email:          userProfile.Email,
					Phone:          userProfile.Phone,
					Country:        userProfile.Country,
					City:           userProfile.City,
					Address:        userProfile.Address,
					PostalCode:     userProfile.PostalCode,
					DrivingLicense: userProfile.DrivingLicense,
					Nationality:    userProfile.Nationality,
					PlaceOfBirth:   userProfile.PlaceOfBirth,
					DateOfBirth:    userProfile.DateOfBirth,
				}).Return(userProfile.ProfileCode, nil)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseSuccess.Code,
				Message: constant.ResponseSuccess.Description,
				Data: response.CreateProfileResponse{
					ProfileCode: userProfile.ProfileCode,
				},
			},
			wantHttpStatus: constant.ResponseSuccess.Status,
		},
		{
			name: "failed on insert db",
			args: args{
				ctx: context.Background(),
				req: request.CreateProfileRequest{
					WantedJobTitle: userProfile.WantedJobTitle,
					FirstName:      userProfile.FirstName,
					LastName:       userProfile.LastName,
					Email:          userProfile.Email,
					Phone:          userProfile.Phone,
					Country:        userProfile.Country,
					City:           userProfile.City,
					Address:        userProfile.Address,
					PostalCode:     userProfile.PostalCode,
					DrivingLicense: userProfile.DrivingLicense,
					Nationality:    userProfile.Nationality,
					PlaceOfBirth:   userProfile.PlaceOfBirth,
					DateOfBirth:    "12-12-2024",
				},
			},
			mock: func(mock fields) {
				mock.profileRepository.EXPECT().InsertProfile(context.Background(), entity.Profile{
					WantedJobTitle: userProfile.WantedJobTitle,
					FirstName:      userProfile.FirstName,
					LastName:       userProfile.LastName,
					Email:          userProfile.Email,
					Phone:          userProfile.Phone,
					Country:        userProfile.Country,
					City:           userProfile.City,
					Address:        userProfile.Address,
					PostalCode:     userProfile.PostalCode,
					DrivingLicense: userProfile.DrivingLicense,
					Nationality:    userProfile.Nationality,
					PlaceOfBirth:   userProfile.PlaceOfBirth,
					DateOfBirth:    userProfile.DateOfBirth,
				}).Return(uint64(0), assert.AnError)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseInternalServerError.Code,
				Message: constant.ResponseInternalServerError.Description,
			},
			wantHttpStatus: constant.ResponseInternalServerError.Status,
		},
		{
			name: "failed on invalid birth date",
			args: args{
				ctx: context.Background(),
				req: request.CreateProfileRequest{
					WantedJobTitle: userProfile.WantedJobTitle,
					FirstName:      userProfile.FirstName,
					LastName:       userProfile.LastName,
					Email:          userProfile.Email,
					Phone:          userProfile.Phone,
					Country:        userProfile.Country,
					City:           userProfile.City,
					Address:        userProfile.Address,
					PostalCode:     userProfile.PostalCode,
					DrivingLicense: userProfile.DrivingLicense,
					Nationality:    userProfile.Nationality,
					PlaceOfBirth:   userProfile.PlaceOfBirth,
					DateOfBirth:    "12-12-20242",
				},
			},
			mock: func(mock fields) {
			},
			wantResp: writer.Response{
				Code:    constant.ResponseValidationError.Code,
				Message: constant.ResponseValidationError.Description,
			},
			wantHttpStatus: constant.ResponseValidationError.Status,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			fields := fields{
				profileRepository: mockprofile.NewMockRepositoryProvider(ctrl),
			}

			tt.mock(fields)

			uc := &UseCase{
				profileRepository: fields.profileRepository,
			}
			gotResp, gotHttpStatus := uc.CreateUserProfile(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UseCase.CreateUserProfile() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if gotHttpStatus != tt.wantHttpStatus {
				t.Errorf("UseCase.CreateUserProfile() gotHttpStatus = %v, want %v", gotHttpStatus, tt.wantHttpStatus)
			}
		})
	}
}

func TestUseCase_UpdateUserProfile(t *testing.T) {
	type fields struct {
		profileRepository *mockprofile.MockRepositoryProvider
	}
	type args struct {
		ctx context.Context
		req request.UpdateProfileRequest
	}
	tests := []struct {
		name           string
		args           args
		wantResp       writer.Response
		wantHttpStatus int
		mock           func(mock fields)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: request.UpdateProfileRequest{
					ProfileCode:    userProfile.ProfileCode,
					WantedJobTitle: userProfile.WantedJobTitle,
					FirstName:      userProfile.FirstName,
					LastName:       userProfile.LastName,
					Email:          userProfile.Email,
					Phone:          userProfile.Phone,
					Country:        userProfile.Country,
					City:           userProfile.City,
					Address:        userProfile.Address,
					PostalCode:     userProfile.PostalCode,
					DrivingLicense: userProfile.DrivingLicense,
					Nationality:    userProfile.Nationality,
					PlaceOfBirth:   userProfile.PlaceOfBirth,
					DateOfBirth:    "12-12-2024",
				},
			},
			mock: func(mock fields) {
				mock.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(userProfile, nil)

				mock.profileRepository.EXPECT().UpdateProfile(context.Background(), entity.Profile{
					ProfileCode:    userProfile.ProfileCode,
					WantedJobTitle: userProfile.WantedJobTitle,
					FirstName:      userProfile.FirstName,
					LastName:       userProfile.LastName,
					Email:          userProfile.Email,
					Phone:          userProfile.Phone,
					Country:        userProfile.Country,
					City:           userProfile.City,
					Address:        userProfile.Address,
					PostalCode:     userProfile.PostalCode,
					DrivingLicense: userProfile.DrivingLicense,
					Nationality:    userProfile.Nationality,
					PlaceOfBirth:   userProfile.PlaceOfBirth,
					DateOfBirth:    userProfile.DateOfBirth,
				}).Return(userProfile.ProfileCode, nil)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseSuccess.Code,
				Message: constant.ResponseSuccess.Description,
				Data: response.CreateProfileResponse{
					ProfileCode: userProfile.ProfileCode,
				},
			},
			wantHttpStatus: constant.ResponseSuccess.Status,
		},
		{
			name: "failed insert db",
			args: args{
				ctx: context.Background(),
				req: request.UpdateProfileRequest{
					ProfileCode:    userProfile.ProfileCode,
					WantedJobTitle: userProfile.WantedJobTitle,
					FirstName:      userProfile.FirstName,
					LastName:       userProfile.LastName,
					Email:          userProfile.Email,
					Phone:          userProfile.Phone,
					Country:        userProfile.Country,
					City:           userProfile.City,
					Address:        userProfile.Address,
					PostalCode:     userProfile.PostalCode,
					DrivingLicense: userProfile.DrivingLicense,
					Nationality:    userProfile.Nationality,
					PlaceOfBirth:   userProfile.PlaceOfBirth,
					DateOfBirth:    "12-12-2024",
				},
			},
			mock: func(mock fields) {
				mock.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(userProfile, nil)

				mock.profileRepository.EXPECT().UpdateProfile(context.Background(), entity.Profile{
					ProfileCode:    userProfile.ProfileCode,
					WantedJobTitle: userProfile.WantedJobTitle,
					FirstName:      userProfile.FirstName,
					LastName:       userProfile.LastName,
					Email:          userProfile.Email,
					Phone:          userProfile.Phone,
					Country:        userProfile.Country,
					City:           userProfile.City,
					Address:        userProfile.Address,
					PostalCode:     userProfile.PostalCode,
					DrivingLicense: userProfile.DrivingLicense,
					Nationality:    userProfile.Nationality,
					PlaceOfBirth:   userProfile.PlaceOfBirth,
					DateOfBirth:    userProfile.DateOfBirth,
				}).Return(uint64(0), assert.AnError)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseInternalServerError.Code,
				Message: constant.ResponseInternalServerError.Description,
			},
			wantHttpStatus: constant.ResponseInternalServerError.Status,
		},
		{
			name: "failed invalid date of birth",
			args: args{
				ctx: context.Background(),
				req: request.UpdateProfileRequest{
					ProfileCode:    userProfile.ProfileCode,
					WantedJobTitle: userProfile.WantedJobTitle,
					FirstName:      userProfile.FirstName,
					LastName:       userProfile.LastName,
					Email:          userProfile.Email,
					Phone:          userProfile.Phone,
					Country:        userProfile.Country,
					City:           userProfile.City,
					Address:        userProfile.Address,
					PostalCode:     userProfile.PostalCode,
					DrivingLicense: userProfile.DrivingLicense,
					Nationality:    userProfile.Nationality,
					PlaceOfBirth:   userProfile.PlaceOfBirth,
					DateOfBirth:    "12-12-20245",
				},
			},
			mock: func(mock fields) {
				mock.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(userProfile, nil)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseValidationError.Code,
				Message: constant.ResponseValidationError.Description,
			},
			wantHttpStatus: constant.ResponseValidationError.Status,
		},
		{
			name: "failed not found",
			args: args{
				ctx: context.Background(),
				req: request.UpdateProfileRequest{
					ProfileCode:    userProfile.ProfileCode,
					WantedJobTitle: userProfile.WantedJobTitle,
					FirstName:      userProfile.FirstName,
					LastName:       userProfile.LastName,
					Email:          userProfile.Email,
					Phone:          userProfile.Phone,
					Country:        userProfile.Country,
					City:           userProfile.City,
					Address:        userProfile.Address,
					PostalCode:     userProfile.PostalCode,
					DrivingLicense: userProfile.DrivingLicense,
					Nationality:    userProfile.Nationality,
					PlaceOfBirth:   userProfile.PlaceOfBirth,
					DateOfBirth:    "12-12-2024",
				},
			},
			mock: func(mock fields) {
				mock.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(entity.Profile{}, constant.ErrorDatabaseNotFound)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseErrorNotFound.Code,
				Message: constant.ResponseErrorNotFound.Description,
			},
			wantHttpStatus: constant.ResponseErrorNotFound.Status,
		},
		{
			name: "failed on get data",
			args: args{
				ctx: context.Background(),
				req: request.UpdateProfileRequest{
					ProfileCode:    userProfile.ProfileCode,
					WantedJobTitle: userProfile.WantedJobTitle,
					FirstName:      userProfile.FirstName,
					LastName:       userProfile.LastName,
					Email:          userProfile.Email,
					Phone:          userProfile.Phone,
					Country:        userProfile.Country,
					City:           userProfile.City,
					Address:        userProfile.Address,
					PostalCode:     userProfile.PostalCode,
					DrivingLicense: userProfile.DrivingLicense,
					Nationality:    userProfile.Nationality,
					PlaceOfBirth:   userProfile.PlaceOfBirth,
					DateOfBirth:    "12-12-2024",
				},
			},
			mock: func(mock fields) {
				mock.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(entity.Profile{}, assert.AnError)
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
				profileRepository: mockprofile.NewMockRepositoryProvider(ctrl),
			}

			tt.mock(fields)

			uc := &UseCase{
				profileRepository: fields.profileRepository,
			}
			gotResp, gotHttpStatus := uc.UpdateUserProfile(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UseCase.UpdateUserProfile() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if gotHttpStatus != tt.wantHttpStatus {
				t.Errorf("UseCase.UpdateUserProfile() gotHttpStatus = %v, want %v", gotHttpStatus, tt.wantHttpStatus)
			}
		})
	}
}
