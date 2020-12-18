package todo

import (
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