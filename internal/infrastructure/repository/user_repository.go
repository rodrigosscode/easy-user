package repository

import (
	domain "github.com/rodrigosscode/easy-user/internal/domain/entity"
	"github.com/rodrigosscode/easy-user/internal/infrastructure/db/source"
	"github.com/rodrigosscode/easy-user/internal/infrastructure/mapper"
)

type UserRepository struct {
	userDb     source.UserDb
	userMapper mapper.UserMapper
}

func NewUserRepository(userDb source.UserDb, userMapper mapper.UserMapper) *UserRepository {
	return &UserRepository{
		userDb:     userDb,
		userMapper: userMapper,
	}
}

func (r *UserRepository) FindById(id string) (*domain.User, error) {
	uDb, err := r.userDb.FindById(id)

	if err != nil {
		return nil, err
	}

	uFound := r.userMapper.ToDomain(&uDb)

	return uFound, nil
}

func (r *UserRepository) Save(u *domain.User) (*domain.User, error) {
	dbUser := r.userMapper.ToEntity(u)

	uDbSaved, err := r.userDb.Save(*dbUser)

	if err != nil {
		return nil, err
	}

	uSaved := r.userMapper.ToDomain(&uDbSaved)

	return uSaved, nil
}

func (r *UserRepository) Delete(id string) error {
	err := r.userDb.Delete(id)
	return err
}

func (r *UserRepository) Update(u *domain.User) (*domain.User, error) {
	dbUser := r.userMapper.ToEntity(u)

	uDbUpdated, err := r.userDb.Update(*dbUser)

	if err != nil {
		return nil, err
	}

	uUpdated := r.userMapper.ToDomain(&uDbUpdated)

	return uUpdated, nil
}
