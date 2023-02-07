package setting_test

import (
	"context"
	"database/sql"
	"errors"

	"sync"

	"time"

	"go-template/client/redis"
	"go-template/controllers"
	settingController "go-template/controllers/setting"
	settingEntity "go-template/entities/setting"
	errHelper "go-template/helpers/errors"
	redisMock "go-template/mocks/redis"
	settingMock "go-template/mocks/setting"
	settingService "go-template/services/setting"
	"reflect"
	"testing"

	sqlxmock "github.com/zhashkevych/go-sqlxmock"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

var (
	caseError = errors.New("error")
)

func TestSettingService_SelectAll(t *testing.T) {
	settingData := []settingEntity.Setting{
		{
			ID:        1,
			Name:      settingEntity.SettingTypeDefault,
			Value:     "{}",
			CreatedAt: "2021-12-01T13:19:12.801+07:00",
			UpdatedAt: "2021-12-01T15:14:38.09019+07:00",
		},
	}
	tests := []struct {
		name    string
		args    settingController.SettingParameter
		doMock  func(mock *settingMock.MockSettingRepo, res []settingEntity.Setting)
		want    []settingEntity.Setting
		wantErr error
	}{
		{
			name: "flow success",
			args: settingController.SettingParameter{},
			doMock: func(mock *settingMock.MockSettingRepo, res []settingEntity.Setting) {
				mock.EXPECT().SelectAll(gomock.Any(), &settingController.SettingParameter{
					DefaultParameter: controllers.DefaultParameter{
						Page:    1,
						Limit:   10,
						OrderBy: "def.created_at",
						Sort:    "asc",
					},
				}).Return(res, nil)
			},
			want:    settingData,
			wantErr: nil,
		},
		{
			name: "flow error",
			args: settingController.SettingParameter{},
			doMock: func(mock *settingMock.MockSettingRepo, res []settingEntity.Setting) {
				mock.EXPECT().SelectAll(gomock.Any(), &settingController.SettingParameter{
					DefaultParameter: controllers.DefaultParameter{
						Page:    1,
						Limit:   10,
						OrderBy: "def.created_at",
						Sort:    "asc",
					},
				}).Return(res, caseError)
			},
			want:    []settingEntity.Setting{},
			wantErr: caseError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRedis := redisMock.NewMockIRedis(mockCtrl)

			mockSettingRepo := settingMock.NewMockSettingRepo(mockCtrl)
			settingService := settingService.NewSettingService(mockSettingRepo, mockRedis)
			tt.doMock(mockSettingRepo, tt.want)

			got, err := settingService.SelectAll(context.Background(), &tt.args)
			if err != tt.wantErr {
				t.Errorf("settingService.SelectAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			alwaysEqual := cmp.Comparer(func(_, _ interface{}) bool { return true })
			opt := cmp.FilterValues(func(x, y interface{}) bool {
				vx, vy := reflect.ValueOf(x), reflect.ValueOf(y)
				return (vx.IsValid() && vy.IsValid() && vx.Type() == vy.Type()) &&
					(vx.Kind() == reflect.Slice || vx.Kind() == reflect.Map) &&
					(vx.Len() == 0 && vy.Len() == 0)
			}, alwaysEqual)
			if !cmp.Equal(got, tt.want, opt) {
				t.Errorf("settingService.SelectAll()  = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSettingService_FindAll(t *testing.T) {
	settingData := []settingEntity.Setting{
		{
			ID:        1,
			Name:      settingEntity.SettingTypeDefault,
			Value:     "{}",
			CreatedAt: "2021-12-01T13:19:12.801+07:00",
			UpdatedAt: "2021-12-01T15:14:38.09019+07:00",
		},
	}
	tests := []struct {
		name    string
		args    settingController.SettingParameter
		doMock  func(mock *settingMock.MockSettingRepo, res []settingEntity.Setting)
		want    []settingEntity.Setting
		wantErr error
	}{
		{
			name: "flow success",
			args: settingController.SettingParameter{},
			doMock: func(mock *settingMock.MockSettingRepo, res []settingEntity.Setting) {
				mock.EXPECT().FindAll(gomock.Any(), &settingController.SettingParameter{
					DefaultParameter: controllers.DefaultParameter{
						Page:    1,
						Limit:   10,
						OrderBy: "def.created_at",
						Sort:    "asc",
					},
				}).Return(res, 1, nil)
			},
			want:    settingData,
			wantErr: nil,
		},
		{
			name: "flow error",
			args: settingController.SettingParameter{},
			doMock: func(mock *settingMock.MockSettingRepo, res []settingEntity.Setting) {
				mock.EXPECT().FindAll(gomock.Any(), &settingController.SettingParameter{
					DefaultParameter: controllers.DefaultParameter{
						Page:    1,
						Limit:   10,
						OrderBy: "def.created_at",
						Sort:    "asc",
					},
				}).Return(res, 1, caseError)
			},
			want:    []settingEntity.Setting{},
			wantErr: caseError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRedis := redisMock.NewMockIRedis(mockCtrl)

			mockSettingRepo := settingMock.NewMockSettingRepo(mockCtrl)
			settingService := settingService.NewSettingService(mockSettingRepo, mockRedis)
			tt.doMock(mockSettingRepo, tt.want)

			got, _, err := settingService.FindAll(context.Background(), &tt.args)
			if err != tt.wantErr {
				t.Errorf("settingService.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			alwaysEqual := cmp.Comparer(func(_, _ interface{}) bool { return true })
			opt := cmp.FilterValues(func(x, y interface{}) bool {
				vx, vy := reflect.ValueOf(x), reflect.ValueOf(y)
				return (vx.IsValid() && vy.IsValid() && vx.Type() == vy.Type()) &&
					(vx.Kind() == reflect.Slice || vx.Kind() == reflect.Map) &&
					(vx.Len() == 0 && vy.Len() == 0)
			}, alwaysEqual)
			if !cmp.Equal(got, tt.want, opt) {
				t.Errorf("settingService.FindAll()  = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSettingService_FindOne(t *testing.T) {
	settingData := settingEntity.Setting{
		ID:        1,
		Name:      settingEntity.SettingTypeDefault,
		Value:     "{}",
		CreatedAt: "2021-12-01T13:19:12.801+07:00",
		UpdatedAt: "2021-12-01T15:14:38.09019+07:00",
	}
	tests := []struct {
		name           string
		args           settingController.SettingParameter
		doMock         func(mock *settingMock.MockSettingRepo, res settingEntity.Setting)
		doMockRedisSet func(mock *redisMock.MockIRedis)
		want           settingEntity.Setting
		wantErr        error
	}{
		{
			name: "flow success",
			args: settingController.SettingParameter{},
			doMock: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{}).Return(res, nil)
			},
			doMockRedisSet: func(mock *redisMock.MockIRedis) {
				mock.EXPECT().Set(gomock.Any(), redis.Setting, &settingData, 60*time.Minute).Return(nil)
			},
			want:    settingData,
			wantErr: nil,
		},
		{
			name: "flow error",
			args: settingController.SettingParameter{},
			doMock: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{}).Return(res, caseError)
			},
			doMockRedisSet: func(mock *redisMock.MockIRedis) {},
			want:           settingEntity.Setting{},
			wantErr:        caseError,
		},
		{
			name: "flow error data not found",
			args: settingController.SettingParameter{},
			doMock: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{}).Return(res, sql.ErrNoRows)
			},
			doMockRedisSet: func(mock *redisMock.MockIRedis) {},
			want:           settingEntity.Setting{},
			wantErr:        errHelper.DataNotFound.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRedis := redisMock.NewMockIRedis(mockCtrl)

			mockSettingRepo := settingMock.NewMockSettingRepo(mockCtrl)
			settingService := settingService.NewSettingService(mockSettingRepo, mockRedis)
			tt.doMock(mockSettingRepo, tt.want)
			tt.doMockRedisSet(mockRedis)

			got, err := settingService.FindOne(context.Background(), &tt.args)
			if err != tt.wantErr {
				t.Errorf("settingService.FindOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			alwaysEqual := cmp.Comparer(func(_, _ interface{}) bool { return true })
			opt := cmp.FilterValues(func(x, y interface{}) bool {
				vx, vy := reflect.ValueOf(x), reflect.ValueOf(y)
				return (vx.IsValid() && vy.IsValid() && vx.Type() == vy.Type()) &&
					(vx.Kind() == reflect.Slice || vx.Kind() == reflect.Map) &&
					(vx.Len() == 0 && vy.Len() == 0)
			}, alwaysEqual)
			if !cmp.Equal(got, tt.want, opt) {
				t.Errorf("settingService.FindOne()  = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSettingService_FindDefaultSetting(t *testing.T) {
	settingData := settingEntity.SettingDefault{}
	tests := []struct {
		name        string
		doMock      func(mock *settingMock.MockSettingRepo, res settingEntity.Setting)
		wantSetting settingEntity.Setting
		want        settingEntity.SettingDefault
		wantErr     error
	}{
		{
			name: "flow success",
			doMock: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeDefault,
				}).Return(res, nil)
			},
			want:    settingData,
			wantErr: nil,
		},
		{
			name: "flow error",
			doMock: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeDefault,
				}).Return(res, caseError)
			},
			want:    settingEntity.SettingDefault{},
			wantErr: caseError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRedis := redisMock.NewMockIRedis(mockCtrl)

			mockSettingRepo := settingMock.NewMockSettingRepo(mockCtrl)
			settingService := settingService.NewSettingService(mockSettingRepo, mockRedis)
			tt.doMock(mockSettingRepo, tt.wantSetting)

			got, err := settingService.FindDefaultSetting(context.Background())
			if err != tt.wantErr {
				t.Errorf("settingService.FindDefaultSetting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			alwaysEqual := cmp.Comparer(func(_, _ interface{}) bool { return true })
			opt := cmp.FilterValues(func(x, y interface{}) bool {
				vx, vy := reflect.ValueOf(x), reflect.ValueOf(y)
				return (vx.IsValid() && vy.IsValid() && vx.Type() == vy.Type()) &&
					(vx.Kind() == reflect.Slice || vx.Kind() == reflect.Map) &&
					(vx.Len() == 0 && vy.Len() == 0)
			}, alwaysEqual)
			if !cmp.Equal(got, tt.want, opt) {
				t.Errorf("settingService.FindDefaultSetting()  = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSettingService_FindIncentiveBudgetSetting(t *testing.T) {
	settingData := settingEntity.SettingIncentiveBudget{
		OriginAccountNumber:      "",
		DestinationAccountNumber: "",
		OriginDescription:        "",
		DestinationDescription:   "",
		MinimumBudget:            0,
	}
	tests := []struct {
		name        string
		doMock      func(mock *settingMock.MockSettingRepo, res settingEntity.Setting)
		wantSetting settingEntity.Setting
		want        settingEntity.SettingIncentiveBudget
		wantErr     error
	}{
		{
			name: "flow success",
			doMock: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeIncentiveBudget,
				}).Return(res, nil)
			},
			want:    settingData,
			wantErr: nil,
		},
		{
			name: "flow error",
			doMock: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeIncentiveBudget,
				}).Return(res, caseError)
			},
			want:    settingEntity.SettingIncentiveBudget{},
			wantErr: caseError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRedis := redisMock.NewMockIRedis(mockCtrl)

			mockSettingRepo := settingMock.NewMockSettingRepo(mockCtrl)
			settingService := settingService.NewSettingService(mockSettingRepo, mockRedis)
			tt.doMock(mockSettingRepo, tt.wantSetting)

			got, err := settingService.FindIncentiveBudgetSetting(context.Background())
			if err != tt.wantErr {
				t.Errorf("settingService.FindIncentiveBudgetSetting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			alwaysEqual := cmp.Comparer(func(_, _ interface{}) bool { return true })
			opt := cmp.FilterValues(func(x, y interface{}) bool {
				vx, vy := reflect.ValueOf(x), reflect.ValueOf(y)
				return (vx.IsValid() && vy.IsValid() && vx.Type() == vy.Type()) &&
					(vx.Kind() == reflect.Slice || vx.Kind() == reflect.Map) &&
					(vx.Len() == 0 && vy.Len() == 0)
			}, alwaysEqual)
			if !cmp.Equal(got, tt.want, opt) {
				t.Errorf("settingService.FindIncentiveBudgetSetting()  = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSettingService_FindIncentiveBudgetApprovalSetting(t *testing.T) {
	settingData := settingEntity.SettingIncentiveBudgetApproval{
		DebitAccountNumber:  "",
		CreditAccountNumber: "",
	}
	tests := []struct {
		name        string
		doMock      func(mock *settingMock.MockSettingRepo, res settingEntity.Setting)
		wantSetting settingEntity.Setting
		want        settingEntity.SettingIncentiveBudgetApproval
		wantErr     error
	}{
		{
			name: "flow success",
			doMock: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeIncentiveBudgetApproval,
				}).Return(res, nil)
			},
			want:    settingData,
			wantErr: nil,
		},
		{
			name: "flow error",
			doMock: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeIncentiveBudgetApproval,
				}).Return(res, caseError)
			},
			want:    settingEntity.SettingIncentiveBudgetApproval{},
			wantErr: caseError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRedis := redisMock.NewMockIRedis(mockCtrl)

			mockSettingRepo := settingMock.NewMockSettingRepo(mockCtrl)
			settingService := settingService.NewSettingService(mockSettingRepo, mockRedis)
			tt.doMock(mockSettingRepo, tt.wantSetting)

			got, err := settingService.FindIncentiveBudgetApprovalSetting(context.Background())
			if err != tt.wantErr {
				t.Errorf("settingService.FindIncentiveBudgetApprovalSetting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			alwaysEqual := cmp.Comparer(func(_, _ interface{}) bool { return true })
			opt := cmp.FilterValues(func(x, y interface{}) bool {
				vx, vy := reflect.ValueOf(x), reflect.ValueOf(y)
				return (vx.IsValid() && vy.IsValid() && vx.Type() == vy.Type()) &&
					(vx.Kind() == reflect.Slice || vx.Kind() == reflect.Map) &&
					(vx.Len() == 0 && vy.Len() == 0)
			}, alwaysEqual)
			if !cmp.Equal(got, tt.want, opt) {
				t.Errorf("settingService.FindIncentiveBudgetApprovalSetting()  = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSettingService_Create(t *testing.T) {
	ctx := context.Background()
	db, mocks, _ := sqlxmock.Newx()
	mocks.ExpectBegin()
	tx, _ := db.Beginx()
	settingData := settingEntity.Setting{
		Name:  settingEntity.SettingTypeDefault,
		Value: "{\"asana_email_assignee\": \"test@mail.com\",\"asana_comment_mention_id\": \"123\",\"asana_email_followers\": [\"test@mail.com\"]}",
	}
	settingDataIncentiveBudget := settingEntity.Setting{
		Name:  settingEntity.SettingTypeIncentiveBudget,
		Value: "{\"origin_account_number\": \"12345\", \"destination_account_number\": \"12345\", \"origin_description\": \"BD ANR\", \"destination_description\": \"BRI Escrow ANR\"}",
	}
	settingDataBudgetApproval := settingEntity.Setting{
		Name:  settingEntity.SettingTypeIncentiveBudgetApproval,
		Value: "{\"debit_account_number\": \"12345\", \"credit_account_number\": \"12345\"}",
	}
	tests := []struct {
		name           string
		args           settingEntity.Setting
		doMockCreateTx func(mock *settingMock.MockSettingRepo)
		doMockFindOne  func(mock *settingMock.MockSettingRepo, res settingEntity.Setting)
		doMock         func(mock *settingMock.MockSettingRepo, res int)
		WantFindOne    settingEntity.Setting
		want           int
		wantErr        error
	}{
		{
			name: "flow success default",
			args: settingData,
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeDefault,
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {
				mock.EXPECT().CreateTx(ctx).Return(tx, nil)
			},
			doMock: func(mock *settingMock.MockSettingRepo, res int) {
				mock.EXPECT().Create(gomock.Any(), tx, &settingData).Return(res, nil)
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "flow success incentive budget",
			args: settingDataIncentiveBudget,
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeIncentiveBudget,
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {
				mock.EXPECT().CreateTx(ctx).Return(tx, nil)
			},
			doMock: func(mock *settingMock.MockSettingRepo, res int) {
				mock.EXPECT().Create(gomock.Any(), tx, &settingDataIncentiveBudget).Return(res, nil)
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "flow success budget approval",
			args: settingDataBudgetApproval,
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeIncentiveBudgetApproval,
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {
				mock.EXPECT().CreateTx(ctx).Return(tx, nil)
			},
			doMock: func(mock *settingMock.MockSettingRepo, res int) {
				mock.EXPECT().Create(gomock.Any(), tx, &settingDataBudgetApproval).Return(res, nil)
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "flow error",
			args: settingData,
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeDefault,
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {
				mock.EXPECT().CreateTx(ctx).Return(tx, nil)
			},
			doMock: func(mock *settingMock.MockSettingRepo, res int) {
				mock.EXPECT().Create(gomock.Any(), tx, &settingData).Return(res, caseError)
			},
			want:    1,
			wantErr: caseError,
		},
		{
			name: "flow error create tx",
			args: settingData,
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeDefault,
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {
				mock.EXPECT().CreateTx(ctx).Return(tx, caseError)
			},
			doMock:  func(mock *settingMock.MockSettingRepo, res int) {},
			want:    0,
			wantErr: caseError,
		},
		{
			name: "flow error find one",
			args: settingData,
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeDefault,
				}).Return(res, caseError)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {},
			doMock:         func(mock *settingMock.MockSettingRepo, res int) {},
			want:           0,
			wantErr:        caseError,
		},
		{
			name: "flow error name registered",
			args: settingData,
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeDefault,
				}).Return(settingEntity.Setting{
					ID: 1,
				}, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {},
			doMock:         func(mock *settingMock.MockSettingRepo, res int) {},
			want:           0,
			wantErr:        errHelper.SettingRegistered.Error,
		},
		{
			name: "flow error asana email assignee",
			args: settingEntity.Setting{
				Name:  settingEntity.SettingTypeDefault,
				Value: "{\"asana_email_assignee\": \"\"}",
			},
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeDefault,
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {},
			doMock:         func(mock *settingMock.MockSettingRepo, res int) {},
			want:           0,
			wantErr:        errHelper.AsanaEmailAssigneeRequired.Error,
		},
		{
			name: "flow error asana comment mention id",
			args: settingEntity.Setting{
				Name:  settingEntity.SettingTypeDefault,
				Value: "{\"asana_email_assignee\": \"test@mail.com\",\"asana_comment_mention_id\": \"\",\"asana_email_followers\": [\"test@mail.com\"]}",
			},
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeDefault,
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {},
			doMock:         func(mock *settingMock.MockSettingRepo, res int) {},
			want:           0,
			wantErr:        errHelper.AsanaCommentMentionIDRequired.Error,
		},
		{
			name: "flow error asana email followers",
			args: settingEntity.Setting{
				Name:  settingEntity.SettingTypeDefault,
				Value: "{\"asana_email_assignee\": \"test@mail.com\",\"asana_comment_mention_id\": \"123\",\"asana_email_followers\": []}",
			},
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeDefault,
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {},
			doMock:         func(mock *settingMock.MockSettingRepo, res int) {},
			want:           0,
			wantErr:        errHelper.AsanaEmailFollowersRequired.Error,
		},
		{
			name: "flow error origin account number",
			args: settingEntity.Setting{
				Name:  settingEntity.SettingTypeIncentiveBudget,
				Value: "{\"origin_account_number\": \"\", \"destination_account_number\": \"12345\"}",
			},
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeIncentiveBudget,
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {},
			doMock:         func(mock *settingMock.MockSettingRepo, res int) {},
			want:           0,
			wantErr:        errHelper.OriginAccountNumberRequired.Error,
		},
		{
			name: "flow error destination account number",
			args: settingEntity.Setting{
				Name:  settingEntity.SettingTypeIncentiveBudget,
				Value: "{\"origin_account_number\": \"12345\", \"destination_account_number\": \"\"}",
			},
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeIncentiveBudget,
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {},
			doMock:         func(mock *settingMock.MockSettingRepo, res int) {},
			want:           0,
			wantErr:        errHelper.DestinationAccountNumberRequired.Error,
		},
		{
			name: "flow error origin decription",
			args: settingEntity.Setting{
				Name:  settingEntity.SettingTypeIncentiveBudget,
				Value: "{\"origin_account_number\": \"12345\", \"destination_account_number\": \"12345\", \"origin_description\": \"\", \"destination_description\": \"desc\"}",
			},
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeIncentiveBudget,
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {},
			doMock:         func(mock *settingMock.MockSettingRepo, res int) {},
			want:           0,
			wantErr:        errHelper.OriginDescriptionRequired.Error,
		},
		{
			name: "flow error destination decription",
			args: settingEntity.Setting{
				Name:  settingEntity.SettingTypeIncentiveBudget,
				Value: "{\"origin_account_number\": \"12345\", \"destination_account_number\": \"12345\", \"origin_description\": \"desc\", \"destination_description\": \"\"}",
			},
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeIncentiveBudget,
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {},
			doMock:         func(mock *settingMock.MockSettingRepo, res int) {},
			want:           0,
			wantErr:        errHelper.DestinationDescriptionRequired.Error,
		},
		{
			name: "flow error debit account number",
			args: settingEntity.Setting{
				Name:  settingEntity.SettingTypeIncentiveBudgetApproval,
				Value: "{\"debit_account_number\": \"\", \"credit_account_number\": \"123\"}",
			},
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeIncentiveBudgetApproval,
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {},
			doMock:         func(mock *settingMock.MockSettingRepo, res int) {},
			want:           0,
			wantErr:        errHelper.DebitAccountNumberRequired.Error,
		},
		{
			name: "flow error credit account number",
			args: settingEntity.Setting{
				Name:  settingEntity.SettingTypeIncentiveBudgetApproval,
				Value: "{\"debit_account_number\": \"123\", \"credit_account_number\": \"\"}",
			},
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeIncentiveBudgetApproval,
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {},
			doMock:         func(mock *settingMock.MockSettingRepo, res int) {},
			want:           0,
			wantErr:        errHelper.CreditAccountNumberRequired.Error,
		},
		{
			name: "flow error incentive budget convert value",
			args: settingEntity.Setting{
				Name:  settingEntity.SettingTypeIncentiveBudget,
				Value: "[]",
			},
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeIncentiveBudget,
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {},
			doMock:         func(mock *settingMock.MockSettingRepo, res int) {},
			want:           0,
			wantErr:        errHelper.InvalidDataSetting.Error,
		},
		{
			name: "flow error budget approval convert value",
			args: settingEntity.Setting{
				Name:  settingEntity.SettingTypeIncentiveBudgetApproval,
				Value: "[]",
			},
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeIncentiveBudgetApproval,
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {},
			doMock:         func(mock *settingMock.MockSettingRepo, res int) {},
			want:           0,
			wantErr:        errHelper.InvalidDataSetting.Error,
		},
		{
			name: "flow error default convert value",
			args: settingEntity.Setting{
				Name:  settingEntity.SettingTypeDefault,
				Value: "[]",
			},
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeDefault,
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {},
			doMock:         func(mock *settingMock.MockSettingRepo, res int) {},
			want:           0,
			wantErr:        errHelper.InvalidDataSetting.Error,
		},
		{
			name: "flow error invalid name",
			args: settingEntity.Setting{
				Name:  "",
				Value: "[]",
			},
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: "",
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {},
			doMock:         func(mock *settingMock.MockSettingRepo, res int) {},
			want:           0,
			wantErr:        errHelper.InvalidName.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRedis := redisMock.NewMockIRedis(mockCtrl)

			mockSettingRepo := settingMock.NewMockSettingRepo(mockCtrl)
			settingService := settingService.NewSettingService(mockSettingRepo, mockRedis)
			tt.doMockFindOne(mockSettingRepo, tt.WantFindOne)
			tt.doMockCreateTx(mockSettingRepo)
			tt.doMock(mockSettingRepo, tt.want)

			got, err := settingService.Create(context.Background(), &tt.args)
			if err != tt.wantErr {
				t.Errorf("settingService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			alwaysEqual := cmp.Comparer(func(_, _ interface{}) bool { return true })
			opt := cmp.FilterValues(func(x, y interface{}) bool {
				vx, vy := reflect.ValueOf(x), reflect.ValueOf(y)
				return (vx.IsValid() && vy.IsValid() && vx.Type() == vy.Type()) &&
					(vx.Kind() == reflect.Slice || vx.Kind() == reflect.Map) &&
					(vx.Len() == 0 && vy.Len() == 0)
			}, alwaysEqual)
			if !cmp.Equal(got, tt.want, opt) {
				t.Errorf("settingService.Create()  = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSettingService_Update(t *testing.T) {
	ctx := context.Background()
	db, mocks, _ := sqlxmock.Newx()
	mocks.ExpectBegin()
	tx, _ := db.Beginx()
	settingData := settingEntity.Setting{
		ID:    1,
		Name:  settingEntity.SettingTypeDefault,
		Value: "{\"asana_email_assignee\": \"test@mail.com\",\"asana_comment_mention_id\": \"123\",\"asana_email_followers\": [\"test@mail.com\"]}",
	}
	tests := []struct {
		name              string
		args              settingEntity.Setting
		doMockCreateTx    func(mock *settingMock.MockSettingRepo)
		doMockFindOne     func(mock *settingMock.MockSettingRepo, res settingEntity.Setting)
		doMockFindOneName func(mock *settingMock.MockSettingRepo, res settingEntity.Setting)
		doMock            func(mock *settingMock.MockSettingRepo)
		doMockDelRedis    func(mock *redisMock.MockIRedis, wg *sync.WaitGroup)
		WantFindOne       settingEntity.Setting
		WantFindOneName   settingEntity.Setting
		wantErr           error
	}{
		{
			name: "flow success",
			args: settingData,
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					DefaultParameter: controllers.DefaultParameter{ID: 1},
				}).Return(settingData, nil)
			},
			doMockFindOneName: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingData.Name,
				}).Return(settingData, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {
				mock.EXPECT().CreateTx(ctx).Return(tx, nil)
			},
			doMock: func(mock *settingMock.MockSettingRepo) {
				mock.EXPECT().Update(gomock.Any(), tx, &settingData).Return(nil)
			},
			doMockDelRedis: func(mock *redisMock.MockIRedis, wg *sync.WaitGroup) {
				mock.EXPECT().Del(context.Background(), redis.Setting+settingData.Name).Do(func(arg0, arg1 interface{}) {
					defer wg.Done()
				})
			},
			WantFindOne:     settingData,
			WantFindOneName: settingData,
			wantErr:         nil,
		},
		{
			name: "flow error",
			args: settingData,
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					DefaultParameter: controllers.DefaultParameter{ID: 1},
				}).Return(res, nil)
			},
			doMockFindOneName: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeDefault,
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {
				mock.EXPECT().CreateTx(ctx).Return(tx, nil)
			},
			doMock: func(mock *settingMock.MockSettingRepo) {
				mock.EXPECT().Update(gomock.Any(), tx, &settingData).Return(caseError)
			},
			doMockDelRedis: func(mock *redisMock.MockIRedis, wg *sync.WaitGroup) {
				defer wg.Done()
			},
			wantErr: caseError,
		},
		{
			name: "flow error create tx",
			args: settingData,
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					DefaultParameter: controllers.DefaultParameter{ID: 1},
				}).Return(res, nil)
			},
			doMockFindOneName: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeDefault,
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {
				mock.EXPECT().CreateTx(ctx).Return(tx, caseError)
			},
			doMock: func(mock *settingMock.MockSettingRepo) {},
			doMockDelRedis: func(mock *redisMock.MockIRedis, wg *sync.WaitGroup) {
				defer wg.Done()
			},
			wantErr: caseError,
		},
		{
			name: "flow error find one",
			args: settingData,
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					DefaultParameter: controllers.DefaultParameter{ID: 1},
				}).Return(res, caseError)
			},
			doMockFindOneName: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {},
			doMockCreateTx:    func(mock *settingMock.MockSettingRepo) {},
			doMock:            func(mock *settingMock.MockSettingRepo) {},
			doMockDelRedis: func(mock *redisMock.MockIRedis, wg *sync.WaitGroup) {
				defer wg.Done()
			},
			wantErr: errHelper.DataNotFound.Error,
		},
		{
			name: "flow error find one name",
			args: settingData,
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					DefaultParameter: controllers.DefaultParameter{ID: 1},
				}).Return(res, nil)
			},
			doMockFindOneName: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					Name: settingEntity.SettingTypeDefault,
				}).Return(res, caseError)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {},
			doMock:         func(mock *settingMock.MockSettingRepo) {},
			doMockDelRedis: func(mock *redisMock.MockIRedis, wg *sync.WaitGroup) {
				defer wg.Done()
			},
			wantErr: caseError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRedis := redisMock.NewMockIRedis(mockCtrl)
			mockSettingRepo := settingMock.NewMockSettingRepo(mockCtrl)
			settingService := settingService.NewSettingService(mockSettingRepo, mockRedis)
			tt.doMockFindOne(mockSettingRepo, tt.WantFindOne)
			tt.doMockFindOneName(mockSettingRepo, tt.WantFindOneName)
			tt.doMockCreateTx(mockSettingRepo)
			tt.doMock(mockSettingRepo)

			wg := new(sync.WaitGroup)
			wg.Add(1)
			tt.doMockDelRedis(mockRedis, wg)

			err := settingService.Update(ctx, &tt.args)
			if err != tt.wantErr {
				t.Errorf("settingService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			wg.Wait()

			alwaysEqual := cmp.Comparer(func(_, _ interface{}) bool { return true })
			opt := cmp.FilterValues(func(x, y interface{}) bool {
				vx, vy := reflect.ValueOf(x), reflect.ValueOf(y)
				return (vx.IsValid() && vy.IsValid() && vx.Type() == vy.Type()) &&
					(vx.Kind() == reflect.Slice || vx.Kind() == reflect.Map) &&
					(vx.Len() == 0 && vy.Len() == 0)
			}, alwaysEqual)
			if !cmp.Equal(nil, nil, opt) {
				t.Errorf("settingService.Update()  = %v, want %v", nil, nil)
			}
		})
	}
}

func TestSettingService_Delete(t *testing.T) {
	ctx := context.Background()
	db, mocks, _ := sqlxmock.Newx()
	mocks.ExpectBegin()
	tx, _ := db.Beginx()
	settingData := settingEntity.Setting{
		ID:   1,
		Name: settingEntity.SettingTypeDefault,
	}
	tests := []struct {
		name           string
		args           int
		doMockCreateTx func(mock *settingMock.MockSettingRepo)
		doMockFindOne  func(mock *settingMock.MockSettingRepo, res settingEntity.Setting)
		doMock         func(mock *settingMock.MockSettingRepo)
		doMockRedisDel func(mock *redisMock.MockIRedis, wg *sync.WaitGroup)
		WantFindOne    settingEntity.Setting
		wantErr        error
	}{
		{
			name: "flow success",
			args: 1,
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					DefaultParameter: controllers.DefaultParameter{ID: 1},
				}).Return(settingData, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {
				mock.EXPECT().CreateTx(ctx).Return(tx, nil)
			},
			doMock: func(mock *settingMock.MockSettingRepo) {
				mock.EXPECT().Delete(gomock.Any(), tx, 1).Return(nil)
			},
			doMockRedisDel: func(mock *redisMock.MockIRedis, wg *sync.WaitGroup) {
				mock.EXPECT().Del(context.Background(), redis.Setting+settingData.Name).Do(func(arg0, arg1 interface{}) {
					defer wg.Done()
				})
			},
			WantFindOne: settingData,
			wantErr:     nil,
		},
		{
			name: "flow error",
			args: 1,
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					DefaultParameter: controllers.DefaultParameter{ID: 1},
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {
				mock.EXPECT().CreateTx(ctx).Return(tx, nil)
			},
			doMock: func(mock *settingMock.MockSettingRepo) {
				mock.EXPECT().Delete(gomock.Any(), tx, 1).Return(caseError)
			},
			doMockRedisDel: func(mock *redisMock.MockIRedis, wg *sync.WaitGroup) {
				defer wg.Done()
			},
			wantErr: caseError,
		},
		{
			name: "flow error create tx",
			args: 1,
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					DefaultParameter: controllers.DefaultParameter{ID: 1},
				}).Return(res, nil)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {
				mock.EXPECT().CreateTx(ctx).Return(tx, caseError)
			},
			doMock: func(mock *settingMock.MockSettingRepo) {},
			doMockRedisDel: func(mock *redisMock.MockIRedis, wg *sync.WaitGroup) {
				defer wg.Done()
			},
			wantErr: caseError,
		},
		{
			name: "flow error find one",
			args: 1,
			doMockFindOne: func(mock *settingMock.MockSettingRepo, res settingEntity.Setting) {
				mock.EXPECT().FindOne(gomock.Any(), &settingController.SettingParameter{
					DefaultParameter: controllers.DefaultParameter{ID: 1},
				}).Return(res, caseError)
			},
			doMockCreateTx: func(mock *settingMock.MockSettingRepo) {},
			doMock:         func(mock *settingMock.MockSettingRepo) {},
			doMockRedisDel: func(mock *redisMock.MockIRedis, wg *sync.WaitGroup) {
				defer wg.Done()
			},
			wantErr: errHelper.DataNotFound.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRedis := redisMock.NewMockIRedis(mockCtrl)
			mockSettingRepo := settingMock.NewMockSettingRepo(mockCtrl)
			settingService := settingService.NewSettingService(mockSettingRepo, mockRedis)
			tt.doMockFindOne(mockSettingRepo, tt.WantFindOne)
			tt.doMockCreateTx(mockSettingRepo)
			tt.doMock(mockSettingRepo)

			wg := new(sync.WaitGroup)
			wg.Add(1)
			tt.doMockRedisDel(mockRedis, wg)

			err := settingService.Delete(ctx, tt.args)
			if err != tt.wantErr {
				t.Errorf("settingService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			wg.Wait()

			alwaysEqual := cmp.Comparer(func(_, _ interface{}) bool { return true })
			opt := cmp.FilterValues(func(x, y interface{}) bool {
				vx, vy := reflect.ValueOf(x), reflect.ValueOf(y)
				return (vx.IsValid() && vy.IsValid() && vx.Type() == vy.Type()) &&
					(vx.Kind() == reflect.Slice || vx.Kind() == reflect.Map) &&
					(vx.Len() == 0 && vy.Len() == 0)
			}, alwaysEqual)
			if !cmp.Equal(nil, nil, opt) {
				t.Errorf("settingService.Delete()  = %v, want %v", nil, nil)
			}
		})
	}
}
