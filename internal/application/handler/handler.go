package handler

import (
	"errors"
	"net/http"

	"github.com/rodrigosscode/easy-user/internal/application/response"
	domainErr "github.com/rodrigosscode/easy-user/internal/domain/error"
)

func HandleError(w http.ResponseWriter, err error) {
	var statusCode int

	badRequestErrors := []any{
		&domainErr.UserNameEmptyErr{},
		&domainErr.UserNameTooShortErr{},
		&domainErr.UserAgeNegativeErr{},
		&domainErr.UserAgeUnrealisticErr{},
		&domainErr.UserEmailUnavailableErr{},
		&domainErr.UserEmailEmptyErr{},
		&domainErr.UserEmailInvalidErr{},
		&domainErr.InvalidIdRequestErr{},
		&domainErr.InvalidPageRequestErr{},
	}

	for _, e := range badRequestErrors {
		if errors.As(err, e) {
			statusCode = http.StatusBadRequest
			response.NewErrorResponse(err, statusCode).Send(w)
			return
		}
	}

	switch {
	case errors.As(err, &domainErr.UserNotFoundErr{}):
		statusCode = http.StatusNotFound
	default:
		statusCode = http.StatusInternalServerError
	}

	response.NewErrorResponse(err, statusCode).Send(w)
}
