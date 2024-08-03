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

// Define a generic Storage interface
type StorageGeneric[T interface{}] interface {
	Create(data T) error
	Find(query interface{}) ([]T, error)
	FindByID(id string) (T, error)
	UpdateByID(id string, data T) error
	DeleteByID(id string) error
}

func NewDBFactory(dbType, uri string) (Storage, error) {

	switch dbType {
	case "postgres":
		return postgres.NewPostgresDB(uri), nil

	default:
		return nil, errors.New("unsupported database type")
	}

}
