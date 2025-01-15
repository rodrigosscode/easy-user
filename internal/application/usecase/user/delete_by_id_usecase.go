package usecase

import (
	"github.com/rodrigosscode/easy-user/internal/application/usecase/user/input"
	"github.com/rodrigosscode/easy-user/internal/domain/repository"
)

type (
	DeleteByIdUseCase interface {
		Execute(i *input.DeleteByIdInput) error
	}
	deleteByIdUseCase struct {
		repository repository.UserRepository
	}
)

func NewDeleteByIdUseCase(repository repository.UserRepository) DeleteByIdUseCase {
	return &deleteByIdUseCase{repository: repository}
}

func (uc *deleteByIdUseCase) Execute(i *input.DeleteByIdInput) error {
	if err := uc.repository.DeleteById(i.Id); err != nil {
		return err
	}

	return nil
}
