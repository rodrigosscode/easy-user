package usecase

import (
	"context"

	"github.com/rodrigosscode/easy-user/adapter/api/response"
	domain "github.com/rodrigosscode/easy-user/core/domain/entity"
	"github.com/rodrigosscode/easy-user/core/repository"
	"github.com/rodrigosscode/easy-user/core/usecase/user/input"
)

type (
	FindByPageUseCase interface {
		Execute(ctx *context.Context, i *input.FindByPageInput) (*response.Page[domain.User], error)
	}
	findByPageUseCase struct {
		repository repository.UserRepository
	}
)

func NewFindByPageUseCase(repository repository.UserRepository) FindByPageUseCase {
	return &findByPageUseCase{repository: repository}
}

func (uc *findByPageUseCase) Execute(ctx *context.Context, i *input.FindByPageInput) (*response.Page[domain.User], error) {
	users, totalRecords, err := uc.repository.FindByPage(ctx, i.Page, i.Limit)

	if err != nil {
		return nil, err
	}

	p := response.NewPage(*users, i.Page, i.Limit, totalRecords)

	return &p, nil
}
