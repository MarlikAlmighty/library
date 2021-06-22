package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/MarlikAlmighty/library/internal/config"
	"github.com/ory/dockertest/v3"
)

func Test_newMigrate(t *testing.T) {

	ctx := context.Background()
	cnf := &config.Config{}
	database := "library"

	cnf.Migrate = true
	cnf.PathToMigrate = "./migrations"
	cnf.HTTPPort = 3000
	cnf.DB = "postgres://postgres:secret@localhost:5432/library?sslmode=disable"

	var (
		db       *sql.DB
		err      error
		pool     *dockertest.Pool
		resource *dockertest.Resource
	)

	if pool, err = dockertest.NewPool(""); err != nil {
		t.Errorf("Could not connect to docker: %s", err)
	}

	if resource, err = pool.Run("postgres", "9.6",
		[]string{"POSTGRES_PASSWORD=secret", "POSTGRES_DB=" + database}); err != nil {
		t.Errorf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		var err error
		db, err = sql.Open("postgres",
			fmt.Sprintf("postgres://postgres:secret@localhost:%s/%s?sslmode=disable", resource.GetPort("5432/tcp"), database))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		t.Errorf("Could not connect to docker: %s", err)
	}

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

	// When you're done, kill and remove the container
	if err = pool.Purge(resource); err != nil {
		t.Errorf("Could not purge resource: %s", err)
	}

}
