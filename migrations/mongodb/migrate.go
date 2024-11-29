package main

import (
	"flag"
	"github.com/joho/godotenv"
	"graphql/config"
	mongoConn "graphql/pkg/dependencies/mongodb"
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
		if err := mongoConn.MigrateMngUp(&cfg.Mongo); err != nil {
			log.Fatal(err)
		}
	case "down":
		if err := mongoConn.MigrateMngDown(&cfg.Mongo); err != nil {
			log.Fatal(err)
		}
	}
}
