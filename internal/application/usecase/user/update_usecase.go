package usecase

import (
	"github.com/rodrigosscode/easy-user/internal/application/usecase/user/input"
	domain "github.com/rodrigosscode/easy-user/internal/domain/entity"
	"github.com/rodrigosscode/easy-user/internal/domain/repository"
)

type (
	UpdateUseCase interface {
		Execute(i *input.UpdateInput) (*domain.User, error)
	}
	updateUseCase struct {
		repository repository.UserRepository
	}
)

func NewUpdateUseCase(repository repository.UserRepository) UpdateUseCase {
	return &updateUseCase{repository: repository}
}

func (uc *updateUseCase) Execute(i *input.UpdateInput) (*domain.User, error) {

	if err := i.Validate(); err != nil {
		return nil, err
	}

	uToUpdate, err := uc.repository.FindById(i.Id)

	if err != nil {
		return nil, err
	}

	uToUpdate.Name = i.Name
	uToUpdate.Email = i.Email
	uToUpdate.Age = i.Age

	uUpdated, err := uc.repository.Update(uToUpdate)

	if err != nil {
		return nil, err
	}

	return uUpdated, nil
}
