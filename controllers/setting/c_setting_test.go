package setting_test

import (
	"errors"
	"go-template/controllers/setting"
	errHelper "go-template/helpers/errors"
	"go-template/services"

	settingEntity "go-template/entities/setting"
	"go-template/middlewares"
	mocks "go-template/mocks/setting"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

type args struct {
	url     string
	reqBody string
}

func TestSettingController_SelectAll(t *testing.T) {
	settingData := []settingEntity.Setting{
		{
			ID:        1,
			Name:      "setting_default",
			Value:     "{\"min_balance\": 5000000, \"limit_transaction\": 20}",
			CreatedAt: "2021-12-01T13:19:12.801+07:00",
			UpdatedAt: "2021-12-01T15:14:38.09019+07:00",
		},
	}
	tests := []struct {
		name     string
		args     args
		doMock   func(mock *mocks.MockSettingService)
		wantRes  string
		wantCode int
	}{
		{
			name: "flow success",
			args: args{
				url: "/v1/apiStatic/setting/all",
			},
			doMock: func(mock *mocks.MockSettingService) {
				mock.EXPECT().SelectAll(gomock.Any(), &setting.SettingParameter{}).
					Return(settingData, nil)
			},
			wantRes:  `{"success":true,"data":[{"id":1,"name":"setting_default","value":{"limit_transaction":20,"min_balance":5000000},"created_at":"2021-12-01T13:19:12.801+07:00","updated_at":"2021-12-01T15:14:38.09019+07:00"}]}`,
			wantCode: http.StatusOK,
		},
		{
			name: "flow error",
			args: args{
				url: "/v1/apiStatic/setting/all",
			},
			doMock: func(mock *mocks.MockSettingService) {
				mock.EXPECT().SelectAll(gomock.Any(), &setting.SettingParameter{}).
					Return([]settingEntity.Setting{}, errors.New("error"))
			},
			wantRes:  `{"success":false,"data":null,"stat_msg":"error"}`,
			wantCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, tt.args.url, strings.NewReader(tt.args.reqBody))
			r.Header = http.Header{"X-Correlation-Id": []string{"testing"}}
			r.Header.Set("Authorization", "secret")
			w := httptest.NewRecorder()

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockService := mocks.NewMockSettingService(mockCtrl)
			settingController := setting.NewSettingController(mockService)

			tt.doMock(mockService)

			router := mux.NewRouter()
			v1 := router.PathPrefix("/v1").Subrouter()
			apiUser := v1.PathPrefix("/apiUser").Subrouter()
			apiAdmin := v1.PathPrefix("/apiAdmin").Subrouter()
			apiStatic := v1.PathPrefix("/apiStatic").Subrouter()
			apiUser.Use(middlewares.VerifyAPIToken("secret"))
			apiAdmin.Use(middlewares.VerifyAPIToken("secret"))
			apiStatic.Use(middlewares.VerifyAPIToken("secret"))
			settingUserRouter := apiUser.PathPrefix("/setting").Subrouter()
			settingAdminRouter := apiAdmin.PathPrefix("/setting").Subrouter()
			settingStaticRouter := apiStatic.PathPrefix("/setting").Subrouter()
			settingController.InitializeRoutes(settingUserRouter, settingAdminRouter, settingStaticRouter)
			router.ServeHTTP(w, r)
			if w.Code != tt.wantCode {
				t.Fatalf("invalid status code. got %v, want %v", w.Code, tt.wantCode)
			}
			body := strings.Trim(w.Body.String(), "\n")
			if body != tt.wantRes {
				t.Fatalf("different response body, got %v want %v", body, tt.wantRes)
			}
		})
	}
}

