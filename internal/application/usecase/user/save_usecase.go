package usecase

import (
	"github.com/rodrigosscode/easy-user/internal/application/usecase/user/input"
	domain "github.com/rodrigosscode/easy-user/internal/domain/entity"
	"github.com/rodrigosscode/easy-user/internal/domain/repository"
)

type (
	SaveUseCase interface {
		Execute(i *input.SaveInput) (*domain.User, error)
	}
	saveUseCase struct {
		repository repository.UserRepository
	}
)

func NewSaveUseCase(repository repository.UserRepository) SaveUseCase {
	return &saveUseCase{repository: repository}
}

func (uc *saveUseCase) Execute(i *input.SaveInput) (*domain.User, error) {

	if err := i.Validate(); err != nil {
		return nil, err
	}

	uToAdd := domain.NewUser(
		domain.WithName(i.Name),
		domain.WithEmail(i.Email),
		domain.WithAge(i.Age),
	)

	uSaved, err := uc.repository.Save(uToAdd)

	if err != nil {
		return nil, err
	}

	return uSaved, nil
}
