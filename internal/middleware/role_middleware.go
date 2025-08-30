package middleware

import (
	"net/http"
	"slices"

	"github.com/VeluraDoc/Velura-Backend-Main/internal/model"
	"github.com/gin-gonic/gin"
)

func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			c.Abort()
			return
		}

		user, err := model.GetUserByID(userID.(string))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			c.Abort()
			return
		}

		userRole := user.GetRole()

		if slices.Contains(roles, userRole) {
			c.Next()
			return
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		c.Abort()
	}
}
