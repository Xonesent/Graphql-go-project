package postgresConn

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"log"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigratePsqlUp(cfg *ConfigPostgresql) error {
	log.Println("Applying up postgresql migrations...")

	m, err := getMigrateInstance(cfg)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	log.Println("Postgresql  migration up completed successfully!")
	return nil
}

func MigratePsqlDown(cfg *ConfigPostgresql) error {
	log.Println("Applying down postgresql migrations...")

	m, err := getMigrateInstance(cfg)
	if err != nil {
		return err
	}
	if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	log.Println("Postgresql migration down completed successfully!")
	return nil
}

func getMigrateInstance(cfg *ConfigPostgresql) (*migrate.Migrate, error) {
	pool, err := NewPostgresPool(cfg)
	if err != nil {
		return nil, err
	}

	db, err := ConvertPoolToDb(pool)
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{DatabaseName: cfg.Database})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations/postgresql", cfg.Database, driver)
	if err != nil {
		return nil, err
	}

	return m, nil
}
