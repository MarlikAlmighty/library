package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	DB       string `required:"true"`
	HTTPPort int    `required:"true" split_words:"true"`
}

func InitConfig(prefix string) (*Config, error) {

	var cfg Config

	if prefix == "" {
		return nil, errors.New("the argument must not be an empty string")
	}

	err := envconfig.Process(prefix, &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "error reading env: ")
	}

	return &cfg, nil
}
