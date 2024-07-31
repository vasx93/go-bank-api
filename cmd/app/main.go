package main

import (
	"log"
	"net/http"
	"os"

	"github.com/vasx93/go-bank-api/config"
	"github.com/vasx93/go-bank-api/internal/db"
)

func main() {

	config.LoadEnv()

	dbType := os.Getenv("DB_TYPE")
	dbURI := os.Getenv("DB_URI")

	db, err := db.NewDBFactory(dbType, dbURI)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("db", db)

	http.ListenAndServe(":5000", nil)

}
