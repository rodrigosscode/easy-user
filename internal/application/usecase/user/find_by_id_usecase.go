package usecase

import (
	"github.com/rodrigosscode/easy-user/internal/application/usecase/user/input"
	domain "github.com/rodrigosscode/easy-user/internal/domain/entity"
	"github.com/rodrigosscode/easy-user/internal/domain/repository"
)

type (
	FindByIdUseCase interface {
		Execute(i *input.FindByIdInput) (*domain.User, error)
	}
	findByIdUseCase struct {
		repository repository.UserRepository
	}
)

func NewFindByIdUseCase(repository repository.UserRepository) FindByIdUseCase {
	return &findByIdUseCase{repository: repository}
}

func (uc *findByIdUseCase) Execute(i *input.FindByIdInput) (*domain.User, error) {
	u, err := uc.repository.FindById(i.Id)

	if err != nil {
		return nil, err
	}

	return u, nil
}
