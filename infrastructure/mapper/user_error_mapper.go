package mapper

import (
	"errors"

	mysqlErrors "github.com/go-mysql/errors"
	errs "github.com/rodrigosscode/easy-user/core/error"
	"gorm.io/gorm"
)

type (
	UserErrorMapper interface {
		ToDomain(err *error, params map[string]interface{}) error
	}
	userErrorMapper struct{}
)

func NewUserErrorMapper() UserErrorMapper {
	return &userErrorMapper{}
}

func (m *userErrorMapper) ToDomain(err *error, params map[string]interface{}) error {

	if ok, err := mysqlErrors.Error(*err); ok {
		return handleMySQLErrors(err, params)
	} else {
		return handleGormErrors(err, params)
	}
}

func handleMySQLErrors(err error, params map[string]interface{}) error {

	if errors.Is(err, mysqlErrors.ErrDupeKey) {
		if email, ok := params[errs.EmailParam].(string); ok {
			return errs.NewUserEmailUnavailableErr(email)
		}
	}

	return err
}

func handleGormErrors(err error, params map[string]interface{}) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if id, ok := params[errs.UserIdParam].(int); ok {
			return errs.NewUserNotFoundErr(id)
		}
	} else if errors.Is(err, gorm.ErrDuplicatedKey) {
		if email, ok := params[errs.EmailParam].(string); ok {
			return errs.NewUserEmailUnavailableErr(email)
		}
	}

	return err
}
