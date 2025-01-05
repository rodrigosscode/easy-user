package source

import (
	db "github.com/rodrigosscode/easy-user/internal/infrastructure/db/entity"
)

type UserDb interface {
	FindById(id string) (db.User, error)
	Save(u db.User) (db.User, error)
	Delete(id string) error
	Update(u db.User) (db.User, error)
}
