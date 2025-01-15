package repository

import (
	domain "github.com/rodrigosscode/easy-user/internal/domain/entity"
)

type UserRepository interface {
	FindById(id int) (*domain.User, error)
	Save(u *domain.User) (*domain.User, error)
	DeleteById(id int) error
	Update(u *domain.User) (*domain.User, error)
	FindByPage(page, pageSize int) (_ *[]domain.User, totalRecords int64, _ error)
}
