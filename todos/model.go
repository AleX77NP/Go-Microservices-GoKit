package todo

import (
	"context"
)

// Todo model
type Todo struct {
	ID string `json:"id"`
	Text string `json:"text"`
	Completed bool `json:"completed"`
}
// Repository type
type Repository interface {
	GetByID(ctx context.Context, id string) (string, error)
	Add(ctx context.Context, todo Todo) error
	Update(ctx context.Context, id string, text string, completed bool) (string, error)
	Delete(ctx context.Context, id string) error
}