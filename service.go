package todo

import (
	"context"
)

// Service for Todos
type Service interface {
	GetByID(ctx context.Context, id string) (string, error)
	Add(ctx context.Context, id string, text string, completed bool) (string, error)
	Update(ctx context.Context, id string, text string, completed bool) (string,error)
	Delete(ctx context.Context, id string) error
}