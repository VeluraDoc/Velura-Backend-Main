package user_usecase

import (
	"errors"

	user_dto "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/dto"
	user_model "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/model"
)

func SignInUser(dto user_dto.UserRequestDTO) (string, error) {

	if err := dto.Validate(); err != nil {
		return "", err
	}

	inputUser := user_model.User{
		Email:    dto.Email,
		Password: dto.Password,
	}

	var user user_model.User
	var err error
	if user, err = user_model.GetUserByEmail(inputUser.Email); err != nil {
		return "", errors.New("invalid email or password")
	}

	if !user.CheckPassword(inputUser.Password) {
		return "", errors.New("invalid email or password")
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
