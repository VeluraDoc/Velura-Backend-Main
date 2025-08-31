package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/VeluraDoc/Velura-Backend-Main/config"
	"github.com/dgrijalva/jwt-go"
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

func VerifyToken(tokenString string) (map[string]interface{}, error) {
	var jwtSecretKey = []byte(config.GetEnv("JWT_SECRET"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userData := make(map[string]interface{})
		userData["id"] = claims["id"]
		return userData, nil
	}

	return nil, errors.New("invalid token")
}
