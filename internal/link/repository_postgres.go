package link

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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

func (r *postgresRepository) Create(ctx context.Context, link Link) error {
	query := `INSERT INTO links (code, target_url, created_at) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, link.Code, link.TargetURL, link.CreatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrCodeAlreadyExists
		}
		return err
	}
	return nil
}

func (r *postgresRepository) Get(ctx context.Context, code string) (Link, error) {
	query := `SELECT code, target_url, created_at FROM links WHERE code = $1`
	var link Link
	err := r.db.QueryRow(ctx, query, code).Scan(&link.Code, &link.TargetURL, &link.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Link{}, ErrNotFound
		}

		return Link{}, err
	}
	return link, nil
}

func (r *postgresRepository) Delete(ctx context.Context, code string) error {
	query := `DELETE FROM links WHERE code = $1`
	result, err := r.db.Exec(ctx, query, code)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}