func TestSettingController_FindAll(t *testing.T) {
	settingData := []settingEntity.Setting{
		{
			ID:        1,
			Name:      "setting_default",
			Value:     "{\"min_balance\": 5000000, \"limit_transaction\": 20}",
			CreatedAt: "2021-12-01T13:19:12.801+07:00",
			UpdatedAt: "2021-12-01T15:14:38.09019+07:00",
		},
	}
	pagination := services.Pagination{
		CurrentPage: 1,
		LastPage:    1,
		Total:       1,
		PerPage:     10,
	}
	tests := []struct {
		name     string
		args     args
		doMock   func(mock *mocks.MockSettingService)
		wantRes  string
		wantCode int
	}{
		{
			name: "flow success",
			args: args{
				url: "/v1/apiStatic/setting",
			},
			doMock: func(mock *mocks.MockSettingService) {
				mock.EXPECT().FindAll(gomock.Any(), &setting.SettingParameter{}).
					Return(settingData, pagination, nil)
			},
			wantRes:  `{"success":true,"data":[{"id":1,"name":"setting_default","value":{"limit_transaction":20,"min_balance":5000000},"created_at":"2021-12-01T13:19:12.801+07:00","updated_at":"2021-12-01T15:14:38.09019+07:00"}],"meta":{"current_page":1,"last_page":1,"total":1,"per_page":10}}`,
			wantCode: http.StatusOK,
		},
		{
			name: "flow error",
			args: args{
				url: "/v1/apiStatic/setting",
			},
			doMock: func(mock *mocks.MockSettingService) {
				mock.EXPECT().FindAll(gomock.Any(), &setting.SettingParameter{}).
					Return([]settingEntity.Setting{}, services.Pagination{}, errors.New("error"))
			},
			wantRes:  `{"success":false,"data":null,"stat_msg":"error"}`,
			wantCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, tt.args.url, strings.NewReader(tt.args.reqBody))
			r.Header = http.Header{"X-Correlation-Id": []string{"testing"}}
			r.Header.Set("Authorization", "secret")
			w := httptest.NewRecorder()

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockService := mocks.NewMockSettingService(mockCtrl)
			settingController := setting.NewSettingController(mockService)

			tt.doMock(mockService)

			router := mux.NewRouter()
			v1 := router.PathPrefix("/v1").Subrouter()
			apiUser := v1.PathPrefix("/apiUser").Subrouter()
			apiAdmin := v1.PathPrefix("/apiAdmin").Subrouter()
			apiStatic := v1.PathPrefix("/apiStatic").Subrouter()
			apiUser.Use(middlewares.VerifyAPIToken("secret"))
			apiAdmin.Use(middlewares.VerifyAPIToken("secret"))
			apiStatic.Use(middlewares.VerifyAPIToken("secret"))
			settingUserRouter := apiUser.PathPrefix("/setting").Subrouter()
			settingAdminRouter := apiAdmin.PathPrefix("/setting").Subrouter()
			settingStaticRouter := apiStatic.PathPrefix("/setting").Subrouter()
			settingController.InitializeRoutes(settingUserRouter, settingAdminRouter, settingStaticRouter)
			router.ServeHTTP(w, r)
			if w.Code != tt.wantCode {
				t.Fatalf("invalid status code. got %v, want %v", w.Code, tt.wantCode)
			}
			body := strings.Trim(w.Body.String(), "\n")
			if body != tt.wantRes {
				t.Fatalf("different response body, got %v want %v", body, tt.wantRes)
			}
		})
	}
}

