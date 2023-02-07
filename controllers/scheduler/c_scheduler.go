package scheduler

import (
	"github.com/gorilla/mux"
)

type SchedulerService interface {
}

type schedulerController struct {
	schedulerService SchedulerService
}

func NewSchedulerController(schedulerService SchedulerService) schedulerController {
	return schedulerController{schedulerService: schedulerService}
}

func (ctrl schedulerController) InitializeRoutes(routerScheduler *mux.Router) {
}
