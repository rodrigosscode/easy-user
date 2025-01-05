package db

import (
	db "github.com/rodrigosscode/easy-user/internal/infrastructure/db/entity"
)

type DbConnection struct {
}

func NewDbConnection() (*DbConnection, error) {
	return &DbConnection{}, nil
}

func (d *DbConnection) FindById(id string) (db.User, error) {

	return db.User{
		ID: "1", Name: "Rodrigo", Email: "email@email.com", Age: 24,
	}, nil
}

func (d *DbConnection) Save(u db.User) (db.User, error) {
	return db.User{
		ID: "2", Name: "Rodrigo sdasd", Email: "email@emaiasasl.com", Age: 25,
	}, nil
}

func (d *DbConnection) Delete(id string) error {
	return nil
}

func (d *DbConnection) Update(u db.User) (db.User, error) {
	return db.User{
		ID: "2", Name: "Rodrigo sdasd", Email: "email@emaiasasl.com", Age: 25,
	}, nil
}
