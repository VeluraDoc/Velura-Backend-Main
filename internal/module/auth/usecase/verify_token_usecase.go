package user_usecase

import (
	"errors"

	"github.com/VeluraDoc/Velura-Backend-Main/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

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
