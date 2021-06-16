package app

import (
	"database/sql"

	"github.com/MarlikAlmighty/library/internal/config"
	"github.com/MarlikAlmighty/library/internal/logger"
	"github.com/MarlikAlmighty/library/internal/repository/postgresql"
	"go.uber.org/zap"
)

type Service struct {
	Logger *zap.Logger    `logger:"-"`
	Conf   *config.Config `config:"-"`
	Pool   *sql.DB        `pool:"-"`
}

// New init app
func New() (*Service, error) {

	// TODO Fix it
	prefix := "LIBRARY"

	var (
		srv Service
		err error
	)

	if srv.Logger, err = logger.InitLogger(); err != nil {
		return &srv, err
	}

	if srv.Conf, err = config.InitConfig(prefix); err != nil {
		return &srv, err
	}

	if srv.Pool, err = postgresql.InitDataBase(srv.Conf.DB); err != nil {
		return &srv, err
	}

	return &srv, nil
}

func (srv *Service) Stop() {

	if err := srv.Pool.Close(); err != nil {
		srv.Logger.Sugar().Infof("error database Close() func: %s", err)
	}
}
