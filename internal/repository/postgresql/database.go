package postgresql

import (
	"context"
	"github.com/MarlikAlmighty/library/internal/config"
	"github.com/MarlikAlmighty/library/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type Booker interface {
	ListAllBookcases() (models.BookCases, error)
}

type DB struct {
	Pool *pgxpool.Pool
}

// InitDatabase database connection init
func InitDatabase(cnf *config.Config) (*pgxpool.Pool, error) {

	var db DB

	if cnf.DB == "" {
		return nil, errors.New("the argument must not be an empty string")
	}

	conf, err := pgxpool.ParseConfig(cnf.DB)
	if err != nil {
		return nil, err
	}

	if db.Pool, err = pgxpool.Connect(context.Background(), conf.ConnConfig.ConnString()); err != nil {
		return nil, err
	}

	if err = db.Pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return db.Pool, nil
}

func (db DB) ListAllBookcases() (models.BookCases, error) {
	var md models.BookCases
	return md, nil
}
