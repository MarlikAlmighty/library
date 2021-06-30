package postgresql

import (
	"context"
	"github.com/MarlikAlmighty/library/internal/config"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/tern/migrate"
	"github.com/pkg/errors"
	"path"
)

func Migrate(ctx context.Context, cfg *config.Config) error {

	var m *migrate.Migrator

	if cfg.DB == "" {
		return errors.New("connect string to db not set in env")
	}

	if cfg.Migrate && cfg.PathToMigrate == "" {
		return errors.New("path to migrate folder not set in env")
	}

	conn, err := pgx.Connect(ctx, cfg.DB)
	if err != nil {
		return errors.Wrap(err, "Unable to connection to database: ")
	}

	defer conn.Close(ctx)

	if m, err = migrate.NewMigrator(ctx, conn, "schema_version"); err != nil {
		return errors.Wrap(err, "Unable to create a migrator: ")
	}

	dir := path.Join("./", cfg.PathToMigrate)

	if err = m.LoadMigrations(dir); err != nil {
		return errors.Wrap(err, "Unable to load migrations: ")
	}

	if err = m.Migrate(ctx); err != nil {
		return errors.Wrap(err, "Unable to migrate: ")
	}

	if _, err = m.GetCurrentVersion(ctx); err != nil {
		return errors.Wrap(err, "Unable to get current schema version: ")
	}

	return nil
}
