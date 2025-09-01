package user_handler

import (
	"net/http"

	user_dto "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/dto"
	user_model "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/model"
	shared_dto "github.com/VeluraDoc/Velura-Backend-Main/internal/shared/dto"
	"github.com/gin-gonic/gin"
)

// @Summary      Get current user profile
// @Description  Returns profile info of authenticated user
// @Tags         User
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  user_dto.UserResponseDTO
// @Failure      401  {object}  shared_dto.ErrorResponseDto
// @Router       /user/me [get]
func GetMe(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, shared_dto.ErrorResponseDto{Error: "user not found"})
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
