package health_test

import (
	"context"
	"errors"
	"testing"

	healthRepo "go-template/repositories/health"

	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

var (
	caseError = errors.New("error")
)

func TestHealthRespository_CheckDB(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Correlation-ID", "testing")

	db, _, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := healthRepo.NewHealthRepository(db)
	t.Log(repo.CheckDB(ctx))
}
