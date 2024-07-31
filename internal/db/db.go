package db

import (
	"errors"

	"github.com/vasx93/go-bank-api/internal/db/postgres"
)

type Storage interface {
	Create()
	Find()
	FindByID()
	UpdateByID()
	DeleteByID()
}

func NewDBFactory(dbType, uri string) (Storage, error) {

	switch dbType {
	case "postgres":
		return postgres.NewPostgresDB(uri), nil

	default:
		return nil, errors.New("unsupported database type")
	}

}
