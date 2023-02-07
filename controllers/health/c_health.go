package health

import (
	"context"
	"errors"
	"go-template/logger"
	"go-template/response"
	"net/http"

	"github.com/gorilla/mux"
)

type HealthController struct {
	HealthService HealthServiceInterface
}

type HealthServiceInterface interface {
	CheckHealthDB(context.Context) (err error)
}

func NewHealthController(healthService HealthServiceInterface) HealthController {
	return HealthController{HealthService: healthService}
}

func (ctrl HealthController) InitializeRoutes(router *mux.Router) {
	router.HandleFunc("", ctrl.CheckHealthDB).Methods(http.MethodGet)
}

func (ctrl HealthController) CheckHealthDB(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	err := ctrl.HealthService.CheckHealthDB(ctx)
	if err != nil {
		logger.Error(ctx, "FAILED_CONNECT_TO_DB")
		response.RespondError(w, http.StatusInternalServerError, errors.New("FAILED_CONNECT_TO_DB"))

		return
	}

	response.RespondSuccess(w, http.StatusOK, "DB_CONNECTED", nil)
}
