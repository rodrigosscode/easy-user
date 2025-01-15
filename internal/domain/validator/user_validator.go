package validator

import (
	userError "github.com/rodrigosscode/easy-user/internal/domain/error"
	"github.com/rodrigosscode/easy-user/internal/util"
)

const (
	minNameLength = 3
	maxAge        = 120
)

func ValidateName(name string) error {
	if name == "" {
		return userError.NewUserNameEmptyErr()
	}
	if len(name) < minNameLength {
		return userError.NewUserNameTooShortErr(name)
	}
	return nil
}

func ValidateEmail(email string) error {
	if email == "" {
		return userError.NewUserEmailEmptyErr()
	}
	if !util.IsValidEmailFormat(email) {
		return userError.NewUserEmailInvalidErr(email)
	}
	return nil
}

func ValidateAge(age int) error {
	if age < 0 {
		return userError.NewUserAgeNegativeErr(age)
	}
	if age > maxAge {
		return userError.NewUserAgeUnrealisticErr(age)
	}
	return nil
}
