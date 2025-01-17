package error

import (
	"fmt"
)

type UserAgeNegativeErr struct {
	age int
}

func NewUserAgeNegativeErr(age int) error {
	return UserAgeNegativeErr{age: age}
}

func (e UserAgeNegativeErr) Error() string {
	return fmt.Sprintf("age cannot be negative, %v", e.age)
}
