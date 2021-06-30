package postgresql

import (
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"log"
	"os"
	"testing"
)

var (
	pool     *dockertest.Pool
	resource *dockertest.Resource
)

func TestMain(m *testing.M) {

	var err error
	database := "library"

	if pool, err = dockertest.NewPool(""); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "latest",
		Env: []string{
			"POSTGRES_USER=postgres",
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_DB=" + database,
			"listen_addresses = '*'",
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "localhost", HostPort: "5433"},
				{HostIP: "localhost", HostPort: "5434"},
			},
		},
	}

	if resource, err = pool.RunWithOptions(&opts); err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	i := m.Run()

//	time.Sleep(200 * time.Second)

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(i)
}
