package error

import (
	"fmt"
)

type UserEmailInvalidErr struct {
	email string
}

func NewUserEmailInvalidErr(email string) error {
	return UserEmailInvalidErr{email: email}
}

func (e UserEmailInvalidErr) Error() string {
	return fmt.Sprintf("email is not in a valid format, %v", e.email)
}
