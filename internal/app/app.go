package app

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/reversersed/taskservice/docs"
	"github.com/reversersed/taskservice/internal/application/services"
	"github.com/reversersed/taskservice/internal/config"
	"github.com/reversersed/taskservice/internal/infrastructure/repository"
	"github.com/reversersed/taskservice/internal/interface/api/rest"
	"github.com/reversersed/taskservice/pkg/logging/logrus"
	"github.com/reversersed/taskservice/pkg/middleware"
	"github.com/reversersed/taskservice/pkg/postgres"
	"github.com/reversersed/taskservice/pkg/shutdown"
	"github.com/reversersed/taskservice/pkg/validator"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	service := services.NewTaskService(repository, log)

	log.Info("setting up endpoint...")
	rest.NewTaskController(router, service, validator.New())

	log.Info("endpoint set up")

	return &app{
		log:    log,
		cfg:    cfg,
		router: router,
		dbPool: databasePool,
	}, nil
}

func (a *app) Run() error {
	if a.cfg.Server.Environment == "debug" {
		a.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	go shutdown.Graceful(a)

	if err := a.router.Run(fmt.Sprintf("%s:%d", a.cfg.Server.Url, a.cfg.Server.Port)); err != nil {
		return err
	}
	go shutdown.Graceful(a)
	return nil
}

func (a *app) Close() error {
	a.dbPool.Close()
	return nil
}
