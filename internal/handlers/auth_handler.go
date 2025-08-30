package handlers

import (
	"net/http"

	"github.com/VeluraDoc/Velura-Backend-Main/internal/model"
	"github.com/VeluraDoc/Velura-Backend-Main/internal/service"
	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	if err := user.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save user"})
		return
	}

	token, err := service.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user registered successfully",
		"token":   token,
	})
}
