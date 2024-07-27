# go-bank-api

\*\*WIP: Bank API built with Go, PostgreSQL

## Setup project

1. Install docker/docker desktop
2. Pull postgres:latest image
3. Run container with through docker desktop or docker compose ( issues with port mappings on wsl)
   docker run --name bank-db -p 5432:5432 -e POSTGRES_USERNAME=root -e POSTGRES_PASSWORD=secret -d postgres

## Enter docker container

docker exec -it containerID

## Enter Postgres container

docker exec -it containerID psql -U root bank(db name)

## Postgres commands

1. createdb --username=root --owner=root bank(db name)
2. psql bank
3. dropdb bank
