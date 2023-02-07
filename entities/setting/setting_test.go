package setting_test

import (
	"context"
	"database/sql"
	settingEntity "go-template/entities/setting"
	"testing"

	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestScanRows(t *testing.T) {
	temp := settingEntity.Setting{}
	_ = temp.ScanRows(nil, &sql.Rows{})
}

func TestScanRow(t *testing.T) {
	ctx := context.Background()
	db, _, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	expQuery := settingEntity.SelectStatement + ` WHERE def.deleted_at IS NULL`
	row := db.QueryRowContext(ctx, expQuery)

	temp := settingEntity.Setting{}
	_ = temp.ScanRows(row, nil)
}
