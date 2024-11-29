package postgresConn

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"time"
)

type ConfigPostgresql struct {
	Host     string `envconfig:"POSTGRES_HOST" validate:"required"`
	Port     string `envconfig:"POSTGRES_PORT" validate:"required"`
	Database string `envconfig:"POSTGRES_DATABASE" validate:"required"`
	User     string `envconfig:"POSTGRES_USER" validate:"required"`
	Password string `envconfig:"POSTGRES_PASSWORD" validate:"required"`
}

func NewPostgresPool(cfg *ConfigPostgresql) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, getConnStr(cfg))
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func ConvertPoolToDb(pool *pgxpool.Pool) (*sql.DB, error) {
	db := stdlib.OpenDBFromPool(pool)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func getConnStr(cfg *ConfigPostgresql) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
}
