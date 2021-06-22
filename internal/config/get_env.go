package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	DB            string `required:"true"`
	HTTPPort      int    `required:"true" split_words:"true"`
	Migrate       bool   `required:"true"`
	PathToMigrate string `required:"true" split_words:"true"`
}

func InitConfig() (*Config, error) {

	var cfg Config

	// TODO Fix it
	prefix := "LIBRARY"

	err := envconfig.Process(prefix, &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "error reading env: ")
	}

	return &cfg, nil
}
