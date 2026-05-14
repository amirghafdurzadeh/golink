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
	shortCodeLength int
}

func NewService(repository Repository, shortCodeLength int) Service {
	return &service{
		repository:      repository,
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

	return link, nil
}

func (s *service) Get(ctx context.Context, code string) (Link, error) {
	return s.repository.Get(ctx, code)
}

func (s *service) Delete(ctx context.Context, code string) error {
	return s.repository.Delete(ctx, code)
}
