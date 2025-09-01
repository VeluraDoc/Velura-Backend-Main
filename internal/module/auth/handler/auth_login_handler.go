package auth_handler

import (
	"net/http"

	auth_dto "github.com/VeluraDoc/Velura-Backend-Main/internal/module/auth/dto"
	user_usecase "github.com/VeluraDoc/Velura-Backend-Main/internal/module/auth/usecase"
	user_dto "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/dto"
	shared_dto "github.com/VeluraDoc/Velura-Backend-Main/internal/shared/dto"
	"github.com/gin-gonic/gin"
)

// @Summary      User Login
// @Description  Authenticates a user and returns a JWT token
// @Tags         Auth
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  auth_dto.AuthResponseDTO
// @Failure      401  {object}  shared_dto.ErrorResponseDto
// @Router       /auth/login [post]
func Login(c *gin.Context) {
	var dto user_dto.UserRequestDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, shared_dto.ErrorResponseDto{Error: "invalid input"})
		return
	}

	token, err := user_usecase.SignInUser(dto)
	if err != nil {
		c.JSON(http.StatusUnauthorized, shared_dto.ErrorResponseDto{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, auth_dto.AuthResponseDTO{
		Message: "Login successful",
		Token:   token,
	})
}
