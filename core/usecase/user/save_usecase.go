package usecase

import (
	"context"

	domain "github.com/rodrigosscode/easy-user/core/domain/entity"
	"github.com/rodrigosscode/easy-user/core/repository"
	"github.com/rodrigosscode/easy-user/core/usecase/user/input"
)

type (
	SaveUseCase interface {
		Execute(ctx *context.Context, i *input.SaveInput) (*domain.User, error)
	}
	saveUseCase struct {
		repository repository.UserRepository
	}
)

func NewSaveUseCase(repository repository.UserRepository) SaveUseCase {
	return &saveUseCase{repository: repository}
}

func (uc *saveUseCase) Execute(ctx *context.Context, i *input.SaveInput) (*domain.User, error) {

	if err := i.Validate(); err != nil {
		return nil, err
	}

	uToAdd := domain.NewUser(
		domain.WithName(i.Name),
		domain.WithEmail(i.Email),
		domain.WithAge(i.Age),
	)

	uSaved, err := uc.repository.Save(ctx, uToAdd)

	if err != nil {
		return nil, err
	}

	return uSaved, nil
}
