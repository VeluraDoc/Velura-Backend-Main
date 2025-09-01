package user

import (
	"net/http"

	user_dto "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/dto"
	user_model "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/model"
	"github.com/gin-gonic/gin"
)

func GetMe(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		c.Abort()
		return
	}

	userStruct := user.(*user_model.User)

	c.JSON(http.StatusOK, user_dto.UserResponseDTO{
		Email: userStruct.Email,
		ID:    userStruct.ID,
		Role:  userStruct.Role.String(),
	})
}
