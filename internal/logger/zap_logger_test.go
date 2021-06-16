package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
	"testing"
)

func setupLogsCapture() (*zap.Logger, *observer.ObservedLogs) {
	core, logs := observer.New(zap.InfoLevel)
	return zap.New(core), logs
}

func TestInitLogger(t *testing.T) {

	logger, logs := setupLogsCapture()

	logger.Warn("This is the warning")

	if logs.Len() != 1 {
		t.Errorf("No logs")
	} else {
		entry := logs.All()[0]
		if entry.Level != zap.WarnLevel || entry.Message != "This is the warning" {
			t.Errorf("Invalid log entry %v", entry)
		}
	}
}
