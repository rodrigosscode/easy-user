package response

import (
	"encoding/json"
	"net/http"

	configs "github.com/rodrigosscode/easy-user/configs/http"
)

type SuccessResponse struct {
	statusCode int
	body       interface{}
}

func NewSuccessResponse(statusCode int, body any) *SuccessResponse {
	return &SuccessResponse{statusCode: statusCode, body: body}
}

func (s *SuccessResponse) Send(w http.ResponseWriter) {
	w.Header().Add(configs.HeaderContentType, configs.MIMEApplicationJSON)
	w.WriteHeader(s.statusCode)
	if errEnconder := json.NewEncoder(w).Encode(s.body); errEnconder != nil {
		return
	}
}
