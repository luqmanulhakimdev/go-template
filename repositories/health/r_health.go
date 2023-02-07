package health

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type HealthRepository struct {
	DbMaster *sqlx.DB
	TxMaster *sqlx.Tx
}

func NewHealthRepository(dbMaster *sqlx.DB) *HealthRepository {
	return &HealthRepository{
		DbMaster: dbMaster,
	}
}

func (repo HealthRepository) CheckDB(ctx context.Context) (err error) {
	err = repo.DbMaster.Ping()
	return
}
