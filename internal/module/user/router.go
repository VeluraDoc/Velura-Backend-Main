package user_router

import (
	auth_middleware "github.com/VeluraOpenSource/Velura_Documents_Service/internal/module/auth/middleware"
	user_handler "github.com/VeluraOpenSource/Velura_Documents_Service/internal/module/user/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	protected := r.Group("/")
	protected.Use(auth_middleware.AuthMiddleware())

	protected.GET("/me", user_handler.GetMe)
}
