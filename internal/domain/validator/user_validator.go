package validator

import (
	"errors"

	"github.com/rodrigosscode/easy-user/internal/util"
)

const (
	minNameLenght = 3
	maxAge        = 120
)

func ValidateName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	if len(name) < minNameLenght {
		return errors.New("name must be at least 3 characters")
	}
	return nil
}

func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}
	if !util.IsValidEmailFormat(email) { // Imagine uma função que valida o formato.
		return errors.New("email is not in a valid format")
	}
	return nil
}

func ValidateAge(age int) error {
	if age < 0 {
		return errors.New("age cannot be negative")
	}
	if age > maxAge {
		return errors.New("age must be realistic")
	}
	return nil
}
