package link

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, link Link) error
	Get(ctx context.Context, code string) (Link, error)
	Delete(ctx context.Context, code string) error
}
