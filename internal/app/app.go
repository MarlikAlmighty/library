package app

import (
	"context"
	"github.com/MarlikAlmighty/library/internal/logger"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/MarlikAlmighty/library/internal/config"
	"github.com/MarlikAlmighty/library/internal/repository/postgresql"
	"go.uber.org/zap"
)

type Service struct {
	Logger *zap.Logger       `logger:"-"`
	Conf   *config.Config    `conf:"-"`
	Booker postgresql.Booker `booker:"-"`
	Pool   *pgxpool.Pool     `pool:"-"`
}

// New return new app
func New() *Service {
	return &Service{}
}

// InitConfig init configuration
func (s *Service) InitConfig() error {

	conf, err := config.InitConfig()
	if err != nil {
		return err
	}
	s.Conf = conf
	return nil
}

func (s *Service) InitLogger() error {

	lg, err := logger.InitLogger()
	if err != nil {
		return err
	}
	s.Logger = lg
	return nil
}

func (s *Service) Migrate() error {

	if err := postgresql.Migrate(context.Background(), s.Conf); err != nil {
		return err
	}

	return nil
}

func (s *Service) InitDatabase() error {

	pool, err := postgresql.InitDatabase(s.Conf)
	if err != nil {
		return err
	}

	s.Pool = pool
	return nil
}

func (s *Service) Stop() {
	s.Pool.Close()
}
