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
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	db, err := db.NewDBFactory()

	if err != nil {
		log.Fatal(err)
	}

	log.Print("db", db)

	http.ListenAndServe(":"+port, nil)

}
