package usecase

import (
	"context"

	"github.com/rodrigosscode/easy-user/core/repository"
	"github.com/rodrigosscode/easy-user/core/usecase/user/input"
)

type (
	DeleteByIdUseCase interface {
		Execute(ctx *context.Context, i *input.DeleteByIdInput) error
	}
	deleteByIdUseCase struct {
		repository repository.UserRepository
	}
)

func NewDeleteByIdUseCase(repository repository.UserRepository) DeleteByIdUseCase {
	return &deleteByIdUseCase{repository: repository}
}

func (uc *deleteByIdUseCase) Execute(ctx *context.Context, i *input.DeleteByIdInput) error {
	if err := uc.repository.DeleteById(ctx, i.Id); err != nil {
		return err
	}

	return nil
}
