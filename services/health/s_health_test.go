package health_test

import (
	"context"
	"errors"

	healthMock "go-template/mocks/health"
	healthService "go-template/services/health"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

var (
	caseError = errors.New("error")
)

func TestHealthService_CheckHealthDB(t *testing.T) {
	tests := []struct {
		name    string
		doMock  func(mock *healthMock.MockHealthRepoInterface)
		wantErr error
	}{
		{
			name: "flow success",
			doMock: func(mock *healthMock.MockHealthRepoInterface) {
				mock.EXPECT().CheckDB(gomock.Any()).Return(nil)
			},
			wantErr: nil,
		},
		{
			name: "flow error",
			doMock: func(mock *healthMock.MockHealthRepoInterface) {
				mock.EXPECT().CheckDB(gomock.Any()).Return(caseError)
			},
			wantErr: caseError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockHealthRepo := healthMock.NewMockHealthRepoInterface(mockCtrl)
			healthService := healthService.NewHealthService(mockHealthRepo)
			tt.doMock(mockHealthRepo)

			err := healthService.CheckHealthDB(context.Background())
			if err != tt.wantErr {
				t.Errorf("healthService.SelectAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			alwaysEqual := cmp.Comparer(func(_, _ interface{}) bool { return true })
			opt := cmp.FilterValues(func(x, y interface{}) bool {
				vx, vy := reflect.ValueOf(x), reflect.ValueOf(y)
				return (vx.IsValid() && vy.IsValid() && vx.Type() == vy.Type()) &&
					(vx.Kind() == reflect.Slice || vx.Kind() == reflect.Map) &&
					(vx.Len() == 0 && vy.Len() == 0)
			}, alwaysEqual)
			if !cmp.Equal(nil, nil, opt) {
				t.Errorf("healthService.SelectAll()  = %v, want %v", nil, nil)
			}
		})
	}
}
