package link

import (
	"context"
	"time"
)

type Service interface {
	Create(ctx context.Context, customCode string, targetURL string) (Link, error)
	Get(ctx context.Context, code string) (Link, error)
	Delete(ctx context.Context, code string) error
}

type service struct {
	repository      Repository
	cache           Cache
	shortCodeLength int
}

func NewService(
	repository Repository,
	cache Cache,
	shortCodeLength int,
) Service {

	return &service{
		repository:      repository,
		cache:           cache,
		shortCodeLength: shortCodeLength,
	}
}

func (s *service) Create(ctx context.Context, customCode string, targetURL string) (Link, error) {
	link := Link{
		Code:      buildCode(customCode, s.shortCodeLength),
		TargetURL: targetURL,
		CreatedAt: time.Now(),
	}

	exists, err := s.repository.CodeExists(ctx, link.Code)
	if err != nil {
		return Link{}, err
	}

	if exists {
		return Link{}, ErrCodeAlreadyExists
	}

	err = s.repository.Create(ctx, link)
	if err != nil {
		return Link{}, err
	}

	_ = s.cache.Set(ctx, link.Code, link.TargetURL)

	return link, nil
}

func (s *service) Get(ctx context.Context, code string) (Link, error) {
	targetURL, err := s.cache.Get(ctx, code)
	if err == nil {
		return Link{
			Code:      code,
			TargetURL: targetURL,
		}, nil
	}

	link, err := s.repository.Get(ctx, code)
	if err != nil {
		return Link{}, err
	}

	_ = s.cache.Set(ctx, link.Code, link.TargetURL)

	return link, nil
}

func (s *service) Delete(ctx context.Context, code string) error {
	err := s.repository.Delete(ctx, code)
	if err != nil {
		return err
	}

	_ = s.cache.Delete(ctx, code)

	return nil
}
