package user_dto

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type UserRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (u *UserRequestDTO) Validate() error {
	if err := validate.Struct(u); err != nil {
		return err
	}

	return nil
}
