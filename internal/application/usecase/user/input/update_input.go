package input

import "github.com/rodrigosscode/easy-user/internal/domain/validator"

type UpdateInput struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func (i *UpdateInput) Validate() error {

	if err := validator.ValidateName(i.Name); err != nil {
		return err
	}

	if err := validator.ValidateEmail(i.Email); err != nil {
		return err
	}

	if err := validator.ValidateAge(i.Age); err != nil {
		return err
	}

	return nil
}
