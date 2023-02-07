package main

import (
	"os/exec"

	"go-template/client/swagger"
	"go-template/config"
	"go-template/logger"
	"go-template/server"

	log "github.com/sirupsen/logrus"
)

var (
	commit         = "N/A"
	build_datetime = "N/A"
)

// @securityDefinitions.apikey 	ApiKeyAuth
// @in 							Header
// @name 						Authorization
// @Description					For /apiUser Fill value with prefix "Bearer "+(user_token)
// @Description					For /apiAdmin Fill value with prefix "Bearer "+(admin_token)
// @Description					For /apiStatic Fill value with (static_token)
func main() {
	conf := config.New()
	logger.Init(conf.LogLevel, conf.LogFormat)

	if conf.Env == "local" || conf.Env == "" {
		// config githook for development only
		if err := exec.Command("git", "config", "core.hooksPath", ".githooks").Run(); err != nil {
			log.Fatal(err)
		}
	}

	swagger.SwaggerInfo(conf)
	srv := server.New()
	srv.Start(commit, build_datetime)
}
