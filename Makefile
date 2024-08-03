ifneq (,$(wildcard ./.env))
    include .env
    export
endif

createdb:
	docker exec -it bank-db createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${DB_NAME}

dropdb:
	docker exec -it bank-db dropdb ${DB_NAME}

migrateup:
	migrate -path internal/db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" --verbose up

migratedown:
	migrate -path internal/db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" --verbose down


run:
	go run cmd/app/main.go

build:
	go build -o build/bank ./cmd/app

format:
	go fmt ./...


.PHONY: createdb dropdb migrateup migratedown run