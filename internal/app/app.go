package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/reversersed/taskservice/internal/config"
	"github.com/reversersed/taskservice/internal/repository"
	"github.com/reversersed/taskservice/pkg/logging/logrus"
	"github.com/reversersed/taskservice/pkg/middleware"
	"github.com/reversersed/taskservice/pkg/postgres"
	"github.com/reversersed/taskservice/pkg/shutdown"
)

// @title API
// @version 1.0

// @host localhost:9000
// @BasePath /

// @scheme http
// @accept json
func New() (*app, error) {
	log, err := logrus.GetLogger()
	if err != nil {
		return nil, err
	}

	log.Info("setting up config...")
	cfg, err := config.Load("config/.env")
	if err != nil {
		return nil, err
	}
	log.Infof("setting up router with %s environment...", cfg.Server.Environment)
	gin.SetMode(cfg.Server.Environment)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PATCH", "DELETE"},
	}))
	router.Use(gin.LoggerWithWriter(log.Writer()))
	router.Use(middleware.ErrorHandler)
	log.Info("router has been set up")

	log.Info("setting up database connection pool...")
	databasePool, err := postgres.NewConnectionPool(cfg.Database, log)
	if err != nil {
		return nil, err
	}
	log.Info("setting up repository...")
	repository := repository.New(databasePool)

	log.Info("setting up service...")
	//service := service.New(log, repository)

	log.Info("setting up endpoint...")
	//handler = endpoint.New(service, log, validator.New())
	log.Info("endpoint set up")

	return &app{
		log:    log,
		cfg:    cfg,
		router: router,
		dbPool: databasePool,
		//handler: handler,
	}, nil
}

func (a *app) Run() error {

	go shutdown.Graceful(a)
	return nil
}

func (a *app) Close() error {
	if err := a.handler.Close(); err != nil {
		return nil
	}
	a.dbPool.Close()
	return nil
}
