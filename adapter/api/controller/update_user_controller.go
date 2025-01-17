package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rodrigosscode/easy-user/adapter/api/handler"
	"github.com/rodrigosscode/easy-user/adapter/api/response"
	userUseCase "github.com/rodrigosscode/easy-user/core/usecase/user"
	"github.com/rodrigosscode/easy-user/core/usecase/user/input"
	"github.com/rodrigosscode/easy-user/infrastructure/logger"
	"go.uber.org/zap"
)

type UpdateUserController struct {
	uc userUseCase.UpdateUseCase
}

func NewUpdateUserController(uc userUseCase.UpdateUseCase) *UpdateUserController {
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

	ctx := r.Context()
	uUpdated, err := c.uc.Execute(&ctx, &i)

	if err != nil {
		logger.Error("Failed to update user", zap.Any("input", i), zap.Error(err))
		handler.HandleError(w, err)
		return
	}

	logger.Info("User updated successfully", zap.Any("user", uUpdated))
	response.NewSuccessResponse(http.StatusAccepted, uUpdated).Send(w)
}
