package config

import (
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {

	prefix := "LIBRARY"

	if err := os.Setenv(prefix + "_" + "DB", "TEST"); err != nil {
		t.Error("setting env DB got failure", err)
	}

	if err := os.Setenv(prefix + "_" + "HTTP_PORT", "8010"); err != nil {
		t.Error("setting env HTTP_PORT got failure", err)
	}

	m, err := LoadEnv(prefix)
	if err != nil {
		t.Error("got an error while testing a function LoadEnv()", err)
	}

	if len(m.DB) <= 0 {
		t.Error("wanted a string, got zero or less", err)
	}

	if m.HTTPPort <= 0 {
		t.Error("wanted a decimal, got zero or less", err)
	}
}
