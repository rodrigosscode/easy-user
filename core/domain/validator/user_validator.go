package validator

import (
	errs "github.com/rodrigosscode/easy-user/core/error"
)

const (
	minNameLength = 3
	maxAge        = 120
)

func ValidateName(name string) error {
	if name == "" {
		return errs.NewUserNameEmptyErr()
	}
	if len(name) < minNameLength {
		return errs.NewUserNameTooShortErr(name)
	}
	return nil
}

func ValidateEmail(email string) error {
	if email == "" {
		return errs.NewUserEmailEmptyErr()
	}
	if !IsValidEmailFormat(email) {
		return errs.NewUserEmailInvalidErr(email)
	}
	return nil
}

func ValidateAge(age int) error {
	if age < 0 {
		return errs.NewUserAgeNegativeErr(age)
	}
	if age > maxAge {
		return errs.NewUserAgeUnrealisticErr(age)
	}
	return nil
}
