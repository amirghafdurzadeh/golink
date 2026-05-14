package link

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, link Link) error
	CodeExists(ctx context.Context, code string) (bool, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

// CodeExists implements [Repository].
func (r *repository) CodeExists(ctx context.Context, code string) (bool, error) {
	panic("unimplemented")
}

// Create implements [Repository].
func (r *repository) Create(ctx context.Context, link Link) error {
	panic("unimplemented")
}
