package logger

import (
	"go.uber.org/zap"
)

func InitLogger() (*zap.Logger, error) {

	var config zap.Config

	config = zap.NewProductionConfig()

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
