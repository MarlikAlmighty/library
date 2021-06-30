package app

import (
	"context"

	"github.com/MarlikAlmighty/library/internal/config"
	"github.com/MarlikAlmighty/library/internal/logger"
	"github.com/MarlikAlmighty/library/internal/repository/postgresql"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type Service struct {
	Logger *zap.Logger    `logger:"-"`
	Conf   *config.Config `conf:"-"`
	Pool   *pgxpool.Pool  `pool:"-"`
}

// New app
func New() (*Service, error) {

	var (
		srv Service
		err error
	)

	if srv.Logger, err = logger.InitLogger(); err != nil {
		return nil, err
	}

	if srv.Conf, err = config.InitConfig(); err != nil {
		return nil, err
	}

	if srv.Conf.Migrate {
		if err := postgresql.Migrate(context.Background(), srv.Conf); err != nil {
			return nil, err
		}
	}

	if srv.Pool, err = postgresql.InitDatabase(srv.Conf); err != nil {
		return nil, err
	}

	return &srv, nil
}

func (srv *Service) Stop() {
	srv.Pool.Close()
}
