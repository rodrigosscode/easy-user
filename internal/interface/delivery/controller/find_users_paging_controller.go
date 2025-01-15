package controller

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	configs "github.com/rodrigosscode/easy-user/configs/http"
	"github.com/rodrigosscode/easy-user/internal/application/handler"
	"github.com/rodrigosscode/easy-user/internal/application/response"
	userUsecase "github.com/rodrigosscode/easy-user/internal/application/usecase/user"
	"github.com/rodrigosscode/easy-user/internal/application/usecase/user/input"
	domainErr "github.com/rodrigosscode/easy-user/internal/domain/error"
	"github.com/rodrigosscode/easy-user/internal/infrastructure/logger"
	"go.uber.org/zap"
)

const (
	DefaultPage  = 1
	DefaultLimit = 5
)

type FindUsersPagingController struct {
	uc userUsecase.FindByPageUseCase
}

func NewFindUsersPagingController(uc userUsecase.FindByPageUseCase) *FindUsersPagingController {
	return &FindUsersPagingController{uc: uc}
}

func (c *FindUsersPagingController) Execute(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get(configs.QueryParamPage)
	limitStr := r.URL.Query().Get(configs.QueryParamPageLimit)

	page, err := strconv.Atoi(pageStr)
	if err != nil || pageStr == "" {
		page = DefaultPage
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limitStr == "" {
		limit = DefaultLimit
	}

	i := &input.FindByPageInput{Page: page, Limit: limit}

	validate := validator.New()
	if err := validate.Struct(i); err != nil {
		logger.Error("Invalid pagination input", zap.Any("input", i), zap.Error(err))
		handler.HandleError(w, domainErr.NewInvalidPageRequestErr(pageStr, limitStr))
		return
	}

	pr, err := c.uc.Execute(i)

	if err != nil {
		logger.Error("Failed to find users paginated", zap.Any("input", i), zap.Error(err))
		handler.HandleError(w, err)
		return
	}

	logger.Info("Successfully retrieved paginated users", zap.Int("limit", limit), zap.Int("totalPages", int(pr.TotalPages)))
	response.NewSuccessResponse(http.StatusOK, pr).Send(w)
}
