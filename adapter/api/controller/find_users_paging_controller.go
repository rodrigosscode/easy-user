package controller

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/rodrigosscode/easy-user/adapter/api/handler"
	"github.com/rodrigosscode/easy-user/adapter/api/response"
	configs "github.com/rodrigosscode/easy-user/config/http"
	errs "github.com/rodrigosscode/easy-user/core/error"
	userUseCase "github.com/rodrigosscode/easy-user/core/usecase/user"
	"github.com/rodrigosscode/easy-user/core/usecase/user/input"
	"github.com/rodrigosscode/easy-user/infrastructure/logger"
	"go.uber.org/zap"
)

const (
	DefaultPage  = 1
	DefaultLimit = 5
)

type FindUsersPagingController struct {
	uc userUseCase.FindByPageUseCase
}

func NewFindUsersPagingController(uc userUseCase.FindByPageUseCase) *FindUsersPagingController {
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
		handler.HandleError(w, errs.NewInvalidPageRequestErr(pageStr, limitStr))
		return
	}

	ctx := r.Context()
	p, err := c.uc.Execute(&ctx, i)

	if err != nil {
		logger.Error("Failed to find users paginated", zap.Any("input", i), zap.Error(err))
		handler.HandleError(w, err)
		return
	}

	logger.Info("Successfully retrieved paginated users", zap.Int("limit", limit), zap.Int("totalPages", int(p.TotalPages)))
	response.NewSuccessResponse(http.StatusOK, p).Send(w)
}
