package link

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, link Link) error
	Get(ctx context.Context, code string) (Link, error)
	Delete(ctx context.Context, code string) error
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

// Create implements [Repository].
func (r *repository) Create(ctx context.Context, link Link) error {
	panic("unimplemented")
}

// Get implements [Repository].
func (r *repository) Get(ctx context.Context, code string) (Link, error) {
	panic("unimplemented")
}

// Delete implements [Repository].
func (r *repository) Delete(ctx context.Context, code string) error {
	panic("unimplemented")
}

// CodeExists implements [Repository].
func (r *repository) CodeExists(ctx context.Context, code string) (bool, error) {
	panic("unimplemented")
}
