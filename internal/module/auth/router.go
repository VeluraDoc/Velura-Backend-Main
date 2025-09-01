package auth_router

import (
	auth_handler "github.com/VeluraDoc/Velura-Backend-Main/internal/module/auth/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/sign-up", auth_handler.SignUp)
	r.POST("/login", auth_handler.Login)
}
