package setting_test

import (
	"context"
	"errors"
	"fmt"
	"go-template/controllers"
	settingController "go-template/controllers/setting"
	settingEntity "go-template/entities/setting"
	"reflect"
	"regexp"
	"testing"

	settingRepo "go-template/repositories/setting"

	"github.com/google/go-cmp/cmp"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

var (
	caseError = errors.New("error")
)

func TestSettingRespository_CreateTx(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Correlation-ID", "testing")

	db, _, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := settingRepo.NewSettingRepository(db, nil)
	t.Log(repo.CreateTx(ctx))
}

func TestSettingRepository_SelectAll(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Correlation-ID", "testing")

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
		name      string
		parameter settingController.SettingParameter
		expQuery  string
		wantRes   []settingEntity.Setting
		wantErr   error
	}{
		{
			name:      "flow success",
			parameter: settingController.SettingParameter{},
			expQuery:  settingEntity.SelectStatement + ` WHERE def.deleted_at IS NULL ` + settingEntity.GroupStatement,
			wantRes:   settingData,
		},
		{
			name:      "flow error",
			parameter: settingController.SettingParameter{},
			expQuery:  settingEntity.SelectStatement + ` WHERE def.deleted_at IS NULL ` + settingEntity.GroupStatement,
			wantRes:   []settingEntity.Setting{},
			wantErr:   caseError,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlxmock.Newx()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			if tc.wantErr == nil {
				rows := sqlxmock.NewRows(settingEntity.Column).AddRow(1, "setting_default", "{\"min_balance\": 5000000, \"limit_transaction\": 20}", "2021-12-01T13:19:12.801+07:00", "2021-12-01T15:14:38.09019+07:00")
				mock.ExpectQuery(regexp.QuoteMeta(tc.expQuery)).WithArgs().WillReturnRows(rows)
			} else {
				mock.ExpectQuery(regexp.QuoteMeta(tc.expQuery)).WithArgs().WillReturnError(tc.wantErr)
			}
			repo := settingRepo.NewSettingRepository(nil, db)
			result, err := repo.SelectAll(ctx, &tc.parameter)
			alwaysEqual := cmp.Comparer(func(_, _ interface{}) bool { return true })
			opt := cmp.FilterValues(func(x, y interface{}) bool {
				vx, vy := reflect.ValueOf(x), reflect.ValueOf(y)
				return (vx.IsValid() && vy.IsValid() && vx.Type() == vy.Type()) &&
					(vx.Kind() == reflect.Slice || vx.Kind() == reflect.Map) &&
					(vx.Len() == 0 && vy.Len() == 0)
			}, alwaysEqual)
			if !cmp.Equal(result, tc.wantRes, opt) {
				t.Fatalf("invalid result. got %v, want %v", result, tc.wantRes)
			}

			if err != tc.wantErr {
				fmt.Println(err.Error())
				t.Fatalf("invalid error. got %v, want %v", err, tc.wantErr)
			}
		})
	}
}

func TestSettingRepository_FindAll(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Correlation-ID", "testing")

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
		name          string
		parameter     settingController.SettingParameter
		expQuery      string
		expCountQuery string
		wantRes       []settingEntity.Setting
		wantErr       error
	}{
		{
			name:      "flow success",
			parameter: settingController.SettingParameter{},
			expQuery:  settingEntity.SelectStatement + ` WHERE def.deleted_at IS NULL ` + settingEntity.GroupStatement,
			wantRes:   settingData,
		},
		{
			name:      "flow error",
			parameter: settingController.SettingParameter{},
			expQuery:  settingEntity.SelectStatement + ` WHERE def.deleted_at IS NULL ` + settingEntity.GroupStatement,
			wantRes:   []settingEntity.Setting{},
			wantErr:   caseError,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlxmock.Newx()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			if tc.wantErr == nil {
				rows := sqlxmock.NewRows(settingEntity.Column).AddRow(1, "setting_default", "{\"min_balance\": 5000000, \"limit_transaction\": 20}", "2021-12-01T13:19:12.801+07:00", "2021-12-01T15:14:38.09019+07:00")
				mock.ExpectQuery(regexp.QuoteMeta(tc.expQuery)).WithArgs().WillReturnRows(rows)
				rowsCount := sqlxmock.NewRows([]string{"count"}).AddRow(1)
				mock.ExpectQuery(regexp.QuoteMeta(tc.expCountQuery)).WithArgs().WillReturnRows(rowsCount)
			} else {
				mock.ExpectQuery(regexp.QuoteMeta(tc.expQuery)).WithArgs().WillReturnError(tc.wantErr)
			}
			repo := settingRepo.NewSettingRepository(nil, db)
			result, _, err := repo.FindAll(ctx, &tc.parameter)
			alwaysEqual := cmp.Comparer(func(_, _ interface{}) bool { return true })
			opt := cmp.FilterValues(func(x, y interface{}) bool {
				vx, vy := reflect.ValueOf(x), reflect.ValueOf(y)
				return (vx.IsValid() && vy.IsValid() && vx.Type() == vy.Type()) &&
					(vx.Kind() == reflect.Slice || vx.Kind() == reflect.Map) &&
					(vx.Len() == 0 && vy.Len() == 0)
			}, alwaysEqual)
			if !cmp.Equal(result, tc.wantRes, opt) {
				t.Fatalf("invalid result. got %v, want %v", result, tc.wantRes)
			}

			if err != tc.wantErr {
				fmt.Println(err.Error())
				t.Fatalf("invalid error. got %v, want %v", err, tc.wantErr)
			}
		})
	}
}

