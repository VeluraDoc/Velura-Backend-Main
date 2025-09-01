package auth_handler

import (
	"net/http"

	user_usecase "github.com/VeluraDoc/Velura-Backend-Main/internal/module/auth/usecase"
	user_dto "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/dto"
	shared_dto "github.com/VeluraDoc/Velura-Backend-Main/internal/shared/dto"
	"github.com/gin-gonic/gin"
)

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

	c.JSON(http.StatusOK, gin.H{
		"message": "user registered successfully",
		"token":   token,
	})
}
