package controller

import (
	"net/http"
	"strconv"

	configs "github.com/rodrigosscode/easy-user/configs/http"
	"github.com/rodrigosscode/easy-user/internal/application/handler"
	"github.com/rodrigosscode/easy-user/internal/application/response"
	userUsecase "github.com/rodrigosscode/easy-user/internal/application/usecase/user"
	userInput "github.com/rodrigosscode/easy-user/internal/application/usecase/user/input"
	domainErr "github.com/rodrigosscode/easy-user/internal/domain/error"
	"github.com/rodrigosscode/easy-user/internal/infrastructure/logger"
	"go.uber.org/zap"
)

type DeleteUserByIdController struct {
	uc userUsecase.DeleteByIdUseCase
}

func NewDeleteUserByIdController(uc userUsecase.DeleteByIdUseCase) *DeleteUserByIdController {
	return &DeleteUserByIdController{uc: uc}
}

func (c *DeleteUserByIdController) Execute(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Query().Get(configs.QueryParamUserId)
	userId, err := strconv.Atoi(userIdStr)

	if err != nil {
		logger.Error("Invalid user ID in request", zap.String("userId", userIdStr), zap.Error(err))
		handler.HandleError(w, domainErr.NewInvalidIdRequestErr(userId))
		return
	}

	i := &userInput.DeleteByIdInput{Id: userId}

	// ctx := r.Context()
	err = c.uc.Execute(i)

	if err != nil {
		logger.Error("Failed to delete user", zap.Int("userId", userId), zap.Error(err))
		handler.HandleError(w, err)
		return
	}

	logger.Info("User deleted successfully", zap.Int("userId", userId))
	response.NewSuccessResponse(http.StatusNoContent, "User deleted successfully").Send(w)
}
