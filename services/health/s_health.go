package health

import (
	"context"
	healthCtrl "go-template/controllers/health"
)

type HealthService struct {
	healthRepo HealthRepoInterface
}

type HealthRepoInterface interface {
	CheckDB(context.Context) (err error)
}

func NewHealthService(healthRepo HealthRepoInterface) healthCtrl.HealthServiceInterface {
	return &HealthService{
		healthRepo: healthRepo,
	}
}

func (service *HealthService) CheckHealthDB(ctx context.Context) (err error) {
	return service.healthRepo.CheckDB(ctx)
}
