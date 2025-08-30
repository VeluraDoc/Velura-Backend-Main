package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/VeluraDoc/Velura-Backend-Main/config"
	"github.com/dgrijalva/jwt-go"
)

var jwtSecretKey = []byte(config.GetEnv("JWT_SECRET"))
var jwtExpiration, _ = strconv.Atoi(config.GetEnv("JWT_EXPIRATION"))

func GenerateToken(id string) (string, error) {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Duration(jwtExpiration) * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

func VerifyToken(tokenString string) (*jwt.Token, map[string]interface{}, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userData := make(map[string]interface{})
		userData["id"] = claims["id"]
		userData["email"] = claims["email"]
		return token, userData, nil
	}

	return nil, nil, errors.New("invalid token")
}
