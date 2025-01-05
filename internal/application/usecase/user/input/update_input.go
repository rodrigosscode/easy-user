package input

import domain "github.com/rodrigosscode/easy-user/internal/domain/entity"

type UpdateInput struct {
	Id   string      `json:"id"`
	User domain.User `json:"user"`
}
