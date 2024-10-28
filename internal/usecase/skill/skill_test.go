package skill

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mockprofile "github.com/ramasuryananda/dummy-cv/gomock/repository/mockProfile"
	mockskill "github.com/ramasuryananda/dummy-cv/gomock/repository/mockSkill"
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

var skillData = entity.Skill{
	ID:          1,
	ProfileCode: userProfile.ProfileCode,
	Skill:       "test skill",
	Level:       "test level",
}

func TestUseCase_GetUserSkill(t *testing.T) {
	type fields struct {
		skillRepository   *mockskill.MockRepositoryProvider
		profileRepository *mockprofile.MockRepositoryProvider
	}
	type args struct {
		ctx context.Context
		req request.GetSkillRequest
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
				req: request.GetSkillRequest{
					ProfileCode: userProfile.ProfileCode,
				},
			},
			mock: func(mocks fields) {
				mocks.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(userProfile, nil)

				mocks.skillRepository.EXPECT().GetSkillByProfileCode(context.Background(), userProfile.ProfileCode).Return([]entity.Skill{
					skillData,
				}, nil)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseSuccess.Code,
				Message: constant.ResponseSuccess.Description,
				Data: []response.SkillDataResponse{
					{
						ID:    skillData.ID,
						Skill: skillData.Skill,
						Level: skillData.Level,
					},
				},
			},
			wantHttpStatus: constant.ResponseSuccess.Status,
		},
		{
			name: "failed get user skill",
			args: args{
				ctx: context.Background(),
				req: request.GetSkillRequest{
					ProfileCode: userProfile.ProfileCode,
				},
			},
			mock: func(mocks fields) {
				mocks.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(userProfile, nil)

				mocks.skillRepository.EXPECT().GetSkillByProfileCode(context.Background(), userProfile.ProfileCode).Return([]entity.Skill{}, assert.AnError)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseInternalServerError.Code,
				Message: constant.ResponseInternalServerError.Description,
			},
			wantHttpStatus: constant.ResponseInternalServerError.Status,
		},
		{
			name: "failed get user profile",
			args: args{
				ctx: context.Background(),
				req: request.GetSkillRequest{
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
				req: request.GetSkillRequest{
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
				profileRepository: mockprofile.NewMockRepositoryProvider(ctrl),
				skillRepository:   mockskill.NewMockRepositoryProvider(ctrl),
			}

			tt.mock(fields)

			uc := &UseCase{
				profileRepository: fields.profileRepository,
				skillRepository:   fields.skillRepository,
			}
			gotResp, gotHttpStatus := uc.GetUserSkill(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UseCase.GetUserSkill() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if gotHttpStatus != tt.wantHttpStatus {
				t.Errorf("UseCase.GetUserSkill() gotHttpStatus = %v, want %v", gotHttpStatus, tt.wantHttpStatus)
			}
		})
	}
}

