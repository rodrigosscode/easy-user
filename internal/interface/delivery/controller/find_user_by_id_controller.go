package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	configs "github.com/rodrigosscode/easy-user/configs/http"
	userUseCase "github.com/rodrigosscode/easy-user/internal/application/usecase/user"
	userInput "github.com/rodrigosscode/easy-user/internal/application/usecase/user/input"
)

type FindUserByIdController struct {
	uc userUseCase.FindByIdUseCase
}

func NewFindUserByIdController(uc userUseCase.FindByIdUseCase) *FindUserByIdController {
	return &FindUserByIdController{uc: uc}
}

func (c *FindUserByIdController) Execute(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get(configs.QueryParamUserId)

	if userId == "" {
		// handler.HandleError(w, erros.NewInvalidRequestErr())
		return
	}

	i := &userInput.FindByIdInput{Id: userId}
	// ctx := r.Context()
	userFound, err := c.uc.Execute(i)

	if err != nil {
		panic(errors.New("deu ruim"))
		// handler.HandleError(w, err)
		//return
	}

	// logger.Info("Response Body", user)
	// response.NewSucess(http.StatusOK, user).Send(w)
	json.NewEncoder(w).Encode(userFound)
}
