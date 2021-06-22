package postgresql

import (
	"context"

	"github.com/MarlikAlmighty/library/internal/config"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/tern/migrate"
	"github.com/pkg/errors"
)

func newMigrate(ctx context.Context, cfg *config.Config) error {

	if cfg.DB == "" {
		return errors.New("connect string to db not set in env")
	}

	if cfg.PathToMigrate == "" {
		return errors.New("path to migrate folder not set in env")
	}

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
