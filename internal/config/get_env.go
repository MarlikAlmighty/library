package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	DB       string `required:"true"`
	HTTPPort int    `required:"true" split_words:"true"`
}

func LoadEnv(prefix string) (*Config, error) {

	var cfg Config

	err := envconfig.Process(prefix, &cfg)
	if err != nil {
		return &cfg, errors.Wrap(err, "error reading env: %v\n")
	}

	return &cfg, nil
}
