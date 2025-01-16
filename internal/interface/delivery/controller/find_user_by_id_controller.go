package controller

import (
	"net/http"
	"strconv"

	configs "github.com/rodrigosscode/easy-user/configs/http"
	"github.com/rodrigosscode/easy-user/internal/application/handler"
	"github.com/rodrigosscode/easy-user/internal/application/response"
	userUseCase "github.com/rodrigosscode/easy-user/internal/application/usecase/user"
	userInput "github.com/rodrigosscode/easy-user/internal/application/usecase/user/input"
	domainErr "github.com/rodrigosscode/easy-user/internal/domain/error"
	"github.com/rodrigosscode/easy-user/internal/infrastructure/logger"
	"go.uber.org/zap"
)

type FindUserByIdController struct {
	uc userUseCase.FindByIdUseCase
}

func NewFindUserByIdController(uc userUseCase.FindByIdUseCase) *FindUserByIdController {
	return &FindUserByIdController{uc: uc}
}

func (c *FindUserByIdController) Execute(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Query().Get(configs.QueryParamUserId)
	userId, err := strconv.Atoi(userIdStr)

	if err != nil {
		logger.Error("Invalid user ID in request", zap.String("userId", userIdStr), zap.Error(err))
		handler.HandleError(w, domainErr.NewInvalidIdRequestErr(userId))
		return
	}

	i := &userInput.FindByIdInput{Id: userId}
	userFound, err := c.uc.Execute(i)

	if err != nil {
		logger.Error("Failed to find user by id", zap.Int("userId", userId), zap.Error(err))
		handler.HandleError(w, err)
		return
	}

	logger.Info("User found successfully", zap.Int("userId", userId), zap.Any("user", userFound))
	response.NewSuccessResponse(http.StatusOK, userFound).Send(w)
}
