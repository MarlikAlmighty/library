package config

import (
	"os"
	"reflect"
	"testing"
)

func TestInitConfig(t *testing.T) {

	prefix := "LIBRARY"

	if err := os.Setenv(prefix+"_"+"DB", "TEST"); err != nil {
		t.Error("setting env DB got failure", err)
	}

	if err := os.Setenv(prefix+"_"+"HTTP_PORT", "8010"); err != nil {
		t.Error("setting env HTTP_PORT got failure", err)
	}

	type args struct {
		prefix string
	}

	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{"test_env", args{prefix}, &Config{"TEST", 8010}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitConfig(tt.args.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
