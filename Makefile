ifneq (,$(wildcard ./.env))
    include .env
    export
endif

createdb:
	docker exec -it bank-db createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${DB_NAME}

dropdb:
	docker exec -it bank-db dropdb ${DB_NAME}

migrateup:
	migrate -path db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migration -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" --verbose down

.PHONY: createdb dropdb migrateup migratedown