package controller

import (
	"net/http"
	"strconv"

	"github.com/rodrigosscode/easy-user/adapter/api/handler"
	"github.com/rodrigosscode/easy-user/adapter/api/response"
	configs "github.com/rodrigosscode/easy-user/config/http"
	errs "github.com/rodrigosscode/easy-user/core/error"
	userUseCase "github.com/rodrigosscode/easy-user/core/usecase/user"
	userInput "github.com/rodrigosscode/easy-user/core/usecase/user/input"
	"github.com/rodrigosscode/easy-user/infrastructure/logger"
	"go.uber.org/zap"
)

type DeleteUserByIdController struct {
	uc userUseCase.DeleteByIdUseCase
}

func NewDeleteUserByIdController(uc userUseCase.DeleteByIdUseCase) *DeleteUserByIdController {
	return &DeleteUserByIdController{uc: uc}
}

func (c *DeleteUserByIdController) Execute(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Query().Get(configs.QueryParamUserId)
	userId, err := strconv.Atoi(userIdStr)

	if err != nil {
		logger.Error("Invalid user ID in request", zap.String("userId", userIdStr), zap.Error(err))
		handler.HandleError(w, errs.NewInvalidIdRequestErr(userId))
		return
	}

	i := &userInput.DeleteByIdInput{Id: userId}
	ctx := r.Context()
	err = c.uc.Execute(&ctx, i)

	if err != nil {
		logger.Error("Failed to delete user", zap.Int("userId", userId), zap.Error(err))
		handler.HandleError(w, err)
		return
	}

	logger.Info("User deleted successfully", zap.Int("userId", userId))
	response.NewSuccessResponse(http.StatusNoContent, "User deleted successfully").Send(w)
}
