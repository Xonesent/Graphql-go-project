package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
	"graphql/pkg/constant"
	fiberConn "graphql/pkg/dependencies/fiber"
	"graphql/pkg/dependencies/mongodb"
	postgresConn "graphql/pkg/dependencies/postgresql"
)

type Config struct {
	Mongo    mongoConn.ConfigMongo         `validate:"required"`
	Postgres postgresConn.ConfigPostgresql `validate:"required"`
	Fiber    fiberConn.FiberConfig         `validate:"required"`
}

func LoadConfig() (cfg Config, err error) {
	if err := envconfig.Process("", &cfg); err != nil {
		return cfg, err
	}

	if err := validator.New().Struct(&cfg); err != nil {
		return cfg, err
	}

	constant.Host = cfg.Fiber.Host

	return
}