func TestSettingController_FindOne(t *testing.T) {
	settingData := settingEntity.Setting{
		ID:        1,
		Name:      "setting_default",
		Value:     "{\"min_balance\": 5000000, \"limit_transaction\": 20}",
		CreatedAt: "2021-12-01T13:19:12.801+07:00",
		UpdatedAt: "2021-12-01T15:14:38.09019+07:00",
	}
	tests := []struct {
		name     string
		args     args
		doMock   func(mock *mocks.MockSettingService)
		wantRes  string
		wantCode int
	}{
		{
			name: "flow success",
			args: args{
				url: "/v1/apiStatic/setting/one",
			},
			doMock: func(mock *mocks.MockSettingService) {
				mock.EXPECT().FindOne(gomock.Any(), &setting.SettingParameter{}).
					Return(settingData, nil)
			},
			wantRes:  `{"success":true,"data":{"id":1,"name":"setting_default","value":{"limit_transaction":20,"min_balance":5000000},"created_at":"2021-12-01T13:19:12.801+07:00","updated_at":"2021-12-01T15:14:38.09019+07:00"}}`,
			wantCode: http.StatusOK,
		},
		{
			name: "flow error",
			args: args{
				url: "/v1/apiStatic/setting/one",
			},
			doMock: func(mock *mocks.MockSettingService) {
				mock.EXPECT().FindOne(gomock.Any(), &setting.SettingParameter{}).
					Return(settingEntity.Setting{}, errors.New("error"))
			},
			wantRes:  `{"success":false,"data":null,"stat_msg":"error"}`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "flow error data not found",
			args: args{
				url: "/v1/apiStatic/setting/one",
			},
			doMock: func(mock *mocks.MockSettingService) {
				mock.EXPECT().FindOne(gomock.Any(), &setting.SettingParameter{}).
					Return(settingEntity.Setting{}, errHelper.DataNotFound.Error)
			},
			wantRes:  `{"success":false,"data":null,"stat_msg":"data_not_found"}`,
			wantCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, tt.args.url, strings.NewReader(tt.args.reqBody))
			r.Header = http.Header{"X-Correlation-Id": []string{"testing"}}
			r.Header.Set("Authorization", "secret")
			w := httptest.NewRecorder()

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockService := mocks.NewMockSettingService(mockCtrl)
			settingController := setting.NewSettingController(mockService)

			tt.doMock(mockService)

			router := mux.NewRouter()
			v1 := router.PathPrefix("/v1").Subrouter()
			apiUser := v1.PathPrefix("/apiUser").Subrouter()
			apiAdmin := v1.PathPrefix("/apiAdmin").Subrouter()
			apiStatic := v1.PathPrefix("/apiStatic").Subrouter()
			apiUser.Use(middlewares.VerifyAPIToken("secret"))
			apiAdmin.Use(middlewares.VerifyAPIToken("secret"))
			apiStatic.Use(middlewares.VerifyAPIToken("secret"))
			settingUserRouter := apiUser.PathPrefix("/setting").Subrouter()
			settingAdminRouter := apiAdmin.PathPrefix("/setting").Subrouter()
			settingStaticRouter := apiStatic.PathPrefix("/setting").Subrouter()
			settingController.InitializeRoutes(settingUserRouter, settingAdminRouter, settingStaticRouter)
			router.ServeHTTP(w, r)
			if w.Code != tt.wantCode {
				t.Fatalf("invalid status code. got %v, want %v", w.Code, tt.wantCode)
			}
			body := strings.Trim(w.Body.String(), "\n")
			if body != tt.wantRes {
				t.Fatalf("different response body, got %v want %v", body, tt.wantRes)
			}
		})
	}
}

