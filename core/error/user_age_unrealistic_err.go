package error

import (
	"fmt"
)

type UserAgeUnrealisticErr struct {
	age int
}

func NewUserAgeUnrealisticErr(age int) error {
	return UserAgeUnrealisticErr{age: age}
}

func (e UserAgeUnrealisticErr) Error() string {
	return fmt.Sprintf("age must be realistic, %v", e.age)
}
