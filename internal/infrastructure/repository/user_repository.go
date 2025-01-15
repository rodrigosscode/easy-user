package repository

import (
	domain "github.com/rodrigosscode/easy-user/internal/domain/entity"
	domainErr "github.com/rodrigosscode/easy-user/internal/domain/error"
	"github.com/rodrigosscode/easy-user/internal/infrastructure/db/source"
	"github.com/rodrigosscode/easy-user/internal/infrastructure/mapper"
)

type UserRepository struct {
	userDb          source.UserDb
	userMapper      mapper.UserMapper
	userErrorMapper mapper.UserErrorMapper
}

func NewUserRepository(userDb source.UserDb, userMapper mapper.UserMapper, userErrorMapper mapper.UserErrorMapper) *UserRepository {
	return &UserRepository{
		userDb:          userDb,
		userMapper:      userMapper,
		userErrorMapper: userErrorMapper,
	}
}

func (r *UserRepository) FindById(id int) (*domain.User, error) {
	uDb, err := r.userDb.FindById(id)

	if err != nil {
		return nil, r.userErrorMapper.ToDomain(&err, map[string]interface{}{
			domainErr.UserIdParam: id,
		})
	}

	uFound := r.userMapper.ToDomain(&uDb)

	return uFound, nil
}

func (r *UserRepository) Save(u *domain.User) (*domain.User, error) {
	dbUser := r.userMapper.ToEntity(u)

	uDbSaved, err := r.userDb.Save(*dbUser)

	if err != nil {
		return nil, r.userErrorMapper.ToDomain(&err, map[string]interface{}{
			domainErr.EmailParam: u.Email,
		})
	}

	uSaved := r.userMapper.ToDomain(&uDbSaved)

	return uSaved, nil
}

func (r *UserRepository) DeleteById(id int) error {
	if err := r.userDb.DeleteById(id); err != nil {
		return r.userErrorMapper.ToDomain(&err, map[string]interface{}{
			domainErr.UserIdParam: id,
		})
	}
	return nil
}

func (r *UserRepository) Update(u *domain.User) (*domain.User, error) {
	dbUser := r.userMapper.ToEntity(u)

	uDbUpdated, err := r.userDb.Update(*dbUser)

	if err != nil {
		return nil, r.userErrorMapper.ToDomain(&err, map[string]interface{}{
			domainErr.UserIdParam: u.ID,
			domainErr.EmailParam:  u.Email,
		})
	}

	uUpdated := r.userMapper.ToDomain(&uDbUpdated)

	return uUpdated, nil
}

func (r *UserRepository) FindByPage(page, limit int) (*[]domain.User, int64, error) {
	dbUsers, totalRecords, err := r.userDb.FindByPage(page, limit)

	if err != nil {
		return nil, 0, r.userErrorMapper.ToDomain(&err, nil)
	}

	domainUsers := r.userMapper.ToDomains(&dbUsers)

	return domainUsers, totalRecords, nil
}
