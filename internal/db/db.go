package db

type Storage interface {
	Create()
	Find()
	FindByID()
	UpdateByID()
	DeleteByID()
}
