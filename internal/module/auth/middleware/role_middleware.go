package auth_middleware

import (
	"net/http"
	"slices"

	user_model "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/model"
	shared_dto "github.com/VeluraDoc/Velura-Backend-Main/internal/shared/dto"
	"github.com/gin-gonic/gin"
)

func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, shared_dto.ErrorResponseDto{Error: "user not found"})
			c.Abort()
			return
		}

		userRole := user.(*user_model.User).GetRole()

		if slices.Contains(roles, userRole) {
			c.Next()
			return
		}

		c.JSON(http.StatusForbidden, shared_dto.ErrorResponseDto{Error: "forbidden"})
		c.Abort()
	}
}
