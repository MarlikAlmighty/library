package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"os"
)

type Config struct {
	DB            string `required:"true"`
	HTTPPort      int    `required:"true" split_words:"true"`
	Migrate       bool   `required:"true"`
	PathToMigrate string `required:"true" split_words:"true"`
}

func InitConfig() (*Config, error) {

	var cfg Config

	prefix := os.Getenv("PREFIX")

	if prefix == "" {
		return nil, errors.New("prefix not set in env")
	}

	err := envconfig.Process(prefix, &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "error reading env: ")
	}

	return &cfg, nil
}
