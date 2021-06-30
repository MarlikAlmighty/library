package app

import (
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {

	t.Parallel()

	prefix := "LIBRARY"
	database := "library"

	if err := os.Setenv(prefix+"_"+"DB",
		"postgres://postgres:secret@0.0.0.0:5433/"  + database +  "?sslmode=disable"); err != nil {
		t.Error("setting env DB got failure", err)
	}

	if err := os.Setenv(prefix+"_"+"HTTP_PORT", "8010"); err != nil {
		t.Error("setting env HTTP_PORT got failure", err)
	}

	if err := os.Setenv(prefix+"_"+"MIGRATE", "true"); err != nil {
		t.Error("setting env MIGRATE got failure", err)
	}

	if err := os.Setenv(prefix+"_"+"PATH_TO_MIGRATE", "migrations"); err != nil {
		t.Error("setting env PATH_TO_MIGRATE got failure", err)
	}

	var srv Service

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
