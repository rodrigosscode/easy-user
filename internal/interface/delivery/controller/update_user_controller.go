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

type UpdateUserController struct {
	uc userUsecase.UpdateUseCase
}

func NewUpdateUserController(uc userUsecase.UpdateUseCase) *UpdateUserController {
	return &UpdateUserController{uc: uc}
}

func (c *UpdateUserController) Execute(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := io.ReadAll(r.Body)

	if err != nil {
		logger.Error("Failed to read request body", zap.Error(err))
		handler.HandleError(w, err)
		return
	}

	var i input.UpdateInput
	if err := json.Unmarshal(jsonBody, &i); err != nil {
		logger.Error("Failed to unmarshal request body", zap.ByteString("requestBody", jsonBody), zap.Error(err))
		handler.HandleError(w, err)
		return
	}

	uUpdated, err := c.uc.Execute(&i)

	if err != nil {
		logger.Error("Failed to update user", zap.Any("input", i), zap.Error(err))
		handler.HandleError(w, err)
		return
	}

	logger.Info("User updated successfully", zap.Any("user", uUpdated))
	response.NewSuccessResponse(http.StatusAccepted, uUpdated).Send(w)
}
