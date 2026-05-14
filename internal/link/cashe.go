package link

import "context"

type Cache interface {
	Get(ctx context.Context, code string) (string, error)
	Set(ctx context.Context, code string, targetURL string) error
	Delete(ctx context.Context, code string) error
}
