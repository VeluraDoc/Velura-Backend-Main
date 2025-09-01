package auth_handler

import (
	"net/http"

	user_usecase "github.com/VeluraDoc/Velura-Backend-Main/internal/module/auth/usecase"
	user_dto "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/dto"
	shared_dto "github.com/VeluraDoc/Velura-Backend-Main/internal/shared/dto"
	"github.com/gin-gonic/gin"
)

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

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
