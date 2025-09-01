package auth

import (
	"errors"
	"strconv"
	"time"

	"github.com/VeluraDoc/Velura-Backend-Main/internal/config"
	user_dto "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/dto"
	user_model "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/model"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id string) (string, error) {
	var jwtSecretKey = []byte(config.GetEnv("JWT_SECRET"))
	var jwtExpiration, _ = strconv.Atoi(config.GetEnv("JWT_EXPIRATION"))

	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Duration(jwtExpiration) * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	var jwtSecretKey = []byte(config.GetEnv("JWT_SECRET"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

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

	if err := user.Save(); err != nil {
		return "", errors.New("failed to save user")
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
