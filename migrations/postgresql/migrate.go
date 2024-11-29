package main

import (
	"flag"
	"github.com/joho/godotenv"
	"graphql/config"
	postgresConn "graphql/pkg/dependencies/postgresql"
	"log"
)

var (
	methodFlag = flag.String("method", "up", "Your migrate method")
)

func main() {
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	switch *methodFlag {
	case "up":
		if err := postgresConn.MigratePsqlUp(&cfg.Postgres); err != nil {
			log.Fatal(err)
		}
	case "down":
		if err := postgresConn.MigratePsqlDown(&cfg.Postgres); err != nil {
			log.Fatal(err)
		}
	}
}
