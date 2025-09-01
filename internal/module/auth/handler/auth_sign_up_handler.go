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
// @Success      201  {object}  auth_dto.AuthResponseDTO
// @Failure      401  {object}  shared_dto.ErrorResponseDto
// @Router       /auth/sign-up [post]
func SignUp(c *gin.Context) {
	var dto user_dto.UserRequestDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, shared_dto.ErrorResponseDto{Error: "invalid input"})
		return
	}

	token, err := user_usecase.RegisterUser(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, shared_dto.ErrorResponseDto{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, auth_dto.AuthResponseDTO{
		Message: "user registered successfully",
		Token:   token,
	})
}
