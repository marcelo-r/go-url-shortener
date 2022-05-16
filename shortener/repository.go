package shortener

import "context"

// Repository for querying and storing data
type Repository interface {
	Query
	Store
}

type Query interface {
	FindByID(ctx context.Context, id string) (Redirect, error)
}

type Store interface {
	Store(ctx context.Context, redirect Redirect) error
}
