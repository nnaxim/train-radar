package app

import (
	"net/http"

	"github.com/nnaxim/train-radar/internal/config"
	"github.com/nnaxim/train-radar/internal/database"
	"github.com/nnaxim/train-radar/internal/handler"
	"github.com/nnaxim/train-radar/internal/logger"
	"github.com/nnaxim/train-radar/internal/repository"
	"github.com/nnaxim/train-radar/internal/server"
	"github.com/nnaxim/train-radar/internal/service"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

type App struct {
	Config            *config.Config
	Logger            *zap.Logger
	DB                *bun.DB
	StationRepository *repository.StationRepository
	StationService    *service.StationService
	StationHandler    *handler.StationHandler
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

	stationRepository := repository.NewStationRepository(
		db,
	)

	stationService := service.NewStationService(
		stationRepository,
	)

	StationHandler := handler.NewStationHandler(
		stationService,
	)

	return &App{
		Config:            cfg,
		Logger:            log,
		DB:                db,
		StationRepository: stationRepository,
		StationService:    stationService,
		StationHandler:    StationHandler,
	}, nil
}

func (a *App) Run() {
	srv := server.New(
		a.StationHandler,
	)

	a.Logger.Info(
		"starting http server",
		zap.String("port", a.Config.ServerPort),
	)

	err := http.ListenAndServe(
		":"+a.Config.ServerPort,
		srv.Router(),
	)

	if err != nil {
		a.Logger.Fatal(
			"failed to start server",
			zap.Error(err),
		)
	}
}
