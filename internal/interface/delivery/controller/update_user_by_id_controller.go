package controller

import (
	"net/http"

	usecase "github.com/rodrigosscode/easy-user/internal/application/usecase/user"
)

type UpdateUserByIdController struct {
	uc usecase.UpdateUseCase
}

func NewUpdateUserByIdController(uc usecase.UpdateUseCase) *UpdateUserByIdController {
	return &UpdateUserByIdController{uc: uc}
}

func (c *UpdateUserByIdController) Execute(w http.ResponseWriter, r *http.Request) {

}
