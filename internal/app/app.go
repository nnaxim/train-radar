package app

import (
	"github.com/nnaxim/train-radar/internal/config"
	"github.com/nnaxim/train-radar/internal/logger"
	"go.uber.org/zap"
)

type App struct {
	Config *config.Config
	Logger *zap.Logger
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

	return &App{
		Config: cfg,
		Logger: log,
	}, nil
}

func (a *App) Run() {
	a.Logger.Info(
		"Application started",
	)

	a.Logger.Info(
		"server port loaded",
		zap.String("port", a.Config.ServerPort),
	)
}
