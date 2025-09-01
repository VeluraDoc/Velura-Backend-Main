package auth_usecase

import (
	"strconv"
	"time"

	"github.com/VeluraDoc/Velura-Backend-Main/internal/config"
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
