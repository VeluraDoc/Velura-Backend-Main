package auth

import (
	"net/http"

	user_usecase "github.com/VeluraDoc/Velura-Backend-Main/internal/module/auth/usecase"
	user_dto "github.com/VeluraDoc/Velura-Backend-Main/internal/module/user/dto"
	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	var dto user_dto.UserRequestDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	token, err := user_usecase.RegisterUser(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user registered successfully",
		"token":   token,
	})
}

func LoginHandler(c *gin.Context) {
	var dto user_dto.UserRequestDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	token, err := user_usecase.SignInUser(dto)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
