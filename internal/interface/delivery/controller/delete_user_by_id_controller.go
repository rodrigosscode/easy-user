package controller

import (
	"encoding/json"
	"net/http"

	configs "github.com/rodrigosscode/easy-user/configs/http"
	usecase "github.com/rodrigosscode/easy-user/internal/application/usecase/user"
	userInput "github.com/rodrigosscode/easy-user/internal/application/usecase/user/input"
)

type DeleteUserByIdController struct {
	uc usecase.DeleteByIdUseCase
}

func NewDeleteUserByIdController(uc usecase.DeleteByIdUseCase) *DeleteUserByIdController {
	return &DeleteUserByIdController{uc: uc}
}

func (c *DeleteUserByIdController) Execute(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get(configs.QueryParamUserId)

	if userId == "" {
		// handler.HandleError(w, erros.NewInvalidRequestErr())
		return
	}

	i := &userInput.DeleteByIdInput{Id: userId}

	// ctx := r.Context()
	err := c.uc.Execute(i)

	if err != nil {
		// handler.HandleError(w, err)
		return
	}

	// logger.Info("Response Body", user)
	//response.NewSuccess(http.StatusNoContent,"Delete successfuly").Send(w)
	json.NewEncoder(w).Encode("Deletado com sucesso")
}
