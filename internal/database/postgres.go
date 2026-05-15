package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(ctx context.Context, connURL string) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(ctx, connURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
