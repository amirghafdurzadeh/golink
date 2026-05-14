package link

import (
	"context"
	"errors"
)

var ErrCodeAlreadyExists = errors.New("code already exists")

type Service interface {
	Create(ctx context.Context, link Link) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, link Link) error {
	exists, err := s.repository.CodeExists(ctx, link.Code)
	if err != nil {
		return err
	}

	if exists {
		return ErrCodeAlreadyExists
	}

	return s.repository.Create(ctx, link)
}
