package controller

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	usecase "github.com/rodrigosscode/easy-user/internal/application/usecase/user"
	"github.com/rodrigosscode/easy-user/internal/application/usecase/user/input"
)

type SaveUserController struct {
	uc usecase.SaveUseCase
}

func NewSaveUserController(uc usecase.SaveUseCase) *SaveUserController {
	return &SaveUserController{uc: uc}
}

func (c *SaveUserController) Execute(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := io.ReadAll(r.Body)

	if err != nil {
		panic(errors.New("deu ruim"))
		//return
	}

	var input input.SaveInput
	if err := json.Unmarshal(jsonBody, &input); err != nil {
		panic(errors.New("deu ruim unmarshal"))
		// return
	}

	uSaved, err := c.uc.Execute(&input)

	if err != nil {
		panic(errors.New("deu ruim execute save usecase"))
		// return
	}

	json.NewEncoder(w).Encode(uSaved)
}
