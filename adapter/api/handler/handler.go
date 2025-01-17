package handler

import (
	"errors"
	"net/http"

	"github.com/rodrigosscode/easy-user/adapter/api/response"
	errs "github.com/rodrigosscode/easy-user/core/error"
)

func HandleError(w http.ResponseWriter, err error) {
	var statusCode int

	badRequestErrors := []any{
		&errs.UserNameEmptyErr{},
		&errs.UserNameTooShortErr{},
		&errs.UserAgeNegativeErr{},
		&errs.UserAgeUnrealisticErr{},
		&errs.UserEmailUnavailableErr{},
		&errs.UserEmailEmptyErr{},
		&errs.UserEmailInvalidErr{},
		&errs.InvalidIdRequestErr{},
		&errs.InvalidPageRequestErr{},
	}

	for _, e := range badRequestErrors {
		if errors.As(err, e) {
			statusCode = http.StatusBadRequest
			response.NewErrorResponse(err, statusCode).Send(w)
			return
		}
	}

	switch {
	case errors.As(err, &errs.UserNotFoundErr{}):
		statusCode = http.StatusNotFound
	default:
		statusCode = http.StatusInternalServerError
	}

	response.NewErrorResponse(err, statusCode).Send(w)
}
