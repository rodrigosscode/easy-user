package source

import (
	db "github.com/rodrigosscode/easy-user/internal/infrastructure/db/entity"
)

type UserDb interface {
	FindById(id int) (db.User, error)
	Save(u db.User) (db.User, error)
	DeleteById(id int) error
	Update(u db.User) (db.User, error)
	FindByPage(page, limit int) (_ []db.User, totalRecords int64, _ error)
}
