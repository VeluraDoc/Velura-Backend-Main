package auth_middleware

import (
	"net/http"
	"strings"

	auth_usecase "github.com/VeluraOpenSource/Velura_Documents_Service/internal/module/auth/usecase"
	user_repository "github.com/VeluraOpenSource/Velura_Documents_Service/internal/module/user/repository"
	shared_dto "github.com/VeluraOpenSource/Velura_Documents_Service/internal/shared/dto"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, shared_dto.ErrorResponseDto{Error: "authorization header is required"})
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			c.JSON(http.StatusUnauthorized, shared_dto.ErrorResponseDto{Error: "invalid token format"})
			c.Abort()
			return
		}

		parsedToken := bearerToken[1]

		claims, err := auth_usecase.VerifyToken(parsedToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, shared_dto.ErrorResponseDto{Error: "invalid or expired token"})
			c.Abort()
			return
		}

		var userRepo user_repository.UserRepository = user_repository.GetRepository()

		user, err := userRepo.FindByID(claims["id"].(string))

		if err != nil {
			c.JSON(http.StatusUnauthorized, shared_dto.ErrorResponseDto{Error: "user not found"})
			c.Abort()
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
