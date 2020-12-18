package todo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

// ErrFoo ...
var ErrFoo = errors.New("Unable to handle Repo Requests")

type repo struct {
	db *sql.DB
	logger log.Logger
}

// NewRepo ...
func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db: db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) Add(ctx context.Context, todo Todo) error {
	sql := `INSERT INTO todos (id, text, completed) VALUES ($1, $2, $3)`

	if todo.Text == "" {
		return ErrFoo
	}
	_, err := repo.db.ExecContext(ctx, sql, todo.ID, todo.Text, todo.Completed)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetByID(ctx context.Context, id string) (string, error) {
	var todoText string

	err := repo.db.QueryRow("SELECT text from todos WHERE id=$1", id).Scan(&todoText)
	if err != nil {
		return "", err
	}
	return todoText, nil
}

func (repo *repo) Delete(ctx context.Context, id string) error {
	_, err := repo.db.Exec("DELETE FROM todos WHERE ID=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) Update(ctx context.Context, id string, text string, completed bool) (string, error) {
	sql := `UPDATE todos SET text=$2, completed=$3 WHERE ID=$1`

	_, err := repo.db.Exec(sql, id,text, completed)

	if err != nil {
		return "",err
	}
	return "updated",nil
}