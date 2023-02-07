package setting

import (
	"context"
	settingRequest "go-template/controllers/setting"

	"database/sql"
	settingEntity "go-template/entities/setting"
	"go-template/helpers/postgres"

	"github.com/jmoiron/sqlx"
)

type settingRepository struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewSettingRepository(dbMaster, dbSlave *sqlx.DB) *settingRepository {
	return &settingRepository{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}

func (repo settingRepository) CreateTx(ctx context.Context) (tx *sqlx.Tx, err error) {
	return repo.DbMaster.BeginTxx(ctx, &sql.TxOptions{})
}

func (repo settingRepository) buildParameters(ctx context.Context, parameters *settingRequest.SettingParameter) (conditionString string, conditionParam []interface{}) {
	if parameters.ID != 0 {
		conditionString += ` AND def.id = ?`
		conditionParam = append(conditionParam, parameters.ID)
	}
	if parameters.Name != "" {
		conditionString += ` AND def.name = ?`
		conditionParam = append(conditionParam, parameters.Name)
	}

	return
}

func (repo settingRepository) SelectAll(ctx context.Context, parameters *settingRequest.SettingParameter) (data []settingEntity.Setting, err error) {
	whereStatement, conditionParam := repo.buildParameters(ctx, parameters)
	query := settingEntity.SelectStatement + ` WHERE def.deleted_at IS NULL ` + whereStatement + ` ` + settingEntity.GroupStatement +
		` ORDER BY ` + parameters.OrderBy + ` ` + parameters.Sort + `, def.id ` + parameters.Sort
	query = postgres.SubstitutePlaceholder(query, 1)
	rows, err := repo.DbSlave.QueryContext(ctx, query, conditionParam...)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		temp := settingEntity.Setting{}
		err = temp.ScanRows(nil, rows)
		if err != nil {
			return
		}
		data = append(data, temp)
	}

	err = rows.Err()
	if err != nil {
		return
	}

	return
}

func (repo settingRepository) FindAll(ctx context.Context, parameters *settingRequest.SettingParameter) (data []settingEntity.Setting, count int, err error) {
	whereStatement, conditionParam := repo.buildParameters(ctx, parameters)
	query := settingEntity.SelectStatement + ` WHERE def.deleted_at IS NULL ` + whereStatement + ` ` + settingEntity.GroupStatement +
		` ORDER BY ` + parameters.OrderBy + ` ` + parameters.Sort + `, def.id ` + parameters.Sort + ` OFFSET ? LIMIT ? `
	query = postgres.SubstitutePlaceholder(query, 1)

	findCondition := append(conditionParam, parameters.Offset, parameters.Limit)
	rows, err := repo.DbSlave.QueryContext(ctx, query, findCondition...)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		temp := settingEntity.Setting{}
		err = temp.ScanRows(nil, rows)
		if err != nil {
			return
		}
		data = append(data, temp)
	}

	err = rows.Err()
	if err != nil {
		return
	}

	query = settingEntity.SelectCountStatement + ` WHERE def.deleted_at IS NULL ` + whereStatement
	query = postgres.SubstitutePlaceholder(query, 1)
	err = repo.DbSlave.QueryRowContext(ctx, query, conditionParam...).Scan(&count)
	return
}

func (repo settingRepository) FindOne(ctx context.Context, parameters *settingRequest.SettingParameter) (data settingEntity.Setting, err error) {
	whereStatement, conditionParam := repo.buildParameters(ctx, parameters)
	if whereStatement == "" {
		return data, sql.ErrNoRows
	}

	query := settingEntity.SelectStatement + ` WHERE def.deleted_at IS NULL ` + whereStatement
	query = postgres.SubstitutePlaceholder(query, 1)
	row := repo.DbSlave.QueryRowContext(ctx, query, conditionParam...)
	err = data.ScanRows(row, nil)
	if err != nil {
		return
	}
	return
}

func (repo settingRepository) Create(ctx context.Context, tx *sqlx.Tx, data *settingEntity.Setting) (res int, err error) {
	query := `INSERT INTO ` + settingEntity.TableName + ` (name, value) VALUES ($1, $2) RETURNING id`
	err = tx.QueryRowContext(ctx, query, data.Name, data.Value).Scan(&res)
	if err != nil {
		return
	}

	return
}

func (repo settingRepository) Update(ctx context.Context, tx *sqlx.Tx, data *settingEntity.Setting) error {
	query := `UPDATE ` + settingEntity.TableName + ` SET name = $1, value = $2 WHERE deleted_at IS NULL AND id = $3`
	res, err := tx.ExecContext(ctx, query, data.Name, data.Value, data.ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count < 1 {
		return sql.ErrNoRows
	}

	return nil
}

func (repo settingRepository) Delete(ctx context.Context, tx *sqlx.Tx, id int) error {
	query := `UPDATE ` + settingEntity.TableName + ` SET deleted_at = NOW() WHERE
	deleted_at IS NULL AND id = $1`
	res, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count < 1 {
		return sql.ErrNoRows
	}

	return nil
}
