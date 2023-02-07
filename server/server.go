package server

import (
	"context"
	"fmt"
	"go-template/config"
	"go-template/middlewares"
	"strings"

	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	log "github.com/sirupsen/logrus"
)

var (
	hashCommit    = "N/A"
	buildDatetime = "N/A"
)

type RestServer struct{}

func New() *RestServer {
	return &RestServer{}
}

func (s *RestServer) Start(inputHashCommit, inputBuildDatetime string) {
	conf := config.New()
	ctx := context.Background()

	hashCommit = inputHashCommit
	buildDatetime = inputBuildDatetime

	corsDomainList := strings.Split(conf.AppCorsDomain, ",")

	router := mux.NewRouter()
	router.Use(middlewares.Logging(), middlewares.Recovery())

	InitApp(router, conf, false)

	c := cors.New(cors.Options{ // CORS setting
		AllowedOrigins: corsDomainList,
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	handler := c.Handler(router)

	// Setup http server
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Port),
		Handler: handler,
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		log.Info(ctx, "We received an interrupt signal, shut down.")
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Error(ctx, "HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
		log.Info(ctx, "Bye.")
	}()

	log.Info(ctx, "Listening on port %d", conf.Port)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(ctx, "%v", err)
	}
	<-idleConnsClosed
}