func TestSettingController_Create(t *testing.T) {
	settingData := settingEntity.Setting{
		Name:  "setting_default",
		Value: "{\"limit_transaction\":20,\"min_balance\":5000000}",
	}
	tests := []struct {
		name     string
		args     args
		doMock   func(mock *mocks.MockSettingService)
		wantRes  string
		wantCode int
	}{
		{
			name: "flow success",
			args: args{
				reqBody: `
				{
				    "name": "setting_default",
				    "value": {"min_balance": 5000000, "limit_transaction": 20}
				}`,
			},
			doMock: func(mock *mocks.MockSettingService) {
				mock.EXPECT().Create(gomock.Any(), &settingData).Return(1, nil)
			},
			wantRes:  `{"success":true,"data":1}`,
			wantCode: http.StatusOK,
		},
		{
			name: "flow error service",
			args: args{
				reqBody: `
				{
				    "name": "setting_default",
				    "value": {"min_balance": 5000000, "limit_transaction": 20}
				}`,
			},
			doMock: func(mock *mocks.MockSettingService) {
				mock.EXPECT().Create(gomock.Any(), &settingData).Return(0, errors.New("error"))
			},
			wantRes:  `{"success":false,"data":null,"stat_msg":"error"}`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "flow error json decode",
			args: args{
				reqBody: `
				{
				    "name": setting_default,
				    "value": {"min_balance": 5000000, "limit_transaction": 20}
				}`,
			},
			doMock:   func(mock *mocks.MockSettingService) {},
			wantRes:  `{"success":false,"data":null,"stat_msg":"invalid character 's' looking for beginning of value"}`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "flow error validate",
			args: args{
				reqBody: `
				{
				    "name": "",
				    "value": {"min_balance": 5000000, "limit_transaction": 20}
				}`,
			},
			doMock:   func(mock *mocks.MockSettingService) {},
			wantRes:  `{"success":false,"data":null,"stat_msg":"Name required"}`,
			wantCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, "/v1/apiStatic/setting", strings.NewReader(tt.args.reqBody))
			r.Header = http.Header{"X-Correlation-Id": []string{"testing"}}
			r.Header.Set("Authorization", "secret")
			w := httptest.NewRecorder()

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockService := mocks.NewMockSettingService(mockCtrl)
			settingController := setting.NewSettingController(mockService)

			tt.doMock(mockService)

			router := mux.NewRouter()
			v1 := router.PathPrefix("/v1").Subrouter()
			apiUser := v1.PathPrefix("/apiUser").Subrouter()
			apiAdmin := v1.PathPrefix("/apiAdmin").Subrouter()
			apiStatic := v1.PathPrefix("/apiStatic").Subrouter()
			apiUser.Use(middlewares.VerifyAPIToken("secret"))
			apiAdmin.Use(middlewares.VerifyAPIToken("secret"))
			apiStatic.Use(middlewares.VerifyAPIToken("secret"))
			settingUserRouter := apiUser.PathPrefix("/setting").Subrouter()
			settingAdminRouter := apiAdmin.PathPrefix("/setting").Subrouter()
			settingStaticRouter := apiStatic.PathPrefix("/setting").Subrouter()
			settingController.InitializeRoutes(settingUserRouter, settingAdminRouter, settingStaticRouter)
			router.ServeHTTP(w, r)
			if w.Code != tt.wantCode {
				t.Fatalf("invalid status code. got %v, want %v", w.Code, tt.wantCode)
			}
			body := strings.Trim(w.Body.String(), "\n")
			if body != tt.wantRes {
				t.Fatalf("different response body, got %v want %v", body, tt.wantRes)
			}
		})
	}
}

func TestSettingController_Update(t *testing.T) {
	settingData := settingEntity.Setting{
		ID:    1,
		Name:  "setting_default",
		Value: "{\"limit_transaction\":20,\"min_balance\":5000000}",
	}
	tests := []struct {
		name     string
		args     args
		doMock   func(mock *mocks.MockSettingService)
		wantRes  string
		wantCode int
	}{
		{
			name: "flow success",
			args: args{
				reqBody: `
				{
				    "name": "setting_default",
				    "value": {"min_balance": 5000000, "limit_transaction": 20}
				}`,
			},
			doMock: func(mock *mocks.MockSettingService) {
				mock.EXPECT().Update(gomock.Any(), &settingData).Return(nil)
			},
			wantRes:  `{"success":true,"data":null}`,
			wantCode: http.StatusOK,
		},
		{
			name: "flow error service",
			args: args{
				reqBody: `
				{
				    "name": "setting_default",
				    "value": {"min_balance": 5000000, "limit_transaction": 20}
				}`,
			},
			doMock: func(mock *mocks.MockSettingService) {
				mock.EXPECT().Update(gomock.Any(), &settingData).Return(errors.New("error"))
			},
			wantRes:  `{"success":false,"data":null,"stat_msg":"error"}`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "flow error json decode",
			args: args{
				reqBody: `
				{
				    "name": setting_default,
				    "value": {"min_balance": 5000000, "limit_transaction": 20}
				}`,
			},
			doMock:   func(mock *mocks.MockSettingService) {},
			wantRes:  `{"success":false,"data":null,"stat_msg":"invalid character 's' looking for beginning of value"}`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "flow error validate",
			args: args{
				reqBody: `
				{
				    "name": "",
				    "value": {"min_balance": 5000000, "limit_transaction": 20}
				}`,
			},
			doMock:   func(mock *mocks.MockSettingService) {},
			wantRes:  `{"success":false,"data":null,"stat_msg":"Name required"}`,
			wantCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPut, "/v1/apiStatic/setting/id/1", strings.NewReader(tt.args.reqBody))
			r.Header = http.Header{"X-Correlation-Id": []string{"testing"}}
			r.Header.Set("Authorization", "secret")
			w := httptest.NewRecorder()

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mocks.NewMockSettingService(mockCtrl)
			settingController := setting.NewSettingController(mockService)

			tt.doMock(mockService)

			router := mux.NewRouter()
			v1 := router.PathPrefix("/v1").Subrouter()
			apiUser := v1.PathPrefix("/apiUser").Subrouter()
			apiAdmin := v1.PathPrefix("/apiAdmin").Subrouter()
			apiStatic := v1.PathPrefix("/apiStatic").Subrouter()
			apiUser.Use(middlewares.VerifyAPIToken("secret"))
			apiAdmin.Use(middlewares.VerifyAPIToken("secret"))
			apiStatic.Use(middlewares.VerifyAPIToken("secret"))
			settingUserRouter := apiUser.PathPrefix("/setting").Subrouter()
			settingAdminRouter := apiAdmin.PathPrefix("/setting").Subrouter()
			settingStaticRouter := apiStatic.PathPrefix("/setting").Subrouter()
			settingController.InitializeRoutes(settingUserRouter, settingAdminRouter, settingStaticRouter)
			router.ServeHTTP(w, r)
			if w.Code != tt.wantCode {
				t.Fatalf("invalid status code. got %v, want %v", w.Code, tt.wantCode)
			}
			body := strings.Trim(w.Body.String(), "\n")
			if body != tt.wantRes {
				t.Fatalf("different response body, got %v want %v", body, tt.wantRes)
			}
		})
	}
}

