package auth_usecase

import (
	"errors"

	user_dto "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/dto"
	user_model "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/model"
	user_repository "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/repository"
)

func SignInUser(dto user_dto.UserRequestDTO) (string, error) {

	if err := dto.Validate(); err != nil {
		return "", err
	}

	inputUser := user_model.User{
		Email:    dto.Email,
		Password: dto.Password,
	}

	var userRepo user_repository.UserRepository = user_repository.GetRepository()

	var user *user_model.User
	var err error
	if user, err = userRepo.FindByEmail(inputUser.Email); err != nil {
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
