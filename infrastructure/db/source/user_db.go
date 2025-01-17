package source

import (
	"context"

	db "github.com/rodrigosscode/easy-user/infrastructure/db/entity"
)

type UserDb interface {
	FindById(ctx *context.Context, id int) (db.User, error)
	Save(ctx *context.Context, u db.User) (db.User, error)
	DeleteById(ctx *context.Context, id int) error
	Update(ctx *context.Context, u db.User) (db.User, error)
	FindByPage(ctx *context.Context, page, limit int) (_ []db.User, totalRecords int64, _ error)
}
