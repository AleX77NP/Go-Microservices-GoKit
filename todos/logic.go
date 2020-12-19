package todo

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	repository Repository
	logger     log.Logger
}

// NewService function
func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s service) Add(ctx context.Context, id string, text string, completed bool) (string, error) {
	logger := log.With(s.logger, "method", "Add")

	todo := Todo{
		ID:        id,
		Text:      text,
		Completed: completed,
	}

	if err := s.repository.Add(ctx, todo); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	logger.Log("create todo", id)
	return "Success", nil
}

func (s service) Delete(ctx context.Context, id string) error {
	logger := log.With(s.logger, "method", "Delete")

	if err := s.repository.Delete(ctx, id); err != nil {
		level.Error(logger).Log("err", err)
		return err
	}
	logger.Log("Delete todo", id)
	return nil
}

func (s service) GetByID(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetByID")

	todo, err := s.repository.GetByID(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Get Todo", id)

	return todo, nil
}

func (s service) Update(ctx context.Context, id string, text string, completed bool) (string, error) {
	logger := log.With(s.logger, "method", "Update")

	updated, err := s.repository.Update(ctx, id, text, completed)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Update Todo", id)

	return updated, nil
}