func TestUseCase_CreateSkillData(t *testing.T) {
	type fields struct {
		skillRepository   *mockskill.MockRepositoryProvider
		profileRepository *mockprofile.MockRepositoryProvider
	}
	type args struct {
		ctx context.Context
		req request.CreateSkillRequest
	}
	tests := []struct {
		name           string
		mock           func(mocks fields)
		args           args
		wantResp       writer.Response
		wantHttpStatus int
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				req: request.CreateSkillRequest{
					ProfileCode: userProfile.ProfileCode,
					Skill:       "test skill",
					Level:       "test level",
				},
			},
			mock: func(mocks fields) {
				mocks.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(userProfile, nil)

				mocks.skillRepository.EXPECT().CreateSkillData(context.Background(), entity.Skill{
					ProfileCode: userProfile.ProfileCode,
					Skill:       skillData.Skill,
					Level:       skillData.Level,
				}).Return(uint64(1), nil)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseSuccess.Code,
				Message: constant.ResponseSuccess.Description,
				Data: response.CreateSkillResponse{
					Id:          skillData.ID,
					ProfileCode: skillData.ProfileCode,
				},
			},
			wantHttpStatus: constant.ResponseSuccess.Status,
		},
		{
			name: "Failed Create Data",
			args: args{
				ctx: context.Background(),
				req: request.CreateSkillRequest{
					ProfileCode: userProfile.ProfileCode,
					Skill:       "test skill",
					Level:       "test level",
				},
			},
			mock: func(mocks fields) {
				mocks.profileRepository.EXPECT().GetUserByProfileCode(context.Background(), userProfile.ProfileCode).Return(userProfile, nil)

				mocks.skillRepository.EXPECT().CreateSkillData(context.Background(), entity.Skill{
					ProfileCode: userProfile.ProfileCode,
					Skill:       skillData.Skill,
					Level:       skillData.Level,
				}).Return(uint64(0), assert.AnError)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseInternalServerError.Code,
				Message: constant.ResponseInternalServerError.Description,
			},
			wantHttpStatus: constant.ResponseInternalServerError.Status,
		},
		{
			name: "Failed Get Profile Data",
			args: args{
				ctx: context.Background(),
				req: request.CreateSkillRequest{
					ProfileCode: userProfile.ProfileCode,
					Skill:       "test skill",
					Level:       "test level",
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
			name: "Failed Get Profile Data not found",
			args: args{
				ctx: context.Background(),
				req: request.CreateSkillRequest{
					ProfileCode: userProfile.ProfileCode,
					Skill:       "test skill",
					Level:       "test level",
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
				profileRepository: mockprofile.NewMockRepositoryProvider(ctrl),
				skillRepository:   mockskill.NewMockRepositoryProvider(ctrl),
			}

			tt.mock(fields)

			uc := &UseCase{
				profileRepository: fields.profileRepository,
				skillRepository:   fields.skillRepository,
			}
			gotResp, gotHttpStatus := uc.CreateSkillData(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UseCase.CreateSkillData() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if gotHttpStatus != tt.wantHttpStatus {
				t.Errorf("UseCase.CreateSkillData() gotHttpStatus = %v, want %v", gotHttpStatus, tt.wantHttpStatus)
			}
		})
	}
}

func TestUseCase_DeleteSkillData(t *testing.T) {
	type fields struct {
		skillRepository   *mockskill.MockRepositoryProvider
		profileRepository *mockprofile.MockRepositoryProvider
	}
	type args struct {
		ctx context.Context
		req request.DeleteSkillRequest
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
				req: request.DeleteSkillRequest{
					ProfileCode: userProfile.ProfileCode,
					ID:          skillData.ID,
				},
			},
			mock: func(mocks fields) {
				mocks.skillRepository.EXPECT().GetFirstSkillByProfileCodeandID(context.Background(), userProfile.ProfileCode, skillData.ID).Return(skillData, nil)

				mocks.skillRepository.EXPECT().DeleteSkillData(context.Background(), skillData.ProfileCode, skillData.ID).Return(nil)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseSuccess.Code,
				Message: constant.ResponseSuccess.Description,
				Data: response.DeleteSkillResponse{
					ProfileCode: userProfile.ProfileCode,
				},
			},
			wantHttpStatus: constant.ResponseSuccess.Status,
		},
		{
			name: "failed delete",
			args: args{
				ctx: context.Background(),
				req: request.DeleteSkillRequest{
					ProfileCode: userProfile.ProfileCode,
					ID:          skillData.ID,
				},
			},
			mock: func(mocks fields) {
				mocks.skillRepository.EXPECT().GetFirstSkillByProfileCodeandID(context.Background(), userProfile.ProfileCode, skillData.ID).Return(skillData, nil)

				mocks.skillRepository.EXPECT().DeleteSkillData(context.Background(), skillData.ProfileCode, skillData.ID).Return(assert.AnError)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseInternalServerError.Code,
				Message: constant.ResponseInternalServerError.Description,
			},
			wantHttpStatus: constant.ResponseInternalServerError.Status,
		},
		{
			name: "failed get data",
			args: args{
				ctx: context.Background(),
				req: request.DeleteSkillRequest{
					ProfileCode: userProfile.ProfileCode,
					ID:          skillData.ID,
				},
			},
			mock: func(mocks fields) {
				mocks.skillRepository.EXPECT().GetFirstSkillByProfileCodeandID(context.Background(), userProfile.ProfileCode, skillData.ID).Return(entity.Skill{}, assert.AnError)
			},
			wantResp: writer.Response{
				Code:    constant.ResponseInternalServerError.Code,
				Message: constant.ResponseInternalServerError.Description,
			},
			wantHttpStatus: constant.ResponseInternalServerError.Status,
		},
		{
			name: "failed data not found",
			args: args{
				ctx: context.Background(),
				req: request.DeleteSkillRequest{
					ProfileCode: userProfile.ProfileCode,
					ID:          skillData.ID,
				},
			},
			mock: func(mocks fields) {
				mocks.skillRepository.EXPECT().GetFirstSkillByProfileCodeandID(context.Background(), userProfile.ProfileCode, skillData.ID).Return(entity.Skill{}, constant.ErrorDatabaseNotFound)
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
				profileRepository: mockprofile.NewMockRepositoryProvider(ctrl),
				skillRepository:   mockskill.NewMockRepositoryProvider(ctrl),
			}

			tt.mock(fields)

			uc := &UseCase{
				profileRepository: fields.profileRepository,
				skillRepository:   fields.skillRepository,
			}

			gotResp, gotHttpStatus := uc.DeleteSkillData(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UseCase.DeleteSkillData() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if gotHttpStatus != tt.wantHttpStatus {
				t.Errorf("UseCase.DeleteSkillData() gotHttpStatus = %v, want %v", gotHttpStatus, tt.wantHttpStatus)
			}
		})
	}
}
