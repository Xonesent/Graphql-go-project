ifneq (,$(wildcard ./.env))
    include .env
    export
endif

CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
APP_MAIN=${CURDIR}/cmd/main.go

init_mongo:
	docker run -d --name ${MONGO_HOST} \
	-e MONGO_INITDB_ROOT_USERNAME=${MONGO_USER} \
	-e MONGO_INITDB_ROOT_PASSWORD=${MONGO_PASSWORD} \
	-p ${MONGO_PORT}:${MONGO_PORT} mongo:latest

init_postgresql:
	docker run -d --name ${POSTGRES_HOST} \
      -e POSTGRES_USER=${POSTGRES_USER} \
      -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
      -e POSTGRES_DB=${POSTGRES_DATABASE} \
      -p ${POSTGRES_PORT}:${POSTGRES_PORT} postgres:latest

install-migrates:
	@test -f ${BINDIR}/migrate-mongodb || GOBIN=${BINDIR} go install -tags 'mongodb' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@mv ${BINDIR}/migrate ${BINDIR}/migrate-mongodb
	@test -f ${BINDIR}/migrate-postgres || GOBIN=${BINDIR} go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@mv ${BINDIR}/migrate ${BINDIR}/migrate-postgres
	@echo "Migrate tools installed successfully in ./bin!"

create-mgrt-files:
	${BINDIR}/migrate-postgres create -ext sql -dir ./migrations/postgresql -seq create_tables

migrate-mongo-up:
	go run ./migrations/mongodb/migrate.go -method=up

migrate-mongo-down:
	go run ./migrations/mongodb/migrate.go -method=down

migrate-postgres-up:
	go run ./migrations/postgresql/migrate.go -method=up

migrate-postgres-down:
	go run ./migrations/postgresql/migrate.go -method=down