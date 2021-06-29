package postgresql

import (
	"context"
	"testing"

	"github.com/MarlikAlmighty/library/internal/config"
	"github.com/jackc/pgx/v4"
)

func TestMigrate(t *testing.T) {

	database := "library"

	var (
		conn *pgx.Conn
		cfg  config.Config
		err error
	)

	cfg.Migrate = true
	cfg.PathToMigrate = "./migrations"
	cfg.DB = "postgres://postgres:secret@localhost:5433/" + database + "?sslmode=disable"

	if err = pool.Retry(func() error {
		var err error
		conn, err = pgx.Connect(context.Background(), cfg.DB)
		if err != nil {
			return err
		}
		return conn.Ping(context.Background())
	}); err != nil {
		t.Errorf("Could not connect to docker: %s", err)
	}

	defer conn.Close(context.Background())

	type args struct {
		ctx context.Context
		cfg *config.Config
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"migrate", args{context.Background(), &cfg}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Migrate(tt.args.ctx, tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("Migrate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
