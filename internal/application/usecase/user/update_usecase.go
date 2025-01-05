package usecase

import (
	"github.com/rodrigosscode/easy-user/internal/application/usecase/user/input"
	domain "github.com/rodrigosscode/easy-user/internal/domain/entity"
	"github.com/rodrigosscode/easy-user/internal/domain/repository"
)

type (
	UpdateUseCase interface {
		Execute(i input.UpdateInput) (*domain.User, error)
	}
	updateUseCase struct {
		repository repository.UserRepository
	}
)

func NewUpdateUseCase(repository repository.UserRepository) UpdateUseCase {
	return &updateUseCase{repository: repository}
}

func (uc *updateUseCase) Execute(i input.UpdateInput) (*domain.User, error) {

	// // Validação inicial do input
	// if i.Id == "" {
	// 	return nil, fmt.Errorf("invalid input: ID cannot be empty")
	// }

	uToUpdate, err := uc.repository.FindById(i.Id)

	if err != nil {
		return nil, err
	}

	// TODO: avaliar para validação
	uToUpdate.Name = i.User.Name
	uToUpdate.Email = i.User.Email
	uToUpdate.Age = i.User.Age

	uUpdated, err := uc.repository.Update(uToUpdate)

	if err != nil {
		return nil, err
	}

	return uUpdated, nil
}
