package mongoConn

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
	"log"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateMngUp(cfg *ConfigMongo) error {
	log.Println("Applying up mongo migrations...")

	m, err := getMigrateInstance(cfg)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	log.Println("Mongo migration up completed successfully!")
	return nil
}

func MigrateMngDown(cfg *ConfigMongo) error {
	log.Println("Applying down mongo migrations...")

	m, err := getMigrateInstance(cfg)
	if err != nil {
		return err
	}
	if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	log.Println("Mongo migration down completed successfully!")
	return nil
}

func getMigrateInstance(cfg *ConfigMongo) (*migrate.Migrate, error) {
	client, err := NewMongoDB(cfg)
	if err != nil {
		return nil, err
	}

	driver, err := mongodb.WithInstance(client, &mongodb.Config{DatabaseName: cfg.Database})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations/mongodb", cfg.Database, driver)
	if err != nil {
		return nil, err
	}

	return m, nil
}
