package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/reversersed/taskservice/internal/config"
	"github.com/reversersed/taskservice/pkg/logging/logrus"
)

type app struct {
	router *gin.Engine
	cfg    *config.Config
	log    *logrus.Logger
	dbPool *pgxpool.Pool
}
