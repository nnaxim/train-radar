package app

import (
	"github.com/nnaxim/train-radar/internal/config"
	"github.com/nnaxim/train-radar/internal/database"
	"github.com/nnaxim/train-radar/internal/logger"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

type App struct {
	Config *config.Config
	Logger *zap.Logger
	DB     *bun.DB
}

func New() (*App, error) {
	cfg, err := config.Load()

	if err != nil {
		return nil, err
	}

	log, err := logger.New()

	if err != nil {
		return nil, err
	}

	db, err := database.New(cfg)

	if err != nil {
		return nil, err
	}

	return &App{
		Config: cfg,
		Logger: log,
		DB:     db,
	}, nil
}

func (a *App) Run() {
	a.Logger.Info(
		"Application started",
	)

	a.Logger.Info(
		"database connected",
	)

	a.Logger.Info(
		"server port loaded",
		zap.String("port", a.Config.ServerPort),
	)
}
