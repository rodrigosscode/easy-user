package usecase

import (
	"context"

	domain "github.com/rodrigosscode/easy-user/core/domain/entity"
	"github.com/rodrigosscode/easy-user/core/repository"
	"github.com/rodrigosscode/easy-user/core/usecase/user/input"
)

type (
	UpdateUseCase interface {
		Execute(ctx *context.Context, i *input.UpdateInput) (*domain.User, error)
	}
	updateUseCase struct {
		repository repository.UserRepository
	}
)

func NewUpdateUseCase(repository repository.UserRepository) UpdateUseCase {
	return &updateUseCase{repository: repository}
}

func (uc *updateUseCase) Execute(ctx *context.Context, i *input.UpdateInput) (*domain.User, error) {

	if err := i.Validate(); err != nil {
		return nil, err
	}

	uToUpdate, err := uc.repository.FindById(ctx, i.Id)

	if err != nil {
		return nil, err
	}

	uToUpdate.Name = i.Name
	uToUpdate.Email = i.Email
	uToUpdate.Age = i.Age

	uUpdated, err := uc.repository.Update(ctx, uToUpdate)

	if err != nil {
		return nil, err
	}

	return uUpdated, nil
}
