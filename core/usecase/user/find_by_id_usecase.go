package usecase

import (
	"context"

	domain "github.com/rodrigosscode/easy-user/core/domain/entity"
	"github.com/rodrigosscode/easy-user/core/repository"
	"github.com/rodrigosscode/easy-user/core/usecase/user/input"
)

type (
	FindByIdUseCase interface {
		Execute(ctx *context.Context, i *input.FindByIdInput) (*domain.User, error)
	}
	findByIdUseCase struct {
		repository repository.UserRepository
	}
)

func NewFindByIdUseCase(repository repository.UserRepository) FindByIdUseCase {
	return &findByIdUseCase{repository: repository}
}

func (uc *findByIdUseCase) Execute(ctx *context.Context, i *input.FindByIdInput) (*domain.User, error) {
	u, err := uc.repository.FindById(ctx, i.Id)

	if err != nil {
		return nil, err
	}

	return u, nil
}
