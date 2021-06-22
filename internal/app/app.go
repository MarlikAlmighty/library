package app

import (
	"github.com/MarlikAlmighty/library/internal/config"
	"github.com/MarlikAlmighty/library/internal/logger"
	"github.com/MarlikAlmighty/library/internal/repository/postgresql"
	"go.uber.org/zap"
)

type Service struct {
	Logger *zap.Logger            `logger:"-"`
	Conf   *config.Config         `config:"-"`
	DB     *postgresql.PostGreSQl `db:"-"`
}

// New init app
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

	if srv.DB, err = postgresql.InitDatabase(srv.Conf); err != nil {
		return nil, err
	}

	return &srv, nil
}

func (srv *Service) Stop() {
	srv.DB.Pool.Close()
}
