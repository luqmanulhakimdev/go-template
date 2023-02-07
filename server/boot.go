package server

import (
	"go-template/client/swagger"
	"go-template/config"
	"go-template/middlewares"

	// Controller
	healthController "go-template/controllers/health"

	schedulerController "go-template/controllers/scheduler"
	settingController "go-template/controllers/setting"

	// Repository
	healthRepository "go-template/repositories/health"
	settingRepository "go-template/repositories/setting"

	// Service
	healthService "go-template/services/health"
	schedulerService "go-template/services/scheduler"
	settingService "go-template/services/setting"

	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/integrations/nrgorilla"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

func InitApp(router *mux.Router, conf config.Config, unitTest bool) {
	// Setup dependency
	newRelicApp := config.SetupNewRelic(conf)
	config.MasterDB = config.SetupMasterDB(conf, newRelicApp, unitTest)
	config.SlaveDB = config.SetupSlaveDB(conf, newRelicApp, unitTest)

	redisClient := config.SetupRedis(conf)

	// swagger
	if conf.Env != "prod" {
		router.PathPrefix("/swagger/").Handler(swagger.WrapHandler)
	}

	// v1 api
	v1 := router.PathPrefix("/v1").Subrouter()

	// v1 api group
	apiUser := v1.PathPrefix("/apiUser").Subrouter()           // aplus user
	apiAdmin := v1.PathPrefix("/apiAdmin").Subrouter()         // aplus admin user
	apiStatic := v1.PathPrefix("/apiStatic").Subrouter()       // static token
	apiScheduler := v1.PathPrefix("/apiScheduler").Subrouter() // scheduler

	// group api middleware setup
	apiUser.Use(middlewares.VerifyAPIToken(conf.APITokenKey))
	apiAdmin.Use(middlewares.VerifyAPIToken(conf.APITokenKey))
	apiStatic.Use(middlewares.VerifyAPIToken(conf.APITokenKey))
	apiScheduler.Use(middlewares.VerifyAPIToken(conf.APITokenKey))

	// setting up setting feature
	settingRepo := settingRepository.NewSettingRepository(config.MasterDB, config.SlaveDB)
	settingRouter := "/setting"
	settingUserRouter := apiUser.PathPrefix(settingRouter).Subrouter()
	settingAdminRouter := apiAdmin.PathPrefix(settingRouter).Subrouter()
	settingStaticRouter := apiStatic.PathPrefix(settingRouter).Subrouter()
	settingServices := settingService.NewSettingService(settingRepo, redisClient)
	settingController := settingController.NewSettingController(settingServices)
	settingController.InitializeRoutes(settingUserRouter, settingAdminRouter, settingStaticRouter)

	// setting up scheduler service
	schedulerServices := schedulerService.NewSchedulerService()
	schedulerController := schedulerController.NewSchedulerController(schedulerServices)
	schedulerController.InitializeRoutes(apiScheduler)

	//register-route
	healthRouter := router.PathPrefix("/health").Subrouter()
	healthRepo := healthRepository.NewHealthRepository(config.MasterDB)
	healthServices := healthService.NewHealthService(healthRepo)
	healthController := healthController.NewHealthController(healthServices)
	healthController.InitializeRoutes(healthRouter)

	if newRelicApp != nil {
		router.Use(nrgorilla.Middleware(newRelicApp))
		_, router.NotFoundHandler = newrelic.WrapHandle(newRelicApp, "NotFoundHandler", NotFoundHandler())
		_, router.MethodNotAllowedHandler = newrelic.WrapHandle(newRelicApp, "MethodNotAllowedHandler", MethodNotAllowedHandler())
	}
}