func TestSettingRepository_FindOne(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Correlation-ID", "testing")

	settingData := settingEntity.Setting{
		ID:        1,
		Name:      "setting_default",
		Value:     "{\"min_balance\": 5000000, \"limit_transaction\": 20}",
		CreatedAt: "2021-12-01T13:19:12.801+07:00",
		UpdatedAt: "2021-12-01T15:14:38.09019+07:00",
	}

	tests := []struct {
		name      string
		parameter settingController.SettingParameter
		expQuery  string
		wantRes   settingEntity.Setting
		wantErr   error
	}{
		{
			name: "flow success",
			parameter: settingController.SettingParameter{
				DefaultParameter: controllers.DefaultParameter{
					ID: 1,
				},
			},
			expQuery: settingEntity.SelectStatement + ` WHERE def.deleted_at IS NULL AND def.id = $1`,
			wantRes:  settingData,
		},
		{
			name: "flow error",
			parameter: settingController.SettingParameter{
				Name: "setting_default",
				DefaultParameter: controllers.DefaultParameter{
					ID: 1,
				},
			},
			expQuery: settingEntity.SelectStatement + ` WHERE def.deleted_at IS NULL AND def.id = $1 AND def.name = $2 `,
			wantRes:  settingEntity.Setting{},
			wantErr:  caseError,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlxmock.Newx()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			if tc.wantErr == nil {
				rows := sqlxmock.NewRows(settingEntity.Column).AddRow(1, "setting_default", "{\"min_balance\": 5000000, \"limit_transaction\": 20}", "2021-12-01T13:19:12.801+07:00", "2021-12-01T15:14:38.09019+07:00")
				mock.ExpectQuery(regexp.QuoteMeta(tc.expQuery)).WithArgs().WillReturnRows(rows)
			} else {
				mock.ExpectQuery(regexp.QuoteMeta(tc.expQuery)).WithArgs().WillReturnError(tc.wantErr)
			}
			repo := settingRepo.NewSettingRepository(nil, db)
			result, err := repo.FindOne(ctx, &tc.parameter)
			if !cmp.Equal(result, tc.wantRes) {
				t.Fatalf("invalid result. got %v, want %v", result, tc.wantRes)
			}

			if err != tc.wantErr {
				t.Fatalf("invalid error. got %v, want %v", err, tc.wantErr)
			}
		})
	}
}

