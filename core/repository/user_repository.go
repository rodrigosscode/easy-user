package repository

import (
	"context"

	domain "github.com/rodrigosscode/easy-user/core/domain/entity"
)

type UserRepository interface {
	FindById(ctx *context.Context, id int) (*domain.User, error)
	Save(ctx *context.Context, u *domain.User) (*domain.User, error)
	DeleteById(ctx *context.Context, id int) error
	Update(ctx *context.Context, u *domain.User) (*domain.User, error)
	FindByPage(ctx *context.Context, page, pageSize int) (_ *[]domain.User, totalRecords int64, _ error)
}
