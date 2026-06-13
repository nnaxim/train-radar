package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nnaxim/train-radar/internal/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func New(cfg *config.Config) (*bun.DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	sqldb := sql.OpenDB(
		pgdriver.NewConnector(
			pgdriver.WithDSN(dsn),
		),
	)

	db := bun.NewDB(
		sqldb,
		pgdialect.New(),
	)

	db.AddQueryHook(
		bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
		),
	)

	if err := db.PingContext(context.Background()); err != nil {
		return nil, err
	}

	return db, nil
}
