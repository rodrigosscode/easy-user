package adapter

import (
	"context"

	domain "github.com/rodrigosscode/easy-user/core/domain/entity"
	errs "github.com/rodrigosscode/easy-user/core/error"
	"github.com/rodrigosscode/easy-user/infrastructure/db/source"
	"github.com/rodrigosscode/easy-user/infrastructure/mapper"
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

func (r *UserRepository) FindById(ctx *context.Context, id int) (*domain.User, error) {
	uDb, err := r.userDb.FindById(ctx, id)

	if err != nil {
		return nil, r.userErrorMapper.ToDomain(&err, map[string]interface{}{
			errs.UserIdParam: id,
		})
	}

	uFound := r.userMapper.ToDomain(&uDb)

	return uFound, nil
}

func (r *UserRepository) Save(ctx *context.Context, u *domain.User) (*domain.User, error) {
	dbUser := r.userMapper.ToEntity(u)

	uDbSaved, err := r.userDb.Save(ctx, *dbUser)

	if err != nil {
		return nil, r.userErrorMapper.ToDomain(&err, map[string]interface{}{
			errs.EmailParam: u.Email,
		})
	}

	uSaved := r.userMapper.ToDomain(&uDbSaved)

	return uSaved, nil
}

func (r *UserRepository) DeleteById(ctx *context.Context, id int) error {
	if err := r.userDb.DeleteById(ctx, id); err != nil {
		return r.userErrorMapper.ToDomain(&err, map[string]interface{}{
			errs.UserIdParam: id,
		})
	}
	return nil
}

func (r *UserRepository) Update(ctx *context.Context, u *domain.User) (*domain.User, error) {
	dbUser := r.userMapper.ToEntity(u)

	uDbUpdated, err := r.userDb.Update(ctx, *dbUser)

	if err != nil {
		return nil, r.userErrorMapper.ToDomain(&err, map[string]interface{}{
			errs.UserIdParam: u.ID,
			errs.EmailParam:  u.Email,
		})
	}

	uUpdated := r.userMapper.ToDomain(&uDbUpdated)

	return uUpdated, nil
}

func (r *UserRepository) FindByPage(ctx *context.Context, page, limit int) (*[]domain.User, int64, error) {
	dbUsers, totalRecords, err := r.userDb.FindByPage(ctx, page, limit)

	if err != nil {
		return nil, 0, r.userErrorMapper.ToDomain(&err, nil)
	}

	domainUsers := r.userMapper.ToDomains(&dbUsers)

	return domainUsers, totalRecords, nil
}
