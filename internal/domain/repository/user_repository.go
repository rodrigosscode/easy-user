package repository

import (
	domain "github.com/rodrigosscode/easy-user/internal/domain/entity"
)

type UserRepository interface {
	FindById(id string) (*domain.User, error)
	Save(u *domain.User) (*domain.User, error)
	Delete(id string) error
	Update(u *domain.User) (*domain.User, error)
}
