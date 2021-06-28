package main_test

import (
	"log"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var (
	Pool     *dockertest.Pool
	Resource *dockertest.Resource
)

func TestMain(m *testing.M) {

	var err error
	database := "library"

	if Pool, err = dockertest.NewPool(""); err != nil {
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

	if Resource, err = Pool.RunWithOptions(&opts); err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	i := m.Run()

	if err := Pool.Purge(Resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(i)
}
