package postgresql

import (
	"context"
	"github.com/MarlikAlmighty/library/internal/config"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/tern/migrate"
	"github.com/pkg/errors"
)

type PostGreSQl struct {
	Pool *pgxpool.Pool
}

// Init database connection
func Init(cnf *config.Config) (*PostGreSQl, error) {

	ctx := context.Background()
	var db PostGreSQl

	if cnf.DB == "" {
		return nil, errors.New("the argument must not be an empty string")
	}

	if cnf.Migrate == true {
		if err := newMigrate(ctx, cnf); err != nil {
			return nil, errors.Wrap(err, "the argument must not be an empty")
		}
	}

	conf, err := pgxpool.ParseConfig(cnf.DB)
	if err != nil {
		return nil, err
	}

	if db.Pool, err = pgxpool.Connect(ctx, conf.ConnConfig.ConnString()); err != nil {
		return nil, err
	}

	if err = db.Pool.Ping(ctx); err != nil {
		return nil, err
	}

	return &db, nil
}

func newMigrate(ctx context.Context, cfg *config.Config) error {

	conn, err := pgx.Connect(ctx, cfg.DB)
	if err != nil {
		return errors.Wrap(err, "Unable to connection to database: ")
	}

	defer conn.Close(ctx)

	migrator, err := migrate.NewMigrator(ctx, conn, "schema_version")
	if err != nil {
		return errors.Wrap(err, "Unable to create a migrator: ")
	}

	err = migrator.LoadMigrations(cfg.PathToMigrate)
	if err != nil {
		return errors.Wrap(err, "Unable to load migrations: ")
	}

	err = migrator.Migrate(ctx)
	if err != nil {
		return errors.Wrap(err, "Unable to migrate: ")
	}

	if _, err := migrator.GetCurrentVersion(ctx); err != nil {
		return errors.Wrap(err, "Unable to get current schema version: ")
	}

	return nil
}
