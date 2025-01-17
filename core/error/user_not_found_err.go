package error

import (
	"fmt"
)

const (
	UserIdParam = "userId"
)

type UserNotFoundErr struct {
	id int
}

func NewUserNotFoundErr(id int) error {
	return UserNotFoundErr{id: id}
}

func (e UserNotFoundErr) Error() string {
	return fmt.Sprintf("user not found, id = %v", e.id)
}