func TestSettingController_Delete(t *testing.T) {
	tests := []struct {
		name     string
		args     args
		doMock   func(mock *mocks.MockSettingService)
		wantRes  string
		wantCode int
	}{
		{
			name: "flow success",
			doMock: func(mock *mocks.MockSettingService) {
				mock.EXPECT().Delete(gomock.Any(), 1).Return(nil)
			},
			wantRes:  `{"success":true,"data":null}`,
			wantCode: http.StatusOK,
		},
		{
			name: "flow error service",
			doMock: func(mock *mocks.MockSettingService) {
				mock.EXPECT().Delete(gomock.Any(), 1).Return(errors.New("error"))
			},
			wantRes:  `{"success":false,"data":null,"stat_msg":"error"}`,
			wantCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodDelete, "/v1/apiStatic/setting/id/1", strings.NewReader(tt.args.reqBody))
			r.Header = http.Header{"X-Correlation-Id": []string{"testing"}}
			r.Header.Set("Authorization", "secret")
			w := httptest.NewRecorder()

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockService := mocks.NewMockSettingService(mockCtrl)
			settingController := setting.NewSettingController(mockService)

			tt.doMock(mockService)

			router := mux.NewRouter()
			v1 := router.PathPrefix("/v1").Subrouter()
			apiUser := v1.PathPrefix("/apiUser").Subrouter()
			apiAdmin := v1.PathPrefix("/apiAdmin").Subrouter()
			apiStatic := v1.PathPrefix("/apiStatic").Subrouter()
			apiUser.Use(middlewares.VerifyAPIToken("secret"))
			apiAdmin.Use(middlewares.VerifyAPIToken("secret"))
			apiStatic.Use(middlewares.VerifyAPIToken("secret"))
			settingUserRouter := apiUser.PathPrefix("/setting").Subrouter()
			settingAdminRouter := apiAdmin.PathPrefix("/setting").Subrouter()
			settingStaticRouter := apiStatic.PathPrefix("/setting").Subrouter()
			settingController.InitializeRoutes(settingUserRouter, settingAdminRouter, settingStaticRouter)
			router.ServeHTTP(w, r)
			if w.Code != tt.wantCode {
				t.Fatalf("invalid status code. got %v, want %v", w.Code, tt.wantCode)
			}
			body := strings.Trim(w.Body.String(), "\n")
			if body != tt.wantRes {
				t.Fatalf("different response body, got %v want %v", body, tt.wantRes)
			}
		})
	}
}
