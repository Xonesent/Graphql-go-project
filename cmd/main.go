package main

import (
	"context"
	"github.com/joho/godotenv"
	"graphql/config"
	"graphql/internal/server"
	"graphql/pkg/dependencies/fiber"
	mongoConn "graphql/pkg/dependencies/mongodb"
	"graphql/pkg/dependencies/postgresql"
	"log"
	"time"
)

func main() {
	time.Sleep(12 * time.Second)

	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	run(&cfg)
}

func run(cfg *config.Config) {
	psqlDB, err := postgresConn.NewPostgresPool(&cfg.Postgres)
	if err != nil {
		log.Fatal("Error connecting psqlDB", err.Error())
	}
	defer psqlDB.Close()

	if err := postgresConn.MigratePsqlUp(&cfg.Postgres); err != nil {
		log.Fatal(err)
	}

	mngDB, err := mongoConn.NewMongoDB(&cfg.Mongo)
	if err != nil {
		log.Fatal("Error connecting mngDB", err.Error())
	}
	defer mngDB.Disconnect(context.Background())

	if err := mongoConn.MigrateMngUp(&cfg.Mongo); err != nil {
		log.Fatal(err)
	}

	fiberApp := fiberConn.NewFiberClient()
	defer fiberApp.Shutdown()

	s := server.NewServer(
		cfg,
		fiberApp,
		psqlDB,
		mngDB.Database(cfg.Mongo.Database),
	)
	if err = s.Run(); err != nil {
		log.Fatal("Cannot start server", err.Error())
	}
}
