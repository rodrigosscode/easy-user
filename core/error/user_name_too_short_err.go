package error

import (
	"fmt"
)

type UserNameTooShortErr struct {
	name string
}

func NewUserNameTooShortErr(name string) error {
	return UserNameTooShortErr{name: name}
}

func (e UserNameTooShortErr) Error() string {
	return fmt.Sprintf("name must be at least 3 characters, %v", e.name)
}
