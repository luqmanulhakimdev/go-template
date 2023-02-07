package health_test

import (
	"errors"
	healthController "go-template/controllers/health"

	mocks "go-template/mocks/health"
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

func TestHealthController_CheckHealthDB(t *testing.T) {
	tests := []struct {
		name     string
		args     args
		doMock   func(mock *mocks.MockHealthServiceInterface)
		wantRes  string
		wantCode int
	}{
		{
			name: "flow success",
			args: args{
				url: "/health",
			},
			doMock: func(mock *mocks.MockHealthServiceInterface) {
				mock.EXPECT().CheckHealthDB(gomock.Any()).Return(nil)
			},
			wantRes:  `{"success":true,"data":"DB_CONNECTED"}`,
			wantCode: http.StatusOK,
		},
		{
			name: "flow error",
			args: args{
				url: "/health",
			},
			doMock: func(mock *mocks.MockHealthServiceInterface) {
				mock.EXPECT().CheckHealthDB(gomock.Any()).Return(errors.New("FAILED_CONNECT_TO_DB"))
			},
			wantRes:  `{"success":false,"data":null,"stat_msg":"FAILED_CONNECT_TO_DB"}`,
			wantCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, tt.args.url, strings.NewReader(tt.args.reqBody))
			r.Header = http.Header{"X-Correlation-Id": []string{"testing"}}
			r.Header.Set("Authorization", "Bearer jwttoken")
			w := httptest.NewRecorder()

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockService := mocks.NewMockHealthServiceInterface(mockCtrl)
			healthController := healthController.NewHealthController(mockService)
			tt.doMock(mockService)

			router := mux.NewRouter()
			healthRouter := router.PathPrefix("/health").Subrouter()
			healthController.InitializeRoutes(healthRouter)
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
