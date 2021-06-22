package postgresql

import (
	"context"
	"github.com/MarlikAlmighty/library/internal/config"
	"testing"
)

func Test_newMigrate(t *testing.T) {

	ctx := context.Background()
	cnf := &config.Config{}

	cnf.Migrate = true
	cnf.PathToMigrate = "./migrations"
	cnf.HTTPPort = 3000
	cnf.DB = "postgres://user:password@localhost:5432/library?sslmode=disable"

	type args struct {
		ctx context.Context
		cfg *config.Config
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"migrate", args{ctx, cnf}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := newMigrate(tt.args.ctx, tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("newMigrate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
