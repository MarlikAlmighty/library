package postgresql_test

import (
	"context"
	"github.com/MarlikAlmighty/library/internal/repository/postgresql"
	"github.com/jackc/pgx/v4/pgxpool"
	"reflect"
	"testing"

	"github.com/MarlikAlmighty/library/internal/config"
)

func TestInitDatabase(t *testing.T) {

	var (
		pgxPool *pgxpool.Pool
		cfg     config.Config
		err     error
	)

	database := "library"
	cfg.DB = "postgres://postgres:secret@localhost:5432/" + database + "?sslmode=disable"

	if err = pool.Retry(func() error {

		var err error

		if pgxPool, err = pgxpool.Connect(context.Background(), cfg.DB); err != nil {
			return err
		}

		return pgxPool.Ping(context.Background())

	}); err != nil {
		t.Errorf("Could not connect to docker: %s", err)
	}

	defer pgxPool.Close()

	type args struct {
		cfg *config.Config
	}

	tests := []struct {
		name    string
		args    args
		want    *pgxpool.Pool
		wantErr bool
	}{
		{"init_database", args{&cfg}, pgxPool, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := postgresql.InitDatabase(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitDatabase() got = %v, want %v", got, tt.want)
			}
		})
	}
}
