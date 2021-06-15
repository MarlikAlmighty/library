package logger

import (
	"go.uber.org/zap"
)

func InitZapLog(loggerType string) (*zap.SugaredLogger, error) {

	var config zap.Config

	switch loggerType {
	case "development":
		config = zap.NewDevelopmentConfig()
	case "production":
		config = zap.NewProductionConfig()
	default:
		config = zap.NewDevelopmentConfig()
	}

	logger, err := config.Build()
	if err != nil {
		return &zap.SugaredLogger{}, err
	}

	return logger.Sugar(), nil
}
