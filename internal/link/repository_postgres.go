package link

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &postgresRepository{
		db: db,
	}
}

// Create implements [Repository].
func (r *postgresRepository) Create(ctx context.Context, link Link) error {
	panic("unimplemented")
}

// Get implements [Repository].
func (r *postgresRepository) Get(ctx context.Context, code string) (Link, error) {
	panic("unimplemented")
}

// Delete implements [Repository].
func (r *postgresRepository) Delete(ctx context.Context, code string) error {
	panic("unimplemented")
}

// CodeExists implements [Repository].
func (r *postgresRepository) CodeExists(ctx context.Context, code string) (bool, error) {
	panic("unimplemented")
}
