package usecase

import (
	"github.com/rodrigosscode/easy-user/internal/application/response"
	"github.com/rodrigosscode/easy-user/internal/application/usecase/user/input"
	domain "github.com/rodrigosscode/easy-user/internal/domain/entity"
	"github.com/rodrigosscode/easy-user/internal/domain/repository"
)

type (
	FindByPageUseCase interface {
		Execute(i *input.FindByPageInput) (*response.PageResponse[domain.User], error)
	}
	findByPageUseCase struct {
		repository repository.UserRepository
	}
)

func NewFindByPageUseCase(repository repository.UserRepository) FindByPageUseCase {
	return &findByPageUseCase{repository: repository}
}

func (uc *findByPageUseCase) Execute(i *input.FindByPageInput) (*response.PageResponse[domain.User], error) {
	users, totalRecords, err := uc.repository.FindByPage(i.Page, i.Limit)

	if err != nil {
		return nil, err
	}

	p := response.NewPageResponse(*users, i.Page, i.Limit, totalRecords)

	return &p, nil
}
