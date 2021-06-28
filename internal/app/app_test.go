package app_test

import (
	"context"
	"github.com/MarlikAlmighty/library/internal/app"
	"os"
	"reflect"
	"testing"

	"github.com/MarlikAlmighty/library/internal/config"
	"github.com/MarlikAlmighty/library/internal/logger"
	"github.com/MarlikAlmighty/library/internal/repository/postgresql"
)

func TestNew(t *testing.T) {

	prefix := "LIBRARY"

	if err := os.Setenv(prefix+"_"+"DB", "LIBRARY"); err != nil {
		t.Error("setting env DB got failure", err)
	}

	if err := os.Setenv(prefix+"_"+"HTTP_PORT", "8010"); err != nil {
		t.Error("setting env HTTP_PORT got failure", err)
	}

	if err := os.Setenv(prefix+"_"+"MIGRATE", "true"); err != nil {
		t.Error("setting env MIGRATE got failure", err)
	}

	if err := os.Setenv(prefix+"_"+"PATH_TO_MIGRATE", "./migrations"); err != nil {
		t.Error("setting env PATH_TO_MIGRATE got failure", err)
	}

	var (
		srv app.Service
		err error
	)

	if srv.Logger, err = logger.InitLogger(); err != nil {
		t.Errorf("init logger fail %v", err)
	}

	if srv.Conf, err = config.InitConfig(); err != nil {
		t.Errorf("init config fail %v", err)
	}

	if srv.Conf.Migrate {
		if err := postgresql.Migrate(context.Background(), srv.Conf); err != nil {
			t.Errorf("init migrate fail %v", err)
		}
	}

	if srv.Pool, err = postgresql.InitDatabase(srv.Conf); err != nil {
		t.Errorf("connect to database fail %v", err)
	}

	tests := []struct {
		name    string
		want    *Service
		wantErr bool
	}{
		{"app", &srv, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}
