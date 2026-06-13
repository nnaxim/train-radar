package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Station struct {
	bun.BaseModel `bun:"table:stations"`

	ID uint64 `bun:",pk,autoincrement"`

	Name    string
	Country string

	Latitude  float64
	Longitude float64

	CreatedAt time.Time
	UpdatedAt time.Time
}
