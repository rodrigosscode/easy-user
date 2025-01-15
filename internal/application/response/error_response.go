package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	statusCode int
	errs       []string
}

func NewErrorResponse(err error, statusCode int) *ErrorResponse {
	return &ErrorResponse{statusCode: statusCode, errs: []string{err.Error()}}
}

func (e *ErrorResponse) Send(w http.ResponseWriter) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(e.statusCode)
	if errEnconder := json.NewEncoder(w).Encode(e.errs); errEnconder != nil {
		return
	}
}
