package service

import (
	"context"

	"github.com/nnaxim/train-radar/internal/model"
	"github.com/nnaxim/train-radar/internal/repository"
)

type StationService struct {
	repository *repository.StationRepository
}

func NewStationService(
	repository *repository.StationRepository,
) *StationService {
	return &StationService{
		repository: repository,
	}
}

func (s *StationService) Create(
	ctx context.Context,
	station *model.Station,
) error {
	return s.repository.Create(
		ctx,
		station,
	)
}

func (s *StationService) GetByID(
	ctx context.Context,
	id uint64,
) (*model.Station, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *StationService) GetAll(
	ctx context.Context,
) ([]model.Station, error) {
	return s.repository.GetAll(ctx)
}