func TestSettingRepository_Create(t *testing.T) {
	ctx := context.Background()

	db, mock, _ := sqlxmock.Newx()
	mock.ExpectBegin()
	tx, _ := db.Beginx()

	settingData := settingEntity.Setting{
		Name:  "setting_default",
		Value: "{\"min_balance\": 5000000, \"limit_transaction\": 20}",
	}

	tests := []struct {
		name    string
		expExec string
		body    settingEntity.Setting
		wantErr error
	}{
		{
			name:    "flow success",
			expExec: `INSERT INTO ` + settingEntity.TableName + ` (name, value) VALUES ($1, $2) RETURNING id`,
			body:    settingData,
			wantErr: nil,
		},
		{
			name:    "flow error",
			expExec: `INSERT INTO ` + settingEntity.TableName + ` (name, value) VALUES ($1, $2) RETURNING id`,
			body:    settingData,
			wantErr: caseError,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			if tc.wantErr == nil {
				row := sqlxmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery(regexp.QuoteMeta(tc.expExec)).WithArgs(tc.body.Name, tc.body.Value).WillReturnRows(row)
			} else {
				mock.ExpectQuery(regexp.QuoteMeta(tc.expExec)).WithArgs(tc.body.Name, tc.body.Value).WillReturnError(tc.wantErr)
			}

			data := settingEntity.Setting{
				Name:  tc.body.Name,
				Value: tc.body.Value,
			}

			repo := settingRepo.NewSettingRepository(db, nil)
			_, err := repo.Create(ctx, tx, &data)

			if e := mock.ExpectationsWereMet(); e != nil {
				t.Fatalf("there were unfulfilled expectations: %s", e)
			}

			if err != tc.wantErr {
				t.Fatalf("invalid error. got %v, want %v", err, tc.wantErr)
			}

		})
	}
}

func TestSettingRepository_Update(t *testing.T) {
	ctx := context.Background()
	settingData := settingEntity.Setting{
		ID:    1,
		Name:  "setting_default",
		Value: "{\"min_balance\": 5000000, \"limit_transaction\": 20}",
	}
	tests := []struct {
		name    string
		expExec string
		body    settingEntity.Setting
		wantErr error
	}{
		{
			name:    "flow success",
			expExec: `UPDATE ` + settingEntity.TableName + ` SET name = $1, value = $2 WHERE deleted_at IS NULL AND id = $3`,
			body:    settingData,
			wantErr: nil,
		},
		{
			name:    "flow error",
			expExec: `UPDATE ` + settingEntity.TableName + ` SET name = $1, value = $2 WHERE deleted_at IS NULL AND id = $3`,
			body:    settingData,
			wantErr: caseError,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			db, mock, err := sqlxmock.Newx()
			mock.ExpectBegin()
			tx, _ := db.Beginx()

			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			if tc.wantErr == nil {
				mock.ExpectExec(regexp.QuoteMeta(tc.expExec)).WithArgs(tc.body.Name, tc.body.Value, tc.body.ID).WillReturnResult(sqlxmock.NewResult(0, 1))
			} else {
				mock.ExpectExec(regexp.QuoteMeta(tc.expExec)).WithArgs(tc.body.Name, tc.body.Value, tc.body.ID).WillReturnError(tc.wantErr)
			}

			repo := settingRepo.NewSettingRepository(db, nil)
			err = repo.Update(ctx, tx, &tc.body)
			if e := mock.ExpectationsWereMet(); e != nil {
				t.Fatalf("there were unfulfilled expectations: %s", e)
			}

			if err != tc.wantErr {
				t.Fatalf("invalid error. got %v, want %v", err, tc.wantErr)
			}

		})
	}
}

func TestSettingRepository_Delete(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name    string
		expExec string
		body    settingEntity.Setting
		wantErr error
	}{
		{
			name:    "flow success",
			expExec: `UPDATE ` + settingEntity.TableName + ` SET deleted_at = NOW() WHERE deleted_at IS NULL AND id = $1`,
			wantErr: nil,
		},
		{
			name:    "flow error",
			expExec: `UPDATE ` + settingEntity.TableName + ` SET deleted_at = NOW() WHERE deleted_at IS NULL AND id = $1`,
			wantErr: caseError,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			db, mock, err := sqlxmock.Newx()
			mock.ExpectBegin()
			tx, _ := db.Beginx()

			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			if tc.wantErr == nil {
				mock.ExpectExec(regexp.QuoteMeta(tc.expExec)).WithArgs(1).WillReturnResult(sqlxmock.NewResult(0, 1))
			} else {
				mock.ExpectExec(regexp.QuoteMeta(tc.expExec)).WithArgs(1).WillReturnError(tc.wantErr)
			}

			repo := settingRepo.NewSettingRepository(db, nil)
			err = repo.Delete(ctx, tx, 1)
			if e := mock.ExpectationsWereMet(); e != nil {
				t.Fatalf("there were unfulfilled expectations: %s", e)
			}

			if err != tc.wantErr {
				t.Fatalf("invalid error. got %v, want %v", err, tc.wantErr)
			}

		})
	}
}
