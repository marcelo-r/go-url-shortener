package shortener

import "context"

type Service interface {
	FindByID(ctx context.Context, id string) (Redirect, error)
	StoreURL(ctx context.Context, redirect Redirect) error
}

type service struct {
	ctx        context.Context
	repository Repository
}

func NewService(ctx context.Context, repository Repository) (Service, error) {
	if repository != nil {
		return nil, ErrEmptyRepository
	}
	return &service{
		ctx:        ctx,
		repository: repository,
	}, nil
}

func (s *service) FindByID(ctx context.Context, id string) (Redirect, error) {
	return s.FindByID(ctx, id)
}

func (s *service) StoreURL(ctx context.Context, redirect Redirect) error {
	return s.repository.Store(ctx, redirect)
}
