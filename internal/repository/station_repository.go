package repository

import (
	"context"

	"github.com/nnaxim/train-radar/internal/model"
	"github.com/uptrace/bun"
)

type StationRepository struct {
	db *bun.DB
}

func NewStationRepository(db *bun.DB) *StationRepository {
	return &StationRepository{
		db: db,
	}
}

func (r *StationRepository) Create(
	ctx context.Context,
	station *model.Station,
) error {
	_, err := r.db.
		NewInsert().
		Model(station).
		Exec(ctx)

	return err
}

func (r *StationRepository) GetByID(
	ctx context.Context,
	id uint64,
) (*model.Station, error) {
	station := new(model.Station)

	err := r.db.
		NewSelect().
		Model(station).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return station, nil
}

func (r *StationRepository) GetAll(
	ctx context.Context,
) ([]model.Station, error) {
	var stations []model.Station

	err := r.db.
		NewSelect().
		Model(&stations).
		Order("id ASC").
		Scan(ctx)

	return stations, err
}

func (r *StationRepository) Update(
	ctx context.Context,
	station *model.Station,
) error {
	_, err := r.db.
		NewUpdate().
		Model(station).
		WherePK().
		Exec(ctx)

	return err
}

func (r *StationRepository) DeleteByID(
	ctx context.Context,
	id uint64,
) error {
	_, err := r.db.
		NewDelete().
		Model((*model.Station)(nil)).
		Where("id = ?", id).
		Exec(ctx)

	return err
}
