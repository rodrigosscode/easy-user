package error

import (
	"fmt"
)

const (
	EmailParam = "email"
)

type UserEmailUnavailableErr struct {
	email string
}

func NewUserEmailUnavailableErr(email string) error {
	return UserEmailUnavailableErr{email: email}
}

func (e UserEmailUnavailableErr) Error() string {
	return fmt.Sprintf("email is already associated with another user, %v", e.email)
}
