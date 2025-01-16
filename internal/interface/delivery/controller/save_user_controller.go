package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rodrigosscode/easy-user/internal/application/handler"
	"github.com/rodrigosscode/easy-user/internal/application/response"
	userUsecase "github.com/rodrigosscode/easy-user/internal/application/usecase/user"
	"github.com/rodrigosscode/easy-user/internal/application/usecase/user/input"
	"github.com/rodrigosscode/easy-user/internal/infrastructure/logger"
	"go.uber.org/zap"
)

type SaveUserController struct {
	uc userUsecase.SaveUseCase
}

func NewSaveUserController(uc userUsecase.SaveUseCase) *SaveUserController {
	return &SaveUserController{uc: uc}
}

func (c *SaveUserController) Execute(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := io.ReadAll(r.Body)

	if err != nil {
		logger.Error("Failed to read request body", zap.Error(err))
		handler.HandleError(w, err)
		return
	}

	var i input.SaveInput
	if err := json.Unmarshal(jsonBody, &i); err != nil {
		logger.Error("Failed to unmarshal request body", zap.ByteString("requestBody", jsonBody), zap.Error(err))
		handler.HandleError(w, err)
		return
	}

	uSaved, err := c.uc.Execute(&i)

	if err != nil {
		logger.Error("Failed to save user", zap.Any("input", i), zap.Error(err))
		handler.HandleError(w, err)
		return
	}

	logger.Info("User saved successfully", zap.Any("user", uSaved))
	response.NewSuccessResponse(http.StatusAccepted, uSaved).Send(w)
}
