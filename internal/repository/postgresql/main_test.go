package postgresql_test

import (
	"github.com/MarlikAlmighty/library/internal/config"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"log"
	"os"
	"testing"
)

var (
	pool     *dockertest.Pool
	resource *dockertest.Resource
	cnf      config.Config
)

func TestMain(m *testing.M) {

	var err error
	database := "library"

	cnf.DB = "postgres://postgres:secret@localhost:5432/" + database + "?sslmode=disable"

	if pool, err = dockertest.NewPool(""); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "12.3",
		Env: []string{
			"POSTGRES_USER=postgres",
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_DB=" + database,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "localhost", HostPort: "5432"},
			},
		},
	}

	if resource, err = pool.RunWithOptions(&opts); err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	i := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(i)
}
