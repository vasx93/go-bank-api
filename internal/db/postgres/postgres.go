package postgres

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB() (*PostgresDB, error) {

	dsn := os.Getenv("POSTGRES_DSN")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Cannot open Postgres DB")
	}

	if err = db.Ping(); err != nil {
		log.Fatal("No response from db")
	}

	return &PostgresDB{db: db}, nil
}

func (p *PostgresDB) Close() error {
	return p.db.Close()
}

func (s *PostgresDB) Create() {
	// Implement create logic
}

func (s *PostgresDB) Find() {
	// Implement find logic
}

func (s *PostgresDB) FindByID() {
	// Implement find by ID logic
}

func (s *PostgresDB) UpdateByID() {
	// Implement update by ID logic
}

func (s *PostgresDB) DeleteByID() {
	// Implement delete by ID logic
}
