package mapper

import (
	domain "github.com/rodrigosscode/easy-user/core/domain/entity"
	db "github.com/rodrigosscode/easy-user/infrastructure/db/entity"
)

type (
	UserMapper interface {
		ToDomain(u *db.User) *domain.User
		ToDomains(us *[]db.User) *[]domain.User
		ToEntity(u *domain.User) *db.User
	}
	userMapper struct{}
)

func NewUserMapper() UserMapper {
	return &userMapper{}
}

func (m *userMapper) ToDomains(us *[]db.User) *[]domain.User {
	domainUsers := make([]domain.User, len(*us))

	for i, u := range *us {
		domainUsers[i] = *m.ToDomain(&u)
	}

	return &domainUsers
}

func (m *userMapper) ToDomain(u *db.User) *domain.User {
	return &domain.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Age:   u.Age,
	}
}

func (m *userMapper) ToEntity(u *domain.User) *db.User {
	return &db.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Age:   u.Age,
	}
}
