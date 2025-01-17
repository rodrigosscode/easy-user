package input

import "github.com/rodrigosscode/easy-user/core/domain/validator"

type SaveInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func (i *SaveInput) Validate() error {

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
