package link

import (
	"context"
	"time"
)

type Service interface {
	Create(ctx context.Context, customCode string, targetURL string) (Link, error)
	Get(ctx context.Context, code string) (Link, error)
	Delete(ctx context.Context, code string) error
	GetBaseURL() string
}

type ServiceConfig struct {
	BaseURL         string
	ShortCodeLength int
}

type service struct {
	cfg        ServiceConfig
	repository Repository
	cache      Cache
}

func NewService(
	cfg ServiceConfig,
	repository Repository,
	cache Cache,
) Service {

	return &service{
		cfg:        cfg,
		repository: repository,
		cache:      cache,
	}
}

func (s *service) Create(ctx context.Context, customCode string, targetURL string) (Link, error) {
	code, err := buildCode(customCode, s.cfg.ShortCodeLength)
	if err != nil {
		return Link{}, ErrRandomNumberGen
	}

	link := Link{
		Code:      code,
		TargetURL: targetURL,
		CreatedAt: time.Now(),
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

func (s *service) GetBaseURL() string {
	return s.cfg.BaseURL
}
