package auth_usecase

import (
	"errors"

	user_dto "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/dto"
	user_model "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/model"
	user_repository "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/repository"
)

func RegisterUser(dto user_dto.UserRequestDTO) (string, error) {

	if err := dto.Validate(); err != nil {

		return "", err
	}

	user := user_model.User{
		Email:    dto.Email,
		Password: dto.Password,
	}

	if err := user.HashPassword(user.Password); err != nil {
		return "", errors.New("failed to hash password")
	}

	var userRepo user_repository.UserRepository = user_repository.GetRepository()

	if err := userRepo.Save(&user); err != nil {
		return "", errors.New("failed to save user")
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
