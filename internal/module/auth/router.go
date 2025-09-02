package auth_router

import (
	auth_handler "github.com/VeluraOpenSource/Velura_Documents_Service/internal/module/auth/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/sign-up", auth_handler.SignUp)
	r.POST("/login", auth_handler.Login)
}
