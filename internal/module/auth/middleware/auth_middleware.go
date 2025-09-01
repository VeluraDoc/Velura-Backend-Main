package auth_middleware

import (
	"net/http"
	"strings"

	"github.com/VeluraDoc/Velura-Backend-Main/internal/module/auth"
	user_model "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/model"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			c.Abort()
			return
		}

		parsedToken := bearerToken[1]

		claims, err := auth.VerifyToken(parsedToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		user, err := user_model.GetUserByID(claims["id"].(string))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			c.Abort()
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
